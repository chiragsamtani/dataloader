package utils

import "golang.org/x/text/cases"

const (
	CountryCodeLength = 2
)

func MergeStringFieldByLength(exist, new string, caseOption *cases.Caser) string {
	if len(exist) > len(new) {
		return caseOption.String(exist)
	}
	return caseOption.String(new)
}

func MergeCountry(exist, new string) string {
	if new != "" && len(new) == CountryCodeLength {
		return new
	} else if exist == "" {
		return new
	}
	return exist
}

func MergingCoordinateFields(exist, new float64) float64 {
	if exist == 0.0 {
		return new
	}
	return exist
}

func MergeStringArrayField(exist, new []string, caseOption *cases.Caser) []string {
	for _, val := range new {
		exist = append(exist, caseOption.String(val))
	}
	return exist
}
