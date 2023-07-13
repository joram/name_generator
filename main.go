package main

import (
	"fmt"
	"github.com/joram/name_generator/name_generator"
)

func main() {
	ng := name_generator.NewNameGenerator()
	for i := 0; i < 10; i++ {
		maleName := ng.MaleName(int64(i))
		femaleName := ng.FemaleName(int64(i))
		fmt.Printf("%s %s\n", maleName, femaleName)
	}
}
