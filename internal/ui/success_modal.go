package ui

import "github.com/rivo/tview"

func ShowSuccessModal(pages *tview.Pages, message string, returnToPage string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(
			func(buttonIndex int, buttonLabel string) {
				pages.SwitchToPage(returnToPage)
			},
		)

	pages.AddAndSwitchToPage("successModal", modal, true)
}
