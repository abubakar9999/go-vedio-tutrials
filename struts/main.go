package main


import(
	"fmt"
)
type Mystructs  struct{
	Name string
	Email string
	dob int
}
func main() {

	user :=Mystructs{"abu","abu@gmail.com",20}
	

	fmt.Println(user)

	 var abu = new(Mystructs)

	abu.Name="abubakaf"

	abu.Email="emeid"
	abu.dob=232341

	fmt.Printf("%+v\n",abu)
	

}

 