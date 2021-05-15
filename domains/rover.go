package domains

import (
	"fmt"
	"strings"

	"github.com/mars-rover-go/models"
	"github.com/mars-rover-go/utils"
)

type RoverDomain struct {
	location       models.LocationDto
	gridDomain     IGridDomain
	obstacleDomain IObstacleDomain
}

func NewRoverDomain(startingLocation models.LocationDto, gridDomain IGridDomain, obstacleDomain IObstacleDomain) (IRoverDomain, error) {
	isPointInGrid := gridDomain.IsPointInGrid(startingLocation.Point)
	if !isPointInGrid {
		return nil, fmt.Errorf("starting location is out the grid")
	}

	return &RoverDomain{
		startingLocation,
		gridDomain,
		obstacleDomain,
	}, nil
}

func (r *RoverDomain) ExecuteCommands(commands []string) (models.LocationDto, error) {
	location := &r.location

	if len(commands) == 0 {
		return *location, nil
	}

	for _, cmd := range commands {
		var err error = nil
		point := location.Point
		direction := location.Direction

		switch strings.ToLower(cmd) {
		case string(models.CommandForward):
			point, err = r.move(*location, models.MoveTypeForward)
		case string(models.CommandBackward):
			point, err = r.move(*location, models.MoveTypeBackward)
		case string(models.CommandLeft):
			direction = r.turnLeft(direction)
		case string(models.CommandRight):
			direction = r.turnRight(direction)
		default:
			return *location, fmt.Errorf("command '%s' unknown", cmd)
		}

		if err != nil {
			return *location, err
		}

		location.Point = point
		location.Direction = direction
	}

	return *location, nil
}

func (r *RoverDomain) move(currentLocation models.LocationDto, moveType models.MoveType) (models.PointDto, error) {
	point := currentLocation.Point

	switch currentLocation.Direction {
	case models.DirectionNorth:
		point.YPoint = point.YPoint + int(moveType)
	case models.DirectionSouth:
		point.YPoint = point.YPoint - int(moveType)
	case models.DirectionEast:
		point.XPoint = point.XPoint + int(moveType)
	case models.DirectionWest:
		point.XPoint = point.XPoint - int(moveType)
	}

	// Detectes obstacle
	if r.obstacleDomain.IsObstacle(point) {
		return currentLocation.Point, fmt.Errorf("obstacle detected - last possible point: %s", utils.LocationToString(r.location))
	}

	// Checks if the point is in the grid
	isPointInGrid := r.gridDomain.IsPointInGrid(point)
	if !isPointInGrid {
		return currentLocation.Point, nil
	}

	return point, nil
}

func (r *RoverDomain) turnLeft(currentDirection models.Direction) models.Direction {
	var newDirection models.Direction

	switch currentDirection {
	case models.DirectionNorth:
		newDirection = models.DirectionWest
	case models.DirectionSouth:
		newDirection = models.DirectionEast
	case models.DirectionEast:
		newDirection = models.DirectionNorth
	case models.DirectionWest:
		newDirection = models.DirectionSouth
	}

	return newDirection
}

func (r *RoverDomain) turnRight(currentDirection models.Direction) models.Direction {
	var newDirection models.Direction

	switch currentDirection {
	case models.DirectionNorth:
		newDirection = models.DirectionEast
	case models.DirectionSouth:
		newDirection = models.DirectionWest
	case models.DirectionEast:
		newDirection = models.DirectionSouth
	case models.DirectionWest:
		newDirection = models.DirectionNorth
	}

	return newDirection
}
