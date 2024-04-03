package main

import (
	//"Groupie-tracker/affichage_cover"
	"Main/main.go/function"
	"image/color"

	//"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	fynex "fyne.io/x/fyne/widget"

	//"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Definition of the theme
type myTheme struct{}

// Color function for custom theme
func (myTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x3b, G: 0x11, B: 0x77, A: 0x90}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0x3b, G: 0x11, B: 0x77, A: 0x90}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0xce, G: 0x9a, B: 0x24, A: 0xfa}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0xa5, G: 0x72, B: 0x34, A: 0xbf}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0xce, G: 0x91, B: 0x24, A: 0xfa}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0xce, G: 0x9a, B: 0x24, A: 0xfa}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xe1, G: 0xba, B: 0x66, A: 0xf9}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xf}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x19}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x19}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0xce, G: 0x9a, B: 0x24, A: 0xfa}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x99}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x66}
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

// Font function for custom theme
func (myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return theme.DefaultTheme().Font(s)
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return theme.DefaultTheme().Font(s)
}

// Icon function for custom theme
func (myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

// Size function for custom theme
func (myTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 2
	default:
		return theme.DefaultTheme().Size(s)
	}
}

func main() {
	// Defining windows:
	// Menu window:
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{})

	icon, _ := fyne.LoadResourceFromPath("icon.png")
	menu := myApp.NewWindow("Menu")
	menu.Resize(fyne.NewSize(400, 500))
	menu.CenterOnScreen()
	menu.SetIcon(icon)

	// Cover window:
	cover := myApp.NewWindow("cover")
	cover.Resize(fyne.NewSize(700, 500))
	cover.CenterOnScreen()
	cover.SetIcon(icon)

	// Creating Total variables that contain the lists of locations/dates:
	var Total_tab [][]string
	var Total_tab_vide [][]string

	Total_tab = append(Total_tab, []string{"Locations :", "Dates :"})

	function.GetRelations()

	Id_nbr := 1

	// Creating the search bar
	input := fynex.NewCompletionEntry([]string{})
	// Creating the function for auto implementation
	input.OnChanged = func(s string) {
		if len(s) == 0 {
			input.HideCompletion()
			return
		}
		var tab []string
		entree := make([]string, len(input.Text))
		for i, r := range input.Text {
			entree[i] = string(r)
		}
		tab = function.Complete(*input, entree, tab)

		input.SetOptions(tab)
		input.ShowCompletion()
	}

	// Creating the input for the search bar:
	input.SetPlaceHolder("Enter name...")
	Nom_entree := ""

	input.Resize(fyne.NewSize(400, 80))
	input.Move(fyne.NewPos(0, 0))

	// Creating the list/scroll of names:
	function.GetArtists()

	var Nom_tab []string

	// Creating the list of names:
	for i := 0; i < 52; i++ {
		Nom_tab = append(Nom_tab, function.Artists[i].Name)
	}

	noms := widget.NewList(
		func() int {
			return 52
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("noms")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(Nom_tab[lii])
		})

	noms.Resize(fyne.NewSize(400, 420))
	noms.Move(fyne.NewPos(0, 80))

	// Making names redirectable to the cover:
	noms.OnSelected = func(id widget.ListItemID) {
		Id_nbr = id

		// Creating the list of locations/dates (according to the id):
		Total_tab = Total_tab_vide

		Total_tab = append(Total_tab, []string{"Locations :", "Dates :"})

		var Locat_tab []string

		function.GetDates()

		function.GetLocations()
		for x := 0; x < len(function.Location.Index[Id_nbr].Location); x++ {
			Locat_tab = append(Locat_tab, function.Location.Index[Id_nbr].Location[x])
		}

		function.GetRelations()
		for y := 0; y < len(Locat_tab); y++ {
			Lt := Locat_tab[y]
			if len(function.Relations.Index[Id_nbr].Dates_Locations[Lt]) > 0 {
				for f := 0; f < len(function.Relations.Index[Id_nbr].Dates_Locations[Lt]); f++ {
					Dt := function.Relations.Index[Id_nbr].Dates_Locations[Lt][f]
					Total_tab = append(Total_tab, []string{Lt, Dt})
				}
			}
		}

		// Populating our cover window with various information:
		function.Affichage_cover(cover, menu, Id_nbr)

		// Hide the menu window and show the cover window:
		menu.Hide()
		cover.Show()
		log.Println(Id_nbr)
	}

	// Creating the search button for the search bar:
	Navbar := container.NewVBox(input, widget.NewButton("Search", func() {
		log.Println("Content was:", input.Text)
		Nom_valide := true

		// Verifying and getting the id corresponding to the entered name:
		Total_tab = Total_tab_vide
		Nom_entree = input.Text
		if strings.Contains(Nom_entree, "-Member") {
			Nom_entree = strings.TrimSuffix(Nom_entree, "-Member")
		} else if strings.Contains(Nom_entree, "-Artists") {
			Nom_entree = strings.TrimSuffix(Nom_entree, "-Artists")
		}

		for t := 0; t < len(function.Artists); t++ {
			if function.Artists[t].Name != Nom_entree {
				for x := 0; x <= len(function.Artists[t].Members)-1; x++ {
					if function.Artists[t].Members[x] == Nom_entree {
						Id_nbr = t
						Nom_valide = true

					}
				}
			} else if function.Artists[t].Name == Nom_entree {
				Id_nbr = t
				Nom_valide = true
				t = len(function.Artists)
			} else if t == len(function.Artists) {
				Nom_valide = false
			}
		}

		// Creating the list of locations/dates (according to the id):
		Total_tab = append(Total_tab, []string{"Locations :", "Dates :"})

		//var Dates_tab []string
		var Locat_tab []string

		function.GetDates()

		function.GetLocations()
		for x := 0; x < len(function.Location.Index[Id_nbr].Location); x++ {
			Locat_tab = append(Locat_tab, function.Location.Index[Id_nbr].Location[x])

		}

		function.GetRelations()
		for y := 0; y < len(Locat_tab); y++ {
			Lt := Locat_tab[y]
			if len(function.Relations.Index[Id_nbr].Dates_Locations[Lt]) > 0 {
				for f := 0; f < len(function.Relations.Index[Id_nbr].Dates_Locations[Lt]); f++ {
					Dt := function.Relations.Index[Id_nbr].Dates_Locations[Lt][f]
					Total_tab = append(Total_tab, []string{Lt, Dt})
				}
			}
		}

		// Populating our cover window with various information:
		function.Affichage_cover(cover, menu, Id_nbr)

		// Verifying that the name is correct (it is in the list of artists):
		if Nom_valide {
			// Hide the menu window and show the cover window:
			menu.Hide()
			cover.Show()
			log.Println(Id_nbr)
		}
		input.Text = ""
		Id_nbr = 0
	}))

	Navbar.Resize(fyne.NewSize(400, 80))
	Navbar.Move(fyne.NewPos(0, 0))

	// Defining what the menu contains:
	cont_menu := container.NewWithoutLayout(Navbar, noms)

	// Running the application by opening the menu window:
	menu.SetContent(cont_menu)
	menu.ShowAndRun()
}
