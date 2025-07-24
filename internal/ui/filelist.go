package ui

import (
	"os"
	"path/filepath"
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/kais-blkc/ascii_art/internal/shared/constants"
	"github.com/kais-blkc/ascii_art/internal/shared/event"
	"github.com/kais-blkc/ascii_art/internal/shared/eventbus"
	"github.com/rivo/tview"
)

type FileListScreen struct {
	App              *tview.Application
	Pages            *tview.Pages
	Layout           *tview.Flex
	fileList         *tview.List
	acceptedExt      []string
	currentItemIndex int
}

func NewFileListScreen(app *tview.Application, pages *tview.Pages) *FileListScreen {
	screen := &FileListScreen{
		App:   app,
		Pages: pages,
	}

	screen.initUI()
	return screen
}

func (s *FileListScreen) GetPrimitive() tview.Primitive {
	return s.Layout
}

func (s *FileListScreen) initUI() {
	s.acceptedExt = []string{".jpg", ".png", ".jpeg"}

	s.fileList = tview.NewList()
	s.fillFileList()

	s.fileList.AddItem("❌ Close list", "", 0, func() {
		closeFileList(s.Pages)
	})

	s.fileList.SetBorder(true).SetTitle("Select image")
	s.initFileSelectHandler()
	s.initCenteredLayout()
	s.initKeyBindings()
}

func (s *FileListScreen) initCenteredLayout() {
	s.Layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false). // spacer top
		AddItem(
			tview.NewFlex().
				AddItem(nil, 0, 1, false).       // spacer left
				AddItem(s.fileList, 0, 1, true). // list in center
				AddItem(nil, 0, 1, false),       // spacer right
						0, 2, true).
		AddItem(nil, 0, 1, false) // spacer bottom
}

func (s *FileListScreen) fillFileList() {
	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())

		if file.IsDir() {
			continue
		}
		if !slices.Contains(s.acceptedExt, ext) {
			continue
		}

		s.fileList.AddItem(file.Name(), "", 0, nil)
	}
}

func (s *FileListScreen) initFileSelectHandler() {
	s.fileList.SetSelectedFunc(
		func(index int, name string, _ string, _ rune) {
			eventbus.UIListener.Emit(
				constants.EventFileSelected,
				event.EventData{
					constants.KeyEventDataFileName: name,
				},
			)
		},
	)
}

func (s *FileListScreen) initKeyBindings() {

	s.Layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlJ:
			index := (s.fileList.GetCurrentItem() + 1) % s.fileList.GetItemCount()
			s.fileList.SetCurrentItem(index)
			return nil
		case tcell.KeyCtrlK:
			index := (s.fileList.GetCurrentItem() - 1 + s.fileList.GetItemCount()) % s.fileList.GetItemCount()
			s.fileList.SetCurrentItem(index)
			return nil
		}

		return event
	})
}

func closeFileList(pages *tview.Pages) {
	pages.HidePage(constants.PageFileList)
	pages.ShowPage(constants.PageMain)
}
