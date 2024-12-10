package tea

import (
	"github.com/charmbracelet/lipgloss"
)

func InitTerminal() {
	_ = lipgloss.HasDarkBackground()
}

// This function is commented out because its causing garbage to be printed to webterminals
// when using coder agent-exec. We call it manually for non-agent-exec calls.
//func init() {
	// XXX: This is a workaround to make assure that Lip Gloss and Termenv
	// query the terminal before any Bubble Tea Program runs and acquires the
	// terminal. Without this, Programs that use Lip Gloss/Termenv might hang
	// while waiting for a a [termenv.OSCTimeout] while querying the terminal
	// for its background/foreground colors.
	//
	// This happens because Bubble Tea acquires the terminal before termenv
	// reads any responses.
	//
	// Note that this will only affect programs running on the default IO i.e.
	// [os.Stdout] and [os.Stdin].
	//
	// This workaround will be removed in v2.
//	_ = lipgloss.HasDarkBackground()
//}
//
