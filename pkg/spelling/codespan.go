// ABOUTME: Inline protected span splitter. Splits a line into segments tagged
// as Code or Text so literal content can bypass sanitization.
package spelling

import (
	"strings"
	"unicode"
)

// SegmentKind distinguishes code spans from normal text.
type SegmentKind int

const (
	// Text is content outside backtick spans, eligible for processing.
	Text SegmentKind = iota
	// Code is content inside backtick spans, exempt from processing.
	Code
)

// Segment is a contiguous piece of a line with a kind tag.
type Segment struct {
	Content string
	Kind    SegmentKind
}

const (
	orgVerbatimOpeningBoundaries = "-([{"
	orgVerbatimClosingBoundaries = ".,;:!?')}*[]"
)

// SplitCodeSpans splits a line into Text and Code segments based on Markdown
// backticks and org-mode verbatim delimiters. Delimiters are included in Code
// segments. Unclosed or invalid delimiters are treated as literal text.
func SplitCodeSpans(line string) []Segment {
	runes := []rune(line)
	var segments []Segment
	i := 0

	for i < len(runes) {
		if end, ok := protectedSpanEnd(runes, i); ok {
			segments = append(segments, Segment{
				Content: string(runes[i : end+1]),
				Kind:    Code,
			})
			i = end + 1
			continue
		}

		j := i + 1
		for j < len(runes) {
			if _, ok := protectedSpanEnd(runes, j); ok {
				break
			}
			j++
		}
		segments = append(segments, Segment{
			Content: string(runes[i:j]),
			Kind:    Text,
		})
		i = j
	}

	return segments
}

func protectedSpanEnd(runes []rune, start int) (int, bool) {
	switch runes[start] {
	case '`':
		for end := start + 1; end < len(runes); end++ {
			if runes[end] == '`' {
				return end, true
			}
		}
	case '=':
		return orgVerbatimSpanEnd(runes, start)
	}
	return 0, false
}

func orgVerbatimSpanEnd(runes []rune, start int) (int, bool) {
	if !isOrgVerbatimOpeningBoundary(runes, start) {
		return 0, false
	}
	if start+1 >= len(runes) || unicode.IsSpace(runes[start+1]) {
		return 0, false
	}

	for end := start + 2; end < len(runes); end++ {
		if runes[end] != '=' || unicode.IsSpace(runes[end-1]) {
			continue
		}
		if isOrgVerbatimClosingBoundary(runes, end) {
			return end, true
		}
	}
	return 0, false
}

func isOrgVerbatimOpeningBoundary(runes []rune, start int) bool {
	if start == 0 {
		return true
	}
	previous := runes[start-1]
	return unicode.IsSpace(previous) || strings.ContainsRune(orgVerbatimOpeningBoundaries, previous)
}

func isOrgVerbatimClosingBoundary(runes []rune, end int) bool {
	if end == len(runes)-1 {
		return true
	}
	next := runes[end+1]
	return unicode.IsSpace(next) || strings.ContainsRune(orgVerbatimClosingBoundaries, next)
}
