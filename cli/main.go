package main

import (
	"flag"
	"fmt"
)

func main() {
	filepath := flag.String("file", "skill-Basics.json", "")
	flag.Parse()

	err := ExamineWatsonSkill(*filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Skill go files created!")
}
