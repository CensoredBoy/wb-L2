package main

import (
	"fmt"
)

/*
Состояние — это поведенческий паттерн проектирования, который позволяет
объектам менять поведение в зависимости от своего состояния. Извне создаётся
впечатление, что изменился класс объекта.

Основная идея в том, что программа может находиться в одном из нескольких
состояний, которые всё время сменяют друг друга. Набор этих состояний,
а также переходов между ними, предопределён и конечен. Находясь в разных
состояниях, программа может по-разному реагировать на одни и те же события,
которые происходят с ней.

Плюсы:
- Избавляет от множества больших условных операторов машины состояний.
- Концентрирует в одном месте код, связанный с определённым состоянием.

Минусы:
- Может неоправданно усложнить код, если состояний мало и они редко меняются.

Пример сделаем на основе торгового автомата, который может пребывать только
в одном из 4 состояний:

-hasItem (имеетПредмет)
-noItem (неИмеетПредмет)
-itemRequested (выдаётПредмет)
-hasMoney (получилДеньги)

И может выполнять только 4 действия:

-Выбрать предмет
-Добавить предмет
-Ввести деньги
-Выдать предмет
*/
type AlertStarter interface {
	Alert() string
}

type AlertBox struct{
	state AlertStarter
}

func (a *AlertBox) Alert() string{
	return a.state.Alert()
}

func (a *AlertBox) SetState(state AlertStarter) {
	a.state = state
}

func NewAlertBox() *AlertBox {
	return &AlertBox{state:&ErrorAlert{}}
}


type ErrorAlert struct {}

func (a *ErrorAlert) Alert() string{
	return "ERROR"
}

type WarningAlert struct {}

func  (a *WarningAlert) Alert() string {
	return "WARNING"
}

type InfoAlert struct {}

func (a *InfoAlert) Alert() string {
	return "INFO"
}

func main() {
	alertBox := NewAlertBox()

	message := alertBox.Alert()

	fmt.Println(message)

	alertBox.SetState(&WarningAlert{})

	message = alertBox.Alert()

	fmt.Println(message)

	alertBox.SetState(&InfoAlert{})

	message = alertBox.Alert()	

	fmt.Println(message)
}
