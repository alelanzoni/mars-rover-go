package models

type ConfigurationDto struct {
	Grid     GridDto
	Obstacle []ObstacleDto
}

type PointDto struct {
	XPoint int
	YPoint int
}

type LocationDto struct {
	Point     PointDto
	Direction Direction
}

type GridDto struct {
	XPointMax int
	YPointMax int
}

type ObstacleDto struct {
	Point PointDto
}
