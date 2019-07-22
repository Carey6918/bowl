package main

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func CommonCount(str string) int {
	score := 0
	strs := strings.Split(str, "|")
	for i := 0; i < FrameNums; i++ {
		if strs[i] == "X" {
			score += 10
			if strs[i+1] == "X" {
				score += 10 + getFirstScore(strs[i+2])
				continue
			}
			score += getFirstScore(strs[i+1]) + getSecondScore(strs[i+1])
			continue
		}
		if strs[i][1:2] == "/" {
			score += 10
			score += getFirstScore(strs[i+1])
			continue
		}
		score += getFirstScore(strs[i]) + getSecondScore(strs[i])
	}
	return score
}

func getFirstScore(str string) int {
	num, _ := str2int(str[0:1])
	return num
}
func getSecondScore(str string) int {
	num, _ := str2int(str[1:2])
	return num
}

func initPins() string {
	strs := make([]string, 0, FrameNums+2)
	for i := 0; i < FrameNums; i++ {
		pins1 := rand.Intn(11)
		if pins1 == 10 {
			strs = append(strs, "X")
			continue
		}
		pins2 := rand.Intn(10 - pins1)
		if pins1+pins2 == 10 {
			strs = append(strs, pins2str(pins1)+"/")
			continue
		}
		strs = append(strs, pins2str(pins1)+pins2str(pins2))
	}
	if strs[FrameNums-1] == "X" {
		strs = append(strs, pins2str(rand.Intn(11))+pins2str(rand.Intn(11)))
	}
	if strings.Contains(strs[FrameNums-1], "/") {
		strs = append(strs, pins2str(rand.Intn(11)))
	}
	return strings.Join(strs, "|")
}

func pins2str(n int) string {
	if n == 0 {
		return "-"
	}
	if n == 10 {
		return "X"
	}
	return strconv.Itoa(n)
}

func TestCounting(t *testing.T) {
	for i := 0; i < 100; i++ {
		input := initPins()
		expected := CommonCount(input)
		result := FrameCount(input)
		if result != expected {
			t.Errorf("result is not expected, input= %v result= %v, expected= %v", input, result, expected)
		}
		t.Logf("result is expected, input= %v result= %v", input, result)
	}
}
