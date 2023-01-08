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
