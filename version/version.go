package version

import (
	"sort"
)

/*
TODO: create parse func
TODO: create compare
*/

// VersionRegex
/*
As defined in https://projects.gentoo.org/pms/8/pms.html#x1-250003.2

[0-9]+(\.[0-9]+)*
		uint followed by 0+ uints prefixed by '.'
[a-z]?
		MAY be followed by 1 lowercase letter (OpenSSL to blame)
(_(alpha|beta|pre|rc|p)[0-9]*)*
		MAY be followed by 0+ of _alpha, _beta, _pre, _rc or _p,
		each of which MAY optionally be followed by an uint.
		Suffix and integer count as separate version components.
(-r[0-9]*)?
		MAY optionally be followed by the suffix -r
		followed immediately by 1 uint (the “revision number”).
		If this suffix is not present, it is assumed to be -r0.￥
*/
const VersionRegex = `[0-9]+(\.[0-9]+)*[a-z]?(_(alpha|beta|pre|rc|p)[0-9]*)*(-r[0-9]*)?`

type suffix struct {
	txt string
	num string
}
type PMSVersion struct {
	number   []string
	letter   string
	suffix   []suffix
	revision suffix
}

type ByVersion []string

func (vs ByVersion) Len() int           { return len(vs) }
func (vs ByVersion) Swap(i, j int)      { vs[i], vs[j] = vs[j], vs[i] }
func (vs ByVersion) Less(i, j int) bool { return false } // TODO

func Sort(list []string) {
	sort.Sort(ByVersion(list))
}

func parseInts(v string) (nums []string, rest string, ok bool) {
	if v == "" {
		return
	}
	rest = v
	for rest != "" && rest[0] != '_' && rest[0] != '-' {
		if rest[0] == '.' {
			rest = rest[1:]
		}
		var num string
		num, rest, ok = parseInt(rest)
		nums = append(nums, num)
	}
	if len(rest) == 0 || rest[0] == '_' || rest[0] == '-' {
		ok = true
	}
	return nums, rest, ok
}
func parseInt(v string) (num, rest string, ok bool) {
	if v == "" {
		return
	}
	i := 0
	for i < len(v) && isNumber(v[i]) {
		i++
	}
	/* Check PMS if it allows it or not
	if v[0] == '0' && i != 1 { // 0 can be only if it's the only digit
		return
	}
	*/
	return v[:i], v[i:], true
}

func parseLetter(v string) (letter, rest string, ok bool) {
	if v == "" {
		return
	}
	if 'a' <= v[0] && v[0] <= 'z' {
		return v[:1], v[1:], true
	}
	return
}

func parseSuffix(v string) (s suffix, rest string, ok bool) {
	// suffixes := [...]string{"alpha", "beta", "pre", "rc", "p"}
	// _ = suffixes TODO: instead just compare to viable suffixes
	if v == "" || v[0] != '_' {
		return
	}
	// Will find prefixes without num
	i := 1
	for i < len(v) && v[i] != '_' && v[i] != '-' && !isNumber(v[i]) {
		i++
	}
	s.txt = v[1:i]
	s.num, rest, _ = parseInt(v[i:])
	return s, rest, true
}
func parseSuffixes(v string) (s []suffix, rest string, ok bool) {
	if v == "" {
		return
	}
	rest = v
	for len(rest) > 0 && rest[0] != '-' {
		var sx suffix
		sx, rest, ok = parseSuffix(rest)
		print("suffix found: ", sx.txt, sx.num, "\n")
		s = append(s, sx)
	}

	if rest == "" || rest[0] == '-' {
		ok = true
	}
	return s, rest, ok
}

func Compare(v, w PMSVersion) {

}

func isNumber(r uint8) bool {
	if '0' <= r && r <= '9' {
		return true
	}
	return false
}
