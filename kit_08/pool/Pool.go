package pool

import (
	"errors"
	"net"
	"sync"
	"time"
)

var (
	ErrInvaliaConfig = errors.New("invalid pool conf")
	ErrPoolClosed    = errors.New("pool closed")
)

type Pool interface {
	Acquire() (net.Conn, error) //获取资源
	Release( net.Conn) error     //释放资源
	Close( net.Conn) error       //关闭资源
	Shutdown() error             //关闭池
}

type factory func() (net.Conn, error)
type GenericPool struct {
	sync.Mutex
	pool        chan net.Conn
	maxOpen     int  //池中最大资源数
	numOpen     int  //当前池中资源数
	minOpen     int  //池中最少资源数
	closed      bool //池是否已关闭
	maxLifetime time.Duration
	factory     factory //创建链接的方法
}

func NewGenericPool(minOpen, maxOpen int, maxLifetime time.Duration, factory factory) (*GenericPool, error) {
	if maxOpen <= 0 || minOpen > maxOpen {
		return nil, ErrInvaliaConfig
	}
	p := &GenericPool{maxOpen: maxOpen,
		minOpen:     minOpen,
		maxLifetime: maxLifetime,
		factory:     factory,
		pool:        make(chan net.Conn, maxOpen),
	}

	for i := 0; i < minOpen; i++ {
		closer, err := factory()
		if err != nil {
			continue
		}
		p.numOpen++
		p.pool <- closer
	}
	return p, nil
}

func (p *GenericPool) Acquire() (net.Conn, error) {
	if p.closed {
		return nil, ErrPoolClosed
	}
	for {
		closer, err := p.getOrCreate()
		if err != nil {
			return nil, err
		}
		return closer, nil
	}
}

func (p *GenericPool) getOrCreate() (net.Conn, error) {
	select {
	case closer := <-p.pool:
		return closer, nil
	default:
	}
	p.Lock()
	defer p.Unlock()
	if p.numOpen > p.maxOpen {
		closer := <-p.pool
		return closer, nil
	}
	//新建链接
	closer, err := p.factory()
	if err != nil {
		return nil, err
	}
	p.numOpen++
	return closer, nil
}

// Release 释放单个资源到链接池
func (p *GenericPool) Release(closer net.Conn) error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	defer p.Unlock()
	p.pool <- closer
	return nil
}

func (p *GenericPool) Close(closer net.Conn) error {
	p.Lock()
	defer p.Unlock()
	closer.Close()
	p.numOpen--
	p.Unlock()
	return nil
}

func (p *GenericPool) Shutdown() error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	defer p.Unlock()

	for closer := range p.pool {
		closer.Close()
		p.numOpen--
	}
	p.closed = true
	return nil
}
