package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// WaifuTheme implements the Theme interface with Sakura pink and lavender colors.
type WaifuTheme struct {
	BaseTheme
}

// NewWaifuTheme creates a new instance of the Waifu (Sakura) theme.
func NewWaifuTheme() *WaifuTheme {
	theme := &WaifuTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{
		Dark:  "#FF6B9D", // Sakura Pink
		Light: "#FF6B9D",
	}
	theme.SecondaryColor = lipgloss.AdaptiveColor{
		Dark:  "#C084FC", // Lavender
		Light: "#C084FC",
	}
	theme.AccentColor = lipgloss.AdaptiveColor{
		Dark:  "#67E8F9", // Pastel Cyan
		Light: "#67E8F9",
	}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{
		Dark:  "#FB7185", // Soft Red
		Light: "#FB7185",
	}
	theme.WarningColor = lipgloss.AdaptiveColor{
		Dark:  "#FDE68A", // Pastel Yellow
		Light: "#FDE68A",
	}
	theme.SuccessColor = lipgloss.AdaptiveColor{
		Dark:  "#86EFAC", // Pastel Green
		Light: "#86EFAC",
	}
	theme.InfoColor = lipgloss.AdaptiveColor{
		Dark:  "#67E8F9", // Pastel Cyan
		Light: "#67E8F9",
	}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{
		Dark:  "#F5E6FF",
		Light: "#3D1F4E",
	}
	theme.TextMutedColor = lipgloss.AdaptiveColor{
		Dark:  "#9B8AAE",
		Light: "#8B6E9E",
	}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{
		Dark:  "#FFB6D9",
		Light: "#FF6B9D",
	}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{
		Dark:  "#1A1025",
		Light: "#FFF0F5",
	}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{
		Dark:  "#2D1F3D",
		Light: "#FFE4EF",
	}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{
		Dark:  "#110A1A",
		Light: "#FFF5F9",
	}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{
		Dark:  "#4A3560",
		Light: "#E8B4D0",
	}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{
		Dark:  "#FF6B9D",
		Light: "#FF6B9D",
	}
	theme.BorderDimColor = lipgloss.AdaptiveColor{
		Dark:  "#2A1F35",
		Light: "#F5D5E5",
	}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{
		Dark:  "#5DA07A",
		Light: "#388E3C",
	}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{
		Dark:  "#D4636E",
		Light: "#E57373",
	}
	theme.DiffContextColor = lipgloss.AdaptiveColor{
		Dark:  "#9B8AAE",
		Light: "#8B6E9E",
	}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{
		Dark:  "#9B8AAE",
		Light: "#8B6E9E",
	}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{
		Dark:  "#C8FAD0",
		Light: "#A5D6A7",
	}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{
		Dark:  "#FAD0D5",
		Light: "#EF9A9A",
	}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{
		Dark:  "#1F2E25",
		Light: "#E8F5E9",
	}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{
		Dark:  "#2E1F22",
		Light: "#FFEBEE",
	}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{
		Dark:  "#1A1025",
		Light: "#FFF0F5",
	}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{
		Dark:  "#9B8AAE",
		Light: "#8B6E9E",
	}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{
		Dark:  "#1A291F",
		Light: "#C8E6C9",
	}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{
		Dark:  "#291A1D",
		Light: "#FFCDD2",
	}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{
		Dark:  "#F5E6FF",
		Light: "#3D1F4E",
	}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{
		Dark:  "#FF6B9D",
		Light: "#E91E63",
	}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{
		Dark:  "#C084FC",
		Light: "#9C27B0",
	}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{
		Dark:  "#FFB6D9",
		Light: "#FF6B9D",
	}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{
		Dark:  "#86EFAC",
		Light: "#4CAF50",
	}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{
		Dark:  "#FDE68A",
		Light: "#F9A825",
	}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{
		Dark:  "#FDE68A",
		Light: "#F9A825",
	}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{
		Dark:  "#FF6B9D",
		Light: "#FF6B9D",
	}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{
		Dark:  "#4A3560",
		Light: "#E8B4D0",
	}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{
		Dark:  "#FF6B9D",
		Light: "#FF6B9D",
	}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{
		Dark:  "#C084FC",
		Light: "#9C27B0",
	}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{
		Dark:  "#67E8F9",
		Light: "#00BCD4",
	}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{
		Dark:  "#FFB6D9",
		Light: "#FF6B9D",
	}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{
		Dark:  "#F5E6FF",
		Light: "#3D1F4E",
	}

	// Syntax highlighting colors (pastel rainbow)
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{
		Dark:  "#7A6B8A",
		Light: "#9E8EB0",
	}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{
		Dark:  "#FF6B9D",
		Light: "#E91E63",
	}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{
		Dark:  "#C084FC",
		Light: "#9C27B0",
	}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{
		Dark:  "#67E8F9",
		Light: "#00BCD4",
	}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{
		Dark:  "#86EFAC",
		Light: "#4CAF50",
	}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{
		Dark:  "#FDE68A",
		Light: "#F9A825",
	}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{
		Dark:  "#FFB6D9",
		Light: "#FF6B9D",
	}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{
		Dark:  "#FB7185",
		Light: "#F44336",
	}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{
		Dark:  "#F5E6FF",
		Light: "#3D1F4E",
	}

	return theme
}

func init() {
	// Register the Waifu theme with the theme manager
	RegisterTheme("waifu", NewWaifuTheme())
}
