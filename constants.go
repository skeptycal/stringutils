package stringutils

import (
	"unicode/utf8"
)

// Numbers fundamental to the encoding.
const (
	RuneError = utf8.RuneError // '\uFFFD'       // the "error" Rune or "Unicode replacement character"
	RuneSelf  = utf8.RuneSelf  // 0x80           // characters below RuneSelf are represented as themselves in a single byte.
	MaxRune   = utf8.MaxRune   // '\U0010FFFF'   // Maximum valid Unicode code point.
	UTFMax    = utf8.UTFMax    // 4              // maximum number of bytes of a UTF-8 encoded Unicode character.
)
