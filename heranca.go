package main
import "fmt"

type  Animal struct {
	peso int
	familia string
}

type Dog struct {
	Animal
	nome string
}

func (a Animal) comer(){
	fmt.Println("comendo ...")
} 
func (d Dog) comer(){
    fmt.Println("comendo racao ...")
}

func (d Dog) latir(){
	fmt.Println("latindo ...")
}

func main(){
	doguinho := Dog{Animal{23,"canidae"}, "doguinho"}
	doguinho.comer()
	doguinho.Animal.comer()
    doguinho.latir()
    fmt.Println ("familia : ", doguinho.familia)
}


