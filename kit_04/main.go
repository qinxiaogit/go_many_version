package main

import (
	"context"
	"fmt"
	"github.com/bamzi/jobrunner"
	"github.com/go-redis/redis/v8"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"time"
)
var (
	redisClient *redis.Client
)
func main()  {

	//go func() {
		profiler.Start(profiler.Config{ApplicationName: "test",ServerAddress: "http://192.168.211.65:4040"})

		//for i:=0;i<1000;i++ {
		//	time.Sleep(time.Second*10)
		//}
	//}()

	redisClient = redis.NewClient(&redis.Options{
		Addr: "172.16.2.121:27005",
		DB: 15,
	})
	fmt.Println(redisClient.Set(context.Background(),"2","1",1000*time.Second).Result())
	//r := gin.Default()
	jobrunner.Start()

	jobrunner.Schedule("@every 5s", ReminderEmails{})
	//jobrunner.Schedule("@every 1m", Input{})
	//jobrunner.Schedule("@reboot", OutPut{})
	//r.GET("/list", func(context *gin.Context) {
	//	context.JSON(http.StatusOK,gin.H{"data":"aaas","list":jobrunner.StatusJson()})
	//})
	//r.GET("/buy", func(context *gin.Context) {
	//	var out = new(Input)
	//	out.Run()
	//	context.JSON(http.StatusOK,gin.H{"data":"aaas","list":jobrunner.StatusJson()})
	//})
	//
	//r.Run()
	Run()
}

// ReminderEmails Job Specific Functions
type ReminderEmails struct {
	// filtered
}

// Run ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	//fmt.Printf("Every 5 sec send reminder emails \n",time.Now().Second())
}
