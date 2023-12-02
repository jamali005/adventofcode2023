package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	TotalRed   = 12
	TotalGreen = 13
	TotalBlue  = 14
)

type game struct {
	red   int
	green int
	blue  int
}

func (g game) validate() bool {
	return g.red <= TotalRed && g.green <= TotalGreen && g.blue <= TotalBlue
}

func fromString(data string) (*game, error) {
	g := new(game)

	data = strings.TrimSpace(data)
	split := strings.Split(data, ",")

	for _, d := range split {
		d = strings.TrimSpace(d)
		spl := strings.Fields(d)

		val, err := strconv.Atoi(spl[0])
		if err != nil {
			return nil, err
		}

		colour := spl[1]
		switch colour {
		case "green":
			g.green += val
		case "red":
			g.red += val
		case "blue":
			g.blue += val
		default:
			return nil, errors.New("invalid input, color could not be identified")
		}
	}

	return g, nil
}

func part1() error {
	file, err := os.Open("inputadventofcodeday2.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		index, games, err := getGames(line)
		if err != nil {
			log.Fatal(err)
		}

		valid := true
		for _, game := range games {
			if !game.validate() {
				valid = false
			}
		}

		if valid {
			result += index
		}
	}

	fmt.Println(result)
	return nil
}

func getGames(input string) (int, []game, error) {
	split := strings.Split(input, ":")
	index, err := strconv.Atoi(strings.Fields(split[0])[1])
	if err != nil {
		log.Fatal(err)
	}

	data := split[1]
	gamesData := strings.Split(data, ";")

	var games []game
	for _, gameData := range gamesData {
		currGame, err := fromString(gameData)
		if err != nil {
			log.Fatal(err)
		}
		games = append(games, *currGame)
	}

	return index, games, nil
}

func main() {
	err := part1()
	if err != nil {
		log.Println("Error:", err)
	}
}
