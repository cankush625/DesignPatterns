package abstract_factory_design_pattern

type Laptop struct {
	LaptopCamera LaptopCamera
	LaptopDisplay LaptopDisplay
}

func (l Laptop) Camera() string {
	return l.LaptopCamera.Click()
}

func (l Laptop) Display() string {
	return l.LaptopDisplay.Show()
}
