package main

import "fmt"

type animal struct {
	food       string
	locomotion string
	noice      string
}
type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (a animal) Eat() {
	fmt.Println(a.food)
	return
}
func (a animal) Move() {
	fmt.Println(a.locomotion)
	return
}
func (a animal) Speak() {
	fmt.Println(a.noice)
	return
}
func main() {

	mapAnimal := make(map[string]animal)
	mapAnimal["cow"] = animal{"grass", "walk", "moo"}
	mapAnimal["bird"] = animal{"worms", "fly", "peep"}
	mapAnimal["snake"] = animal{"mice", "slither", "hsss"}
	// creat temp variable
	var temp animalInterface
	for {
		var command, requestAni, requestType string
		fmt.Print(">")
		fmt.Scan(&command, &requestAni, &requestType)
		if command == "query" {
			temp = mapAnimal[requestAni]
			switch requestType {
			case "eat":
				temp.Eat()
			case "move":
				temp.Move()
			case "speak":
				temp.Speak()
			}
		}
		if command == "newanimal" {
			mapAnimal[requestAni] = mapAnimal[requestType]
			fmt.Println("Created it!")
		}
	}
}
