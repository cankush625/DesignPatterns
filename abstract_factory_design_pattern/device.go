package abstract_factory_design_pattern

type IDeviceAbstractFactory interface {
	Camera() string
	Display() string
}
