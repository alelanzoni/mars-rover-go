package domains

import "github.com/mars-rover-go/models"

type IRoverDomain interface {
	// ExecuteCommands executes the mars rover commands.
	// Commands:
	//  - f: forward
	//  - b: backward
	//  - l: left
	//  - r: right
	// Returns the rover location (x, y and direction).
	ExecuteCommands(commands []string) (models.LocationDto, error)
}

type IGridDomain interface {
	// IsPointInGrid checks if the point is in the grid
	IsPointInGrid(point models.PointDto) bool
}

type IObstacleDomain interface {
	// IsObstacle checks if the point is an obstacle
	IsObstacle(point models.PointDto) bool
}
