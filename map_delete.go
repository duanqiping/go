package main

import "fmt"
import "sync"

func main()  {
	scene := make(map[string]int)

	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	delete(scene, "brazil")

	//for k, v := range scene {
	//	fmt.Println(k, v)
	//}

	var scene1 sync.Map
	// 将键值对保存到sync.Map
	scene1.Store("greece", 97)
	scene1.Store("london", 100)
	scene1.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene1.Load("london"))
	// 根据键删除对应的键值对
	scene1.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene1.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}