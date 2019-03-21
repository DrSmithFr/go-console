package style

import "github.com/MrSmith777/go-console/pkg/output"

type StylerInterface interface {
	// retrieve OutputInterface
	GetOutput() output.OutputInterface

	// Formats a command title.
	Title(message string)

	// Formats a section title.
	Section(message string)

	// Formats a list.
	Listing(messages []string)

	// Formats informational text.
	Text(message string)

	// Formats a success result bar.
	Success(message string)

	// Formats an error result bar.
	Error(message string)

	// Formats an warning result bar.
	Warning(message string)

	// Formats a note admonition.
	Note(message string)

	// Formats a caution admonition.
	Caution(message string)

	// Formats a table.
	Table(headers []string, rows [][]string)

	// Add newline(s).
	NewLine(count int)

	// TODO add ask(), askHidden(), confirm() and choice() when inputInterface is ready

	// TODO add progressStart(), progressAdvance(), and progressFinish() when helper.ProgressBar is ready

}
