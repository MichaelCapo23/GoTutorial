package main

import "fmt"
import "reflect"

func main() {
	//arrays (NOT A REFERENCE DATATYPE)
	// use [...] to not have to specific a length the array when instatiation
	b := [...]string{"yo", "hey", "later"}

	//slices (REFERENCE DATATYPE)
	// dont put anything in the [] to specify a slice.
	c := []int{1,2,3,4,5,6}

	//maps (REFERENCE DATATYPE)
	//keys have to be able to be tested by equality. (strings, ints, arrays, bools) DOESNT NOT COME BACK IN ANY SPECIFIC ORDER
	//delete(a["Farewells"])
	//farewellsVal, ok := a["drgdrg"] <-- reutrns 0 fasle because ok will return true if found and false if not found
	//len works for arrays, slices and maps (maybe structs?)
	a := map[string]string{
		"Greetings":  "1",
		"Farewells": "2",
	}
	
	//structs (NOT A REFERENCE DATATYPE)
	type StructName struct {
		name string
		number int
		company []string
	}

	//dont have to type out the field names, you dont use the field names you need to match the order of the field names within the decloration.
	//
	d := StructName{
		name: "mike",
		number: 3,
		company: []string{
			"Coastline Micro",
			"Fusion of Ideas",
		},
	}
	e := d.company

	//this is an anyonomous struct doesnt need a decloration but this cant be used outside this package
	f := struct{nickname string}{nickname: "Capo"}

	//EMBEDDING STRUCTS

	type Animal struct {
		name string `required max:"100"`  //this is a tag, its literally just a string attached to this structs field. needs other logic to implement the tag. (needs reflect package)
		origin string
	}

	type Bird struct{
		Animal //embedding the animal struct here and can change the values from later declared variables if needbe
		speedMPG float32
		canFly bool
	}

	g:= Bird{}
	g.name = "penguin"
	g.origin = "Antarctic"
	g.speedMPG = 1.3
	g.canFly = false
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)



	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	//GENERAL NOTES
	//go doesnt use inheritance, go uses compostition (through embedding)
}