package multiselect

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/components/selectd"
)

// Option the option of Select
type Option func(s *Select)

// WithRowRender default is
//
// `
// fmt.Sprintf("%s [%s] %s", cursorSymbol, hintSymbol, choice)
// `
func WithRowRender(rowRender func(string, string, string) string) Option {
	return func(s *Select) {
		s.inner.RowRender = rowRender
	}
}

// WithPageSize default is 5
func WithPageSize(pageSize int) Option {
	return func(s *Select) {
		s.inner.PageSize = pageSize
	}
}

// WithKeyBinding replace key map.
func WithKeyBinding(keymap selectd.KeyMap) Option {
	return func(s *Select) {
		s.inner.Keymap = keymap
	}
}

// WithCursorSymbol default is ">"
func WithCursorSymbol(symbol string) Option {
	return func(s *Select) {
		s.inner.CursorSymbol = symbol
	}
}

// WithCursorSymbolStyle default is theme.DefaultTheme.CursorSymbolStyle.
func WithCursorSymbolStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.CursorSymbolStyle = style
	}
}

// WithChoiceTextStyle default is theme.DefaultTheme.ChoiceTextStyle.
func WithChoiceTextStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.ChoiceTextStyle = style
	}
}

// WithHintSymbol default is "✓".
func WithHintSymbol(selectedStr string) Option {
	return func(s *Select) {
		s.inner.HintSymbol = selectedStr
	}
}

// WithHintSymbolStyle default is Theme.MultiSelectedHintSymbolStyle.
func WithHintSymbolStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.HintSymbolStyle = style
	}
}

// WithUnHintSymbol default is "✗".
func WithUnHintSymbol(unSelectedStr string) Option {
	return func(s *Select) {
		s.inner.UnHintSymbol = unSelectedStr
	}
}

// WithUnHintSymbolStyle default is Theme.UnHintSymbolStyle.
func WithUnHintSymbolStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.UnHintSymbolStyle = style
	}
}

// WithPromptStyle default is Theme.PromptStyle.
func WithPromptStyle(style lipgloss.Style) Option {
	return func(s *Select) {
		s.inner.PromptStyle = style
	}
}

// WithPrompt default is "Please selectd your options:"
func WithPrompt(prompt ...string) Option {
	return func(s *Select) {
		if len(prompt) >= 1 && len(prompt[0]) > 0 {
			s.inner.Prompt = prompt[0]
		}
	}
}

// WithDisableOutputResult disable output result.
func WithDisableOutputResult() Option {
	return func(s *Select) {
		s.inner.DisableOutPutResult = true
	}
}