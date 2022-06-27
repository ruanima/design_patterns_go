package main

import (
	"fmt"
	"os"
)

type Pet interface {
	say() string
}

type ChinaCat struct{}

type ChinaDog struct{}

type USACat struct{}

type USADog struct{}

// 此处没法细分用Cat和Dog接口, 因为都只有一个方法
type PetFactory interface {
	create_dog() Pet
	create_cat() Pet
}

type ChinaPetFactory struct{}
type USAPetFactory struct{}

func (p *ChinaCat) say() string {
	return "喵喵喵"
}

func (p *ChinaDog) say() string {
	return "汪汪汪"
}

func (p *USACat) say() string {
	return "miaow"
}

func (p *USADog) say() string {
	return "bark"
}

func (p *ChinaPetFactory) create_dog() Pet {
	return &ChinaDog{}
}
func (p *ChinaPetFactory) create_cat() Pet {
	return &ChinaCat{}
}

func (p *USAPetFactory) create_dog() Pet {
	return &USADog{}
}
func (p *USAPetFactory) create_cat() Pet {
	return &USACat{}
}

func main() {
	var factory PetFactory
	if pid := os.Getpid(); pid%2 == 0 {
		factory = &ChinaPetFactory{}
	} else {
		factory = &USAPetFactory{}
	}
	cat := factory.create_cat()
	dog := factory.create_dog()
	fmt.Println(cat.say())
	fmt.Println(dog.say())
}
