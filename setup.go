package main

import (
	"os"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func LayoutToWidget(layout widgets.QLayout_ITF) widgets.QWidget_ITF {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	return widget
}

func VerifyAdbPath(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	// Qt process needed because of threading issues
	process := core.NewQProcess(nil)
	process.Start(path, []string{ "--version" }, 1)
	process.WaitForFinished(-1)
	out := process.ReadAllStandardOutput().Data()

	return strings.HasPrefix(out, "Android Debug Bridge")
}

func OpenSetup() {
	// Dialog content
	dialog := widgets.NewQDialog(nil, 0)
	layout := widgets.NewQVBoxLayout()

	// ADB path container
	adbPath := widgets.NewQHBoxLayout()
	adbFilePath := widgets.NewQLineEdit(nil)
	adbPath.AddWidget(adbFilePath, 1, 0)
	adbBrowse := widgets.NewQPushButton2("Browse", nil)
	adbBrowse.ConnectReleased(func() {
		path := widgets.QFileDialog_GetOpenFileName(dialog, "Select ADB Executable",
			"/usr/bin", "", "", 0)
		if len(path) > 0 {
			if VerifyAdbPath(path) {
				adbFilePath.SetText(path)
			} else {
				widgets.QMessageBox_Warning(dialog, "Error", "Selected file is not a valid ADB executable",
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			}
		}
	})
	adbPath.AddWidget(adbBrowse, 0, 0)

	// Device container
	device := widgets.NewQHBoxLayout()
	deviceSelect := widgets.NewQComboBox(nil)
	deviceSelect.SetEnabled(false)
	device.AddWidget(deviceSelect, 1, 0)
	device.AddWidget(widgets.NewQPushButton2("Refresh", nil), 0, 0)

	layout.AddWidget(widgets.NewQLabel2("ADB path", nil, 0), 0, 0)
	layout.AddWidget(LayoutToWidget(adbPath), 0, 0)
	layout.AddWidget(widgets.NewQLabel2("Device", nil, 0), 0, 0)
	layout.AddWidget(LayoutToWidget(device), 0, 0)

	// Buttons
	buttons := widgets.NewQHBoxLayout2(nil)
	buttons.AddWidget(widgets.NewQPushButton2("OK", nil), 0, 2)
	buttons.AddWidget(widgets.NewQPushButton2("Cancel", nil), 0, 2)
	layout.AddWidget(LayoutToWidget(buttons), 0, 2)

	// Create dialog
	dialog.SetWindowTitle("Setup")
	dialog.SetLayout(layout)
	dialog.SetModal(true)
	dialog.SetFixedSize2(350, 250)

	// Show it
	dialog.Show()
}