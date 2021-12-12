package strategy_design_pattern

import "fmt"

type WaveFly struct {
}

func (w *WaveFly) fly() {
	fmt.Println("Wave fly")
}
