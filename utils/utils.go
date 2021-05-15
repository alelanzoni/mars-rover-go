package utils

import (
	"fmt"

	"github.com/mars-rover-go/models"
)

func LocationToString(location models.LocationDto) string {
	return fmt.Sprintf("(%d,%d) %s", location.Point.XPoint, location.Point.YPoint, location.Direction)
}
