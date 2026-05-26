// Copyright 2014 Dario Castañé. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on src/pkg/reflect/deepequal.go from official
// golang's stdlib.

package mergo

import (
	"reflect"
)

func changeInitialCase(s string, mapper func(rune) rune) string {
	_ = "STUB: not implemented"
	return ""
}

func isExported(field reflect.StructField) bool { _ = "STUB: not implemented"; return false }

// Traverses recursively both values, assigning src's fields values to dst.
// The map argument tracks comparisons that have already been seen, which allows
// short circuiting on recursive types.
func deepMap(dst, src reflect.Value, visited map[uintptr]*visit, depth int, config *Config) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Remember, remember...

// We discard it because the field doesn't exist.

// Can this work? I guess it can't.

// Map sets fields' values in dst from src.
// src can be a map with string keys or a struct. dst must be the opposite:
// if src is a map, dst must be a valid pointer to struct. If src is a struct,
// dst must be map[string]interface{}.
// It won't merge unexported (private) fields and will do recursively
// any exported field.
// If dst is a map, keys will be src fields' names in lower camel case.
// Missing key in src that doesn't match a field in dst will be skipped. This
// doesn't apply if dst is a map.
// This is separated method from Merge because it is cleaner and it keeps sane
// semantics: merging equal types, mapping different (restricted) types.
func Map(dst, src interface{}, opts ...func(*Config)) error { _ = "STUB: not implemented"; return nil }

// MapWithOverwrite will do the same as Map except that non-empty dst attributes will be overridden by
// non-empty src attribute values.
// Deprecated: Use Map(…) with WithOverride
func MapWithOverwrite(dst, src interface{}, opts ...func(*Config)) error {
	_ = "STUB: not implemented"
	return nil
}

func _map(dst, src interface{}, opts ...func(*Config)) error { _ = "STUB: not implemented"; return nil }

// To be friction-less, we redirect equal-type arguments
// to deepMerge. Only because arguments can be anything.
