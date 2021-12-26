package abstract_factory_design_pattern

type Mobile struct {
	MobileCamera MobileCamera
	MobileDisplay MobileDisplay
}

func (m Mobile) Camera() string {
	return m.MobileCamera.Click()
}

func (m Mobile) Display() string {
	return m.MobileDisplay.Show()
}
