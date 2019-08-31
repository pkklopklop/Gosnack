package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var TennisScore map[int64]string = map[int64]string{
	0: "Love",
	1: "15",
	2: "30",
	3: "40",
	4: "Game",
}

type Player struct {
	Name  string
	Score int64
}

type Game struct {
	PlayerA Player
	PlayerB Player
}

//PlayerAGetPoint get 1 point.
func (g *Game) PlayerAGetPoint() {
	g.PlayerA.Score++
}

//PlayerBGetPoint get 1 point.
func (g *Game) PlayerBGetPoint() {
	g.PlayerB.Score++
}

//CurrentScore print Tennis score.
func (g *Game) CurrentScore() {
	switch {

	//One player wins.
	case (g.PlayerA.Score >= 4 || g.PlayerB.Score >= 4) && (math.Abs(float64(g.PlayerA.Score-g.PlayerB.Score)) >= 2):
		if g.PlayerA.Score > g.PlayerB.Score {
			fmt.Printf("[%s vs %s] Score: %s wins [%01d - %01d].\n", g.PlayerA.Name, g.PlayerB.Name, g.PlayerA.Name, g.PlayerA.Score, g.PlayerB.Score)
		} else {
			fmt.Printf("[%s vs %s] Score: %s wins [%01d - %01d].\n", g.PlayerA.Name, g.PlayerB.Name, g.PlayerB.Name, g.PlayerA.Score, g.PlayerB.Score)
		}

	//One player got advantage.
	case (g.PlayerA.Score >= 4 || g.PlayerB.Score >= 4) && (math.Abs(float64(g.PlayerA.Score-g.PlayerB.Score)) == 1):
		if g.PlayerA.Score > g.PlayerB.Score {
			fmt.Printf("[%s vs %s] Score: %s got advantage [%01d - %01d].\n", g.PlayerA.Name, g.PlayerB.Name, g.PlayerA.Name, g.PlayerA.Score, g.PlayerB.Score)
		} else {
			fmt.Printf("[%s vs %s] Score: %s got advantage [%01d - %01d].\n", g.PlayerA.Name, g.PlayerB.Name, g.PlayerB.Name, g.PlayerA.Score, g.PlayerB.Score)
		}

	//Deuce
	case (g.PlayerA.Score == g.PlayerB.Score) && g.PlayerA.Score >= 3:
		fmt.Printf("[%s vs %s] Score: Deuce [%01d - %01d].\n", g.PlayerA.Name, g.PlayerB.Name, g.PlayerA.Score, g.PlayerB.Score)

	//Same score but below 3 points.
	case (g.PlayerA.Score == g.PlayerB.Score) && g.PlayerA.Score != 0:
		fmt.Printf("[%s vs %s] Score: %s - All.\n", g.PlayerA.Name, g.PlayerB.Name, TennisScore[g.PlayerA.Score])

	//Normal score announcement.
	default:
		fmt.Printf("[%s vs %s] Score: %s - %s.\n", g.PlayerA.Name, g.PlayerB.Name, TennisScore[g.PlayerA.Score], TennisScore[g.PlayerB.Score])
	}
}

func main() {
	//Initial game.
	g := Game{
		PlayerA: Player{
			Name:  "โรเจอร์ เฟเดอเรอร์",
			Score: 0,
		},
		PlayerB: Player{
			Name:  "ราฟาเอล นาดัล",
			Score: 0,
		},
	}

	var random int
	//Seeding the random number by current time.
	rand.Seed(time.Now().UTC().UnixNano())

	//Random loop until one player wins.
	for !((g.PlayerA.Score >= 4 || g.PlayerB.Score >= 4) && (math.Abs(float64(g.PlayerA.Score-g.PlayerB.Score)) >= 2)) {
		g.CurrentScore()
		random = rand.Int()
		if random%2 == 0 {
			g.PlayerAGetPoint()
		} else {
			g.PlayerBGetPoint()
		}
	}

	g.CurrentScore()

}
