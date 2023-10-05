package main

import "fmt"

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Строитель позволяет не использовать длинные конструкторы.

Если потребуется объединить шаги создания создается объект директор, определяющий
порядок шагов, так процесс создания будет полностью скрыт от клиента
(директор - не является обязательным в реализации паттерна)

Плюсы:
- Создает объекты пошагово (если это требуется, например, деревья)
- Позволяет создавать несколько представлений одного объекта, переиспользуя код
Минусы:
- Усложняет код, требуется написание дополнительных классов, может усложняться внедрнее зависимостей

*/

const (
	apple = "apple"
	xiaomi ="xiaomi"
)


type Computer struct{
	maker string
	cpu int
	ram int
}

type ComputerBuilder interface{
	SetMaker(maker string) ComputerBuilder
	SetCpu(cpu int) ComputerBuilder
	SetRam(ram int) ComputerBuilder
	Build() *Computer
}

type computerBuilder struct {
	computer *Computer
}

func NewComputerBuilder() ComputerBuilder{
	return &computerBuilder{&Computer{cpu: 0, ram: 0}}
}

func (b *computerBuilder) SetMaker(maker string) ComputerBuilder {
	b.computer.maker = maker 
	return b
}

func (b *computerBuilder) SetCpu(cpu int) ComputerBuilder {
	b.computer.cpu = cpu 
	return b
}

func (b *computerBuilder) SetRam(ram int) ComputerBuilder {
	b.computer.ram = ram 
	return b
}

func (b *computerBuilder) Build() *Computer {
	return b.computer
}

type ComputerDirector interface {
	MakeComputer(builder ComputerBuilder, maker string) *Computer
}

type computerDirector struct{}

func NewComputerDirector() ComputerDirector{
	return &computerDirector{}
}

func (b *computerDirector) MakeComputer(builder ComputerBuilder, maker string) *Computer {
	switch maker{
	case apple:
		return  builder.SetMaker(maker).SetCpu(2).SetRam(16).Build()
	case xiaomi:
		return builder.SetMaker(maker).SetCpu(0).SetRam(0).Build()
	default:
		return nil
	}
}

func main() {
	computerBuilder := NewComputerBuilder()
	computerDirector := NewComputerDirector()
	mac := computerDirector.MakeComputer(computerBuilder, "apple")

	fmt.Printf("Computer: %+v\n", mac)

	redmiBook := computerDirector.MakeComputer(computerBuilder, "xiaomi")
	fmt.Printf("Computer: %+v\n", redmiBook)

}
