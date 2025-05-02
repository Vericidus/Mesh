package core

import (
	"encoding/json"
	"fmt"
)

const MESH_VERSION = "0.10"

const DEFAULT_SERVER_CONFIG_PATH = ".mesh.server.config"
const DEFAULT_COMPUTE_CONFIG_PATH = "./mesh.compute.config"

func Logln(v ...any) {
	fmt.Println(fmt.Sprint(v...))
}

func Logf(s string, a ...any) {
	fmt.Printf(s, a...)
}

// Pretty print structs
type Struct any

func PrettyPrint(a Struct) string {
	b, _ := json.MarshalIndent(a, "", "  ")
	return string(b)
}

type GenericCol[T any] struct {
	Header string
	Rows   []T
}

type TableColumn interface {
	header() string
	len() int
	cell(i int) string
}

func (c GenericCol[T]) header() string {
	return c.Header
}

func (c GenericCol[T]) len() int {
	return len(c.Rows)
}

func (c GenericCol[T]) cell(i int) string {
	
	return fmt.Sprint(c.Rows[i])
}

// Prints a table with variable number of columns.
func LogTable(v ...TableColumn) {
	colWidths := make([]int, len(v))
	totalWidth := 0

	for j := range len(v) {
		colWidths[j] = 0
		for i := range v[j].len() {
			colWidths[j] = max(colWidths[j], len(v[j].cell(i)))
		}

		totalWidth += colWidths[j]
	}

	for range totalWidth + (len(colWidths))*3 - 1 {
		Logf("-")
	}
	Logln()
	for i := range v[0].len() {
		Logf("|")
		for j := range len(v) {
			Logf("%-*s | ", colWidths[j], v[j].cell(i))
		}
		Logln()
	}
	for range totalWidth + (len(colWidths))*3 - 1 {
		Logf("-")
	}
	Logln()
}
