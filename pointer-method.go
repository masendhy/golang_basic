package main

import "fmt"

type Man struct {
	Name string
}

// gunakan pointer * dalam membuat struct function
func (man *Man) Married() {
	man.Name = "Hai Mr." + man.Name
}

func main() {
	sendhy := Man{"sendhy"}
	sendhy.Married()

	fmt.Println(sendhy)

}
