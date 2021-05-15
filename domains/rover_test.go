package domains

import (
	"reflect"
	"testing"

	"github.com/mars-rover-go/models"
)

func TestNewRoverDomain(t *testing.T) {
	type args struct {
		startingLocation models.LocationDto
		gridDomain       IGridDomain
		obstacleDomain   IObstacleDomain
	}
	tests := []struct {
		name    string
		args    args
		want    IRoverDomain
		wantErr bool
	}{
		{
			name: "Starting location is out the grid",
			args: args{
				startingLocation: models.LocationDto{Point: models.PointDto{XPoint: 11, YPoint: 5}, Direction: models.DirectionNorth},
				gridDomain:       &GridDomain{models.GridDto{XPointMax: 10, YPointMax: 10}},
				obstacleDomain:   &ObstacleDomain{[]models.ObstacleDto{}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "New rover domain ok",
			args: args{
				startingLocation: models.LocationDto{Point: models.PointDto{XPoint: 3, YPoint: 5}, Direction: models.DirectionNorth},
				gridDomain:       &GridDomain{models.GridDto{XPointMax: 10, YPointMax: 10}},
				obstacleDomain:   &ObstacleDomain{[]models.ObstacleDto{}},
			},
			want: &RoverDomain{
				models.LocationDto{Point: models.PointDto{XPoint: 3, YPoint: 5}, Direction: models.DirectionNorth},
				&GridDomain{models.GridDto{XPointMax: 10, YPointMax: 10}},
				&ObstacleDomain{[]models.ObstacleDto{}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRoverDomain(tt.args.startingLocation, tt.args.gridDomain, tt.args.obstacleDomain)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRoverDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoverDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoverDomain_ExecuteCommands(t *testing.T) {
	type args struct {
		commands []string
	}

	rdCmdEmpty := newRoverDomainMocked()
	rdCmdUnknown := newRoverDomainMocked()
	rdObstacle := newRoverDomainMocked()
	rd := newRoverDomainMocked()

	tests := []struct {
		name    string
		r       *RoverDomain
		args    args
		want    models.LocationDto
		wantErr bool
	}{
		{
			name: "Commands empty",
			r:    &rdCmdEmpty,
			args: args{
				commands: []string{},
			},
			want:    rdCmdEmpty.location,
			wantErr: false,
		},
		{
			name: "Command unknown",
			r:    &rdCmdUnknown,
			args: args{
				commands: []string{"f", "T", "r"},
			},
			want:    models.LocationDto{Point: models.PointDto{XPoint: 1, YPoint: 2}, Direction: models.DirectionNorth},
			wantErr: true,
		},
		{
			name: "Command with obstacle",
			r:    &rdObstacle,
			args: args{
				commands: []string{"r", "f", "l", "f", "f", "f", "f", "f"},
			},
			want:    models.LocationDto{Point: models.PointDto{XPoint: 2, YPoint: 5}, Direction: models.DirectionNorth},
			wantErr: true,
		},
		{
			name: "Commands ok",
			r:    &rd,
			args: args{
				commands: []string{"f", "F", "r", "f", "f", "l", "b"},
			},
			want:    models.LocationDto{Point: models.PointDto{XPoint: 3, YPoint: 2}, Direction: models.DirectionNorth},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ExecuteCommands(tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoverDomain.ExecuteCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoverDomain.ExecuteCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoverDomain_move(t *testing.T) {
	type args struct {
		currentLocation models.LocationDto
		moveType        models.MoveType
	}

	roverDomain := newRoverDomainMocked()

	tests := []struct {
		name    string
		r       *RoverDomain
		args    args
		want    models.PointDto
		wantErr bool
	}{
		{
			name: "Move east",
			r:    &roverDomain,
			args: args{
				currentLocation: models.LocationDto{
					Point:     models.PointDto{XPoint: 1, YPoint: 1},
					Direction: models.DirectionEast,
				},
				moveType: models.MoveTypeForward,
			},
			want:    models.PointDto{XPoint: 2, YPoint: 1},
			wantErr: false,
		},
		{
			name: "Move north",
			r:    &roverDomain,
			args: args{
				currentLocation: models.LocationDto{
					Point:     models.PointDto{XPoint: 1, YPoint: 1},
					Direction: models.DirectionNorth,
				},
				moveType: models.MoveTypeForward,
			},
			want:    models.PointDto{XPoint: 1, YPoint: 2},
			wantErr: false,
		},
		{
			name: "Move south",
			r:    &roverDomain,
			args: args{
				currentLocation: models.LocationDto{
					Point:     models.PointDto{XPoint: 1, YPoint: 1},
					Direction: models.DirectionSouth,
				},
				moveType: models.MoveTypeForward,
			},
			want:    models.PointDto{XPoint: 1, YPoint: 0},
			wantErr: false,
		},
		{
			name: "Move west",
			r:    &roverDomain,
			args: args{
				currentLocation: models.LocationDto{
					Point:     models.PointDto{XPoint: 1, YPoint: 1},
					Direction: models.DirectionWest,
				},
				moveType: models.MoveTypeForward,
			},
			want:    models.PointDto{XPoint: 0, YPoint: 1},
			wantErr: false,
		},
		{
			name: "Move obstacle detected",
			r:    &roverDomain,
			args: args{
				currentLocation: models.LocationDto{
					Point:     models.PointDto{XPoint: 2, YPoint: 5},
					Direction: models.DirectionNorth,
				},
				moveType: models.MoveTypeForward,
			},
			want:    models.PointDto{XPoint: 2, YPoint: 5},
			wantErr: true,
		},
		{
			name: "Move point out grid",
			r:    &roverDomain,
			args: args{
				currentLocation: models.LocationDto{
					Point:     models.PointDto{XPoint: 10, YPoint: 10},
					Direction: models.DirectionNorth,
				},
				moveType: models.MoveTypeForward,
			},
			want:    models.PointDto{XPoint: 10, YPoint: 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.move(tt.args.currentLocation, tt.args.moveType)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoverDomain.move() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoverDomain.move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoverDomain_turnLeft(t *testing.T) {
	type args struct {
		currentDirection models.Direction
	}

	roverDomain := newRoverDomainMocked()

	tests := []struct {
		name string
		r    *RoverDomain
		args args
		want models.Direction
	}{
		{
			name: "North to west",
			r:    &roverDomain,
			args: args{
				currentDirection: models.DirectionNorth,
			},
			want: models.DirectionWest,
		},
		{
			name: "West to south",
			r:    &roverDomain,
			args: args{
				currentDirection: models.DirectionWest,
			},
			want: models.DirectionSouth,
		},
		{
			name: "South to east",
			r:    &roverDomain,
			args: args{
				currentDirection: models.DirectionSouth,
			},
			want: models.DirectionEast,
		},
		{
			name: "East to north",
			r:    &roverDomain,
			args: args{
				currentDirection: models.DirectionEast,
			},
			want: models.DirectionNorth,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.turnLeft(tt.args.currentDirection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoverDomain.turnLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoverDomain_turnRight(t *testing.T) {
	type args struct {
		currentDirection models.Direction
	}

	roverDomain := newRoverDomainMocked()

	tests := []struct {
		name string
		r    *RoverDomain
		args args
		want models.Direction
	}{
		{
			name: "North to east",
			r:    &roverDomain,
			args: args{
				currentDirection: models.DirectionNorth,
			},
			want: models.DirectionEast,
		},
		{
			name: "East to south",
			r:    &roverDomain,
			args: args{
				currentDirection: models.DirectionEast,
			},
			want: models.DirectionSouth,
		},
		{
			name: "South to west",
			r:    &roverDomain,
			args: args{
				currentDirection: models.DirectionSouth,
			},
			want: models.DirectionWest,
		},
		{
			name: "West to north",
			r:    &roverDomain,
			args: args{
				currentDirection: models.DirectionWest,
			},
			want: models.DirectionNorth,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.turnRight(tt.args.currentDirection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoverDomain.turnRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newRoverDomainMocked() RoverDomain {
	startingLocation := models.LocationDto{
		Point: models.PointDto{
			XPoint: 1,
			YPoint: 1,
		},
		Direction: models.DirectionNorth,
	}

	gridDomain := GridDomain{
		grid: models.GridDto{
			XPointMax: 10,
			YPointMax: 10,
		},
	}

	obstacleDomain := ObstacleDomain{
		obstacles: []models.ObstacleDto{
			{Point: models.PointDto{XPoint: 2, YPoint: 6}},
			{Point: models.PointDto{XPoint: 8, YPoint: 5}},
		},
	}

	return RoverDomain{
		location:       startingLocation,
		gridDomain:     &gridDomain,
		obstacleDomain: &obstacleDomain,
	}
}
