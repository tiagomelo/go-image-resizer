// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imageresizer

import "gopkg.in/gographics/imagick.v3/imagick"

type FilterType int

const (
	FILTER_UNDEFINED      FilterType = FilterType(imagick.FILTER_UNDEFINED)
	FILTER_POINT          FilterType = FilterType(imagick.FILTER_POINT)
	FILTER_BOX            FilterType = FilterType(imagick.FILTER_BOX)
	FILTER_TRIANGLE       FilterType = FilterType(imagick.FILTER_TRIANGLE)
	FILTER_HERMITE        FilterType = FilterType(imagick.FILTER_HERMITE)
	FILTER_HANNING        FilterType = FilterType(imagick.FILTER_HANNING)
	FILTER_HAMMING        FilterType = FilterType(imagick.FILTER_HAMMING)
	FILTER_BLACKMAN       FilterType = FilterType(imagick.FILTER_BLACKMAN)
	FILTER_GAUSSIAN       FilterType = FilterType(imagick.FILTER_GAUSSIAN)
	FILTER_QUADRATIC      FilterType = FilterType(imagick.FILTER_QUADRATIC)
	FILTER_CUBIC          FilterType = FilterType(imagick.FILTER_CUBIC)
	FILTER_CATROM         FilterType = FilterType(imagick.FILTER_CATROM)
	FILTER_MITCHELL       FilterType = FilterType(imagick.FILTER_MITCHELL)
	FILTER_JINC           FilterType = FilterType(imagick.FILTER_JINC)
	FILTER_SINC           FilterType = FilterType(imagick.FILTER_SINC)
	FILTER_SINC_FAST      FilterType = FilterType(imagick.FILTER_SINC_FAST)
	FILTER_KAISER         FilterType = FilterType(imagick.FILTER_KAISER)
	FILTER_WELSH          FilterType = FilterType(imagick.FILTER_WELSH)
	FILTER_PARZEN         FilterType = FilterType(imagick.FILTER_PARZEN)
	FILTER_BOHMAN         FilterType = FilterType(imagick.FILTER_BOHMAN)
	FILTER_BARTLETT       FilterType = FilterType(imagick.FILTER_BARTLETT)
	FILTER_LAGRANGE       FilterType = FilterType(imagick.FILTER_LAGRANGE)
	FILTER_LANCZOS        FilterType = FilterType(imagick.FILTER_LANCZOS)
	FILTER_LANCZOS_SHARP  FilterType = FilterType(imagick.FILTER_LANCZOS_SHARP)
	FILTER_LANCZOS2       FilterType = FilterType(imagick.FILTER_LANCZOS2)
	FILTER_LANCZOS2_SHARP FilterType = FilterType(imagick.FILTER_LANCZOS2_SHARP)
	FILTER_ROBIDOUX       FilterType = FilterType(imagick.FILTER_ROBIDOUX)
	FILTER_ROBIDOUX_SHARP FilterType = FilterType(imagick.FILTER_ROBIDOUX_SHARP)
	FILTER_COSINE         FilterType = FilterType(imagick.FILTER_COSINE)
	FILTER_SPLINE         FilterType = FilterType(imagick.FILTER_SPLINE)
	FILTER_SENTINEL       FilterType = FilterType(imagick.FILTER_SENTINEL)
	FILTER_LANCZOS_RADIUS FilterType = FilterType(imagick.FILTER_LANCZOS_RADIUS)
)
