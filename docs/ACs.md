<!-- Version: 1.2 | Last updated: 2026-07-09 -->

# Acceptance Criteria

This is the canonical spec. ACs introduced from 2026-06-12 onward live here.
Pre-cutover ACs remain in their originating issues until cited or migrated.

Last migrated: AC17.2 from #17 on 2026-07-09

---

## Symbol Sanitization

### AC14.1 - Given symbol sanitization is active outside protected code regions, common horizontal Unicode arrows have corresponding ASCII arrow notation in stdout.
- Introduced: #14 (closed 2026-06-12)
- Tests:
  - ✅ RT-14.1: Short arrows through the CLI: `← → ↔ ⇒ ⇐ ⇔` become `<- -> <-> => <= <=>`.
  - ✅ RT-14.2: Long arrows and mapsto arrows through the CLI: `⟵ ⟶ ⟷ ⟸ ⟹ ⟺ ↦ ↤` become `<-- --> <--> <== ==> <==> \|-> <-\|`.
  - ✅ RT-14.3: Multiple arrow forms in one input line retain surrounding text and report the correct symbol replacement count.

**Key:** ✅ passing · ⏳ pending · ❌ failing · ~~🚫 removed~~

### AC14.2 - Given symbol sanitization is active, existing quote, dash, ellipsis, and bullet mappings retain their established ASCII forms.
- Introduced: #14 (closed 2026-06-12)
- Tests:
  - ✅ RT-14.4: Em dash, en dash, smart quotes, and ellipsis through the CLI become `---`, `--`, straight quotes, and `...`.
  - ✅ RT-14.5: A bullet-prefixed line through the CLI remains a hyphen list item while an arrow elsewhere in the same line uses ASCII arrow notation.

**Key:** ✅ passing · ⏳ pending · ❌ failing · ~~🚫 removed~~

### AC14.3 - Given text is protected as Markdown fenced code, inline code, or org-mode source, Unicode arrows remain unchanged in protected text.
- Introduced: #14 (closed 2026-06-12)
- Tests:
  - ✅ RT-14.6: A Markdown fenced block containing Unicode arrows preserves those arrows while arrows outside the fence use ASCII notation.
  - ✅ RT-14.7: An inline backtick span containing Unicode arrows preserves those arrows while arrows outside the span use ASCII notation.
  - ✅ RT-14.8: An org-mode source block containing Unicode arrows preserves those arrows while arrows outside the block use ASCII notation.

**Key:** ✅ passing · ⏳ pending · ❌ failing · ~~🚫 removed~~

## OED Spelling Sanitization

### AC15.1 - Given OED spelling conversion is active, common misspelled recognizable-family variants have the OED spelling in stdout.
- Introduced: #15 (closed 2026-07-08)
- Tests:
  - ✅ RT-15.1: `recognizeable`, `recognizeably`, and `recognizeability` through the CLI become `recognizable`, `recognizably`, and `recognizability`.
  - ✅ RT-15.2: `recogniseable`, `recogniseably`, and `recogniseability` through the CLI become `recognizable`, `recognizably`, and `recognizability`.
  - ✅ RT-15.3: `unrecognizeable`, `unrecogniseably`, and `Unrecogniseability` through the CLI become `unrecognizable`, `unrecognizably`, and `Unrecognizability`.

**Key:** ✅ passing · ⏳ pending · ❌ failing · ~~🚫 removed~~

### AC17.1 - Given OED spelling conversion is active, possessive suffixes do not prevent mapped base words from using OED spelling in stdout.
- Introduced: #17 (closed 2026-07-09)
- Tests:
  - ✅ RT-17.1: `color's` and `neighbor's` through the CLI become `colour's` and `neighbour's`.
  - ✅ RT-17.2: `organise's` and `Recognise's` through the CLI become `organize's` and `Recognize's`.

**Key:** ✅ passing · ⏳ pending · ❌ failing · ~~🚫 removed~~

### AC17.2 - Given OED spelling conversion is active, apostrophes inside words that are not possessive suffixes remain part of the lookup token and do not create partial-word conversions.
- Introduced: #17 (closed 2026-07-09)
- Tests:
  - ✅ RT-17.3: A contraction-like token containing a mapped prefix remains unchanged unless the full token is explicitly mapped.
  - ✅ RT-17.4: A plural possessive suffix after a mapped plural base preserves the trailing apostrophe.

**Key:** ✅ passing · ⏳ pending · ❌ failing · ~~🚫 removed~~
