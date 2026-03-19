package facade_design_pattern

func Application() {
	converter := VideoConverter{}
	file, err := converter.Convert("demo_file", "mp4")
	if err != nil {
		return
	}
	file.Save()
}
