package snake

import (
	"math/rand"

	"github.com/kamilGie/snake-golang/snake/point"
)

type directions [4]int

type Snake struct {
	body      []point.Point
	direction directions
	fruit     point.Point
	areaWidth int
	areaHight int
	GameOver  bool
}

func New(areaWidth, areaHight int) *Snake {
	return &Snake{
		body:      []point.Point{{X: 1, Y: 1}, {X: 1, Y: 2}, {X: 1, Y: 3}},
		direction: directions{0, 0, 0, 1},
		fruit:     point.Point{X: 0, Y: 0},
		areaWidth: areaWidth,
		areaHight: areaHight,
		GameOver:  false,
	}
}

func (s *Snake) isEndGame(newPoint point.Point) bool {
	if newPoint.X < 0 || newPoint.X >= s.areaWidth || newPoint.Y < 0 || newPoint.Y >= s.areaHight {
		return true
	}

	for _, p := range s.body {
		if p.X == newPoint.X && p.Y == newPoint.Y {
			return true
		}
	}
	return false
}

// return coordinates of ( body of snake , fruit )
func (s *Snake) GetState() ([]point.Point, point.Point, bool) {
	return s.body, s.fruit, s.GameOver
}

// todo this can never end repair it to ending func
func (s *Snake) newFruitLocation() {
	for {
		randomX := rand.Intn(s.areaWidth)
		randomY := rand.Intn(s.areaHight)

		collision := false
		for _, p := range s.body {
			if p.X == randomX && p.Y == randomY {
				collision = true
				break
			}
		}

		if !collision {
			s.fruit = point.Point{X: randomX, Y: randomY}
			break
		}
	}
}

func (s *Snake) TakeAction(newDirection directions) {
	//check is newDirection present and possible
	for index, value := range newDirection {
		if value == 1 && s.direction[(index+2)%4] == 0 {
			s.direction = newDirection
			break
		}
	}

	newPoint := point.NewPointAtDir(s.body[len(s.body)-1], s.direction)
	if s.isEndGame(newPoint) {
		s.GameOver = true
		return
	}
	s.body = append(s.body, newPoint)

	if newPoint != s.fruit {
		s.body = s.body[1:]
	} else {
		s.newFruitLocation()
	}

	return
}
