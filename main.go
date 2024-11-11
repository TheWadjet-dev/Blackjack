package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

type Card struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
}

type Game struct {
    Deck        []Card `json:"-"`
    PlayerHand  []Card `json:"playerHand"`
    DealerHand  []Card `json:"dealerHand"`
    PlayerScore int    `json:"playerScore"`
    DealerScore int    `json:"dealerScore"`
    Message     string `json:"message"`
}

var game Game

func main() {
    rand.Seed(time.Now().UnixNano())
    http.HandleFunc("/", serveHTML)
    http.HandleFunc("/newgame", newGameHandler)
    http.HandleFunc("/hit", hitHandler)
    http.HandleFunc("/stand", standHandler)
    fmt.Println("Server started at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func newGameHandler(w http.ResponseWriter, r *http.Request) {
    game = Game{}
    game.createDeck()
    game.PlayerHand = []Card{game.drawCard(), game.drawCard()}
    game.DealerHand = []Card{game.drawCard(), game.drawCard()}
    game.updateScores()
    game.Message = "Good luck!"
    json.NewEncoder(w).Encode(game)
}

func hitHandler(w http.ResponseWriter, r *http.Request) {
    game.PlayerHand = append(game.PlayerHand, game.drawCard())
    game.updateScores()
    if game.PlayerScore > 21 {
        game.Message = "You busted! Dealer wins."
    }
    json.NewEncoder(w).Encode(game)
}

func standHandler(w http.ResponseWriter, r *http.Request) {
    for game.DealerScore < 17 {
        game.DealerHand = append(game.DealerHand, game.drawCard())
        game.updateScores()
    }
    if game.DealerScore > 21 || game.PlayerScore > game.DealerScore {
        game.Message = "You win!"
    } else if game.PlayerScore < game.DealerScore {
        game.Message = "Dealer wins."
    } else {
        game.Message = "It's a tie!"
    }
    json.NewEncoder(w).Encode(game)
}

func (g *Game) createDeck() {
    names := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
    values := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10, "A": 11}
    g.Deck = []Card{}
    for _, name := range names {
        for i := 0; i < 4; i++ {
            g.Deck = append(g.Deck, Card{Name: name, Value: values[name]})
        }
    }
    rand.Shuffle(len(g.Deck), func(i, j int) { g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i] })
}

func (g *Game) drawCard() Card {
    card := g.Deck[len(g.Deck)-1]
    g.Deck = g.Deck[:len(g.Deck)-1]
    return card
}

func (g *Game) updateScores() {
    g.PlayerScore = calculateScore(g.PlayerHand)
    g.DealerScore = calculateScore(g.DealerHand)
}

func calculateScore(hand []Card) int {
    score := 0
    aces := 0
    for _, card := range hand {
        score += card.Value
        if card.Name == "A" {
            aces++
        }
    }
    for score > 21 && aces > 0 {
        score -= 10
        aces--
    }
    return score
}
