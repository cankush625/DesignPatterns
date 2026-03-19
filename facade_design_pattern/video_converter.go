package facade_design_pattern

type VideoConverter struct{}

func (vd VideoConverter) Convert(fileName string, format string) (MP4File, error) {
	file := MP4File{FileName: fileName}
	cf := CodeFactory{}
	sourceCodec := cf.Extract(file)
	var destinationCodec CompressionCodec
	if format == "mp4" {
		destinationCodec = MPEG4CompressionCodec{}
	} else {
		destinationCodec = OGGCompressionCodec{}
	}

	buffer := BitReader{}.Read(fileName, sourceCodec)
	result := BitReader{}.Convert(buffer, destinationCodec, fileName)
	result = AudioMixer{}.Fix(result)
	return result, nil
}
