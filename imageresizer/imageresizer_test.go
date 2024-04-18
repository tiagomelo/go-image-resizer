// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imageresizer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name                       string
		options                    []Option
		expectedNewWidth           *int
		expectedNewHeight          *int
		expectedCompressionQuality int
		expectedOutputDir          string
		expectedFilterType         FilterType
	}{
		{
			name: "with all options",
			options: []Option{
				WithDimensions(800, 600),
				WithCompressionQuality(50),
				WithFilterType(FILTER_LANCZOS),
				WithOutputDir("path/to/some/dir"),
			},
			expectedNewWidth:           IntPtr(800),
			expectedNewHeight:          IntPtr(600),
			expectedCompressionQuality: 50,
			expectedOutputDir:          "path/to/some/dir",
			expectedFilterType:         FILTER_LANCZOS,
		},
		{
			name: "no options",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			imgResizer := New(tc.options...)
			defer Terminate()
			ir, ok := imgResizer.(*imageResizer)
			require.True(t, ok)
			assert.Equal(t, tc.expectedNewWidth, ir.newWidth)
			assert.Equal(t, tc.expectedNewHeight, ir.newHeight)
			assert.Equal(t, tc.expectedCompressionQuality, ir.compressionQuality)
			assert.Equal(t, tc.expectedOutputDir, ir.outputDir)
			assert.Equal(t, tc.expectedFilterType, ir.filterType)
			imgResizer.Destroy()
		})
	}
}

func TestResize(t *testing.T) {
	testCases := []struct {
		name           string
		newWidth       *int
		mockClosure    func(m *mockMagickWand)
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "happy path",
			mockClosure:    func(m *mockMagickWand) {},
			expectedOutput: "/path/to/dir/someImage_resized.jpg",
		},
		{
			name: "error when reading image",
			mockClosure: func(m *mockMagickWand) {
				m.errReadImage = errors.New("read image error")
			},
			expectedError: errors.New("reading image someImage.jpg: read image error"),
		},
		{
			name:          "error when ensuring dimensions",
			mockClosure:   func(m *mockMagickWand) {},
			newWidth:      IntPtr(500),
			expectedError: errors.New("both width and height must be set or both must be nil"),
		},
		{
			name: "error when resizing",
			mockClosure: func(m *mockMagickWand) {
				m.errResizeImage = errors.New("resize image error")
			},
			expectedError: errors.New("resizing image: resize image error"),
		},
		{
			name: "error when setting image compression quality",
			mockClosure: func(m *mockMagickWand) {
				m.errSetImageCompressionQuality = errors.New("set image compression quality error")
			},
			expectedError: errors.New("setting image compression quality to 50: set image compression quality error"),
		},
		{
			name: "error when writing the resized image",
			mockClosure: func(m *mockMagickWand) {
				m.errWriteImage = errors.New("write image error")
			},
			expectedError: errors.New("writing image /path/to/dir/someImage_resized.jpg: write image error"),
		},
	}
	for _, tc := range testCases {
		m := new(mockMagickWand)
		t.Run(tc.name, func(t *testing.T) {
			tc.mockClosure(m)
			ir := &imageResizer{
				mw:                 m,
				newWidth:           tc.newWidth,
				compressionQuality: 50,
				outputDir:          "/path/to/dir",
			}
			output, err := ir.Resize("someImage.jpg")
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error %v, got nil", tc.expectedError)
				}
				require.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}

func Test_ensureDimensions(t *testing.T) {
	testCases := []struct {
		name              string
		newWidth          *int
		newHeight         *int
		expectedNewWidth  *int
		expectedNewHeight *int
		expectedError     error
	}{
		{
			name:              "no dimensions provided",
			expectedNewWidth:  IntPtr(1200),
			expectedNewHeight: IntPtr(850),
		},
		{
			name:          "only width was provided",
			newHeight:     IntPtr(850),
			expectedError: errors.New("both width and height must be set or both must be nil"),
		},
		{
			name:          "only height was provided",
			newWidth:      IntPtr(1200),
			expectedError: errors.New("both width and height must be set or both must be nil"),
		},
		{
			name:          "witdh is zero",
			newWidth:      IntPtr(0),
			newHeight:     IntPtr(850),
			expectedError: errors.New("width and height must both be greater than zero"),
		},
		{
			name:          "height is zero",
			newWidth:      IntPtr(1200),
			newHeight:     IntPtr(0),
			expectedError: errors.New("width and height must both be greater than zero"),
		},
	}
	for _, tc := range testCases {
		m := new(mockMagickWand)
		t.Run(tc.name, func(t *testing.T) {
			ir := &imageResizer{
				newWidth:  tc.newWidth,
				newHeight: tc.newHeight,
				mw:        m,
			}
			err := ir.ensureDimensions()
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error %v, got nil", tc.expectedError)
				}
				assert.Equal(t, *tc.expectedNewWidth, *ir.newWidth)
				assert.Equal(t, *tc.expectedNewHeight, *ir.newHeight)
			}
		})
	}
}

func Test_resizedImageFilePath(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		outputDir      string
		expectedOutput string
	}{
		{
			name:           "no output dir, with extension",
			input:          "path/to/some/file.jpg",
			expectedOutput: "path/to/some/file_resized.jpg",
		},
		{
			name:           "no output dir, without extension",
			input:          "path/to/some/file",
			expectedOutput: "path/to/some/file_resized",
		},
		{
			name:           "with output dir, with extension",
			input:          "path/to/some/file.jpg",
			outputDir:      "newpath/to/some",
			expectedOutput: "newpath/to/some/file_resized.jpg",
		},
		{
			name:           "with output dir, without extension",
			input:          "path/to/some/file",
			outputDir:      "newpath/to/some",
			expectedOutput: "newpath/to/some/file_resized",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ir := &imageResizer{
				outputDir: tc.outputDir,
			}
			output := ir.resizedImageFilePath(tc.input)
			require.Equal(t, tc.expectedOutput, output)
		})
	}
}

type mockMagickWand struct {
	errReadImage                  error
	errResizeImage                error
	errSetImageCompressionQuality error
	errWriteImage                 error
}

func (m *mockMagickWand) ReadImage(filename string) error {
	return m.errReadImage
}

func (m *mockMagickWand) ResizeImage(cols uint, rows uint, filter imagick.FilterType) error {
	return m.errResizeImage
}

func (m *mockMagickWand) GetImageWidth() uint {
	return uint(1200)
}

func (m *mockMagickWand) GetImageHeight() uint {
	return uint(850)
}

func (m *mockMagickWand) SetImageCompressionQuality(quality uint) error {
	return m.errSetImageCompressionQuality
}

func (m *mockMagickWand) WriteImage(filename string) error {
	return m.errWriteImage
}

func (m *mockMagickWand) Destroy() {}
