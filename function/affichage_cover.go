package function

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Affichage_cover(cover fyne.Window, menu fyne.Window, Id_nbr int) {
	// Create dropdown lists of locations containing dates (cover window)
	var Relation_lieu_tab []string

	// Create the exit button that returns from the cover window to the menu
	exit := widget.NewButton("Return", func() {
		cover.Hide()
		menu.Show()
	})

	exit.Resize(fyne.NewSize(450, 50))
	exit.Move(fyne.NewPos(250, 450))

	// Selection bar containing all locations
	for i := range Relations.Index[Id_nbr].Dates_Locations {
		Relation_lieu_tab = append(Relation_lieu_tab, i)
	}
	lieu := widget.NewSelect(Relation_lieu_tab, func(s string) {})

	lieu.Selected = "Concerts"
	lieu.Resize(fyne.NewSize(250, 50))
	lieu.Move(fyne.NewPos(0, 0))

	// List containing the dates
	date := widget.NewList(
		func() int {
			return len(Relations.Index[Id_nbr].Dates_Locations[lieu.Selected])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("dates")
		}, func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(Relations.Index[Id_nbr].Dates_Locations[lieu.Selected][lii])
		})

	date.Resize(fyne.NewSize(250, 200))
	date.Move(fyne.NewPos(0, 50))

	// Selection bar for various information
	infos := widget.NewSelect([]string{"Group member(s)", "Creation date", "First album"}, func(s string) {})

	infos.Selected = "Information"
	infos.Resize(fyne.NewSize(250, 50))
	infos.Move(fyne.NewPos(0, 250))

	// List containing the information
	var vide []string

	liste_infos := widget.NewList(
		func() int {
			return len(vide)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(y widget.ListItemID, o fyne.CanvasObject) {

		})

	liste_infos.Resize(fyne.NewSize(250, 200))
	liste_infos.Move(fyne.NewPos(0, 300))

	// Download the artist's image (cover window)
	r, _ := fyne.LoadResourceFromURLString(Artists[Id_nbr].Image)
	img := canvas.NewImageFromResource(r)
	img.Resize(fyne.NewSize(450, 450))
	img.Move(fyne.NewPos(250, 0))
	cover.SetIcon(img.Resource)
	cont_cover := container.NewWithoutLayout(lieu, date, infos, liste_infos, img, exit)

	// Update the lists of dates when changing the location
	lieu.OnChanged = func(s string) {
		date = widget.NewList(
			func() int {
				return len(Relations.Index[Id_nbr].Dates_Locations[s])
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("dates")
			}, func(lii widget.ListItemID, co fyne.CanvasObject) {
				co.(*widget.Label).SetText(Relations.Index[Id_nbr].Dates_Locations[s][lii])
			})
		date.Resize(fyne.NewSize(250, 200))
		date.Move(fyne.NewPos(0, 50))
		cover.SetIcon(img.Resource)
		cont_cover = container.NewWithoutLayout(lieu, date, infos, liste_infos, img, exit)
		cover.SetContent(cont_cover)
	}

	// Update the lists of information when changing the type of information
	infos.OnChanged = func(d string) {
		if d == "Group member(s)" {
			// Create the list of group member names
			str_nbr_membres := "There are " + string(len(Artists[Id_nbr].Members)+48)
			if len(Artists[Id_nbr].Members) == 1 {
				str_nbr_membres = str_nbr_membres + " member in the group:"
			} else if len(Artists[Id_nbr].Members) > 1 {
				str_nbr_membres = str_nbr_membres + " members in the group:"
			}
			infos_membre := []string{str_nbr_membres}
			for nm := 0; nm < len(Artists[Id_nbr].Members); nm++ {
				infos_membre = append(infos_membre, Artists[Id_nbr].Members[nm])
			}

			liste_infos = widget.NewList(
				func() int {
					return len(infos_membre)
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("template")
				},
				func(y widget.ListItemID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(infos_membre[y])
				})

			liste_infos.Resize(fyne.NewSize(250, 200))
			liste_infos.Move(fyne.NewPos(0, 300))
			cover.SetIcon(img.Resource)
			cont_cover = container.NewWithoutLayout(lieu, date, infos, liste_infos, img, exit)
			cover.SetContent(cont_cover)
		}
		if d == "Creation date" {
			var data = []string{fmt.Sprint(Artists[Id_nbr].CreationDate)}
			liste_infos = widget.NewList(
				func() int {
					return len(data)
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("template")
				},
				func(i widget.ListItemID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i])
				})

			liste_infos.Resize(fyne.NewSize(250, 200))
			liste_infos.Move(fyne.NewPos(0, 300))
			cover.SetIcon(img.Resource)
			cont_cover = container.NewWithoutLayout(lieu, date, infos, liste_infos, img, exit)
			cover.SetContent(cont_cover)
		}
		if d == "First album" {
			var data = []string{fmt.Sprint(Artists[Id_nbr].FirstAlbum)}
			liste_infos = widget.NewList(
				func() int {
					return len(data)
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("template")
				},
				func(i widget.ListItemID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i])
				})

			liste_infos.Resize(fyne.NewSize(250, 200))
			liste_infos.Move(fyne.NewPos(0, 300))
			cover.SetIcon(img.Resource)
			cont_cover = container.NewWithoutLayout(lieu, date, infos, liste_infos, img, exit)
			cover.SetContent(cont_cover)
		}
	}
	cover.SetContent(cont_cover)
}
