package abstract_factory_design_pattern

func CreateDevice(device string) IDeviceAbstractFactory {
	if  device == "mobile" {
		return Mobile{
			MobileCamera:  MobileCamera{},
			MobileDisplay: MobileDisplay{},
		}
	} else if device == "laptop" {
		return Laptop{
			LaptopCamera: LaptopCamera{},
			LaptopDisplay: LaptopDisplay{},
		}
	}
	return nil
}
