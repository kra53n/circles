package main

import (
	"fmt"
)

type Statistic struct {
	iters              int
	currOpenNodesNum   int
	currClosedNodesNum int
	maxOpenNodesNum    int
	maxClosedNodesNum  int
	maxNodesNum        int
}

func (s *Statistic) Collect(openNodes, closeNodes []State) {
	s.iters += 1
	s.currOpenNodesNum = len(openNodes)
	s.currClosedNodesNum = len(closeNodes)
	s.maxOpenNodesNum = max(s.maxOpenNodesNum, s.currOpenNodesNum)
	s.maxClosedNodesNum = max(s.maxClosedNodesNum, s.currClosedNodesNum)
	s.maxNodesNum = max(s.maxNodesNum, s.currOpenNodesNum+s.currClosedNodesNum)
}

func (s *Statistic) Print(name string) {
	r := fmt.Sprintf("\n\tРезультат поиска в %s\n\n", name)
	r += fmt.Sprintf("Итераций: %d\n", s.iters)
	r += "Открытые узлы:\n"
	r += fmt.Sprintf("\tКоличество при завершении: %d\n", s.currOpenNodesNum)
	r += fmt.Sprintf("\tМаксимальное количество: %d\n", s.maxOpenNodesNum)
	r += "Закрытые узлы:\n"
	r += fmt.Sprintf("\tКоличество при завершении: %d\n", s.currClosedNodesNum)
	r += fmt.Sprintf("\tМаксимальное количество: %d\n", s.maxClosedNodesNum)
	r += fmt.Sprintf("Максимальное количество хранимых в памяти узлов: %d\n", s.maxNodesNum)
	fmt.Println(r)
}
