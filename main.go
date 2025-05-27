package main

import "fmt"

// 父类
type Animal struct {
	Name string
}

// 父类方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

// 子类继承 Animal
type Dog struct {
	Animal // 继承（嵌套）Animal
	Breed  string
}

// 子类新增方法
func (d *Dog) Bark() {
	fmt.Println(d.Name, "barks loudly!")
}

func main() {
	// 创建 Dog 实例
	dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}

	// 调用父类方法
	dog.Speak() // Buddy makes a sound

	// 调用子类方法
	dog.Bark() // Buddy barks loudly!
}
