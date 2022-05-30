package testing

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestCoverage(t *testing.T) {
	cmd := exec.Command("go", "test", "-coverprofile", "cover.out", "./task/...")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}

	cmd = exec.Command("go", "tool", "cover", "-func", "cover.out")

	out, err := cmd.Output()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(out))
	re, err := regexp.Compile("total:.*")
	if err != nil {
		t.Error(err)
	}
	res := re.FindAll(out, -1)

	if len(res) > 0 {
		re, err := regexp.Compile("\\d*\\.\\d*%")
		if err != nil {
			t.Error(err)
		}
		procent := re.Find(res[0])

		v := strings.TrimRight(string(procent), "%")
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			t.Error(err)
		}

		if f < 80 {
			t.Errorf("coverage is below 80 percent: got %.2f", f)
		}
		return
	}

}
