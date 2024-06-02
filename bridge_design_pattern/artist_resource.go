package bridge_design_pattern

// ArtistResource is a resource class for Artist resource
// using IResource interface
type ArtistResource struct {
	Artist Artist
}

func (ar ArtistResource) getSnippet() string {
	return ar.Artist.Bio
}

func (ar ArtistResource) getTitle() string {
	return ar.Artist.Name
}

func (ar ArtistResource) getImage() string {
	return ar.Artist.Image
}

func (ar ArtistResource) getUrl() string {
	return ar.Artist.Url
}
