package facade_design_pattern

type CodeFactory struct {
}

func (cf CodeFactory) Extract(file MP4File) CompressionCodec {
	print("File is extracted")
	return MPEG4CompressionCodec{}
}
