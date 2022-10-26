package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	/*
		fmt.Println("start")

		newUrl := "/wcs-doctor/api/doctor/findList?version=1.0.4&sort=sort&client=ios&offlinePositionIds=74499fb094a341ca92e6afb8777fa65d_117&appointmentEndDate=2022-04-28&noncestr=6DF07F55-9EBF-4567-98BC-1B3280C402AA&appointmentEnabled=1&timestamp=1650855331741&hospitalId=74499fb094a341ca92e6afb8777fa65d&appointmentStartDate=2022-04-28"

		obj, _ := url.Parse(newUrl)
		//fmt.Println(obj)
		m ,_ := url.ParseQuery(obj.RawQuery)

		fmt.Println(obj.RawQuery)
		for k,v := range m{
			fmt.Println(k,v)
		}

		fmt.Println(fmt.Sprintf("%x",md5.Sum([]byte(obj.RawQuery))))

	*/
	//numbs := []int{1,2,3,4,5,6,99,66,980,98,9800}

	//fmt.Println(largestNumber(numbs))

	//newNumber := [] int{1,8,6,2,5,4,8,3,7}

	//fmt.Println(singleNumbers(newNumber))

	//job1需添加标记 tag1，tag2，tag3
	//job2需添加标记tag3，tag4，tag5
	//job3需添加标记tag2，tag3
	//job4需添加标记tag1

	//每次处理只能给job添加相同的标记或者标记相同的job 且不能重复添加标记和添加不存在的标记，上例直接执行需要四次处理
	tasks := make(map[string][]string)
	taskJob1 := []string{"tag1", "tag2", "tag3"}
	tasks["job1"] = taskJob1
	tasks["job2"] = []string{"tag3", "tag4", "tag5"}
	tasks["job3"] = []string{"tag2", "tag3"}
	tasks["job4"] = []string{"tag1","tag2"}
	tasks["job5"] = []string{"tag3", "tag4"}
	tasks["job6"] = []string{"tag5", "tag6"}
	tasks["job7"] = []string{"tag5", "tag4"}
	tasks["job8"] = []string{"tag6", "tag2"}

	initCoordinate(tasks)

}

//
//func maxArea(height []int) int {
//	maxLen := len(height)
//	for i {
//
//	}
//}
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}

		for sy <= y {
			sy *= 10
		}
		fmt.Println(sy*x+y, sx*y+x, x, y)
		return sy*x+y > sx*y+x
	})

	if nums[0] == 0 {
		return "0"
	}
	fmt.Println(nums)
	ret := ""
	for i := 0; i < len(nums); i++ {
		ret += strconv.Itoa(nums[i])
	}

	return ret
}

func singleNumbers(nums []int) []int {

	if len(nums) == 2 {
		return nums
	}

	x := nums[0]
	for i := 1; i < len(nums); i++ {
		x ^= nums[i]
	}
	c := x & (-x)
	fmt.Println(x, -x, c)

	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		x = nums[i]
		if x&c == c {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}

	return []int{a, b}
}

/**
 * 有这样一个批任务标记
比如
	job1需添加标记 tag1，tag2，tag3
	job2需添加标记tag3，tag4，tag5
	job3需添加标记tag2，tag3
	job4需添加标记tag1

 job1 job2 job3 tag3
 job
	每次处理只能给job添加相同的标记或者标记相同的job 且不能重复添加标记和添加不存在的标记，上例直接执行需要四次处理
	如果合并处理成如下
	job1 job3  一同添加标记 tag2，tag3
	job1 job4  一同添加标记 tag1
	job2 添加标记 tag3，tag4，tag5
	则只需执行三次
*/
func findMinTask(task map[string][]string) {
	myNode := Node{}
	for k, v := range task {
		for _, v0 := range v {
			incrNode(&myNode, k, v0)
			fmt.Println(myNode)
		}
	}
	fmt.Println(myNode)
}

func sortNode(node *Node) {
	node = node.Next
	for true {
		if node == nil || node.Next == nil {
			break
		}
		if node.Count < node.Next.Count {
			tempCount := node.Count
			tempJobs := node.Jobs
			tempTags := node.Tag

			node.Count = node.Next.Count
			node.Jobs = node.Next.Jobs
			node.Tag = node.Next.Tag

			node.Next.Count = tempCount
			node.Next.Jobs = tempJobs
			node.Next.Tag = tempTags
		}

		node = node.Next
	}
}

func incrNode(node *Node, job string, tag string) bool {
	//新增节点
	myNode := Node{Next: nil, Jobs: []string{job}, Tag: tag, Count: 1}
	for true {
		if node.Jobs == nil {
			node.Jobs = myNode.Jobs
			node.Tag = myNode.Tag
			node.Count = 1
			return true
		}
		if node.Tag == tag {
			node.Jobs = append(node.Jobs, job)
			node.Count += 1
			return false
		}
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	node.Next = &myNode
	return true
}

type Node struct {
	Tag   string
	Jobs  []string
	Count int
	Next  *Node
}

//坐标点
type Coordinate struct {
	TagName string
	//TagIndex int
	JobName string
	//JobIndex int
}

/**
 * 初始化成坐标点 （tag1，job1）
 */
func initCoordinate(task map[string][]string) []Coordinate {
	var coor []Coordinate
	for k, v := range task {
		for _, v0 := range v {
			coor = append(coor, Coordinate{JobName: k, TagName: v0})
		}
	}
	cateCoordinate := map[string][]Coordinate{}
	for _, v := range coor {
		current := cateCoordinate[v.JobName]
		current = append(current, v)
		cateCoordinate[v.JobName] = current
	}
	//fmt.Println(cateCoordinate)

	getNodes(coor, cateCoordinate)

	return coor
}

func getNodes(coor []Coordinate, cateCoordinate map[string][]Coordinate) {
	maxLeng := 0
	currentMax := 0
	maxNode := []Coordinate{}
	nodes := []Coordinate{}
	for _, v := range coor {
		currentMax, nodes = currentNodeMax(v, cateCoordinate)
		if maxLeng < currentMax {
			maxLeng = currentMax
			maxNode = nodes
		}
	}
	//fmt.Println("组合元素：",maxLeng,"组合节点：",maxNode)
	display(maxNode)
	//删除
	for _, v := range maxNode {
		newTags := []Coordinate{}
		//删除
		for _, k := range cateCoordinate[v.JobName] {
			if v.TagName != k.TagName {
				newTags = append(newTags, k)
			}
		}
		cateCoordinate[v.JobName] = newTags
		//删除图列表
	}
	newCoor := []Coordinate{}
	for _, v := range coor {
		flag := true
		for _, v0 := range maxNode {
			if v.TagName == v0.TagName && v.JobName == v0.JobName {
				flag = false
			}
		}
		if flag {
			newCoor = append(newCoor, v)
		}
	}
	if len(newCoor) > 0 {
		getNodes(newCoor, cateCoordinate)
	}
}

func display(coordinate []Coordinate)  {
	displatText := fmt.Sprintf("组合算是数量：%d",len(coordinate))
	jobNames := map[string]bool{}
	jobName,tagName := "",""

	tagNames := map[string]bool{}
	for _,v:= range coordinate {
		if ok,_:= jobNames[v.JobName];!ok{
			jobName += v.JobName+","
			jobNames[v.JobName] = true
		}
		if ok,_:= tagNames[v.TagName];!ok{
			tagName += v.TagName+","
			tagNames[v.TagName] = true
		}
	}
	fmt.Println(displatText,"\t 任务列表",strings.TrimRight(jobName,","),"\t标签",strings.TrimRight(tagName,","))
}

func currentNodeMax(node Coordinate, cateCoordinate map[string][]Coordinate) (int, []Coordinate) {
	fmt.Println(cateCoordinate)
	currentNodeArr := cateCoordinate[node.JobName]
	maxLeng := 0
	for _, tags := range cateCoordinate {
		tempCurrentNodeArr := calculate(currentNodeArr, tags)
		if maxLeng < len(tempCurrentNodeArr) {
			maxLeng = len(tempCurrentNodeArr)
			if len(tempCurrentNodeArr) > len(currentNodeArr) {
				currentNodeArr = tempCurrentNodeArr
			}
		}
	}
	return maxLeng, currentNodeArr
}

/**
* 交集
 */
func calculate(one1 []Coordinate, one2 []Coordinate) []Coordinate {
	afterCoordinate := []Coordinate{}
	for _, v1 := range one1 {
		for _, v2 := range one2 {
			if v1.TagName == v2.TagName {
				afterCoordinate = append(afterCoordinate, v1)
				afterCoordinate = append(afterCoordinate, v2)
			}
		}
	}
	return unqiue(afterCoordinate)
}

func unqiue(nums []Coordinate) []Coordinate {

	coordinateArr := map[string]bool{}
	newCoor := []Coordinate{}
	for _, v := range nums {
		k := v.JobName + "_" + v.TagName
		if ok, _ := coordinateArr[k]; !ok {
			newCoor = append(newCoor, v)
			coordinateArr[k] = true
		}
	}
	return newCoor
}
