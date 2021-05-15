package domains

import (
	"testing"

	"github.com/mars-rover-go/models"
)

func TestGridDomain_IsPointInGrid(t *testing.T) {
	type args struct {
		point models.PointDto
	}

	gridDomain := &GridDomain{
		grid: models.GridDto{XPointMax: 10, YPointMax: 10},
	}

	tests := []struct {
		name string
		g    *GridDomain
		args args
		want bool
	}{
		{
			name: "Point Ok",
			g:    gridDomain,
			args: args{
				point: models.PointDto{XPoint: 1, YPoint: 2},
			},
			want: true,
		},
		{
			name: "Point KO - XPoint higher than XPointMax",
			g:    gridDomain,
			args: args{
				point: models.PointDto{XPoint: 11, YPoint: 6},
			},
			want: false,
		},
		{
			name: "Point KO - YPoint higher than YPointMax",
			g:    gridDomain,
			args: args{
				point: models.PointDto{
					XPoint: 6,
					YPoint: 12,
				},
			},
			want: false,
		},
		{
			name: "Point KO - XPoint minor to 0",
			g:    gridDomain,
			args: args{
				point: models.PointDto{XPoint: -1, YPoint: 10},
			},
			want: false,
		},
		{
			name: "Point KO - YPoint minor to 0",
			g:    gridDomain,
			args: args{
				point: models.PointDto{XPoint: 0, YPoint: -5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.IsPointInGrid(tt.args.point); got != tt.want {
				t.Errorf("GridDomain.IsPointInGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
