package main

import "fmt"

/*
Стратегия — это поведенческий паттерн проектирования,
который определяет семейство схожих алгоритмов и помещает каждый из них
в собственный класс, после чего алгоритмы можно взаимозаменять прямо во
время исполнения программы.

Другие объекты содержат ссылку на объект-стратегию и делегируют ей работу.
Программа может подменить этот объект другим, если требуется иной способ решения задачи.

Плюсы:
- Горячая замена алгоритмов на лету.
- Изолирует код и данные алгоритмов от остальных классов.
- Уход от наследования к делегированию.

Минусы:
-  Усложняет программу за счёт дополнительных классов.
*/

type StrategySearch interface{
	Search([]int, int) int
}

type BinSearch struct {
}

func (s *BinSearch) Search(a []int, x int) int{
	l := 0
	r := len(a) - 1
	for ;l < r;{
		m := l + (r-1)/2
		if a[m] <= x{
			return m
		}
		if a[m] < x{
			l = m + 1
		} else {
			r = m + 1
		}
	}
	return -1
}

type NaiveSearch struct{
}

func (s *NaiveSearch) Search(a []int, x int) int{
	for i, v := range a{
		if v == x{
			return i
		}
	}
	return -1
}

type Context struct {
	strategy StrategySearch
}

func (c *Context) Algorithm(a StrategySearch){
	c.strategy = a
}

func (c *Context) Search(a []int, x int) int{
	return c.strategy.Search(a, x)
}


func main() {

	data := []int{1,2,3,4,5,6}

	ctx := new(Context)
	ctx.Algorithm(&NaiveSearch{})
	fmt.Println(ctx.Search(data, 1))

	ctx.Algorithm(&BinSearch{})
	fmt.Println(ctx.Search(data, 2))

}
