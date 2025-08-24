package basics

import (
	"fmt"
	"maps"
)

func goMaps() {
	myMap := make(map[string]int)

	fmt.Println("Initial map:", myMap)

	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println("Map after adding elements:", myMap)

	fmt.Println("Value for not existing key:", myMap["code"])

	myMap["one"] = 11

	fmt.Println("Map after updating 'one':", myMap)

	// delete(myMap, "two")

	delete(myMap, "two")

	fmt.Println("Map after deleting 'two':", myMap)

	// clear(myMap)
	// fmt.Println("Map after clearing:", myMap)

	value, unknownValue:= myMap["three"]

	fmt.Println("Value for 'three':", value, "Unknown value:", unknownValue)

	map2 := map[string]int{
		"four": 4,
		"five": 5,		
	}
	if maps.Equal(myMap, map2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	var nilMap map[string]int // Declare a nil map
	if nilMap == nil {
		fmt.Println("Nil map is nil")
	} else {
		fmt.Println("Nil map is not nil")
	}

    // nilMap["test"] = 100	// we can't do like this to do it proprly we need to initialize it
	
	niceMap := make(map[string]int, 5) 
	
	niceMap["test"] = 100 // Now we can add elements to the map after initializing it
	fmt.Println("Nice map after adding 'test':", niceMap)

	// nested map
	nestedMap := make(map[string]map[string]int)

	nestedMap["Map1"] = map2

	fmt.Println("Nested map:", nestedMap)
}
