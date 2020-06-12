package main

import "fmt"

func main() {

	cards := []string{"First Card", getCardName()}
	cards = append(cards, "Third Card")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func getCardName() string {
	return "New Card Name"
}
