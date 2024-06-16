package composite_design_pattern

// TaskGroup is a collection of leaf's in component design pattern terms
// The type of collection is ITaskList
// It acts as a composite in the composite design pattern terms
// TaskGroup implements the ITaskList interface
type TaskGroup struct {
	Title string
	Todos []ITaskList
}

func (tg TaskGroup) ToHtml() string {
	html := "<h1>"
	html += tg.Title
	html += "</h1><ul>"
	for _, task := range tg.Todos {
		html += "<li>"
		html += task.ToHtml()
		html += "</li>"
	}
	html += "</ul>"
	return html
}
