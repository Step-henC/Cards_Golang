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

		fileName := "dealt_hand"

		if len(cardsData) <= 7 {

			a.Quit()
		} else {

			dealt_hand, newCards := deal(cardsData, 7)
			cardsData = newCards

			tableData = newCards.getIndices()

			dealt_hand.saveToFile(fileName)

			cardTable.Refresh()

		}

	})

	shuff.Resize(fyne.NewSize(100, 50))
	shuff.Move(fyne.NewPos(400, 100))

	deal_button.Resize(fyne.NewSize(100, 50))
	deal_button.Move(fyne.NewPos(400, 180))

	cardTable.Resize(fyne.NewSize(300, 500))
	cardTable.Move(fyne.NewPos(20, 20))

	content := container.NewWithoutLayout(cardTable, deal_button, shuff)

	w.SetContent(content)
	w.ShowAndRun()

}
