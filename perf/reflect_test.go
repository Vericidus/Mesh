// benchmark_test.go
package main

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"testing"
)

// Sample types for testing
type Point struct {
	X, Y int
}

type SimpleType int

// Generic function to convert to string using type specialization
func GenericToString[T any](v T) string {
	// Type-specific handling
	switch any(v).(type) {
	case int:
		return strconv.Itoa(any(v).(int))
	case float64:
		return strconv.FormatFloat(any(v).(float64), 'f', 2, 64)
	case string:
		return any(v).(string)
	case Point:
		p := any(v).(Point)
		return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
	case SimpleType:
		return strconv.Itoa(int(any(v).(SimpleType)))
	case net.Listener:
		l := any(v).(net.Listener)
		if l == nil {
			return "<nil>"
		}
		return "Listener:" + l.Addr().String()
	case error:
		err := any(v).(error)
		if err == nil {
			return "<nil>"
		}
		return "Error:" + err.Error()
	default:
		// Fallback to reflection for unknown types
		return fmt.Sprint(v)
	}
}

// Reflection-based function (similar to fmt.Sprint)
func ReflectionToString(v any) string {
	// Simulate what fmt.Sprint does with reflection
	if v == nil {
		return "<nil>"
	}

	// Check if the value can be converted to string
	val := reflect.ValueOf(v)

	// Try type-specific handling via reflection
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', 2, 64)
	case reflect.String:
		return val.String()
	case reflect.Struct:
		// This is a simplified version, fmt.Sprint is much more complex
		return fmt.Sprint(v)
	case reflect.Ptr:
		if val.IsNil() {
			return "<nil>"
		}
		// Dereference and continue
		return ReflectionToString(val.Elem().Interface())
	default:
		// Fallback for other types
		return fmt.Sprint(v)
	}
}

// Specialized functions for each type (baseline for comparison)
func IntToString(v int) string {
	return strconv.Itoa(v)
}

func FloatToString(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}

func PointToString(p Point) string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

func SimpleTypeToString(s SimpleType) string {
	return strconv.Itoa(int(s))
}

// Benchmarks
func BenchmarkSpecializedInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IntToString(42)
	}
}

func BenchmarkGenericInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenericToString(42)
	}
}

func BenchmarkReflectionInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReflectionToString(42)
	}
}

func BenchmarkFmtSprintInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(42)
	}
}

func BenchmarkSpecializedFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FloatToString(3.14159)
	}
}

func BenchmarkGenericFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenericToString(3.14159)
	}
}

func BenchmarkReflectionFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReflectionToString(3.14159)
	}
}

func BenchmarkFmtSprintFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(3.14159)
	}
}

func BenchmarkSpecializedStruct(b *testing.B) {
	p := Point{10, 20}
	for i := 0; i < b.N; i++ {
		_ = PointToString(p)
	}
}

func BenchmarkGenericStruct(b *testing.B) {
	p := Point{10, 20}
	for i := 0; i < b.N; i++ {
		_ = GenericToString(p)
	}
}

func BenchmarkReflectionStruct(b *testing.B) {
	p := Point{10, 20}
	for i := 0; i < b.N; i++ {
		_ = ReflectionToString(p)
	}
}

func BenchmarkFmtSprintStruct(b *testing.B) {
	p := Point{10, 20}
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(p)
	}
}

func BenchmarkSpecializedCustomType(b *testing.B) {
	var s SimpleType = 42
	for i := 0; i < b.N; i++ {
		_ = SimpleTypeToString(s)
	}
}

func BenchmarkGenericCustomType(b *testing.B) {
	var s SimpleType = 42
	for i := 0; i < b.N; i++ {
		_ = GenericToString(s)
	}
}

func BenchmarkReflectionCustomType(b *testing.B) {
	var s SimpleType = 42
	for i := 0; i < b.N; i++ {
		_ = ReflectionToString(s)
	}
}

func BenchmarkFmtSprintCustomType(b *testing.B) {
	var s SimpleType = 42
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(s)
	}
}

// Run this benchmark with:
// go test -bench=. -benchmem
