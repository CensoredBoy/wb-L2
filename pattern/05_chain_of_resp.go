package main

import "fmt"

/*
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли
передавать запрос дальше по цепи.

Как и многие другие поведенческие паттерны, Цепочка обязанностей базируется на том,
чтобы превратить отдельные поведения в объекты. В нашем случае каждая проверка переедет в
отдельный класс с единственным методом выполнения. Данные запроса, над которым происходит проверка,
будут передаваться в метод как аргументы.

Плюсы:
-Уменьшает зависимость между клиентом и обработчиками.
-Реализует принцип единственной обязанности.

Минусы:
Запрос может остаться никем не обработанным.
*/




type Login interface {
	SendNext(message int) string
}

type Identification struct{
	next Login
}

func (i *Identification) SendNext(message int) (result string){
	if message == 1 {
		result = "Identification done"
	} else if i.next != nil {
		result = i.next.SendNext(message)
	}
	return
} 

type Authentification struct{
	next Login
}

func (i *Authentification) SendNext(message int) (result string){
	if message == 2 {
		result = "Authentification done"
	} else if i.next != nil {
		result = i.next.SendNext(message)
	}
	return
}

type Authorize struct{
	next Login
}

func (i *Authorize) SendNext(message int) (result string){
	if message == 3 {
		result = "Authorize done"
	} else if i.next != nil {
		result = i.next.SendNext(message)
	}
	return
}





func main() {
	login := &Identification{ next: &Authentification{ next: &Authorize{} } }
	result := login.SendNext(2)
	fmt.Println(result)
	
}
