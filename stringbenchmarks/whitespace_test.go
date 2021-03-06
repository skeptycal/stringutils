package stringbenchmarks

import (
	"fmt"
	"math"
	"testing"
	"unicode"
)

var (
	testBytes = []struct {
		name string
		f    func(c byte) bool
		want func(c byte) bool
	}{
		// TODO: Add test cases.
		{"isASCIISpace", isASCIISpace, isASCIISpace},
		{"isWhiteSpace", isWhiteSpace, isASCIISpace},
		{"isWhiteSpace2", isWhiteSpace2, isASCIISpace},
		{"unicodeIsSpace", unicodeIsSpace, unicodeIsSpace},

		// {"isSpaceMask", isSpaceMask, isASCIISpace},
		{"Slice", isWhiteSpaceStringSliceBytes, isASCIISpace},
		{"Regex", isWhiteSpaceRegexByte, isASCIISpace},
	}

	testRunes = []struct {
		name string
		f    func(r rune) bool
		want func(r rune) bool
	}{
		// TODO: Add test cases.
		{"isWhiteSpaceLogicChainNoCheck", isWhiteSpaceLogicChainNoCheck, unicode.IsSpace},
		{"isWhiteSpaceLogicChain", isWhiteSpaceLogicChain, unicode.IsSpace},
		{"isWhiteRuneSlice", isWhiteSpaceRuneSlice, unicode.IsSpace},
		{"isWhiteSpaceBoolMap", isWhiteSpaceBoolMap, unicode.IsSpace},
		{"IsUnicodeWhiteSpaceMap", isUnicodeWhiteSpaceMap, unicode.IsSpace},
		{"unicodeIsSpace", unicodeIsSpaceRune, unicode.IsSpace},
		{"IsUnicodeWhiteSpaceMapSwitch", isUnicodeWhiteSpaceMapSwitch, unicode.IsSpace},
		{"isWhiteSpaceStringMap", isWhiteSpaceStringMap, unicode.IsSpace},
		{"isWhiteSpaceStringSlice", isWhiteSpaceStringSlice, unicode.IsSpace},
		{"(stlib) unicode.IsSpace", unicode.IsSpace, unicode.IsSpace},
	}

	testByteStrings = []struct {
		name string
		f    func(s string) bool
		want func(c byte) bool
	}{
		// TODO: Add test cases.
		{"isWhiteSpaceContainsByte", isWhiteSpaceContainsByte, isASCIISpace},
		{"isWhiteSpaceIndexByte", isWhiteSpaceIndexByte, isASCIISpace},
		// {"isWhiteSpaceTrim", isWhiteSpaceTrim, unicodeIsSpace}, // erratic results
	}

	testRuneStrings = []struct {
		name string
		f    func(s string) bool
		want func(c byte) bool
	}{
		// TODO: Add test cases.
		{"isWhiteSpaceIndexRune", isWhiteSpaceIndexRune, nil},
		{"isWhiteSpaceContainsRune", isWhiteSpaceContainsRune, nil},
	}
)

func BenchmarkByte(b *testing.B) {
	benchmarks := testBytes
	for _, bb := range benchmarks {
		name := fmt.Sprintf("Benchmark: %s", bb.name)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range ByteSamples() {
					for j := 0; j < 10; j++ {
						bb.f(c)
					}
				}
			}
		})
	}
}

func BenchmarkRunes(b *testing.B) {
	benchmarks := testRunes
	for _, bb := range benchmarks {
		name := fmt.Sprintf("Benchmark: %s", bb.name)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, r := range RuneSamples() {
					bb.f(r)
				}
			}
		})
	}
}

func BenchmarkStringByte(b *testing.B) {
	benchmarks := testByteStrings
	for _, bb := range benchmarks {
		name := fmt.Sprintf("Benchmark: %s", bb.name)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, s := range runeStringSamples() {
					bb.f(s)
				}
			}
		})
	}
}

const expCount = 4

// func BenchmarkSamples(b *testing.B) {
//     for benchmark := range benchmarkList {
//         runBenchmark(b, benchmark)
//     }

// }

func BenchmarkStringRune(b *testing.B) {
	benchmarks := testRuneStrings
	for ec := 0; ec < expCount; ec++ {
		for _, bb := range benchmarks {
			name := fmt.Sprintf("Benchmark: %s (run %v times)", bb.name, int(math.Pow(2, float64(ec))))
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for j := 0; j < int(math.Pow(2, float64(ec))); j++ {
						for _, s := range runeStringSamples() {
							bb.f(s)
						}
					}
				}
			})
		}
	}
}

func TestWhiteSpaceBytes(t *testing.T) {
	tests := testBytes

	// 	case '\u00a0', '\u0085':
	// return true

	for _, tt := range tests {
		for _, c := range SmallByteSamples() {
			// want := unicode.IsSpace(rune(c)) // standard library
			want := tt.want(c)
			name := fmt.Sprintf("%s(%q)", tt.name, c)
			t.Run(name, func(t *testing.T) {
				got := tt.f(c)
				if got != want {
					t.Errorf("%s = %v, want %v", name, got, want)
				}
			})
		}
	}
}

func TestWhiteSpaceRunes(t *testing.T) {
	tests := testRunes
	for _, tt := range tests {
		for _, r := range SmallRuneSamples() {
			want := unicode.IsSpace(r) // standard library
			name := fmt.Sprintf("Whitespace check(%q): %s", r, tt.name)
			t.Run(name, func(t *testing.T) {
				got := tt.f(r)
				if got != want {
					t.Errorf("%s = %v, want %v", name, got, want)
				}
			})
		}
	}
}

func TestWhiteSpaceStringsByte(t *testing.T) {
	tests := testByteStrings
	for _, tt := range tests {
		for _, s := range SmallByteStringSamples() {
			// want := unicode.IsSpace(rune(s[0])) // standard library
			for _, r := range []byte(s) {
				want := tt.want(r)
				name := fmt.Sprintf("Testing %s(0x%x)", tt.name, r)
				t.Run(name, func(t *testing.T) {
					got := tt.f(s)
					if got != want {
						t.Errorf("%s = %t, want %t", name, got, want)
					}
				})
			}
		}
	}
}

func TestWhiteSpaceStringsRune(t *testing.T) {
	tests := testRuneStrings
	for _, tt := range tests {
		for _, s := range SmallRuneStringSamples() {
			for _, r := range s {
				// r, _ := utf8.DecodeRuneInString(s)
				want := unicode.IsSpace(r) // standard library "is this rune a whitespace rune"
				name := fmt.Sprintf("Whitespace check(%q): %s", r, tt.name)
				t.Run(name, func(t *testing.T) {
					got := tt.f(s)
					if got != want {
						t.Errorf("%s = %v, want %v", name, got, want)
					}
				})
			}
		}
	}
}

func TestDedupeWhitespace(t *testing.T) {
	tests := []struct {
		name           string
		arg            string
		want           string
		ignoreNewlines bool
	}{
		{"double space", "  ", " ", true},
		{"double space", "  ", " ", false},
		{"lots  \t \n \v  \r      of   whitespace", "lots  \t \n \v  \r      of   whitespace", "lots of whitespace", false},
		{"lots  \t \n \v  \r      of   whitespace", "lots  \t \n \v  \r      of   whitespace", "lots \nof whitespace", true},
		{"tab\ttab_space\t newline\nspaces     nothing", "tab\ttab_space\t newline\nspaces     nothing", "tab tab_space newline\nspaces nothing", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DedupeWhitespace(tt.arg, tt.ignoreNewlines); got != tt.want {
				t.Errorf("DedupeWhitespace() = %X, want %X", got, tt.want)
			}
		})
	}
}

/* Initial Benchmarks with complete set:

///Byte input (ASCII only)
BenchmarkByte/Benchmark:_isWhiteSpace2-8      	               7233 ns/op	     768 B/op	       2 allocs/op
BenchmarkByte/Benchmark:_IsASCIISpace-8       	               7706 ns/op	     768 B/op	       2 allocs/op
BenchmarkByte/Benchmark:_isWhiteSpace-8       	               7580 ns/op	     768 B/op	       2 allocs/op
BenchmarkByte/Benchmark:_isWhiteSpaceStringSliceBytes-8        8362 ns/op	     768 B/op	       2 allocs/op
BenchmarkByte/Benchmark:_isWhiteSpaceRegexByte-8              36780 ns/op	    1032 B/op	     256 allocs/op
///Notes: regex not recommended for simple tasks ... other methods are similar

///Rune Input (All Unicode whitespace characters)
*** (all required 1024 B/op	       1 allocs/op)
BenchmarkRunes/Benchmark:_isWhiteSpaceLogicChainNoCheck-8    7652 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteSpaceLogicChain-8           7592 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteSpaceBoolMap-8             11497 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_IsUnicodeWhiteSpaceMapSwitch-8    12605 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteRuneSlice-8                12436 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_IsUnicodeWhiteSpaceMap-8          12737 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteSpaceStringMap-8           12901 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteSpaceStringSlice-8         24725 ns/op	    1024 B/op	       1 allocs/op
/// Notes: byte operations are more efficient than rune operations.
/// No surprise (4 bytes compared to 1)
/// But ... they are not 4 times slower ...
/// In fact, the best performing rune operation (isWhiteSpaceLogicChain)
/// clocked in ~6% slower ...

/// Standard Library Unicode: unicode.IsSpace(r rune) bool
BenchmarkRunes/Benchmark:_(stlib)_unicode.IsSpace-8          11495 ns/op	    1024 B/op	       1 allocs/op

/// Byte Strings (none of these are suggested)
BenchmarkStringByte/Benchmark:_isWhiteSpaceContainsByte-8     20083 ns/op	   10224 B/op	     266 allocs/op
BenchmarkStringByte/Benchmark:_isWhiteSpaceIndexByte-8        20000 ns/op	   10224 B/op	     266 allocs/op
BenchmarkStringByte/Benchmark:_isWhiteSpaceIndexByte#01-8     20112 ns/op	   10224 B/op	     266 allocs/op
BenchmarkStringByte/Benchmark:_isWhiteSpaceTrim-8             36203 ns/op	   10224 B/op	     266 allocs/op

Rune Strings (none of these are suggested)
BenchmarkStringRune/Benchmark:_isWhiteSpaceIndexRune-8        20210 ns/op	   10224 B/op	     266 allocs/op
BenchmarkStringRune/Benchmark:_isWhiteSpaceContainsRune-8     19831 ns/op	   10224 B/op	     266 allocs/op
/// Notes: string operations (strings package) appear to be 2 to 3 times slower



ok  	github.com/skeptycal/util/stringutils	0.023s	coverage: 75.8% of statements
goos: darwin
goarch: amd64
pkg: github.com/skeptycal/util/stringutils
BenchmarkByte/Benchmark:_IsASCIISpace-8         	  159380	      8353 ns/op	     768 B/op	       2 allocs/op
BenchmarkByte/Benchmark:_isWhiteSpace-8         	  143262	      8199 ns/op	     768 B/op	       2 allocs/op
BenchmarkByte/Benchmark:_isWhiteSpace2-8        	  157026	      7368 ns/op	     768 B/op	       2 allocs/op
BenchmarkByte/Benchmark:_isWhiteSpaceStringSliceBytes-8         	  144788	      8946 ns/op	     768 B/op	       2 allocs/op
BenchmarkByte/Benchmark:_isWhiteSpaceRegexByte-8                	   32184	     38450 ns/op	    1031 B/op	     256 allocs/op

BenchmarkRunes/Benchmark:_isWhiteSpaceLogicChainNoCheck-8       	  157653	      7763 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteSpaceLogicChain-8              	  150936	      7789 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteRuneSlice-8                    	  108244	     11109 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteSpaceBoolMap-8                 	  108823	     11480 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_IsUnicodeWhiteSpaceMap-8              	   94245	     12600 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_IsUnicodeWhiteSpaceMapSwitch-8        	   94190	     12768 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteSpaceStringMap-8               	   95710	     12755 ns/op	    1024 B/op	       1 allocs/op
BenchmarkRunes/Benchmark:_isWhiteSpaceStringSlice-8             	   46706	     26900 ns/op	    1024 B/op	       1 allocs/op

BenchmarkRunes/Benchmark:_(stlib)_unicode.IsSpace-8             	  101884	     11734 ns/op	    1024 B/op	       1 allocs/op

BenchmarkStringByte/Benchmark:_isWhiteSpaceContainsByte-8       	   60166	     29206 ns/op	   10224 B/op	     266 allocs/op
BenchmarkStringByte/Benchmark:_isWhiteSpaceIndexByte-8          	   57104	     29050 ns/op	   10224 B/op	     266 allocs/op
BenchmarkStringByte/Benchmark:_isWhiteSpaceIndexByte#01-8       	   58418	     20268 ns/op	   10224 B/op	     266 allocs/op
BenchmarkStringByte/Benchmark:_isWhiteSpaceTrim-8               	   32604	     37100 ns/op	   10224 B/op	     266 allocs/op

BenchmarkStringRune/Benchmark:_isWhiteSpaceIndexRune-8          	   56760	     20971 ns/op	   10224 B/op	     266 allocs/op
BenchmarkStringRune/Benchmark:_isWhiteSpaceContainsRune-8       	   58694	     20418 ns/op	   10224 B/op	     266 allocs/op
*/
