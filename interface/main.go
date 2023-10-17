package main

import "fmt"

type Eapp struct {
	Name   string
	Post   string
	salary float32
}

func(e Eapp) Display()  {

	fmt.Println("Eapp Emplaye ",e.Name,e.Post,e.salary)



}

type Eon struct{
	Name   string
	Post   string
	salary float32
}

func (o Eon) Display(){
	fmt.Println("Eon Emplayer ",o.Name,o.Post,o.salary)
}


type myconection interface{
Display()
}

func Output(w myconection){
	w.Display()
}

func main() {
eApp :=Eapp{Name:"Abu bakar",Post:"Software",salary:5000}
Output(eApp)

eon :=Eon{Name:"Abu Mahadi",Post:"50",salary:5000}
Output(eon)


}