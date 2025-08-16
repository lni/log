// Copyright 2025 Lei Ni (nilei81@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"fmt"
	"time"
)

// Field represents a key-value pair for structured logging.
// It wraps zap.Field internally but provides a zap-independent interface.
type Field struct {
	Key       string
	Type      FieldType
	Integer   int64
	String    string
	Interface interface{}
}

// FieldType represents the type of a field.
type FieldType uint8

const (
	UnknownType FieldType = iota
	StringType
	Int64Type
	IntType
	Int32Type
	Int16Type
	Int8Type
	Uint64Type
	Uint32Type
	Uint16Type
	Uint8Type
	UintType
	Float64Type
	Float32Type
	BoolType
	ErrorType
	DurationType
	TimeType
	ByteStringType
	StringerType
	AnyType
)

// Logger is the interface that wraps the basic logging methods.
// It is designed to be a drop-in replacement for *zap.Logger.
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)

	// Sugar returns a sugared logger for printf-style APIs
	Sugar() SugaredLogger
}

// SugaredLogger provides a more ergonomic, printf-style API.
type SugaredLogger interface {
	Panicf(template string, args ...interface{})
}

func String(key string, val string) Field {
	return Field{Key: key, Type: StringType, String: val}
}

func Int64(key string, val int64) Field {
	return Field{Key: key, Type: Int64Type, Integer: val}
}

func Int(key string, val int) Field {
	return Field{Key: key, Type: IntType, Integer: int64(val)}
}

func Float64(key string, val float64) Field {
	return Field{Key: key, Type: Float64Type, Interface: val}
}

func Float32(key string, val float32) Field {
	return Field{Key: key, Type: Float32Type, Interface: val}
}

func Bool(key string, val bool) Field {
	return Field{Key: key, Type: BoolType, Interface: val}
}

func Error(err error) Field {
	if err == nil {
		return Field{Key: "error", Type: ErrorType, Interface: nil}
	}
	return Field{Key: "error", Type: ErrorType, Interface: err}
}

func Int32(key string, val int32) Field {
	return Field{Key: key, Type: Int32Type, Integer: int64(val)}
}

func Int16(key string, val int16) Field {
	return Field{Key: key, Type: Int16Type, Integer: int64(val)}
}

func Int8(key string, val int8) Field {
	return Field{Key: key, Type: Int8Type, Integer: int64(val)}
}

func Uint64(key string, val uint64) Field {
	return Field{Key: key, Type: Uint64Type, Interface: val}
}

func Uint32(key string, val uint32) Field {
	return Field{Key: key, Type: Uint32Type, Interface: val}
}

func Uint16(key string, val uint16) Field {
	return Field{Key: key, Type: Uint16Type, Interface: val}
}

func Uint8(key string, val uint8) Field {
	return Field{Key: key, Type: Uint8Type, Interface: val}
}

func Uint(key string, val uint) Field {
	return Field{Key: key, Type: UintType, Interface: val}
}

func Duration(key string, val time.Duration) Field {
	return Field{Key: key, Type: DurationType, Integer: int64(val)}
}

func Time(key string, val time.Time) Field {
	return Field{Key: key, Type: TimeType, Interface: val}
}

func ByteString(key string, val []byte) Field {
	return Field{Key: key, Type: ByteStringType, Interface: val}
}

func Stringer(key string, val fmt.Stringer) Field {
	return Field{Key: key, Type: StringerType, Interface: val}
}

func Any(key string, val interface{}) Field {
	return Field{Key: key, Type: AnyType, Interface: val}
}
