package domains

import "github.com/mars-rover-go/models"

type GridDomain struct {
	grid models.GridDto
}

func NewGridDomain(grid models.GridDto) IGridDomain {
	return &GridDomain{grid}
}

func (g *GridDomain) IsPointInGrid(point models.PointDto) bool {
	if point.XPoint < 0 || point.XPoint > g.grid.XPointMax ||
		point.YPoint < 0 || point.YPoint > g.grid.YPointMax {
		return false
	}

	return true
}
