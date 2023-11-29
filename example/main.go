package main

import (
	"fmt"
	"os"

	"github.com/tiagomelo/go-image-resizer/imageresizer"
)

func main() {
	const originalImgFile = "originalFile.jpg"
	imageResizer := imageresizer.New(
		imageresizer.WithDimensions(800, 600),
		imageresizer.WithCompressionQuality(50),
		imageresizer.WithFilterType(imageresizer.FILTER_LANCZOS),
		imageresizer.WithOutputDir("/path/to/dir"),
	)
	defer imageResizer.Terminate()
	if err := imageResizer.Resize(originalImgFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
