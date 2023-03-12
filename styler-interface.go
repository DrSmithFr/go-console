package go_console

import (
	"github.com/DrSmithFr/go-console/verbosity"
)

type StylerInterface interface {
	MaxLineLength() int

	// Write Implements io.Writer
	Write(p []byte) (n int, err error)

	// PrintTitle formats and print a command title.
	PrintTitle(message string)

	// PrintSection formats and print a section title.
	PrintSection(message string)

	// PrintListing formats and print a list.
	PrintListing(messages []string)

	// PrintText formats and print informational text.
	PrintText(message string)

	// PrintTexts formats and print informational text array.
	PrintTexts(message []string)

	// PrintComment formats and print a comment bar.
	PrintComment(message string)

	// PrintComments formats and print a comment bar.
	PrintComments(message []string)

	// PrintSuccess formats and print a success result bar.
	PrintSuccess(message string)

	// PrintSuccesses formats and print a success result bar.
	PrintSuccesses(message []string)

	// PrintError formats and print an error result bar.
	PrintError(message string)

	// PrintErrors formats and print an error result bar.
	PrintErrors(message []string)

	// PrintWarning formats and print an warning result bar.
	PrintWarning(message string)

	// PrintWarnings formats and print an warning result bar.
	PrintWarnings(message []string)

	// PrintNote formats and print a note admonition.
	PrintNote(message string)

	// PrintNotes formats and print a note admonition.
	PrintNotes(message []string)

	// PrintCaution formats and print a caution admonition.
	PrintCaution(message string)

	// PrintCautions formats and print a caution admonition.
	PrintCautions(message []string)

	// PrintNewLine Print n newline(n).
	PrintNewLine(count int)

	// Verbosity current verbosity of the Output.
	Verbosity() verbosity.Level

	// IsQuiet Returns whether verbosity is quiet (-q)
	IsQuiet() bool

	// IsVerbose Returns whether verbosity is verbose (-v)
	IsVerbose() bool

	// IsVeryVerbose Returns whether verbosity is very verbose (-vv)
	IsVeryVerbose() bool

	// IsDebug Returns whether verbosity is debug (-vvv)
	IsDebug() bool

	// TODO Formats a table.
	// Table(headers []string, rows [][]string)

	// TODO add ask(), askHidden(), confirm() and choice() when questionInterface is ready

	// TODO add progressStart(), progressAdvance(), and progressFinish() when helper.ProgressBar is ready

}
