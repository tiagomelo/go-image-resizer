# go-image-resizer

[![Go Reference](https://pkg.go.dev/badge/github.com/tiagomelo/go-image-resizer.svg)](https://pkg.go.dev/github.com/tiagomelo/go-image-resizer)

![logo](logo.png)

A tiny utility for resizing images with [ImageMagick](https://imagemagick.org/).

It uses [github.com/gographics/imagick](https://github.com/gographics/imagick) as the Go bind to ImageMagick's MagickWand C API.

## installation

```
go get github.com/tiagomelo/go-image-resizer
```

## available options

- `WithDimensions` sets the width and height. If none are specified, image's width and height will be preserved.
- `WithCompressionQuality` sets the compression quality. The quality is an integer value typically ranging from 0 (low quality, high compression) to 100 (high quality, low compression)
- `WithFilterType` sets the filter type. It determines the algorithm used for image resizing. See the available filter types [here](./imageresizer/filters.go).
- `WithOutputDir` sets the output directory. If not set, images will be saved in the same directory as the original.


## example

```
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

```

## running unit tests

```
make test
```