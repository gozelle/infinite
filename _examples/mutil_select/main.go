package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	inf "github.com/gozelle/infinite"
	"github.com/gozelle/infinite/color"
	"github.com/gozelle/infinite/components"
	"github.com/gozelle/infinite/components/selection/multiselect"
	"github.com/gozelle/infinite/style"
)

func main() {
	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.LightBlue)
	
	keymap := components.DefaultMultiKeyMap()
	keymap.Choice = key.NewBinding(
		key.WithKeys(tea.KeySpace.String()),
	)
	_, _ = inf.NewMultiSelect([]string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
	},
		multiselect.WithKeyMap(keymap),
		multiselect.WithHintSymbol("x"),
		multiselect.WithUnHintSymbol("√"),
		//multiselect.WithDisableOutputResult(),
		//multiselect.WithCursorSymbol(emoji.PointRight),
		//multiselect.WithDisableFilter(),
		//multiselect.WithFilterInput(input),
		multiselect.WithDisableFilter(),
	).
		Display("select your items!")
	
	//_, _ = inf.
	//	NewMultiSelect([]string{"f1", "f2", "f3"}).
	//	Display()
}
