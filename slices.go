package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Println(loons)
	fmt.Println(len(loons))

	fmt.Printf("loons= %v (type %T)\n", loons, loons)

	fmt.Printf("%q\n", loons[1])
	fmt.Printf("%v\n", loons[1:3])

	for i := range loons {
		fmt.Printf("%v\n", i)
	}

	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	for _, name := range loons {
		fmt.Printf("%s\n", name)
	}

	loons = append(loons, "zelmer")
	fmt.Println(loons)
}
