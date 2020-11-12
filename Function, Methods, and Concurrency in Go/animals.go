package main

import "fmt"

type Animal struct {
	food, locomotion, sound string
}

func (e Animal) Eat() {
	fmt.Println(e.food)
}
func (e Animal) Move() {
	fmt.Println(e.locomotion)
}
func (e Animal) Speak() {
	fmt.Println(e.sound)
}
func main() {

	m := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	var name string
	var action string
	for true {

		fmt.Println(">")
		fmt.Scan(&name, &action)
		if action == "eat" {
			m[name].Eat()
		} else if action == "move" {
			m[name].Move()
		} else if action == "speak" {
			m[name].Speak()
		} else {
			fmt.Println("Invalid input")
		}
	}
}
