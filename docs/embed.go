// ABOUTME: Embeds executable help documentation so installed binaries remain
// self-contained while docs/sanitize-help.md stays the editable source.
package docs

import _ "embed"

//go:embed sanitize-help.md
var SanitizeHelp string
