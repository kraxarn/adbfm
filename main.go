package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	// Create basic Qt application
	app := widgets.NewQApplication(len(os.Args), os.Args)
	// Create window
	window := widgets.NewQMainWindow(nil, 0)
	// Set minimum and initial size
	window.SetMinimumSize2(960, 540)
	window.Resize2(1280, 720)
	// Set a window title
	window.SetWindowTitle("adbfm")

	// Sidebar
	sidebar := widgets.NewQVBoxLayout()
	bookmarksLayout := widgets.NewQVBoxLayout()
	bookmarks := widgets.NewQGroupBox2("Bookmarks", nil)
	noBookmarks := widgets.NewQLabel2("No bookmarks", nil, 0)
	noBookmarks.SetEnabled(false)
	bookmarksLayout.AddWidget(noBookmarks, 0, 0)
	bookmarks.SetLayout(bookmarksLayout)
	sidebar.AddWidget(bookmarks, 0, 0)

	// Testing
	text1 := widgets.NewQLabel2("Bookmarks", nil, 0)
	text2 := widgets.NewQLabel2("/", nil, 0)

	// Main layout splitter
	splitter := widgets.NewQSplitter(nil)
	splitter.AddWidget(text1)
	splitter.AddWidget(text2)
	window.SetCentralWidget(splitter)

	// Show the main window
	window.Show()
	OpenSetup()
	// Main Qt event loop
	app.Exec()
}
