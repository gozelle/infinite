package components

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gozelle/infinite/style"
	"github.com/gozelle/infinite/theme"
)

type (
	/*
		Components, You can use these components directly:
			 	1. Input
			 	2. Selection
			 	3. Spinner
				4. Autocomplete
				5. Progress
		Or use them inline in your custom component,
		for how to embed them, you can refer to the implementation of `Confirm`.
	*/
	Components interface {
		tea.Model
		
		// SetProgram this method will be called back when the tea.Program starts.
		// please keep passing this method
		SetProgram(program *tea.Program)
	}
)

// NewAutocomplete constructor
func NewAutocomplete(suggester Suggester) *Autocomplete {
	return &Autocomplete{
		Suggester:            suggester,
		Completer:            DefaultCompleter(),
		Input:                NewInput(),
		KeyMap:               DefaultAutocompleteKeyMap(),
		ShowSelection:        true,
		ShouldNewSelection:   true,
		SelectionCreator:     DefaultSelectionCreator,
		SuggestionViewRender: NewLineSuggestionRender,
		//SuggestionViewRender: TabSuggestionRender,
	}
}

// NewInput constructor
func NewInput() *Input {
	c := &Input{
		Model:                    textinput.New(),
		Required:                 InputDefaultRequired,
		RequiredMsg:              InputDefaultRequiredMsg,
		RequiredMsgKeepAliveTime: InputDefaultRequiredMsgKeepTime,
		Status:                   InputDefaultStatus,
		Prompt:                   InputDefaultPrompt,
		DefaultValue:             InputDefaultValue,
		BlinkSpeed:               InputDefaultBlinkSpeed,
		EchoMode:                 InputDefaultEchoMode,
		EchoCharacter:            InputDefaultEchoCharacter,
		CharLimit:                InputDefaultCharLimit,
		KeyMap:                   InputDefaultKeyMap(),
		DefaultValueStyle:        InputDefaultPlaceholderStyle,
		PromptStyle:              InputDefaultPromptStyle,
		TextStyle:                InputDefaultTextStyle,
		CursorStyle:              InputDefaultCursorStyle,
		FocusSymbolStyle:         style.New(),
		UnFocusSymbolStyle:       style.New(),
		FocusIntervalStyle:       style.New(),
		UnFocusIntervalStyle:     style.New(),
		OutputResult:             true,
	}
	return c
}

// NewPrintHelper constructor
func NewPrintHelper(program *tea.Program) *PrintHelper {
	return &PrintHelper{program: program}
}

// NewProgress constructor
func NewProgress() *Progress {
	p := &Progress{
		Id:              nextID(),
		Total:           ProgressDefaultTotal,
		Current:         ProgressDefaultCurrent,
		PercentAgeFunc:  ProgressDefaultPercentAgeFunc,
		PercentAgeStyle: ProgressDefaultPercentAgeStyle,
		Width:           ProgressDefaultWidth,
		Full:            ProgressDefaultFull,
		FullColor:       ProgressDefaultFullColor,
		Empty:           ProgressDefaultEmpty,
		EmptyColor:      ProgressDefaultEmptyColor,
		ShowPercentage:  ProgressDefaultShowPercentage,
		ShowCost:        ProgressDefaultShowCost,
		prevAmount:      ProgressDefaultPrevAmount,
		CostView:        ProgressDefaultCostView,
		TickCostDelay:   ProgressDefaultTickCostDelay,
		Quit:            InterruptKey,
	}
	
	return p
}

func newDefaultPager() paginator.Model {
	model := paginator.New()
	p := paginator.New()
	p.Type = paginator.Dots
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.PerPage = SelectionDefaultPageSize
	return model
}

// NewSelection constructor
func NewSelection(choices []string) *Selection {
	items := slice.Map[string, SelectionItem](choices, func(idx int, item string) SelectionItem {
		return SelectionItem{idx, item}
	})
	
	c := &Selection{
		Paginator:            newDefaultPager(),
		ShowPaginator:        true,
		Choices:              items,
		Selected:             make(map[int]bool),
		CursorSymbol:         SelectionDefaultCursorSymbol,
		UnCursorSymbol:       SelectionDefaultUnCursorSymbol,
		CursorSymbolStyle:    SelectionDefaultCursorSymbolStyle,
		ChoiceTextStyle:      SelectionDefaultChoiceTextStyle,
		Prompt:               SelectionDefaultPrompt,
		PromptStyle:          SelectionDefaultPromptStyle,
		HintSymbol:           SelectionDefaultHintSymbol,
		HintSymbolStyle:      SelectionDefaultHintSymbolStyle,
		UnHintSymbol:         SelectionDefaultUnHintSymbol,
		UnHintSymbolStyle:    SelectionDefaultUnHintSymbolStyle,
		DisableOutPutResult:  SelectionDefaultDisableOutPutResult,
		Keymap:               DefaultMultiKeyMap(),
		Help:                 SelectionDefaultHelp,
		RowRender:            SelectionDefaultRowRender,
		EnableFilter:         SelectionDefaultEnableFilter,
		FilterInput:          SelectionDefaultFilterInput,
		FilterFunc:           SelectionDefaultFilterFunc,
		ShowHelp:             SelectionDefaultShowHelp,
		FocusSymbol:          theme.DefaultTheme.FocusSymbol,
		UnFocusSymbol:        theme.DefaultTheme.UnFocusSymbol,
		FocusInterval:        theme.DefaultTheme.FocusInterval,
		UnFocusInterval:      theme.DefaultTheme.UnFocusInterval,
		FocusSymbolStyle:     theme.DefaultTheme.FocusSymbolStyle,
		UnFocusSymbolStyle:   theme.DefaultTheme.UnFocusSymbolStyle,
		FocusIntervalStyle:   theme.DefaultTheme.FocusIntervalStyle,
		UnFocusIntervalStyle: theme.DefaultTheme.UnFocusIntervalStyle,
		ValueStyle:           theme.DefaultTheme.ChoiceTextStyle.Underline(),
		status:               Normal,
	}
	
	return c
}

// NewSpinner constructor
func NewSpinner() *Spinner {
	c := &Spinner{
		Model:               SpinnerDefaultModel,
		Shape:               SpinnerDefaultShape,
		ShapeStyle:          SpinnerDefaultShapeStyle,
		Prompt:              SpinnerDefaultPrompt,
		DisableOutPutResult: SpinnerDefaultDisableOutPutResult,
		Status:              SpinnerDefaultStatus,
		Quit:                InterruptKey,
	}
	return c
}

// NewSelectionWithList WIP
func NewSelectionWithList[T list.DefaultItem](items []T) *SelectionWithList[T] {
	keymap := &SelectionWithListKeyMap{
		Choice: key.NewBinding(
			key.WithKeys("tab"),
		),
	}
	del := &selectionWithListDelegate[T]{
		height:   2,
		spacing:  1,
		Selected: make(map[string]list.DefaultItem),
		Styles:   list.NewDefaultItemStyles(),
	}

	listModel := list.New([]list.Item{}, del, 0, 0)

	c := &SelectionWithList[T]{
		del:    del,
		List:   &listModel,
		KeyMap: keymap,
	}

	c.SetItems(items)
	return c
}
