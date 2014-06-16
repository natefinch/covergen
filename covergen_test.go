// +build !covergen

package foo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"testing"
)

var (
	pctReg    = regexp.MustCompile(`.*coverage: ([0-9\.]+)% of statements.*`)
	shieldReg = regexp.MustCompile(
		`http://img\.shields\.io/badge/coverage-[0-9\.]+%25-[a-z]+\.svg`)
	replFormat = "http://img.shields.io/badge/coverage-%.1f%%25-%s.svg"
)

func TestCovergen(t *testing.T) {
	cmd := exec.Command("go", "test", "-cover", "-tags", "covergen")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Fatal("Error running go test -cover")
	}
	matches := pctReg.FindSubmatch(out.Bytes())
	if len(matches) < 2 {
		t.Fatal("Failed to find regex match in cover output")
	}

	percent, err := strconv.ParseFloat(string(matches[1]), 64)
	if err != nil {
		t.Fatalf("Failed to parse percent string: %s", err)
	}

	var color string
	if percent > 90.0 {
		color = "brightgreen"
	} else if percent > 80.0 {
		color = "green"
	} else if percent > 70.0 {
		color = "yellowgreen"
	} else if percent > 60.0 {
		color = "yellow"
	} else if percent > 50.0 {
		color = "orange"
	} else {
		color = "red"
	}

	repl := []byte(fmt.Sprintf(replFormat, percent, color))

	src, err := ioutil.ReadFile("README.md")
	if err != nil {
		t.Fatalf("Failed to read README.md: %s", err)
	}

	data := shieldReg.ReplaceAllLiteral(src, repl)

	if !bytes.Equal(data, src) {
		err = ioutil.WriteFile("README.md", data, 0644)
		if err != nil {
			t.Fatalf("Failed to write out README.md: %s", err)
		}
	}
}
