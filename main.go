package main

import "fmt"

func main() {
	//khởi tạo mảng cards
	cards := []string{"Ace of Diamonds", newCard()}
	//thêm phần tử vào mảng
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
