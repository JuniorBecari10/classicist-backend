package query

import (
	"strings"
	"unicode"
)

// normalize converts accented Latin letters to their ASCII equivalents
func normalize(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if mapped := removeDiacritics(r); mapped != 0 {
			sb.WriteRune(mapped)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// removeDiacritics maps all Latin letters with accents to ASCII
func removeDiacritics(r rune) rune {
	switch r {
		// a variants
		case 'à', 'á', 'â', 'ã', 'ä', 'å', 'ā', 'ă', 'ą', 'ǎ', 'ǟ', 'ȁ', 'ȃ', 'ȧ', 'ǻ':
			return 'a'
		case 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Ā', 'Ă', 'Ą', 'Ǎ', 'Ǟ', 'Ȁ', 'Ȃ', 'Ȧ', 'Ǻ':
			return 'A'
		// c variants
		case 'ç', 'ć', 'ĉ', 'č', 'ċ':
			return 'c'
		case 'Ç', 'Ć', 'Ĉ', 'Č', 'Ċ':
			return 'C'
		// d variants
		case 'ď', 'đ':
			return 'd'
		case 'Ď', 'Đ':
			return 'D'
		// e variants
		case 'è', 'é', 'ê', 'ë', 'ē', 'ĕ', 'ė', 'ę', 'ě', 'ȅ', 'ȇ', 'ǝ':
			return 'e'
		case 'È', 'É', 'Ê', 'Ë', 'Ē', 'Ĕ', 'Ė', 'Ę', 'Ě', 'Ȅ', 'Ȇ', 'Ǝ':
			return 'E'
		// f variants
		case 'ƒ':
			return 'f'
		case 'Ƒ':
			return 'F'
		// g variants
		case 'ĝ', 'ğ', 'ġ', 'ģ':
			return 'g'
		case 'Ĝ', 'Ğ', 'Ġ', 'Ģ':
			return 'G'
		// h variants
		case 'ĥ', 'ħ':
			return 'h'
		case 'Ĥ', 'Ħ':
			return 'H'
		// i variants
		case 'ì', 'í', 'î', 'ï', 'ī', 'ĭ', 'į', 'ı', 'ǐ':
			return 'i'
		case 'Ì', 'Í', 'Î', 'Ï', 'Ī', 'Ĭ', 'Į', 'I', 'Ǐ':
			return 'I'
		// j variants
		case 'ĵ':
			return 'j'
		case 'Ĵ':
			return 'J'
		// k variants
		case 'ķ':
			return 'k'
		case 'Ķ':
			return 'K'
		// l variants
		case 'ĺ', 'ļ', 'ľ', 'ŀ', 'ł':
			return 'l'
		case 'Ĺ', 'Ļ', 'Ľ', 'Ŀ', 'Ł':
			return 'L'
		// m variants
		case 'ḿ':
			return 'm'
		case 'Ḿ':
			return 'M'
		// n variants
		case 'ñ', 'ń', 'ņ', 'ň', 'ŉ', 'ŋ':
			return 'n'
		case 'Ñ', 'Ń', 'Ņ', 'Ň', 'Ŋ':
			return 'N'
		// o variants
		case 'ò', 'ó', 'ô', 'õ', 'ö', 'ø', 'ō', 'ŏ', 'ő', 'ǒ', 'ǫ', 'ǭ', 'ȍ', 'ȏ', 'ɵ':
			return 'o'
		case 'Ò', 'Ó', 'Ô', 'Õ', 'Ö', 'Ø', 'Ō', 'Ŏ', 'Ő', 'Ǒ', 'Ǫ', 'Ǭ', 'Ȍ', 'Ȏ', 'Ɵ':
			return 'O'
		// r variants
		case 'ŕ', 'ŗ', 'ř':
			return 'r'
		case 'Ŕ', 'Ŗ', 'Ř':
			return 'R'
		// s variants
		case 'ś', 'ŝ', 'ş', 'š', 'ș':
			return 's'
		case 'Ś', 'Ŝ', 'Ş', 'Š', 'Ș':
			return 'S'
		// t variants
		case 'ţ', 'ť', 'ŧ', 'ț':
			return 't'
		case 'Ţ', 'Ť', 'Ŧ', 'Ț':
			return 'T'
		// u variants
		case 'ù', 'ú', 'û', 'ü', 'ū', 'ŭ', 'ů', 'ű', 'ų', 'ǔ', 'ǖ', 'ǘ', 'ǚ', 'ǜ':
			return 'u'
		case 'Ù', 'Ú', 'Û', 'Ü', 'Ū', 'Ŭ', 'Ů', 'Ű', 'Ų', 'Ǔ', 'Ǖ', 'Ǘ', 'Ǚ', 'Ǜ':
			return 'U'
		// w variants
		case 'ŵ':
			return 'w'
		case 'Ŵ':
			return 'W'
		// y variants
		case 'ÿ', 'ý', 'ŷ':
			return 'y'
		case 'Ÿ', 'Ý', 'Ŷ':
			return 'Y'
		// z variants
		case 'ź', 'ż', 'ž':
			return 'z'
		case 'Ź', 'Ż', 'Ž':
			return 'Z'
		// special
		case 'ß':
			return 's'
		case 'ð':
			return 'd'
		case 'Ð':
			return 'D'
		case 'þ':
			return 't'
		case 'Þ':
			return 'T'
		default:
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				return r
			}
			return r
		}
}
