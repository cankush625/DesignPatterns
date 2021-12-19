package factory_design_pattern

func SignInButton() IButton {
	return ButtonFactory{}.createButton()
}
