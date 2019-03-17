package main

func main() {
	cards := newDeck()

	newDeckFromFile("my_deck")

	cards.display()

	//hand, remaingCards := dealFunction(cards, 4)
	//hand, remaingCards := cards.dealMethod(4)

	//fmt.Println("This is the dealt hand:")
	//hand.display()

	//fmt.Println("")

	//fmt.Println("These are the remaining cards:")
	//remaingCards.display()

}
