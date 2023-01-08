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
// put commas between items, and python wants commas. So we write this helper function
func str_array(data []int) string {
	var sb strings.Builder
	sb.WriteString("[")
	for _, v := range data {
		sb.WriteString(fmt.Sprintf("%d,", v))
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
