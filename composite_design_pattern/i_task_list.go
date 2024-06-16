package composite_design_pattern

// ITaskList is an interface for creating composite and leafs
// in the composite design pattern terms
// It acts as the component in the composite design pattern terms
type ITaskList interface {
	ToHtml() string
}
