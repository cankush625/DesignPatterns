package facade_design_pattern

type AudioMixer struct {
}

func (am AudioMixer) Fix(convertedFile MP4File) MP4File {
	return MP4File{FileName: convertedFile.FileName}
}
