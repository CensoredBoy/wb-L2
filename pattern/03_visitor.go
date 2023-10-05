package main

import "fmt"

/*
Посетитель — это поведенческий паттерн,
который позволяет добавить новую операцию для целой иерархии классов,
не изменяя код этих классов.

Плюсы:
-Нет необходимости изменять классы
-Похожие операции над разными объектами хранятся в одном месте

Минусы:
-Лишний код
-Может привести к нарушению инкапсуляции элементов.
*/

 type Visitor interface {
 	VisitDobriyColaFactory(f *DobriyColaFactory) string
 	VisitCoolColaFactory(f *CoolColaFactory) string
 	VisitRedPriceColaFactory(f *RedPriceColaFactory) string
}

type Factory interface{
	Check(v Visitor) string
}

type Rospotrebnadzor struct{
}

func (r *Rospotrebnadzor) VisitDobriyColaFactory(f *DobriyColaFactory) string{
	return f.CheckDobriyCola()
}
func (r *Rospotrebnadzor) VisitCoolColaFactory(f *CoolColaFactory) string{
	return f.CheckCoolCola()
}

func (r *Rospotrebnadzor) VisitRedPriceColaFactory(f *RedPriceColaFactory) string{
	return f.CheckRedPriceCola()
}

type City struct {
	factories []Factory
}


func (c *City) Add(f Factory) {
	c.factories = append(c.factories, f)
}

func (c *City) Check(v Visitor){
	for _, f := range c.factories {
		fmt.Println(f.Check(v))
	}

}

type DobriyColaFactory struct{}

func (f *DobriyColaFactory) Check(v Visitor) string {
	return v.VisitDobriyColaFactory(f)
}

func (f *DobriyColaFactory) CheckDobriyCola() string {
	return "Dobriy cola is OK"
}

type CoolColaFactory struct{}

func (f *CoolColaFactory) Check(v Visitor) string {
	return v.VisitCoolColaFactory(f)
}

func (f *CoolColaFactory) CheckCoolCola() string {
	return "Cool Cola is OK but so so"
}

type RedPriceColaFactory struct{}

func (f *RedPriceColaFactory) Check(v Visitor) string {
	return v.VisitRedPriceColaFactory(f)
}

func (f *RedPriceColaFactory) CheckRedPriceCola() string {
	return "Red Price Cola is TERRIBLE"
}

func main() {
	c := new(City)

	c.Add(&DobriyColaFactory{})
	c.Add(&RedPriceColaFactory{})
	c.Add(&CoolColaFactory{})


	c.Check(&Rospotrebnadzor{})
	

}
