package rubiks_cube

import (
	"bufio"
	"bytes"
	"embed"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

//go:embed run-test-files/*
var testFiles embed.FS

func TestRubiksCube_RunTestFiles(t *testing.T) {
	files := make(chan string, 10)
	go listFiles(files)

	for i := range files {
		open, err := testFiles.Open(i)
		assert.NoError(t, err)
		assert.NoError(t, startScanningText(open, runTests))
	}
}

func runTests(r io.Reader) {
	cube, err := ParseCube(r)
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(r)
	line, err := rd.ReadBytes('\n')
	if err != nil {
		return
	}
	if line[0] == '#' && line[1] == ' ' {
		NewMoveScanner(bytes.NewReader(line[2:]))
	}
	cube.Move()
}

func startScanningText(open io.Reader, cb func(r io.Reader)) error {
	rd := bufio.NewReader(open)
	for {
		bytes, err := rd.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		line := string(bytes)
		if strings.HasPrefix(line, "=== ") && strings.HasSuffix(line, " ===") {
			cb(rd)
		}
	}
}

func listFiles(files chan<- string) {
	dirNames := make([]string, 0)

	for _, i := range dirNames {
		dir, err := fs.ReadDir(testFiles, i)
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
