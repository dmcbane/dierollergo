package main

import (
	"regexp"
	"strconv"
)

type GenerationType int

const (
	GenerationTypeClassic GenerationType = iota
	GenerationTypeStandard
	GenerationTypeHeroic
	GenerationTypePool
	GenerationTypePurchase
)

const GenerationTypePoolRegex = "\\A(\\d{1})/(\\d{1})/(\\d{1})/(\\d{1})/(\\d{1})/(\\d{1})\\z"
const GenerationTypePurchaseRegex = "(?i)\\A(low|standard|high|epic)\\z"

func IsValidPoolSpecification(value string) bool {
	if match, _ := regexp.MatchString(GenerationTypePoolRegex, value); !match {
		return false
	} else {
		re1, _ := regexp.Compile(GenerationTypePoolRegex)
		matches := re1.FindStringSubmatch(value)
		total := 0
		for i, v := range matches {
			if i == 0 {
				continue
			}
			iv, err := strconv.Atoi(v)
			if iv < 3 || err != nil {
				return false
			}
			total += iv
		}
		if total > 24 {
			return false
		} else {
			return true
		}
	}
}

func IsValidPurchaseSpecification(value string) bool {
	match, _ := regexp.MatchString(GenerationTypePurchaseRegex, value)
	return match
}

func ParsePoolSpecification(value string) *[]int {
	result := make([]int, 6)
	if IsValidPoolSpecification(value) {
		re1, _ := regexp.Compile(GenerationTypePoolRegex)
		matches := re1.FindStringSubmatch(value)
		for i, v := range matches {
			if i == 0 {
				continue
			}
			result[i-1], _ = strconv.Atoi(v)
		}
	}
	return &result
}
