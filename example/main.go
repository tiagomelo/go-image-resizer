package main

import (
	"fmt"
	"os"

	"github.com/tiagomelo/go-image-resizer/imageresizer"
)

func main() {
	const originalImgFile = "originalFile.jpg"
	ir := imageresizer.New(
		imageresizer.WithDimensions(800, 600),
		imageresizer.WithCompressionQuality(50),
		imageresizer.WithFilterType(imageresizer.FILTER_LANCZOS),
		imageresizer.WithOutputDir("/path/to/dir"),
	)
	if err := ir.Resize(originalImgFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Destroy should be called after Resize() completes.
	ir.Destroy()
	// Terminate() should be called when your program exits.
	imageresizer.Terminate()
}
