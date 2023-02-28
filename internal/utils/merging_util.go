package utils

import (
	"golang.org/x/text/cases"
	"strings"
	"unicode"
)

const (
	CountryCodeLength = 2
)

// MergeStringFieldByLength  will prioritize data with a bigger string length
// CaseOption will define what case to apply on each string, if caseOption
// is nil, no further modification is done to elements in the string
// Supported Case Options: lower (lowercase), upper (uppercase), title (capitalizing first char)
func MergeStringFieldByLength(exist, new string, caseOption *cases.Caser) string {
	if len(exist) > len(new) {
		return caseOption.String(exist)
	}
	return caseOption.String(new)
}

// MergeCountry will prioritize data that adheres to ISO-3166 country codes
// if no country code is provided, it will defer to using any valid string
// that's provided
func MergeCountry(exist, new string) string {
	if new != "" && len(new) == CountryCodeLength {
		return new
	} else if exist == "" {
		return new
	}
	return exist
}

// MergingCoordinateFields will return the existing data if and only if the
// existing data is greater than 0
func MergingCoordinateFields(exist, new float64) float64 {
	if exist == 0.0 {
		return new
	}
	return exist
}

// MergeStringArrayField will merge two string arrays by combining the elements of
// both arrays. CaseOption will define what case to apply on each string, if caseOption
// is nil, no further modification is done to elements in the string
// Supported Case Options: lower (lowercase), upper (uppercase), title (capitalizing first char)
func MergeStringArrayField(exist, new []string, caseOption *cases.Caser) []string {
	for _, val := range new {
		if caseOption != nil {
			val = caseOption.String(val)
		}
		exist = append(exist, strings.TrimSpace(val))
	}
	return exist
}

// MergeStringArrayWithNoDuplicates will take a union of existing and new data array
// but eliminates duplicates by using a Set
// CaseOption will define what case to apply on each string, if caseOption
// is nil, no further modification is done to elements in the string
// Supported Case Options: lower (lowercase), upper (uppercase), title (capitalizing first char)
func MergeStringArrayWithNoDuplicates(exist, new []string, caseOption *cases.Caser) []string {
	set := NewSet()
	for _, val := range exist {
		if caseOption != nil {
			val = caseOption.String(val)
		}
		set.Add(strings.TrimSpace(val))
	}
	for _, val := range new {
		if caseOption != nil {
			val = caseOption.String(val)
		}
		set.Add(strings.TrimSpace(val))
	}
	return set.ConvertToArray()
}

func AddSpaceBetweenUpperCaseCharacters(s string) string {
	result := ""
	for _, c := range s {
		if unicode.IsUpper(c) {
			result += " "
		} else if unicode.IsSpace(c) {
			continue
		}
		result += string(c)
	}
	result = strings.TrimSpace(result)
	return result
}
