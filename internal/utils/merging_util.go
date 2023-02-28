package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

const (
	CountryCodeLength = 2
)

func MergeStringFieldByLength(exist, new string) string {
	if len(exist) > len(new) {
		return strings.ToTitle(exist)
	}
	return strings.ToTitle(new)
}

func MergeCountry(exist, new string) string {
	country := ""
	if exist == "" {
	} else {
		if new != "" && len(new) == CountryCodeLength {
			country = new
		} else {
			country = exist
		}
	}
	return country
}

func MergingCoordinateFields(exist, new float64) float64 {
	if exist == 0.0 {
		return new
	} else {
		return exist
	}
}

func MergeStringArrayField(exist, new []string) []string {
	for _, val := range new {
		exist = append(exist, cases.Title(language.English).String(val))
	}
	return exist
}
