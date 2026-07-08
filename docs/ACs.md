<!-- Version: 1.1 | Last updated: 2026-07-08 -->

# Acceptance Criteria

This is the canonical spec. ACs introduced from 2026-06-12 onward live here.
Pre-cutover ACs remain in their originating issues until cited or migrated.

Last migrated: AC15.1 from #15 on 2026-07-08

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
