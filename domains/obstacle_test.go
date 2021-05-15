package domains

import (
	"testing"

	"github.com/mars-rover-go/models"
)

func TestObstacleDomain_isObstacle(t *testing.T) {
	type args struct {
		point models.PointDto
	}

	obstacleDomain := &ObstacleDomain{
		[]models.ObstacleDto{
			{Point: models.PointDto{XPoint: 2, YPoint: 6}},
			{Point: models.PointDto{XPoint: 8, YPoint: 5}},
		},
	}

	tests := []struct {
		name string
		o    *ObstacleDomain
		args args
		want bool
	}{
		{
			name: "Obstacle no detected",
			o:    obstacleDomain,
			args: args{
				point: models.PointDto{XPoint: 2, YPoint: 4},
			},
			want: false,
		},
		{
			name: "Obstacle detected",
			o:    obstacleDomain,
			args: args{
				point: models.PointDto{XPoint: 8, YPoint: 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsObstacle(tt.args.point); got != tt.want {
				t.Errorf("ObstacleDomain.isObstacle() = %v, want %v", got, tt.want)
			}
		})
	}
}
