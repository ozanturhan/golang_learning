package main

import "fmt"

// Map

func main() {

	// 1.
	myMap := make(map[int]string)
	fmt.Println(myMap)

	myMap[43] = "foo"
	myMap[12] = "bar"

	fmt.Println(myMap)

	// 2.
	states := make(map[string]string)

	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["YLV"] = "Yalova"

	fmt.Println(states)

	// get ANK
	antalya := states["ANK"]
	fmt.Println("Selected city:", antalya)

	// delete ANT
	delete(states, "ANT")
	fmt.Println(states)

	// add ERZ
	states["ERZ"] = "Erzurum"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k , v)
	}

	// create keys list
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

	for _, key := range keys {
		fmt.Println(states[key])
	}
}
