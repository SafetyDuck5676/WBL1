package main

import "fmt"

// Определение структуры Human
type Human struct {
	Name    string
	Age     int
	Height  float64
	Weight  float64
	Hobbies []string
}

// Метод Speak() для структуры Human
func (h *Human) Speak() {
	fmt.Printf("Привет, меня зовут %s и мне %d год(а)\n", h.Name, h.Age)
}

// Метод Eat() для структуры Human
func (h *Human) Eat() {
	fmt.Printf("%s ест\n", h.Name)
}

// Метод Sleep() для структуры Human
func (h *Human) Sleep() {
	fmt.Printf("%s спит\n", h.Name)
}

// Метод Exercise() для структуры Human
func (h *Human) Exercise() {
	fmt.Printf("%s занимается физическими упражнениями\n", h.Name)
}

// Метод AddHobby() для структуры Human
func (h *Human) AddHobby(hobby string) {
	h.Hobbies = append(h.Hobbies, hobby)
	fmt.Printf("%s добавил новое хобби: %s\n", h.Name, hobby)
}

// Определение вложенной структуры Action
type Action struct {
	Human // Встраивание структуры Human в структуру Action
}

// Метод Work() для структуры Action
func (a *Action) Work() {
	fmt.Printf("%s занимается работой\n", a.Name)
}

func main() {
	// Создание объекта структуры Action
	action := Action{
		Human: Human{
			Name:    "Иван",
			Age:     25,
			Height:  180.0,
			Weight:  75.0,
			Hobbies: []string{"футбол", "плавание"},
		},
	}

	// Вызов методов
	action.Speak()          // Метод Speak() унаследован от структуры Human
	action.Eat()            // Метод Eat() унаследован от структуры Human
	action.Sleep()          // Метод Sleep() унаследован от структуры Human
	action.Exercise()       // Метод Exercise() унаследован от структуры Human
	action.AddHobby("кино") // Метод AddHobby() унаследован от структуры Human
	action.Work()           // Метод Work() определен в структуре Action
}
