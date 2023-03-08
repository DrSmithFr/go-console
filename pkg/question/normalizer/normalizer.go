package normalizer

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Normalizer func(string) string

func MakeChainedNormalizer(normalizers ...Normalizer) Normalizer {
	return func(answer string) string {
		for _, normalizer := range normalizers {
			answer = normalizer(answer)
		}

		return answer
	}
}

func Ucfirst(answer string) string {
	return cases.Title(language.English, cases.Compact).String(answer)
}
