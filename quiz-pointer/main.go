package main

type Gamer struct {
	Name  string
	Games []string
}

func (g *Gamer) AddGame(game string) {
	g.Games = append(g.Games, game)
}

func main() {
	gamer := Gamer{
		Name: "Alice"}

	gamer.AddGame("Chess")
	gamer.AddGame("Checkers")
	gamer.AddGame("Poker")

	for _, game := range gamer.Games {
		println("Game:", game)
	}
}
