package bridge_design_pattern

import (
	"fmt"
)

func main() {
	// Create artist
	artist := Artist{Name: "ABC", Age: 40, Bio: "Musician", Url: "https://example.com", Image: "https://example.com"}
	// Create artist resource
	artistResource := ArtistResource{Artist: artist}
	// Create a large view using artist resource.
	largeView := LargeView{Resource: artistResource}
	fmt.Println(largeView.Show())
	// Output:
	// <data>Musician</data>

	// Note:
	// Since View requires any resource of type IResource, we can use
	// any resource that implements IResource
	// So, View is independent of the resource
}
