package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/mars-rover-go/domains"
	"github.com/mars-rover-go/models"
	"github.com/mars-rover-go/utils"
)

var startXPoint int
var startYPoint int
var startDirection string

func init() {
	flag.IntVar(&startXPoint, "sx", 0, "Starting X point")
	flag.IntVar(&startYPoint, "sy", 0, "Starting Y point")
	flag.StringVar(&startDirection, "d", string(models.DirectionNorth), "Starting direction")
}

func main() {
	flag.Parse()

	startingLocation := models.LocationDto{
		Point:     models.PointDto{XPoint: startXPoint, YPoint: startYPoint},
		Direction: models.Direction(strings.ToUpper(startDirection)),
	}

	config, err := getConfiguration()
	if err != nil {
		fmt.Printf("ERROR - %+v\n", err)
		return
	}

	roverDomain, err := initComponents(startingLocation, config)
	if err != nil {
		fmt.Printf("ERROR - %+v\n", err)
		return
	}

	startExecution(*config, startingLocation, roverDomain)
}

func startExecution(config models.ConfigurationDto, startingLocation models.LocationDto, roverDomain domains.IRoverDomain) {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("--------------------------")
	fmt.Println("------- MARS ROVER -------")
	fmt.Println("--------------------------")

	fmt.Println("- Commands available:\n\t. f: forward\n\t. b: backward\n\t. l: left\n\t. r: right")
	fmt.Print("- Press ESC to quit\n\n")

	fmt.Println("Grid")
	fmt.Printf("\tXPointMax: %d, YPointMax: %d\n\n", config.Grid.XPointMax, config.Grid.XPointMax)

	fmt.Println("Obstacles")
	if len(config.Obstacle) == 0 {
		fmt.Println("\tNo obstacles")
	}
	for _, o := range config.Obstacle {
		fmt.Printf("\tXPoint: %d, YPoint: %d \n", o.Point.XPoint, o.Point.YPoint)
	}

	fmt.Printf("\nStart location: %s\n\n", utils.LocationToString(startingLocation))

	commands := []string{}
	fmt.Print("Write commands: ")

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		if event.Key == keyboard.KeyEsc {
			break
		}
		if event.Key == keyboard.KeyEnter {
			location, err := roverDomain.ExecuteCommands(commands)
			if err != nil {
				fmt.Printf("\n\tERROR - %+v\n", err)
			}

			commands = []string{}
			fmt.Printf("\n\tLocation: %s\n\n", utils.LocationToString(location))
			fmt.Print("Write commands: ")
			continue
		}

		switch string(event.Rune) {
		case "f", "b", "l", "r":
			commands = append(commands, string(event.Rune))
		default:
			continue
		}

		fmt.Printf("%s ", string(event.Rune))
	}
}

func getConfiguration() (*models.ConfigurationDto, error) {
	var config models.ConfigurationDto

	configFilePath, err := os.Open("./config.json")
	if err != nil {
		return nil, err
	}
	defer configFilePath.Close()
	configValue, _ := ioutil.ReadAll(configFilePath)

	err = json.Unmarshal(configValue, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func initComponents(startingPosition models.LocationDto, config *models.ConfigurationDto) (domains.IRoverDomain, error) {
	gridDomain := domains.NewGridDomain(config.Grid)
	obstacleDomain := domains.NewObstacleDomain(config.Obstacle)

	rover, err := domains.NewRoverDomain(startingPosition, gridDomain, obstacleDomain)

	return rover, err
}
