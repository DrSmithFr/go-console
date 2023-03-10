package go_console

import (
	"github.com/DrSmithFr/go-console/output"
	"github.com/DrSmithFr/go-console/verbosity"
)

type StylerInterface interface {
	// retrieve OutputInterface
	Output() output.OutputInterface

	SetMaxLineLength(length int)

	MaxLineLength() int

	// Formats a command title.
	PrintTitle(message string)

	// Formats a section title.
	PrintSection(message string)

	// Formats a list.
	PrintListing(messages []string)

	// Formats informational text.
	PrintText(message string)

	// Formats informational text array.
	PrintTexts(message []string)

	// Formats a comment bar.
	PrintComment(message string)

	// Formats a comment bar.
	PrintComments(message []string)

	// Formats a success result bar.
	PrintSuccess(message string)

	// Formats a success result bar.
	PrintSuccesses(message []string)

	// Formats an error result bar.
	PrintError(message string)

	// Formats an error result bar.
	PrintErrors(message []string)

	// Formats an warning result bar.
	PrintWarning(message string)

	// Formats an warning result bar.
	PrintWarnings(message []string)

	// Formats a note admonition.
	PrintNote(message string)

	// Formats a note admonition.
	PrintNotes(message []string)

	// Formats a caution admonition.
	PrintCaution(message string)

	// Formats a caution admonition.
	PrintCautions(message []string)

	// Add newline(s).
	PrintNewLine(count int)

	// Gets the current verbosity of the output.
	Verbosity() verbosity.Level

	// Returns whether verbosity is quiet (-q)
	IsQuiet() bool

	// Returns whether verbosity is verbose (-v)
	IsVerbose() bool

	// Returns whether verbosity is very verbose (-vv)
	IsVeryVerbose() bool

	// Returns whether verbosity is debug (-vvv)
	IsDebug() bool

	// TODO Formats a table.
	// Table(headers []string, rows [][]string)

	// TODO add ask(), askHidden(), confirm() and choice() when questionInterface is ready

	// TODO add progressStart(), progressAdvance(), and progressFinish() when helper.ProgressBar is ready

}
