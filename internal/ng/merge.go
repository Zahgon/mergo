package ng

import "reflect"

// Copyright 2025 Dario Castañé. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on src/reflect/deepequal.go from official
// golang's stdlib.

// Don't use this package. It's a work in progress.
// This is the next generation of mergo, and it's not ready for production.

type NilArgumentsError struct{}

func (*NilArgumentsError) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidDestinationError struct{}

func (*InvalidDestinationError) Error() string { _ = "STUB: not implemented"; return "" }

type DifferentArgumentTypesError struct{}

func (*DifferentArgumentTypesError) Error() string { _ = "STUB: not implemented"; return "" }

// Merge sets any [zero-value](https://go.dev/ref/spec#The_zero_value) field
// in dst with the same field's value in src.
// Both dst and src must have the same type, being dst a pointer to the type.
// Merge returns NilArgumentsError if dst is a nil pointer and/or src is a nil
// value.
// If dst and src are values of predefined types, and dst is the type's zero
// value, src is assigned to dst.
// Merge is a convenient wrapper around the more compiler-friendly MergeValue
// and MergePtr functions.
func Merge(dst, src any) error {
	_ = "STUB: not implemented"

	// As dst pointer is a copy; assigning src to a nil pointer is an
	// ineffective assignment.
	return nil
}

// Nothing to do here.

func MergeValue[T any](dst *T, src T) error {
	_ = "STUB: not implemented"

	// As dst pointer is a copy; assigning src to a nil pointer is an
	// ineffective assignment.
	return nil
}

func MergePtr[T any](dst, src *T) error {
	_ = "STUB: not implemented"

	// As dst pointer is a copy; assigning src to a nil pointer is an
	// ineffective assignment.
	return nil
}

// Nothing to do here.

func merge(dst, src reflect.Value, typ reflect.Type) { _ = "STUB: not implemented"; return }

// TODO: handle maps and slices
// TODO: handle pointers and interfaces
// TODO: cover all potential empty cases (as in isEmptyValue from v1)

func mergeStruct(dst, src reflect.Value, typ reflect.Type) { _ = "STUB: not implemented"; return }
