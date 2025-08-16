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

	"go.uber.org/zap"
)

// zapLoggerAdapter wraps a *zap.Logger to implement the Logger interface.
type zapLoggerAdapter struct {
	logger *zap.Logger
}

var _ Logger = (*zapLoggerAdapter)(nil)

// NewZapAdapter creates a new Logger that wraps the given *zap.Logger.
func NewZapAdapter(logger *zap.Logger) Logger {
	return &zapLoggerAdapter{logger: logger}
}

func (z *zapLoggerAdapter) Debug(msg string, fields ...Field) {
	z.logger.Debug(msg, convertFields(fields)...)
}

func (z *zapLoggerAdapter) Info(msg string, fields ...Field) {
	z.logger.Info(msg, convertFields(fields)...)
}

func (z *zapLoggerAdapter) Warn(msg string, fields ...Field) {
	z.logger.Warn(msg, convertFields(fields)...)
}

func (z *zapLoggerAdapter) Error(msg string, fields ...Field) {
	z.logger.Error(msg, convertFields(fields)...)
}

func (z *zapLoggerAdapter) Panic(msg string, fields ...Field) {
	z.logger.Panic(msg, convertFields(fields)...)
}

func (z *zapLoggerAdapter) Fatal(msg string, fields ...Field) {
	z.logger.Fatal(msg, convertFields(fields)...)
}

func (z *zapLoggerAdapter) Sugar() SugaredLogger {
	return z.logger.Sugar()
}

// convertFields converts our Field type to zap.Field.
func convertFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = convertField(f)
	}
	return zapFields
}

// convertField converts a single Field to zap.Field.
func convertField(f Field) zap.Field {
	switch f.Type {
	case StringType:
		return zap.String(f.Key, f.String)
	case Int64Type:
		return zap.Int64(f.Key, f.Integer)
	case IntType:
		return zap.Int(f.Key, int(f.Integer))
	case Int32Type:
		return zap.Int32(f.Key, int32(f.Integer))
	case Int16Type:
		return zap.Int16(f.Key, int16(f.Integer))
	case Int8Type:
		return zap.Int8(f.Key, int8(f.Integer))
	case Uint64Type:
		return zap.Uint64(f.Key, f.Interface.(uint64))
	case Uint32Type:
		return zap.Uint32(f.Key, f.Interface.(uint32))
	case Uint16Type:
		return zap.Uint16(f.Key, f.Interface.(uint16))
	case Uint8Type:
		return zap.Uint8(f.Key, f.Interface.(uint8))
	case UintType:
		return zap.Uint(f.Key, f.Interface.(uint))
	case Float64Type:
		return zap.Float64(f.Key, f.Interface.(float64))
	case Float32Type:
		return zap.Float32(f.Key, f.Interface.(float32))
	case BoolType:
		return zap.Bool(f.Key, f.Interface.(bool))
	case ErrorType:
		if f.Interface == nil {
			return zap.Error(nil)
		}
		return zap.Error(f.Interface.(error))
	case DurationType:
		return zap.Duration(f.Key, time.Duration(f.Integer))
	case TimeType:
		return zap.Time(f.Key, f.Interface.(time.Time))
	case ByteStringType:
		return zap.ByteString(f.Key, f.Interface.([]byte))
	case StringerType:
		return zap.Stringer(f.Key, f.Interface.(fmt.Stringer))
	case AnyType:
		return zap.Any(f.Key, f.Interface)
	default:
		return zap.Any(f.Key, f.Interface)
	}
}
