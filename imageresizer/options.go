// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imageresizer

// Option is a function that configures an imageResizer.
// It is used in the functional options pattern for initializing imageResizer instances.
type Option func(*imageResizer)

// WithDimensions returns an Option that sets the width and height for an imageResizer.
// If either width or height is not set (zero value), the original dimension of the image is retained.
func WithDimensions(width, height int) Option {
	return func(i *imageResizer) {
		i.newWidth = IntPtr(width)   // Set the new width. If zero, original width is retained.
		i.newHeight = IntPtr(height) // Set the new height. If zero, original height is retained.
	}
}

// WithCompressionQuality returns an Option that sets the compression quality for an imageResizer.
// The quality is an integer value typically ranging from 0 (low quality, high compression)
// to 100 (high quality, low compression).
func WithCompressionQuality(cq int) Option {
	return func(i *imageResizer) {
		i.compressionQuality = cq // Set the compression quality.
	}
}

// WithFilterType returns an Option that sets the filter type for an imageResizer.
// FilterType determines the algorithm used for image resizing.
func WithFilterType(ft FilterType) Option {
	return func(i *imageResizer) {
		i.filterType = ft // Set the filter type.
	}
}

// WithOutputDir returns an Option that sets the output directory for an imageResizer.
// If an output directory is provided, resized images will be saved to this directory.
// If not set, images will be saved in the same directory as the original.
func WithOutputDir(outputDir string) Option {
	return func(i *imageResizer) {
		i.outputDir = outputDir // Set the output directory.
	}
}
