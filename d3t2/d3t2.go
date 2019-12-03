package main

import (
	"bufio"
	"common"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	x1, y1, x2, y2, cumulSteps int
}

func (l line) IsVertical() bool {
	return l.x1-l.x2 == 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (l1 line) IntersectsSteps(l2 line) (bool, int) {
	var vLine, hLine line
	if l1.IsVertical() && !l2.IsVertical() {
		vLine, hLine = l1, l2
	} else if !l1.IsVertical() && l2.IsVertical() {
		vLine, hLine = l2, l1
	} else {
		return false, 0
	}

	if (vLine.x1 >= hLine.x1 && vLine.x1 <= hLine.x2) || (vLine.x1 >= hLine.x2 && vLine.x1 <= hLine.x1) {
		if (hLine.y1 >= vLine.y1 && hLine.y1 <= vLine.y2) || (hLine.y1 >= vLine.y2 && hLine.y1 <= vLine.y1) {
			stepsVLine := vLine.cumulSteps + Abs(vLine.y1-hLine.y1)
			stepsHLine := hLine.cumulSteps + Abs(hLine.x1-vLine.x1)
			return true, stepsVLine + stepsHLine
		}
	}

	// Parallel lines or no overlap
	return false, 0
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readInputLine(inputLine string) []line {
	lineStartX, lineStartY, steps := 0, 0, 0
	lines := make([]line, 0)

	scanner := bufio.NewReader(strings.NewReader(inputLine))
	for {
		tokenStr, scanErr := scanner.ReadSlice(',')
		offset := 1
		if scanErr == io.EOF {
			offset = 0
		}
		if scanErr != nil && scanErr != io.EOF {
			log.Panic(scanErr)
		}

		lineLength, err := strconv.Atoi(string(tokenStr[1 : len(tokenStr)-offset]))
		panicIfErr(err)

		lineDirection := tokenStr[0]
		lineEndX, lineEndY := lineStartX, lineStartY
		switch lineDirection {
		case 'U':
			lineEndY += lineLength
		case 'R':
			lineEndX += lineLength
		case 'D':
			lineEndY -= lineLength
		case 'L':
			lineEndX -= lineLength
		default:
			log.Panic("%s is not a valid direction", lineDirection)
		}

		lines = append(lines, line{lineStartX, lineStartY, lineEndX, lineEndY, steps})
		steps += lineLength
		lineStartX, lineStartY = lineEndX, lineEndY

		if scanErr == io.EOF {
			return lines
		}

	}
}

func readInput(path string) [2][]line {
	inputLines, err := common.ReadInputFile(path)
	panicIfErr(err)

	return [2][]line{readInputLine(inputLines[0]), readInputLine(inputLines[1])}

}

func findClosestIntersect(lineSet1, lineSet2 []line) (bestSteps int) {

	bestSteps = 100000000000
	for _, l1 := range lineSet1 {
		for _, l2 := range lineSet2 {
			intersects, steps := l1.IntersectsSteps(l2)
			if intersects {
				if steps != 0 && steps < bestSteps {
					bestSteps = steps
				}
			}
		}
	}
	return
}

func main() {
	var args = os.Args

	lines := readInput(args[1])
	steps := findClosestIntersect(lines[0], lines[1])

	fmt.Printf("Best steps: %d\n", steps)
}
