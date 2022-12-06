package style

import "DrSmithFr/go-console/pkg/output"

type StylerInterface interface {
	// retrieve OutputInterface
	GetOutput() output.OutputInterface

	SetMaxLineLength(length int)

	GetMaxLineLength() int

	// Formats a command title.
	Title(message string)

	// Formats a section title.
	Section(message string)

	// Formats a list.
	Listing(messages []string)

	// Formats informational text.
	Text(message string)

	// Formats informational text array.
	TextArray(message []string)

	// Formats a comment bar.
	Comment(message string)

	// Formats a comment bar.
	CommentArray(message []string)

	// Formats a success result bar.
	Success(message string)

	// Formats a success result bar.
	SuccessArray(message []string)

	// Formats an error result bar.
	Error(message string)

	// Formats an error result bar.
	ErrorArray(message []string)

	// Formats an warning result bar.
	Warning(message string)

	// Formats an warning result bar.
	WarningArray(message []string)

	// Formats a note admonition.
	Note(message string)

	// Formats a note admonition.
	NoteArray(message []string)

	// Formats a caution admonition.
	Caution(message string)

	// Formats a caution admonition.
	CautionArray(message []string)

	// Formats a table.
	Table(headers []string, rows [][]string)

	// Add newline(s).
	NewLine(count int)

	// TODO add ask(), askHidden(), confirm() and choice() when questionInterface is ready

	// TODO add progressStart(), progressAdvance(), and progressFinish() when helper.ProgressBar is ready

}
