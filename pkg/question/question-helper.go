package question

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/output"
	"github.com/DrSmithFr/go-console/pkg/question/answers"
	"io"
	"strings"
)

type Helper struct {
	in  io.Reader
	out output.OutputInterface
}

func NewHelper(input io.Reader, output output.OutputInterface) *Helper {
	return &Helper{
		in:  input,
		out: output,
	}
}

func (h *Helper) Ask(question QuestionBasicInterface) string {
	run := func() (string, error) {
		answer, err := h.doAsk(question)

		if err == nil {
			return answer, nil
		}

		return "", err
	}

	if question.GetMaxAttempts() == 0 {
		for {
			answer, err := run()

			if err == nil {
				return answer
			}

			h.out.Writeln(fmt.Sprintf("<error>%s</error>", err.Error()))
		}
	} else {
		for attempts := 0; attempts < question.GetMaxAttempts(); attempts++ {
			answer, err := run()

			if err == nil {
				return answer
			}

			h.out.Writeln(fmt.Sprintf("<error>%s</error>", err.Error()))
		}
	}

	h.out.Writeln("<error>maximum number of maxAttempts reached</error>")
	panic(errors.New("aborting"))
}

func (h *Helper) doAsk(question QuestionBasicInterface) (string, error) {
	h.writePrompt(question)

	rawText, _ := bufio.NewReader(h.in).ReadString('\n')
	answer := strings.TrimSpace(rawText)

	if len(answer) == 0 {
		answer = question.GetDefaultAnswer()
	}

	if question.GetNormalizer() != nil {
		answer = question.GetNormalizer()(answer)
	}

	if question.GetValidator() != nil {
		if err := question.GetValidator()(answer); err != nil {
			return "", err
		}
	}

	return answer, nil
}

func (h *Helper) writePrompt(question QuestionBasicInterface) {
	if choices, ok := question.(QuestionChoicesInterface); ok {
		h.out.Writeln(choices.GetQuestion())

		for _, line := range h.formatChoiceQuestionChoices(choices, "info") {
			h.out.Writeln(line)
		}

		h.out.Writeln(choices.GetPrompt())
		return
	}

	if confirmation, ok := question.(QuestionConfirmationInterface); ok {
		message := fmt.Sprintf(
			"%s [<info>%s</info>/<info>%s</info>] ",
			confirmation.GetQuestion(),
			answers.YES,
			answers.NO,
		)

		h.out.Writeln(message)
		return
	}

	h.out.Writeln(question.GetQuestion())
}

func (h *Helper) formatChoiceQuestionChoices(question QuestionChoicesInterface, tag string) []string {
	result := make([]string, len(question.GetChoices()))

	for index, choice := range question.GetChoices() {
		result[index] = fmt.Sprintf("  [<%s>%d</%s>] %s", tag, index, tag, choice)
	}

	return result
}
