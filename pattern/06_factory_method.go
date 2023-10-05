package main

import "fmt"
import "log"

/*Фабричный метод — это порождающий паттерн проектирования, который определяет
общий интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

Плюсы:
- Избавляет от привязки к конкретному объекту
- Отделяет производство объектов

Минусы:
- Может привести к созданию больших параллельных иерархий объектов
- Один "божественный" конструктор
- Дополнительный код

В этом примере мы будем создавать разные типы оружия при помощи структуры фабрики.
*/

type carModel string

const(
	Lada carModel = "Lada"
	Kia carModel = "Kia"
)

type Factory interface {
	CreateCar( model carModel ) Car
}

type Car interface {
	Drive() string
}

type ConcreteFactory struct {}

func NewFactory() Factory {
	return &ConcreteFactory{}
}

func (f *ConcreteFactory) CreateCar(model carModel) Car {
	var car Car

	switch model{
	case Lada:
		car = &LadaCar{string(model)}
	case Kia:
		car = &KiaCar{string(model)}
	default:
		log.Fatalln("Unknown car")
	}
	return car
}

type LadaCar struct{
	model string
}

func (c *LadaCar) Drive() string {
	return c.model
}

type KiaCar struct{
	model string
}

func (c *KiaCar) Drive() string{
	return c.model
}


func main() {
	factory := NewFactory()
	cars := []Car{factory.CreateCar(Lada), factory.CreateCar(Kia)}
	for _, car := range cars {
		fmt.Println("I driving ", car.Drive())
	}
}

