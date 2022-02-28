package main

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		Time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), toRadian(0)},
		{simpleTime(0, 0, 15), toRadian(90)},
		{simpleTime(0, 0, 30), toRadian(180)},
		{simpleTime(0, 0, 45), toRadian(270)},
		{simpleTime(0, 0, 50), toRadian(300)},
	}

	for _, test := range cases {
		t.Run(testName(test.Time), func(t *testing.T) {
			got := secondsInRadians(test.Time)
			if !roughlyEqualFloat64(got, test.angle) {
				t.Errorf("want %g, but got %g", test.angle, got)
			}
		})
	}
}

func TestSecondVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
		{simpleTime(0, 0, 50), Point{-0.866025404, 0.5}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondsInVector(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("want %v, but got %v", test.point, got)
			}
		})
	}
}

func TestSecondHead(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{150, 60}},
		{simpleTime(0, 0, 15), Point{240, 150}},
		{simpleTime(0, 0, 30), Point{150, 240}},
		{simpleTime(0, 0, 45), Point{60, 150}},
		{simpleTime(0, 0, 50), Point{72.05771364, 105}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondHandHead(test.time, 90, Point{150, 150})
			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("want %v, but got %v", test.point, got)
			}
		})
	}
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(1337, 4, 2, hour, minute, second, 0, time.UTC)
}

func toRadian(angle float64) float64 {
	return math.Pi * angle / 180
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	precision := 1e-7
	return math.Abs(a-b) < precision
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
