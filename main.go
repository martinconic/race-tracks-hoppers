package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Grid struct {
	Gx  int
	Gy  int
	Obs map[Point]bool
}

type State struct {
	Pos   Point
	Speed Point
	steps int
}

func isValid(point Point, grid Grid) bool {
	return point.X >= 0 && point.X < grid.Gx && point.Y >= 0 && point.Y < grid.Gy && !grid.Obs[point]
}

func findPath(start, end Point, grid Grid) string {
	queue := []State{{start, Point{0, 0}, 0}}
	visited := make(map[State]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Pos.X == end.X && current.Pos.Y == end.Y {
			return fmt.Sprintf("Optimal solution takes %d hops", current.steps)

		}
		if visited[current] {
			continue
		}

		visited[current] = true

		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				speed := Point{current.Speed.X + dx, current.Speed.Y + dy}
				if abs(speed.X) <= 3 && abs(speed.Y) <= 3 {
					newX := current.Pos.X + speed.X
					newY := current.Pos.Y + speed.Y
					newPoint := Point{newX, newY}
					if isValid(newPoint, grid) {
						queue = append(queue, State{newPoint, speed, current.steps + 1})
					}
				}
			}
		}

		if current.Pos.X == queue[0].Pos.X && current.Pos.Y == queue[0].Pos.Y &&
			current.Speed.X == queue[0].Speed.X && current.Speed.Y == queue[0].Speed.Y {
			return "No solution."
		}
	}

	return "No solution."
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {
	if len(os.Args) < 2 {
		log.Println("input data file missing!")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("can not open file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	cases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("erorr reading test cases: %d -> %v", cases, err)
	}

	for i := 0; i < cases; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		gx, _ := strconv.Atoi(line[0])
		gy, _ := strconv.Atoi(line[1])

		grid := Grid{
			Gx:  gx,
			Gy:  gy,
			Obs: make(map[Point]bool),
		}

		scanner.Scan()
		line = strings.Fields(scanner.Text())
		if len(line) != 4 {
			log.Fatalf("error defining start and end points: %s", line)
		}

		x1, _ := strconv.Atoi(line[0])
		y1, _ := strconv.Atoi(line[1])
		x2, _ := strconv.Atoi(line[2])
		y2, _ := strconv.Atoi(line[3])

		if x1 == x2 && y1 == y2 {
			log.Println("Optimal solution takes 0 hops")
			continue
		}

		start := Point{
			x1,
			y1,
		}

		end := Point{
			x2,
			y2,
		}

		scanner.Scan()
		p, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < p; j++ {
			scanner.Scan()
			line := strings.Fields(scanner.Text())

			if len(line) != 4 {
				log.Fatalf("obstacles not defined correctly: %s ", line)
			}

			x1, _ := strconv.Atoi(line[0])
			x2, _ := strconv.Atoi(line[1])
			y1, _ := strconv.Atoi(line[2])
			y2, _ := strconv.Atoi(line[3])

			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					o := Point{x, y}
					grid.Obs[o] = true
				}
			}

		}
		log.Println(grid)
		log.Println(findPath(start, end, grid))
	}
}
