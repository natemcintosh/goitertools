package goitertools

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// pythonCount runs python's `itertools.count(start, step)` `n` times and returns the
// result as a `[]int`
func pythonCount(start, step, n int) []int {
	py_str := fmt.Sprintf("from py_versions import n_count; n_count(%d, %d, %d)", start, step, n)
	cmd := exec.Command("python", "-c", py_str)

	out, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	s := strings.Split(strings.TrimSpace(string(out)), "\n")
	result := make([]int, len(s))

	for idx, sn := range s {
		res, err := strconv.Atoi(sn)
		if err != nil {
			log.Fatalln(err)
		}
		result[idx] = res
	}
	return result
}

// str_array is a helper function that converts a slice into a string representation
// of the slice. You might ask why not just `fmt.Sprintf` the slice, but that does not
// put commas between items, and python wants commas. So we write this helper function.
//
// This is basically a copy of part of the source code for `strings.Join`
func str_array[T any](data []T) string {
	switch len(data) {
	case 0:
		return "[]"
	case 1:
		return "[" + fmt.Sprint(data[0]) + "]"
	}

	var sb strings.Builder
	// Make a rough estimate of the lower bound of the number of runes needed
	sb.Grow(len(data) * 2)
	sb.WriteString("[")
	sb.WriteString(fmt.Sprint(data[0]))
	for _, v := range data[1:] {
		sb.WriteString(",")
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteString("]")
	return sb.String()
}

func pythonAccumulate(data []int) []int {
	str_data := str_array(data)
	// We're just going to do the most simple type of accumulation here, adding each item
	// to the sum
	py_str := fmt.Sprintf("from py_versions import accumulate; accumulate(%s)", str_data)
	cmd := exec.Command("python", "-c", py_str)

	out, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	s := strings.Split(strings.TrimSpace(string(out)), "\n")
	result := make([]int, len(s))

	for idx, sn := range s {
		res, err := strconv.Atoi(sn)
		if err != nil {
			log.Fatalln(err)
		}
		result[idx] = res
	}
	return result

}
