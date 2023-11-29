// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imageresizer

import "gopkg.in/gographics/imagick.v3/imagick"

// magickWand defines an interface for working with ImageMagick's MagickWand API.
// It abstracts the operations needed for resizing images, allowing for easier testing
// and potential future extensions or replacements of the underlying ImageMagick library.
type magickWand interface {
	ReadImage(filename string) error                                   // ReadImage loads an image from the specified file.
	ResizeImage(cols uint, rows uint, filter imagick.FilterType) error // ResizeImage resizes the image using the specified dimensions and filter.
	GetImageWidth() uint                                               // GetImageWidth returns the width of the current image.
	GetImageHeight() uint                                              // GetImageHeight returns the height of the current image.
	SetImageCompressionQuality(quality uint) error                     // SetImageCompressionQuality sets the compression quality of the image.
	WriteImage(filename string) error                                  // WriteImage writes the image to the specified file.
	Destroy()                                                          // Destroy releases resources associated with the MagickWand.
}

// magickWandWrapper implements the magickWand interface and serves as a wrapper
// around the *imagick.MagickWand type provided by the ImageMagick library.
// This wrapper allows for the convenient use of MagickWand's methods while
// adhering to the magickWand interface, facilitating easier testing and modularity.
type magickWandWrapper struct {
	*imagick.MagickWand // Embedding *imagick.MagickWand to provide direct access to its methods.
}
