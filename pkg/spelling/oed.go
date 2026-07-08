// ABOUTME: OED spelling engine. Loads US→UK and -ise→-ize word lists into
// separate maps, performs case-preserving whole-word replacement, and tracks
// spelling and -ize correction counts independently.
package spelling

import (
	"strings"
	"unicode"
)

// OEDEngine holds separate lookup maps for US→UK spelling and -ise→-ize
// corrections, tracking replacement counts independently.
type OEDEngine struct {
	spelling        map[string]string
	ize             map[string]string
	SpellingChanges int
	IzeChanges      int
}

// NewOEDEngine creates an engine from two word list data strings:
// the first for US→UK spelling, the second for -ise→-ize corrections.
func NewOEDEngine(spellingData, izeData string) (*OEDEngine, error) {
	e := &OEDEngine{
		spelling: make(map[string]string),
		ize:      make(map[string]string),
	}
	if err := e.loadWordList(e.spelling, spellingData); err != nil {
		return nil, err
	}
	if err := e.loadWordList(e.ize, izeData); err != nil {
		return nil, err
	}
	return e, nil
}

// loadWordList parses lines of "wrong=correct" into the given map.
func (e *OEDEngine) loadWordList(dest map[string]string, data string) error {
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		if key != "" && val != "" {
			dest[strings.ToLower(key)] = strings.ToLower(val)
		}
	}
	return nil
}

// ProcessLine replaces words in a single line, preserving case.
func (e *OEDEngine) ProcessLine(line string) string {
	runes := []rune(line)
	var result strings.Builder
	result.Grow(len(line))

	i := 0
	for i < len(runes) {
		if isWordChar(runes[i]) {
			// Extract the whole word
			j := i
			for j < len(runes) && isWordChar(runes[j]) {
				j++
			}
			word := string(runes[i:j])
			replaced := e.replaceWord(word)
			result.WriteString(replaced)
			i = j
		} else {
			result.WriteRune(runes[i])
			i++
		}
	}
	return result.String()
}

// isWordChar returns true for letters and apostrophes (word-internal).
func isWordChar(r rune) bool {
	return unicode.IsLetter(r) || r == '\'' || r == '\u2019'
}

// replaceWord looks up a word in both maps and applies case-preserving
// replacement, incrementing the appropriate counter.
func (e *OEDEngine) replaceWord(word string) string {
	if replacement, category, ok := e.lookupWord(word); ok {
		e.incrementChanges(category)
		return applyCase(word, replacement)
	}

	base, suffix, hasPossessive := splitPossessive(word)
	if !hasPossessive {
		return word
	}

	if replacement, category, ok := e.lookupWord(base); ok {
		e.incrementChanges(category)
		return applyCase(base, replacement) + suffix
	}

	return word
}

// lookupWord returns a replacement and counter category for an exact word.
func (e *OEDEngine) lookupWord(word string) (string, string, bool) {
	lower := strings.ToLower(word)
	if replacement, ok := e.spelling[lower]; ok {
		return replacement, "spelling", true
	}
	if replacement, ok := e.ize[lower]; ok {
		return replacement, "ize", true
	}
	return "", "", false
}

// incrementChanges records a replacement in the matching change counter.
func (e *OEDEngine) incrementChanges(category string) {
	if category == "spelling" {
		e.SpellingChanges++
		return
	}
	if category == "ize" {
		e.IzeChanges++
	}
}

// splitPossessive separates a trailing possessive suffix from a word.
func splitPossessive(word string) (string, string, bool) {
	for _, suffix := range []string{"'s", "\u2019s", "'", "\u2019"} {
		if strings.HasSuffix(word, suffix) && len(word) > len(suffix) {
			return word[:len(word)-len(suffix)], suffix, true
		}
	}
	return word, "", false
}

// applyCase transfers the case pattern of orig onto replacement.
func applyCase(orig, replacement string) string {
	if orig == strings.ToLower(orig) {
		return replacement
	}
	if orig == strings.ToUpper(orig) {
		return strings.ToUpper(replacement)
	}
	// Title Case: first letter uppercase, rest lowercase
	origRunes := []rune(orig)
	if len(origRunes) > 0 && unicode.IsUpper(origRunes[0]) {
		runes := []rune(replacement)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		return string(runes)
	}
	// Mixed case fallback: return lowercase
	return replacement
}
