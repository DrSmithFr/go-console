package output

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/formatter"
)

func NewNullOutput() NullOutput {
	out := new(NullOutput)
	out.doWrite = out.Write
	return NullOutput{}
}

type NullOutput struct {
	doWrite func (string)
	formatter *formatter.OutputFormatterInterface
}

func (o *NullOutput) format(message string) string {
	if nil == o.formatter {
		return message
	}

	return (*o.formatter).Format(message)
}

func (o *NullOutput) Write(message string) {
	// do nothing
}

func (o *NullOutput) Writeln(message string) {
	o.doWrite(fmt.Sprintf("%s\n", message))
}

func (o *NullOutput) SetDecorated(decorated bool) {
	if nil == o.formatter {
		return
	}

	(*o.formatter).SetDecorated(decorated)
}

func (o *NullOutput) IsDecorated() bool {
	if nil == o.formatter {
		return false
	}

	return (*o.formatter).IsDecorated()
}

func (o *NullOutput) SetFormatter(formatter *formatter.OutputFormatterInterface) {
	o.formatter = formatter
}

func (o *NullOutput) GetFormatter() *formatter.OutputFormatterInterface {
	return o.formatter
}

