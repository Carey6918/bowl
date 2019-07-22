package main

import (
	"strconv"
)

const TotalPins = 10

type Frame interface {
	CountingScore([]Frame) int
	FirstScore() int
	SecondScore() int
	ToString() string
}

type Strike struct {
}
type Spare struct {
	FirstPins int
}
type Remain struct {
	FirstPins  int
	SecondPins int
}

func (s *Strike) CountingScore(next []Frame) int {
	if len(next) == 0 {
		return TotalPins
	}
	if next[0].SecondScore() == -1 {
		return TotalPins + TotalPins + next[1].FirstScore()
	}
	return TotalPins + next[0].SecondScore()
}

func (s *Strike) FirstScore() int {
	return TotalPins
}

func (s *Strike) SecondScore() int {
	return -1
}

func (s *Strike) ToString() string {
	return "X"
}

func (s *Spare) CountingScore(next []Frame) int {
	if len(next) == 0 {
		return TotalPins
	}
	return TotalPins + next[0].FirstScore()
}

func (s *Spare) FirstScore() int {
	return s.FirstPins
}

func (s *Spare) SecondScore() int {
	return TotalPins
}

func (s *Spare) ToString() string {
	return strconv.Itoa(s.FirstPins) + "/"
}

func (r *Remain) CountingScore(next []Frame) int {
	return r.FirstPins + r.SecondPins
}

func (r *Remain) FirstScore() int {
	return r.FirstPins
}

func (r *Remain) SecondScore() int {
	return r.FirstPins + r.SecondPins
}

func (r *Remain) ToString() string {
	return strconv.Itoa(r.FirstPins) + strconv.Itoa(r.SecondPins)
}
