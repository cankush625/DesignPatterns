package strategy_design_pattern

import "fmt"

type Quack struct {
}

func (q *Quack) speak() {
	fmt.Println("Quack Quack")
}
