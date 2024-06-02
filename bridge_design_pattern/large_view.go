package bridge_design_pattern

// LargeView is the implementation of view from abstract View class
type LargeView struct {
	Resource IResource
}

func (lv LargeView) Show() string {
	return "<data>" + lv.Resource.getSnippet() + "</data>"
}
