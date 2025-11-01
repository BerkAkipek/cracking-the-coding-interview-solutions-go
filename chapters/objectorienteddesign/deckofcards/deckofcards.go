package deckofcards

import (
	"math/rand"
)

/*
Deck of Cards: Design the data structures for a generic deck of cards.
Explain how you would subclass the data structures to implement blackjack.
*/

type Card struct {
	suite int
	rank  int
}

type Deck struct {
	cards []*Card
}

func CreateDeck() *Deck {
	deck := Deck{
		cards: []*Card{},
	}
	for i := range 4 {
		for j := range 13 {
			deck.cards = append(deck.cards, &Card{
				suite: i + 1,
				rank:  j + 1,
			})
		}
	}
	return &deck
}

func (d *Deck) Shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) Reset() *Deck {
	newDeck := CreateDeck()
	d.cards = newDeck.cards
	return d
}

func (d *Deck) Draw() *Card {
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card
}

type Player struct {
	hand   []*Card
	score  int
	busted bool
}

func CreatePlayer() *Player {
	return &Player{
		hand:   []*Card{},
		busted: false,
	}
}

func (p *Player) CalculateHand() int {
	score := 0
	numAces := 0

	for i := range p.hand {
		if p.hand[i].rank == 1 {
			score += 11
			numAces++
		} else if p.hand[i].rank >= 10 {
			score += 10
		} else {
			score += p.hand[i].rank
		}
	}

	for score > 21 && numAces > 0 {
		score -= 10
		numAces--
	}

	return score
}

func Game(numPlayers int) int {

	// Setup
	d := CreateDeck()
	d.Shuffle()
	players := []*Player{}
	for i := range numPlayers {
		if i == 0 {
			dealer := CreatePlayer()
			players = append(players, dealer)
		}
		players = append(players, CreatePlayer())
	}
	for range 2 {
		for i := range len(players) {
			players[i].hand = append(players[i].hand, d.Draw())
		}
	}

	// Game phase
	for i := range len(players) {
		players[i].score = players[i].CalculateHand()
		for {
			if players[i].score >= 17 || players[i].busted {
				break
			}
			players[i].hand = append(players[i].hand, d.Draw())
			players[i].score = players[i].CalculateHand()
			if players[i].score > 21 {
				players[i].busted = true
			}
		}
	}

	player := players[1]
	dealer := players[0]
	if dealer.busted {
		if player.busted {
			return 0
		} else {
			return 1
		}
	}
	if player.busted {
		if !dealer.busted {
			return -1
		}
	}

	if !dealer.busted && !player.busted {
		if player.score > dealer.score {
			return 1
		} else if player.score == dealer.score {
			return 0
		} else {
			return -1
		}
	}

	return 0
}

func MonteCarlo(sampleSize int) (float64, float64, float64) {
	dealer := 0
	player := 0
	tie := 0
	for range sampleSize {
		result := Game(2)
		if result > 0 {
			player++
		}
		if result < 0 {
			dealer++
		}
		if result == 0 {
			tie++
		}
	}

	dealerRate := float64(dealer) / float64(sampleSize)
	playerRate := float64(player) / float64(sampleSize)
	tieRate := float64(tie) / float64(sampleSize)
	return playerRate, dealerRate, tieRate
}
