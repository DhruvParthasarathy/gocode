package main

import (
	"fmt"
	"reflect"
)

type Planet struct {
	Name string `required max:"50"`
}

type Animal struct {
	Origin string
	Weight int
}

type Bird struct {
	Animal // We embed the animal struct here
	name string
	canFly bool
}

type Doctor struct {
	number int
	actorName string
	companions []string
	dogs []int
}



func main(){

// MAPS
	fmt.Println("Maps")

	mapA := map[string]int {
		"California" : 123,
		"Berkley" : 345,
		"New jersey" : 456,
	}

	fmt.Printf("Map a: %v ", mapA)
	mapA["California"] = 007
	fmt.Printf("\nMap a: %v ", mapA)

	fmt.Println("\n\nMaps are memory referenced")
	mapB := mapA
	delete(mapA, "New jersey")
	fmt.Printf("Map a: %v ", mapA)
	fmt.Printf("\nMap b: %v ", mapB)
	fmt.Printf("\nValue of Chennai in map is %v even if the key Chennai is not a part of the map ", mapB["Chennai"])


// Structs
	fmt.Println("\n\nStructs")

	aDoctor := Doctor{
		number: 3,
		actorName: "lolol", 
		companions: [] string {
			"comp1", "comp2", "comp3",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println("The second companion: ", aDoctor.companions[1])

	fmt.Printf("\nAnonymous struct\n")

	bDoctor := struct{name string}{name: "Dhruv Parthasarathy"}
	fmt.Printf("%v",bDoctor)

	fmt.Printf("\nCopying structs: Structs are not memory referenced, \nwe can copy data into a new struct by using the '=' symbol")

	cDoctor := bDoctor
	cDoctor.name = "Aditya"

	fmt.Printf("\nValue of name in cDoctor is %v", cDoctor.name)

	fmt.Println("\n\nStructs are value referenced")
	dDoctor := &cDoctor
	dDoctor.name = "Bedhi"
	fmt.Printf("Value of name in cDoctor is %v", cDoctor.name)

// Struct composition	
	fmt.Println("\n\nScruct composition")
	fmt.Println("Struct composition is like class inheritence in JAVA")

	bird := Bird{} // declaring the object and manupulating it from outside
	bird.Origin = "Australia"
	bird.Weight = 45
	bird.canFly = false
	bird.name = "Emu"

	fmt.Printf("This is the bird %v", bird)

	bird2 := Bird{
		Animal: Animal{Origin:"India", Weight: 60},
		name: "Crow",
		canFly: true,
	}

	fmt.Printf("\nThis is the bird crow %v", bird2)

// Tags

	fmt.Println("\n\nTags")

	t := reflect.TypeOf(Planet{})
	field, _ := t.FieldByName("Name")
	fmt.Println("Tag of the field 'name' in the struct planet : " + field.Tag)

}
