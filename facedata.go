package rubiks_cube

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	parseTopRegex    = regexp.MustCompile("^ {3}[wyogrb]{3}( {6})?$")
	parseMiddleRegex = regexp.MustCompile("^[wyogrb]{12}$")
)

type CubeFaceData [6]FaceData

type FaceData []Color

var ErrInvalidCubeString = errors.New("invalid cube string")

func ParseFaces(v string) (CubeFaceData, error) {
	var faces CubeFaceData

	scanner := bufio.NewScanner(strings.NewReader(v))
	i := 0
	for i < 9 && scanner.Scan() {
		fmt.Println(i)
		line := scanner.Text()
		switch {
		case i < 3:
			if !parseTopRegex.MatchString(line) {
				return faces, ErrInvalidCubeString
			}
			faces[FaceUp][i*3] = ParseColor(line[3])
			faces[FaceUp][i*3+1] = ParseColor(line[4])
			faces[FaceUp][i*3+2] = ParseColor(line[5])
		case i >= 3 && i < 6:
			if !parseMiddleRegex.MatchString(line) {
				return faces, ErrInvalidCubeString
			}
			j := i - 3
			faces[FaceLeft][j*3] = ParseColor(line[0])
			faces[FaceLeft][j*3+1] = ParseColor(line[1])
			faces[FaceLeft][j*3+2] = ParseColor(line[2])
			faces[FaceFront][j*3] = ParseColor(line[3])
			faces[FaceFront][j*3+1] = ParseColor(line[4])
			faces[FaceFront][j*3+2] = ParseColor(line[5])
			faces[FaceRight][j*3] = ParseColor(line[6])
			faces[FaceRight][j*3+1] = ParseColor(line[7])
			faces[FaceRight][j*3+2] = ParseColor(line[8])
			faces[FaceBack][j*3] = ParseColor(line[9])
			faces[FaceBack][j*3+1] = ParseColor(line[10])
			faces[FaceBack][j*3+2] = ParseColor(line[11])
		case i >= 6 && i < 9:
			if !parseTopRegex.MatchString(line) {
				return faces, ErrInvalidCubeString
			}
			j := i - 6
			faces[FaceDown][j*3] = ParseColor(line[3])
			faces[FaceDown][j*3+1] = ParseColor(line[4])
			faces[FaceDown][j*3+2] = ParseColor(line[5])
		default:
			return faces, ErrInvalidCubeString
		}
		i++
	}
	return faces, scanner.Err()
}
