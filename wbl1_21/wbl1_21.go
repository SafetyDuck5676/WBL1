package main

import "fmt"

// Target интерфейс, описывающий нужный нам интерфейс, который Adapter будет адаптировать
type Target interface {
	Request() string
}

// Adaptee реализует устаревший или несовместимый интерфейс, который необходимо адаптировать
type Adaptee struct{}

// SpecificRequest представляет метод Adaptee, который необходимо адаптировать
func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter адаптирует интерфейс Adaptee к интерфейсу Target
type Adapter struct {
	adaptee *Adaptee
}

// Request реализует метод интерфейса Target и вызывает устаревший метод SpecificRequest у Adaptee
func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	// Создаем объект Adapter, адаптирующий интерфейс Adaptee к интерфейсу Target
	adapter := &Adapter{
		adaptee: &Adaptee{},
	}

	// Используем адаптированный интерфейс Target для выполнения запроса
	fmt.Println(adapter.Request())
}
