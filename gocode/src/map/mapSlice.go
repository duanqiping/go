package main

import "fmt"

func main()  {
	var monsters []map[string]string

	monsters = make([]map[string]string,2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] =  "牛魔王"
		monsters[0]["age"] =  "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] =  "铁扇公主"
		monsters[1]["age"] =  "600"
	}

	newMonster := map[string]string{
		"name" : "新的妖怪~火云邪神",
		"age" : "200",
	}

	fmt.Println(monsters)
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
}
