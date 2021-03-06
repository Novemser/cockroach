// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package types

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
)

// T represents an exec physical type - a bytes representation of a particular
// column type.
type T int

const (
	// Bool is a column of type bool
	Bool T = iota
	// Bytes is a column of type []byte
	Bytes
	// Int8 is a column of type int8
	Int8
	// Int16 is a column of type int16
	Int16
	// Int32 is a column of type int32
	Int32
	// Int64 is a column of type int64
	Int64
	// Float32 is a column of type float32
	Float32
	// Float64 is a column of type float64
	Float64

	// Unhandled is a temporary value that represents an unhandled type.
	// TODO(jordan): this should be replaced by a panic once all types are
	// handled.
	Unhandled
)

// FromColumnType returns the T that corresponds to the input ColumnType.
func FromColumnType(ct sqlbase.ColumnType) T {
	switch ct.SemanticType {
	case sqlbase.ColumnType_BOOL:
		return Bool
	case sqlbase.ColumnType_BYTES, sqlbase.ColumnType_STRING, sqlbase.ColumnType_NAME:
		return Bytes
	case sqlbase.ColumnType_INT:
		switch ct.Width {
		case 8:
			return Int8
		case 16:
			return Int16
		case 32:
			return Int32
		case 0, 64:
			return Int64
		}
		panic(fmt.Sprintf("integer with unknown width %d", ct.Width))
	case sqlbase.ColumnType_OID:
		return Int64
	case sqlbase.ColumnType_FLOAT:
		return Float64
	}
	return Unhandled
}
