package factory_design_pattern

import "runtime"

// ButtonFactory use this factory to create new Buttons independent of the OS
type ButtonFactory struct {

}

func (b ButtonFactory) createButton() IButton {
	if runtime.GOOS == "windows" {
		return WindowsButton{}
	} else if runtime.GOOS == "darwin" {
		return MacButton{}
	} else {
		return GenericButton{}
	}
}
