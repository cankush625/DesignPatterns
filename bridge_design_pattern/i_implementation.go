package bridge_design_pattern

// IResource is an interface defined for implementing resources
type IResource interface {
	getSnippet() string
	getTitle() string
	getImage() string
	getUrl() string
}
