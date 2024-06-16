package composite_design_pattern

// Task acts as a leaf in the composite design pattern terms
// Task implements the ITaskList interface
type Task struct {
	Title string
}

func (t Task) ToHtml() string {
	return t.Title
}
