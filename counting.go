package main

import (
	"fmt"
	"strconv"
	"strings"
)

const FrameNums = 10

func Counting(frames []Frame) int {
	score := 0
	for i := 0; i < FrameNums; i++ {
		score += frames[i].CountingScore(frames[i+1:])
	}
	return score
}

func main() {
	input := "X|7/|90|X|08|8/|06|X|X|X|81"
	fmt.Printf("Total score == %v", FrameCount(input))
}

func FrameCount(str string) int {
	strs := strings.Split(str, "|")
	frames := make([]Frame, 0, TotalPins+2)
	for _, str := range strs {
		frame, err := str2frames(str)
		if err != nil {
			fmt.Printf("Err = %v", err)
			return 0
		}
		frames = append(frames, frame)
	}
	return Counting(frames)
}

func str2frames(str string) (Frame, error) {
	if str == "X" {
		return &Strike{}, nil
	}
	if strings.HasSuffix(str, "/") {
		firstPins, err := str2int(str[0:1])
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}
		if firstPins >= TotalPins {
			return nil, fmt.Errorf("invalid input, pins out of max")
		}
		return &Spare{
			FirstPins: firstPins,
		}, nil
	}
	firstPins, err := str2int(str[0:1])
	if err != nil {
		return nil, fmt.Errorf("invalid input")
	}
	if len(str) < 2 {
		return &Remain{
			FirstPins:  firstPins,
			SecondPins: 0,
		}, nil
	}
	secondPins, err := str2int(str[1:2])
	if err != nil {
		return nil, fmt.Errorf("invalid input")
	}

	return &Remain{
		FirstPins:  firstPins,
		SecondPins: secondPins,
	}, nil
}

func str2int(str string) (int, error) {
	if str == "-" {
		return 0, nil
	}
	if str == "X" {
		return 10, nil
	}
	return strconv.Atoi(str)
}
