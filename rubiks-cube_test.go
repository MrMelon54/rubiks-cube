package rubiks_cube

import (
	"bufio"
	"bytes"
	"embed"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

//go:embed run-test-files/*
var testFiles embed.FS

func TestRubiksCube(t *testing.T) {
	files := make(chan string, 10)
	go listFiles("run-test-files", files)

	for i := range files {
		open, err := testFiles.Open(filepath.Join("run-test-files", i))
		assert.NoError(t, err)
		t.Run(i, func(t *testing.T) {
			assert.NoError(t, startScanningText(t, open, runTests))
		})
	}
}

func runTests(t *testing.T, r io.Reader) {
	rd := bufio.NewReader(r)
	var s strings.Builder
	s.Grow(13 * 9)
	for i := 0; i < 9; i++ {
		line, err := rd.ReadBytes('\n')
		if err != nil {
			return
		}
		s.Write(line)
	}
	cube, err := ParseCube(s.String())
	if err != nil {
		t.Error(err)
	}
	line, err := rd.ReadBytes('\n')
	if err != nil {
		return
	}
	if line[0] == '#' && line[1] == ' ' {
		scanner := NewMoveScanner(bytes.NewReader(line[2:]))
		for scanner.Scan() {
			fmt.Println(scanner.Current())
			cube = cube.Move(scanner.Current())
		}
		fmt.Println(cube)
	}
}

func startScanningText(t *testing.T, open io.Reader, cb func(t *testing.T, r io.Reader)) error {
	rd := bufio.NewReader(open)
	for {
		b, err := rd.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		line := string(b)
		if strings.HasPrefix(line, "=== ") && strings.HasSuffix(line, " ===\n") {
			t.Run(line[4:len(line)-5], func(t *testing.T) {
				cb(t, rd)
			})
		}
	}
}

func listFiles(start string, files chan<- string) {
	dirNames := []string{""}
	for _, i := range dirNames {
		dir, err := fs.ReadDir(testFiles, filepath.Join(start, i))
		if err != nil {
			panic(err)
		}
		for _, j := range dir {
			n := filepath.Join(i, j.Name())
			if j.IsDir() {
				dirNames = append(dirNames, n)
			} else {
				files <- n
			}
		}
	}
	close(files)
}
