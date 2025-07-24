package ui

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/kais-blkc/ascii_art/internal/shared/constants"
	"github.com/kais-blkc/ascii_art/internal/shared/event"
	"github.com/kais-blkc/ascii_art/internal/shared/eventbus"
	"github.com/kais-blkc/ascii_art/internal/shared/helpers"
	"github.com/rivo/tview"
)

type MainScreen struct {
	App               *tview.Application
	Pages             *tview.Pages
	Layout            *tview.Flex
	buttonSelectFile  *tview.Button
	buttonConvert     *tview.Button
	buttonExit        *tview.Button
	buttonPreview     *tview.Button
	form              *tview.Form
	fieldOutputName   *tview.InputField
	fieldImgWidth     *tview.InputField
	fieldAsciiRamp    *tview.InputField
	filename          string
	focusables        []tview.Primitive
	currentFocusIndex int
}

func NewMainScreen(app *tview.Application, pages *tview.Pages) *MainScreen {
	screen := &MainScreen{
		App:   app,
		Pages: pages,
	}

	screen.initUi()
	return screen
}

func (s *MainScreen) GetPrimitive() tview.Primitive {
	return s.Layout
}

func (s *MainScreen) initUi() {
	// Colors
	colorFormFieldBG, _ := helpers.HexToCell("#777777")

	// FORM
	s.createFormOptions(colorFormFieldBG)

	// Button Select File
	s.buttonSelectFile = tview.NewButton("📁 Select file (ctrl+f)").
		SetSelectedFunc(func() {
			s.Pages.HidePage(constants.PageMain)
			s.Pages.ShowPage(constants.PageFileList)
		})

	// Button Convert
	s.buttonConvert = tview.NewButton("🚀 Convert (Ctrl+d)").
		SetSelectedFunc(s.emitListenerButtonConvert)

	// Button Preview
	s.buttonPreview = tview.NewButton("👀 Preview (Ctrl+p)").
		SetSelectedFunc(s.buttonPreviewHandler)

	// Button Exit
	s.buttonExit = tview.NewButton("❌ Exit (Ctrl+c)").
		SetSelectedFunc(func() {
			s.App.Stop()
		})

	// Flex
	optionsFlex := tview.NewFlex().
		SetDirection(tview.FlexRow)

	optionsFlex.
		AddItem(s.buttonSelectFile, 1, 0, false).
		AddItem(s.form, 0, 1, false).
		AddItem(s.buttonConvert, 1, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(s.buttonPreview, 1, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(s.buttonExit, 1, 0, false)

	optionsFlex.
		SetTitle("Ascii Art").
		SetBorder(true)

	_, _, width, _ := optionsFlex.GetInnerRect()
	optionsFlex.SetBorderPadding(0, 0, width/2, width/2)

	// index management
	s.focusables = []tview.Primitive{optionsFlex, s.form, s.fieldOutputName, s.buttonSelectFile, s.buttonConvert, s.buttonPreview, s.buttonExit}
	s.currentFocusIndex = 0

	// Layout
	s.Layout = tview.NewFlex()
	s.Layout.
		AddItem(optionsFlex, 0, 1, true)

	s.initKeyBindings()
	s.initListeners()
}

func (s *MainScreen) initListeners() {
	eventbus.UIListener.Subscribe(constants.EventFileSelected, func(e event.EventData) {
		filename, ok := e[constants.KeyEventDataFileName].(string)

		if !ok {
			ShowErrorModal(s.Pages, "File name not found or invalid type", constants.PageMain)
			return
		}

		outputFilename := helpers.ProcessFilename(filename)
		s.fieldOutputName.SetText(outputFilename)
		s.filename = filename

		s.Pages.HidePage(constants.PageFileList)
		s.Pages.ShowPage(constants.PageMain)
		s.App.SetFocus(s.form)
	})
}

func (s *MainScreen) initKeyBindings() {
	s.Layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlF:
			s.App.SetFocus(s.buttonSelectFile)
			return nil
		case tcell.KeyCtrlO:
			s.App.SetFocus(s.form)
			return nil
		case tcell.KeyCtrlD:
			s.App.SetFocus(s.buttonConvert)
			return nil
		case tcell.KeyCtrlP:
			s.App.SetFocus(s.buttonPreview)
			return nil
		case tcell.KeyCtrlJ:
			s.currentFocusIndex = (s.currentFocusIndex + 1) % len(s.focusables)
			s.App.SetFocus(s.focusables[s.currentFocusIndex])
			return nil
		case tcell.KeyCtrlK:
			s.currentFocusIndex = (s.currentFocusIndex - 1 + len(s.focusables)) % len(s.focusables)
			s.App.SetFocus(s.focusables[s.currentFocusIndex])
			return nil
		}

		return event
	})
}

func (s *MainScreen) createFormOptions(colorFormFieldBG tcell.Color) {
	s.fieldOutputName = tview.NewInputField().
		SetLabel("Output name").
		SetText("")

	s.fieldImgWidth = tview.NewInputField().
		SetLabel("Width").
		SetText("1080")

	s.fieldAsciiRamp = tview.NewInputField().
		SetLabel("Ramp (dark to light)").
		SetText(constants.AsciiRampDefault)

	s.form = tview.NewForm().
		AddTextView("Params (Ctrl+o)", "", 0, 1, false, false).
		SetFieldBackgroundColor(colorFormFieldBG).
		SetFieldTextColor(tcell.ColorWhiteSmoke).
		SetLabelColor(tcell.ColorWhiteSmoke)

	s.form.
		AddFormItem(s.fieldOutputName).
		AddFormItem(s.fieldImgWidth).
		AddFormItem(s.fieldAsciiRamp)
}

func (s *MainScreen) emitListenerButtonConvert() {
	eventbus.UIListener.Emit(constants.EventConvertToAscii, event.EventData{
		constants.KeyEventDataFileName:       s.filename,
		constants.KeyEventDataOutputFileName: s.fieldOutputName.GetText(),
		constants.KeyEventDataWidth:          s.fieldImgWidth.GetText(),
		constants.KeyEventDataAsciiRamp:      s.fieldAsciiRamp.GetText(),
	})
}

func (s *MainScreen) buttonPreviewHandler() {
	path := s.fieldOutputName.GetText()

	_, err := os.Stat(path)
	if err != nil {
		ShowErrorModal(s.Pages, "File not found: "+path, constants.PageMain)
		return
	}

	helpers.OpenImage(path)
}
