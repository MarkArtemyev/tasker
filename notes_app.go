package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
)

var notes []string
var noteList *widget.List

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Приложение для заметок")
    noteEntry := widget.NewEntry()
    noteEntry.SetPlaceHolder("Введите заметку...")

    addNote := func() {
        if noteEntry.Text != "" {
            notes = append(notes, noteEntry.Text)
            noteEntry.SetText("")
            noteList.Refresh()
        }
    }

    addButton := widget.NewButton("Добавить", func() {
        addNote()
    })

    noteList = widget.NewList(
        func() int { return len(notes) },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.ListItemID, obj fyne.CanvasObject) { obj.(*widget.Label).SetText(notes[id]) },
    )

    noteList.OnSelected = func(id widget.ListItemID) {
        dialog.ShowConfirm("Удаление", "Удалить эту заметку?", func(confirm bool) {
            if confirm {
                notes = append(notes[:id], notes[id+1:]...)
                noteList.UnselectAll()
                noteList.Refresh()
            }
        }, myWindow)
    }

    noteEntry.OnSubmitted = func(text string) {
        addNote()
    }

    myWindow.SetContent(container.NewVBox(
        noteEntry,
        addButton,
        noteList,
    ))

    myWindow.ShowAndRun()
}
