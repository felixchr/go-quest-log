package main

import "fmt"

func main() {
	func1()
	func2()
	func3()
}

func func1() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map literal at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	val, isPresent := mapLit["key1"]
	if isPresent {
		fmt.Println(val)
	} else {
		fmt.Println("Non-exist")
		mapLit["key1"] = 9
	}
	delete(mapLit, "two")
	fmt.Println(mapLit)

	for key, value := range mapLit {
		fmt.Printf("Key %s: %d\n", key, value)
	}

	var weekdays map[string]int
	weekdays = map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednseday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
	}
	for _, day := range []string{"Tuesday", "Hollyday"} {
		_, ok := weekdays[day]
		if ok {
			fmt.Printf("%s exists\n", day)
		} else {
			fmt.Printf("%s not exists\n", day)
		}
	}
}

func func2() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
}

func func3() {
	var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}
