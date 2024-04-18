// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imageresizer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func init() {
	// Set the CGO_CFLAGS_ALLOW environment variable.
	if err := os.Setenv("CGO_CFLAGS_ALLOW", "-Xpreprocessor"); err != nil {
		fmt.Println("Failed to set CGO_CFLAGS_ALLOW:", err)
		os.Exit(1)
	}
}

type ImageResizer interface {
	// Resize resizes the image located at imageFilePath according to the settings of the imageResizer.
	// It returns an error if the resizing process fails.
	Resize(imageFilePath string) error
	// Destroy releases resources associated with the MagickWand.
	// It is the responsibility of the caller to invoke this function
	// on each ImageMagick object after the resize is complete to free up the memory.
	Destroy()
}

// imageResizer encapsulates the settings and operations for resizing images.
type imageResizer struct {
	newWidth           *int       // Target width of the image; nil to keep original width.
	newHeight          *int       // Target height of the image; nil to keep original height.
	compressionQuality int        // Compression quality of the resized image.
	filterType         FilterType // Filter type used for the resizing process.
	outputDir          string     // Directory where the resized image will be saved.
	mw                 magickWand // Wrapper around MagickWand, the ImageMagick API handler.
}

// New initializes a new imageResizer with provided options.
func New(options ...Option) ImageResizer {
	imagick.Initialize() // Initialize the ImageMagick environment.
	resizer := new(imageResizer)
	for _, option := range options {
		option(resizer) // Apply each option to the resizer.
	}
	resizer.mw = &magickWandWrapper{imagick.NewMagickWand()}
	return resizer
}

// ensureDimensions validates and sets the dimensions for the image resizing.
// It returns an error if the dimensions are not set correctly.
func (ir *imageResizer) ensureDimensions() error {
	if ir.newWidth == nil && ir.newHeight == nil {
		// Default to the current dimensions if both width and height are not set.
		currentWidth := int(ir.mw.GetImageWidth())
		currentHeight := int(ir.mw.GetImageHeight())
		ir.newWidth = IntPtr(currentWidth)
		ir.newHeight = IntPtr(currentHeight)
		return nil
	}
	if (ir.newWidth == nil) != (ir.newHeight == nil) {
		return fmt.Errorf("both width and height must be set or both must be nil")
	}
	if *ir.newWidth <= 0 || *ir.newHeight <= 0 {
		return fmt.Errorf("width and height must both be greater than zero")
	}
	return nil
}

func (i *imageResizer) Resize(imageFilePath string) error {
	if err := i.mw.ReadImage(imageFilePath); err != nil {
		return errors.Wrapf(err, "reading image %s", imageFilePath)
	}
	if err := i.ensureDimensions(); err != nil {
		return err
	}
	if err := i.mw.ResizeImage(uint(*i.newWidth), uint(*i.newHeight), imagick.FilterType(i.filterType)); err != nil {
		return errors.Wrap(err, "resizing image")
	}
	if err := i.mw.SetImageCompressionQuality(uint(i.compressionQuality)); err != nil {
		return errors.Wrapf(err, "setting image compression quality to %d", i.compressionQuality)
	}
	resizedImageFilePath := i.resizedImageFilePath(imageFilePath)
	if err := i.mw.WriteImage(resizedImageFilePath); err != nil {
		return errors.Wrapf(err, "writing image %s", resizedImageFilePath)
	}
	return nil
}

func (i *imageResizer) Destroy() {
	i.mw.Destroy()
}

// Terminate releases resources used by imageResizer and ImageMagick. It is the responsibility
// of the caller to invoke this function after completing image resizing operations. Failing to
// call Terminate can lead to resource leaks as it cleans up the MagickWand instance and
// terminates the ImageMagick environment. This is crucial especially in long-running
// applications or those processing large numbers of images, to avoid excessive memory usage.
func Terminate() {
	imagick.Terminate()
}

// resizedImageFilePath generates the file path for the resized image. It uses the output directory
// specified in the imageResizer. If no output directory is specified, the original image file path
// is used as the base path. This ensures that the resized image is saved either in a specified
// location or alongside the original image if no specific output location is provided.
func (i *imageResizer) resizedImageFilePath(imageFilePath string) string {
	basePath := filepath.Dir(imageFilePath)
	if i.outputDir != "" {
		basePath = i.outputDir
	}
	fileName := filepath.Base(imageFilePath)
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex == -1 {
		return filepath.Join(basePath, fileName+"_resized")
	}
	newFileName := fileName[:dotIndex] + "_resized" + fileName[dotIndex:]
	return filepath.Join(basePath, newFileName)
}
