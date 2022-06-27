package main

import (
	"fmt"
	"os"
)

type Pet interface {
	say() string
}

type Cat struct{}

type Dog struct{}

func (p *Cat) say() string {
	return "喵喵喵"
}

func (p *Dog) say() string {
	return "汪汪汪"
}

func pet_factory() Pet {
	if pid := os.Getpid(); pid%2 == 0 {
		return &Dog{}
	} else {
		return &Cat{}
	}
}

func main() {
	pet := pet_factory()
	fmt.Println(pet.say())
}
