package domains

import "github.com/mars-rover-go/models"

type ObstacleDomain struct {
	obstacles []models.ObstacleDto
}

func NewObstacleDomain(obstacles []models.ObstacleDto) IObstacleDomain {
	return &ObstacleDomain{obstacles}
}

func (o *ObstacleDomain) IsObstacle(point models.PointDto) bool {
	for _, ob := range o.obstacles {
		if ob.Point.XPoint == point.XPoint && ob.Point.YPoint == point.YPoint {
			// obstacle detected
			return true
		}
	}

	return false
}
