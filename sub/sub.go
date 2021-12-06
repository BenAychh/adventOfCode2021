package sub

import (
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidMovement = fmt.Errorf("movements must in the format `Movement Number`")
type Movement string
const (
	Forward Movement = "forward"
	Down Movement = "down"
	Up Movement = "up"
)

type Sub struct {
	HorizontalPos int
	VerticalPos int
	Aim int
	BasicMode bool
}

func (s *Sub) Move(commands []string) error {
	fn := s.moveAim
	if s.BasicMode {
		fn = s.move
	}
	for _, cmd := range commands {
		err := fn(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Sub) Reset() {
	s.VerticalPos = 0
	s.Aim = 0
	s.HorizontalPos = 0
	s.BasicMode = false
}

func (s *Sub) move(command string) error {
	mvt, count, err := splitMovement(command)
	if err != nil {
		return err
	}
	switch mvt {
	case Forward:
		s.HorizontalPos += count
	case Down:
		s.VerticalPos += count
	case Up:
		s.VerticalPos -= count
	}
	return nil
}

func (s *Sub) moveAim(command string) error {
	mvt, count, err := splitMovement(command)
	if err != nil {
		return err
	}
	switch mvt {
	case Forward:
		s.HorizontalPos += count
		s.VerticalPos += count * s.Aim
	case Down:
		s.Aim += count
	case Up:
		s.Aim -= count
	}
	return nil
}

func splitMovement(str string) (mvt Movement, count int, err error) {
	split := strings.Split(str, " ")
	if len(split)!= 2 {
		err = ErrInvalidMovement
		return
	}
	mvt = Movement(split[0])
	count, err = strconv.Atoi(split[1])
	return
}