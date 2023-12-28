package main

type strs struct {
	s1 string
	s2 string
	s3 string
}

func main() {
	st1 := &strs{s1: "abc", s2: "", s3: "xyz"}
	st2 := &strs{s1: "", s2: "def", s3: "zyx"}
	mod := false
	mergeItem(&st1.s1, &st2.s1, &mod)
	mergeItem(&st1.s2, &st2.s2, &mod)
	mergeItem(&st1.s3, &st2.s3, &mod)
}

// mergeItem merges a string with another string
// respecting if one is set or the other
// if both are set, we overwrite to with from
// and setting the modified flag to true if anything was changed
func mergeItem(to *string, from *string, modified *bool) {
	if *to == "" {
		if *from == "" {
			// both are blank
		} else {
			// to is blank, from is not
			*to = *from
			*modified = true
		}
	} else if *from == "" {
		// from is blank, to is not
	} else {
		// neither is blank; overwrite
		*to = *from
		*modified = true
	}
}
