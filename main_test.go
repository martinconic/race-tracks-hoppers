package main

import "testing"

func TestFindPathWithSolution(t *testing.T) {

	start := Point{4, 0}
	end := Point{4, 4}
	grid := Grid{
		5,
		5,
		map[Point]bool{
			{1, 2}: true,
			{1, 3}: true,
			{2, 2}: true,
			{2, 3}: true,
			{3, 2}: true,
			{3, 3}: true,
			{4, 2}: true,
			{4, 3}: true,
		},
	}

	got := findPath(start, end, grid)

	want := "Optimal solution takes 7 hops"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func TestFindPathNoSolution(t *testing.T) {

	start := Point{0, 0}
	end := Point{2, 2}
	grid := Grid{
		5,
		5,
		map[Point]bool{
			{0, 1}: true, {1, 0}: true, {1, 1}: true, {1, 2}: true, {2, 1}: true,
		},
	}

	got := findPath(start, end, grid)

	want := "No solution."

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
