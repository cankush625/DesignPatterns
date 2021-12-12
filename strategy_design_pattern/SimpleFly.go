package strategy_design_pattern

import "fmt"

type SimpleFly struct {
}

func (s *SimpleFly) fly() {
	fmt.Println("Simple fly")
}
