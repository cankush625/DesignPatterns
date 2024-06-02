package bridge_design_pattern

type IView interface {
	Show() string
}

// View is an abstract class
type View struct {
	IView
	Resource IResource
}
