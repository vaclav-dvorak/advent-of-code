package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type poi struct {
	x, y, h int
}

type trail struct {
	points []poi
}

func searchTrails(trails []trail, mountain map[poi]bool) []trail {
	for height := 1; height < 10; height++ {
		newTrails := []trail{}
		for _, t := range trails {
			lastPoi := t.points[len(t.points)-1]
			l := make([]poi, len(t.points))
			copy(l, t.points)
			if ok := mountain[poi{lastPoi.x - 1, lastPoi.y, height}]; ok {
				newTrails = append(newTrails, trail{points: append(l, poi{lastPoi.x - 1, lastPoi.y, height})})
			}
			if ok := mountain[poi{lastPoi.x + 1, lastPoi.y, height}]; ok {
				newTrails = append(newTrails, trail{points: append(l, poi{lastPoi.x + 1, lastPoi.y, height})})
			}
			if ok := mountain[poi{lastPoi.x, lastPoi.y - 1, height}]; ok {
				newTrails = append(newTrails, trail{points: append(l, poi{lastPoi.x, lastPoi.y - 1, height})})
			}
			if ok := mountain[poi{lastPoi.x, lastPoi.y + 1, height}]; ok {
				newTrails = append(newTrails, trail{points: append(l, poi{lastPoi.x, lastPoi.y + 1, height})})
			}
		}
		trails = newTrails
	}
	return trails
}

func calcTrialhead(trails []trail) int {
	unique := make(map[[2]poi]bool)
	for _, t := range trails {
		unique[[2]poi{t.points[0], t.points[len(t.points)-1]}] = true
	}
	return len(unique)
}

func processTrails(input io.Reader) []trail {
	scanner := bufio.NewScanner(input)
	y := 0
	mountain := map[poi]bool{}
	trails := []trail{}
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			h := int(line[x]) - 48
			mountain[poi{x, y, h}] = true
			if h == 0 { // all points with height 0 are trailheads
				trails = append(trails, trail{points: []poi{{x, y, h}}})
			}
		}
		y++
	}
	foundTrails := searchTrails(trails, mountain)
	return foundTrails
}

func fs1(input io.Reader) int {
	foundTrails := processTrails(input)
	return calcTrialhead(foundTrails)
}

func fs2(input io.Reader) int {
	foundTrails := processTrails(input)
	return len(foundTrails)
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs1(f)
}
