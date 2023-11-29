// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imageresizer

// IntPtr takes an integer and returns a pointer to a newly allocated copy of it.
// This is useful when you need to pass an integer by reference, especially in
// scenarios where integer values are optional and a nil pointer is a valid state.
func IntPtr(n int) *int {
	value := n
	return &value
}
