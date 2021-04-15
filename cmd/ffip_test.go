package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestFfipCommand(t *testing.T) {
	cases := []struct {
		args     []string
		expected []string
	}{
		{
			args:     []string{"ffip", "testdata/ffip-test-dir"},
			expected: execFindTypeF("testdata/ffip-test-dir"),
		},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewToolCommand()
		cmd.SetOut(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		fmt.Println(buf.String())

		actual := strings.Split(buf.String(), "\n")

		sort.SliceStable(actual, func(i, j int) bool {
			return actual[i] > actual[j]
		})

		sort.SliceStable(c.expected, func(i, j int) bool {
			return c.expected[i] > c.expected[j]
		})

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("actual and expected not the same\nactual: %#v\nexpected: %#v\n", actual, c.expected)
		}
	}
}

func execFindTypeF(path string) []string {
	out, err := exec.Command("find", path, "-type", "f").Output()
	if err != nil {
		panic(err)
	}

	return strings.Split(string(out), "\n")
}
