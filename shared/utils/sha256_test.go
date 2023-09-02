package utils

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func TestSHA256(t *testing.T) {
	paths := []string{
		"/home/afeather/Videos/2022-05-01_18-28-05.mkv",
		"/home/afeather/Videos/2022-07-30 20-25-16.mkv",
		"/home/afeather/Videos/2022-07-30 20-46-17.mkv",
		"/home/afeather/Videos/1129484104-1-208.mp4",
		"/home/afeather/Videos/1134387778-1-208.mp4",
		"/home/afeather/Videos/mean-girls.mp4",
	}
	for _, s := range paths {
		bytes, err := ioutil.ReadFile(s)
		if err != nil {
			t.Error("file read failed")
			continue
		}
		out, err := exec.Command("sha256sum", s).Output()
		if err != nil {
			t.Error("call failed")
			continue
		}
		r := SHA256(bytes)
		w := string(out)
		fmt.Println(r, w)
		if !strings.Contains(w, r) {
			t.Errorf("different result: %s %s", w, r)
		}
	}
}
