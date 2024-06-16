package composite_design_pattern

import (
	"fmt"
)

func main() {
	first_task := Task{Title: "Explore OAuth providers"}
	second_task := TaskGroup{Title: "Implement Sign In", Todos: []ITaskList{first_task}}
	third_task := Task{Title: "Implement Sign Out"}
	fourth_task := TaskGroup{Title: "Implement Authentication", Todos: []ITaskList{second_task, third_task}}
	fmt.Println(first_task.ToHtml())
	// Output:
	// Explore OAuth providers
	fmt.Println(second_task.ToHtml())
	// Output:
	// <h1>Implement Sign In</h1><ul><li>Explore OAuth providers</li></ul>
	fmt.Println(third_task.ToHtml())
	// Output:
	// Implement Sign Out
	fmt.Println(fourth_task.ToHtml())
	// Output:
	// <h1>Implement Authentication</h1><ul><li><h1>Implement Sign In</h1><ul><li>Explore OAuth providers</li></ul></li><li>Implement Sign Out</li></ul>

	// Note:
	// Composite design patterns executes the composites in recursive way
}
