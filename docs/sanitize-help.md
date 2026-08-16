sanitize converts English text from stdin and writes the transformed text to
stdout. It normalizes selected US and British -ise spellings to Oxford spelling
and replaces common typographic symbols with ASCII equivalents. It is not a
spell checker.

With no subcommands, both oed and symbols are applied.

usage: sanitize <subcommand> [<subcommand>...] [flags]

Subcommands:
  oed       Convert US spellings and non-OED -ise spellings to Oxford spelling
  symbols   Convert typographic characters to ASCII

Protected content:
  Markdown fenced code blocks   ``` ... ```
  Markdown inline code spans    `...`
  Org source blocks             #+BEGIN_SRC ... #+END_SRC
  Org verbatim spans            =...=

Text inside protected content passes through unchanged.

Flags:
  -q          Suppress change summary on stderr
  -h, --help  Print this help message
  --version   Print version
