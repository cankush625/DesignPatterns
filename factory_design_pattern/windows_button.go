package factory_design_pattern

type WindowsButton struct {

}

func (w WindowsButton) Onclick() string {
	return "Windows button is clicked"
}
