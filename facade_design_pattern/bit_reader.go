package facade_design_pattern

type BitReader struct {
}

func (br BitReader) Read(filename string, sourceCodec CompressionCodec) string {
	return "BitReader read"
}

func (br BitReader) Convert(buffer string, destinationCodec CompressionCodec, fileName string) MP4File {
	return MP4File{FileName: fileName}
}
