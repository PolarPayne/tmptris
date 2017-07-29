package parseutils

import (
	"errors"
	"regexp"
	"strconv"
)

func runeInRunes(r rune, rs []rune) bool {
	for _, a := range rs {
		if r == a {
			return true
		}
	}
	return false
}

func EatRunes(s string, rs ...rune) string {
	count := 0
	for _, r := range s {
		if runeInRunes(r, rs) {
			count++
		} else {
			break
		}
	}

	return s[count:]
}

// EatWhitespace reads 's' until a non-whitespace rune is encountered,
// and return rest.
func EatWhitespace(s string) string {
	return EatRunes(s, ' ', '\t', '\n')
}

// SplitFrom reads 's' until one of the runes in 'rs' is encountered
// and returns (lhs, rest).
func SplitFrom(s string, rs ...rune) (string, string) {
	count := 0
	for _, r := range s {
		if !runeInRunes(r, rs) {
			count++
		} else {
			break
		}
	}

	return s[:count], s[count:]
}

// SplitFromWhitespace reads 's' until a whitespace is encountered,
// and then returns (lhs, rest).
func SplitFromWhitespace(s string) (string, string) {
	return SplitFrom(s, ' ', '\t', '\n')
}

// re is always prepended with a '^' to only match from the start of the string
func ReadRegexp(s, re string) (string, string) {
	reg := regexp.MustCompile("^" + re)
	res := reg.FindStringIndex(s)

	if res == nil {
		return "", s // no matches
	}
	return s[:res[1]], s[res[1]:]
}

func ReadString(s, v string) (string, string) {
	return ReadRegexp(s, regexp.QuoteMeta(v))
}

func ReadInt(s string) (int, string, error) {
	lhs, rest := ReadRegexp(s, "-?(0x?)?[0-9]+")
	if len(lhs) == 0 {
		return 0, s, errors.New("string didn't contain an integer")
	}

	out, err := strconv.ParseInt(lhs, 0, 0)
	if err != nil {
		return 0, s, err
	}

	return int(out), rest, nil
}
