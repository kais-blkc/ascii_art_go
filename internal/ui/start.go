package ui

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/kais-blkc/ascii_art/internal/convert"
	imageutils "github.com/kais-blkc/ascii_art/internal/image_utils"
	"github.com/kais-blkc/ascii_art/internal/shared/constants"
	"github.com/kais-blkc/ascii_art/internal/shared/event"
	"github.com/kais-blkc/ascii_art/internal/shared/eventbus"
	"github.com/kais-blkc/ascii_art/internal/shared/helpers"
	"github.com/rivo/tview"
)

func StartUI() {
	// Colors
	// Blue
	// colorAccent, _ := helpers.HexToCell("#4C80A5")
	// colorBG, _ := helpers.HexToCell("#2B3438")
	// Green
	// colorAccent, _ := helpers.HexToCell("#4CA582")
	// colorBG, _ := helpers.HexToCell("#2B382D")
	// New Blue
	colorAccent, _ := helpers.HexToCell("#4C76A5")
	colorBG, _ := helpers.HexToCell("#262D33")
	colorText, _ := helpers.HexToCell("#fafafa")

	// TView Colors
	tview.Styles.PrimaryTextColor = tcell.ColorFloralWhite
	tview.Styles.ContrastBackgroundColor = colorAccent
	tview.Styles.PrimitiveBackgroundColor = colorBG
	tview.Styles.TitleColor = colorText
	tview.Styles.BorderColor = colorText
	tview.Styles.InverseTextColor = colorAccent

	// Create Application
	app := tview.NewApplication()
	pages := tview.NewPages()
	main := NewMainScreen(app, pages)
	fileList := NewFileListScreen(app, pages)

	pages.AddPage(constants.PageMain, main.GetPrimitive(), true, true)
	pages.AddPage(constants.PageFileList, fileList.GetPrimitive(), true, false)

	eventbus.UIListener.Subscribe(constants.EventConvertToAscii, func(ed event.EventData) {
		outputFilename := ed[constants.KeyEventDataOutputFileName].(string)
		filename := ed[constants.KeyEventDataFileName].(string)
		width := ed[constants.KeyEventDataWidth].(string)
		asciiRamp := ed[constants.KeyEventDataAsciiRamp].(string)

		err := helpers.ValidateFormFields(filename, outputFilename, width)
		if err != nil {
			ShowErrorModal(pages, err.Error(), constants.PageMain)
			return
		}

		img := convert.LoadImage("./" + filename)
		if img == nil {
			ShowErrorModal(pages, "Error loading image!", constants.PageMain)
			return
		}

		widthUint, err := strconv.ParseUint(width, 10, 64)
		if err != nil {
			ShowErrorModal(pages, "Error parsing width!", constants.PageMain)
			return
		}

		img = convert.ResizeImage(img, uint(widthUint), false)
		err = imageutils.AsciiToImageRGB(img, outputFilename, asciiRamp)
		if err != nil {
			ShowErrorModal(pages, err.Error(), constants.PageMain)
			log.Fatalf("Failed to save image: %v", err)
		}

		app.QueueUpdateDraw(func() {
			ShowSuccessModal(pages, "Img saved!", constants.PageMain)
		})
	})

	err := app.SetRoot(pages, true).EnableMouse(true).Run()
	if err != nil {
		panic(err)
	}

}
