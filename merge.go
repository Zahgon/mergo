// Copyright 2013 Dario Castañé. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on src/pkg/reflect/deepequal.go from official
// golang's stdlib.

package mergo

import (
	"reflect"
)

func hasMergeableFields(dst reflect.Value) (exported bool) { _ = "STUB: not implemented"; return false }

func isExportedComponent(field *reflect.StructField) bool { _ = "STUB: not implemented"; return false }

type Config struct {
	Transformers                 Transformers
	Overwrite                    bool
	ShouldNotDereference         bool
	AppendSlice                  bool
	TypeCheck                    bool
	overwriteWithEmptyValue      bool
	overwriteSliceWithEmptyValue bool
	sliceDeepCopy                bool
}

type Transformers interface {
	Transformer(reflect.Type) func(dst, src reflect.Value) error
}

// Traverses recursively both values, assigning src's fields values to dst.
// The map argument tracks comparisons that have already been seen, which allows
// short circuiting on recursive types.
func deepMerge(dst, src reflect.Value, visited map[uintptr]*visit, depth int, config *Config) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Remember, remember...

// Ensure that all keys in dst are deleted if they are not in src.

// Merge will fill any empty for value type attributes on the dst struct using corresponding
// src attributes if they themselves are not empty. dst and src must be valid same-type structs
// and dst must be a pointer to struct.
// It won't merge unexported (private) fields and will do recursively any exported field.
func Merge(dst, src interface{}, opts ...func(*Config)) error {
	_ = "STUB: not implemented"
	return nil
}

// MergeWithOverwrite will do the same as Merge except that non-empty dst attributes will be overridden by
// non-empty src attribute values.
// Deprecated: use Merge(…) with WithOverride
func MergeWithOverwrite(dst, src interface{}, opts ...func(*Config)) error {
	_ = "STUB: not implemented"
	return nil
}

// WithTransformers adds transformers to merge, allowing to customize the merging of some types.
func WithTransformers(transformers Transformers) func(*Config) {
	_ = "STUB: not implemented"
	return nil
}

// WithOverride will make merge override non-empty dst attributes with non-empty src attributes values.
func WithOverride(config *Config) { _ = "STUB: not implemented"; return }

// WithOverwriteWithEmptyValue will make merge override non empty dst attributes with empty src attributes values.
func WithOverwriteWithEmptyValue(config *Config) { _ = "STUB: not implemented"; return }

// WithOverrideEmptySlice will make merge override empty dst slice with empty src slice.
func WithOverrideEmptySlice(config *Config) { _ = "STUB: not implemented"; return }

// WithoutDereference prevents dereferencing pointers when evaluating whether they are empty
// (i.e. a non-nil pointer is never considered empty).
func WithoutDereference(config *Config) { _ = "STUB: not implemented"; return }

// WithAppendSlice will make merge append slices instead of overwriting it.
func WithAppendSlice(config *Config) { _ = "STUB: not implemented"; return }

// WithTypeCheck will make merge check types while overwriting it (must be used with WithOverride).
func WithTypeCheck(config *Config) { _ = "STUB: not implemented"; return }

// WithSliceDeepCopy will merge slice element one by one with Overwrite flag.
func WithSliceDeepCopy(config *Config) { _ = "STUB: not implemented"; return }

func merge(dst, src interface{}, opts ...func(*Config)) error {
	_ = "STUB: not implemented"
	return nil
}

// First allocation

// IsReflectNil is the reflect value provided nil
func isReflectNil(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// Both interface and slice are nil if first word is 0.
// Both are always bigger than a word; assume flagIndir.
