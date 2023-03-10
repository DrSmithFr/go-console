package main

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/option"
	"github.com/DrSmithFr/go-console/question"
	"github.com/DrSmithFr/go-console/table"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"regexp"
)

func main() {

	//
	// Easy way to create a command with arguments and options
	//

	cmd := go_console.Cli{
		Args: []go_console.Arg{
			{
				Name: "name",
				Mode: argument.Required,
			},
		},
		Opts: []go_console.Opt{
			{
				Name:     "foo",
				Shortcut: "f",
				Mode:     option.None,
			},
		},
	}

	cmd.Build()

	//cmd := go_console.NewCli().
	//	AddInputArgument(
	//		argument.New("name", argument.Required),
	//	).
	//	AddInputOption(
	//		option.New("foo", option.None).
	//			SetShortcut("f"),
	//	).
	//	Build()

	name := cmd.Input().Argument("name")

	if option.Defined == cmd.Input().Option("foo") {
		cmd.PrintSuccess(fmt.Sprintf("Hello %s, foo option is set", name))
	} else {
		cmd.PrintWarning(fmt.Sprintf("Hello %s, foo option not set", name))
	}

	//
	// Easy way to style your program
	//

	cmd.PrintTitle("Starting console")

	cmd.PrintNote(
		fmt.Sprintf(
			"name argument value '%s'",
			cmd.Input().Argument("name"),
		),
	)

	cmd.PrintTexts([]string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"<comment>Lorem ipsum dolor sit amet</comment>, <info>consectetur adipiscing elit.</info>",
		"<b>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</b>",
		"<b>Lorem ipsum dolor sit amet, <u>consectetur adipiscing elit</u>.</b>",
	})

	cmd.PrintCaution("This is a caution")

	//
	// Easy User Interaction
	//

	qh := question.NewHelper(os.Stdin, cmd.Output())

	// Simple question with default answer
	firstname := qh.Ask(
		question.NewQuestion("What is your firstname?").
			SetDefaultAnswer("Doe"),
	)

	// Simple question with normalizer
	lastname := qh.Ask(
		question.NewQuestion("What is your lastname?").
			SetNormalizer(func(answer string) string {
				return cases.Title(language.English, cases.Compact).String(answer)
			}),
	)

	// Simple question with custom validator
	nickname := qh.Ask(
		question.NewQuestion("What is your nickname?").
			SetValidator(func(answer string) error {
				regex := regexp.MustCompile("^(\\w|_|-)*$")
				if match := regex.MatchString(answer); !match {
					return errors.New("nickname must be alphanumeric")
				}

				return nil
			}),
	)

	cmd.PrintTexts([]string{
		" ",
		" ",
		" ",
	})

	//
	// Render a tables
	//

	tab := table.NewTable().
		SetHeaderTitle("My Table").
		SetFooterTitle("By DrSmith").
		SetColumnPadding(3, table.PadToRight).
		AddHeader(
			&table.TableRow{
				Columns: map[int]table.TableColumnInterface{
					0: &table.TableColumn{
						Cell: &table.TableCell{
							Value:   "Centred Header Cell",
							Colspan: 3,
							PadType: table.PadToCenter,
						},
					},
				},
			},
		).
		AddHeadersFromString(
			[][]string{
				{"==="},
				{"FIRSTNAME", "LASTNAME", "NICKNAME"},
			},
		)

	tab.
		AddRowsFromString(
			[][]string{
				{firstname, lastname, nickname},
				{"---"}, // shortcut for TableSeparator
				{"9971-5-0210-0", "A Tale of Two Cities", "Charles\nDickens"},
			},
		).
		AddTableSeparator().
		AddRow(
			table.NewTableRow().
				AddColumn(
					table.NewTableColumn().
						SetCell(
							table.NewTableCell("This value spans 2 columns.").
								SetPadType(table.PadToCenter).
								SetColspan(2),
						),
				).
				AddColumn(
					table.NewTableColumn().
						SetCell(
							table.NewTableCell("stand alone value"),
						),
				),
		).
		AddTableSeparator().
		AddRowsFromString(
			[][]string{
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)

	tab.
		AddTableSeparator().
		AddRow(
			table.NewTableRow().
				AddColumn(
					table.NewTableColumn().
						SetCell(
							table.NewTableCell("This value spans 2 columns.").
								SetColspan(2),
						),
				).
				AddColumn(
					table.NewTableColumn().
						SetCell(
							table.NewTableCell("stand alone value"),
						),
				),
		)

	tab.
		AddTableSeparator().
		AddRow(
			&table.TableRow{
				Columns: map[int]table.TableColumnInterface{
					0: &table.TableColumn{
						Cell: &table.TableCell{
							Value:   "<question>This value spans use <b>3 lines</b> to get fully displayed and now to long to feet inside the table.</question>",
							Colspan: 3,
							PadType: table.PadToCenter,
						},
					},
				},
			},
		)

	tab.
		AddTableSeparator().
		AddRowsFromString(
			[][]string{
				{"<b>960-425-059-0</b>", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)

	render := table.NewRender(cmd.Output()).
		SetContent(tab)

	render.SetColumnsMinWidths(map[int]int{
		0: 10,
		1: 0,
		2: 30,
	})

	render.SetStyleFromName("box-double")
	render.Render()

	//
	// enable stylish errors on need
	//

	defer cmd.HandleRuntimeException()
	panic("this error will be stylish!")
}
