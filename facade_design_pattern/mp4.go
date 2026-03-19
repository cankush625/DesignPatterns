package facade_design_pattern

type MP4File struct {
	FileName string
}

func (m MP4File) Save() string {
	return "MP4 file saved"
}
