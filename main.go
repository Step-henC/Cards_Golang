package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()

	w := a.NewWindow("Deck of Cards")

	w.Resize(fyne.NewSize(500, 500))

	cardsData := newDeck()

	tableData := cardsData.getIndices()

	cardTable := widget.NewTable(
		func() (int, int) {
			return len(cardsData), 2

		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Deck of Cards")
		},

		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tableData[i.Col][i.Row])
		})

	shuff := widget.NewButton("shuffle", func() {
		cardsData.shuffleDeck()
		tableData = cardsData.getIndices()

		cardTable.Refresh()

	})

	deal_button := widget.NewButton("deal", func() {

		dealt_hand, newCards := deal(cardsData, 7)
		cardsData = newCards

		tableData = newCards.getIndices()

		fileName := "dealt_hand"

		dealt_hand.saveToFile(fileName)

		cardTable.Refresh()

	})

	shuff.Resize(fyne.NewSize(100, 100))

	content := container.NewAdaptiveGrid(3, cardTable, deal_button, shuff)

	w.SetContent(content)
	w.ShowAndRun()

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("my_cards")

	// otherCards := newDeckFromFile("./my_cards")

	// otherCards.shuffleDeck()

	// otherCards.print()
}
