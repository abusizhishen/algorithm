package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(s) {
		return false
	}

	var smap, tmap = map[uint8]uint8{}, map[uint8]uint8{}
	for i := 0; i < len(s); i++ {
		_, oks := smap[s[i]]
		_, okt := tmap[t[i]]

		if oks != okt {
			return false
		}

		if !oks {
			smap[s[i]] = t[i]
			tmap[t[i]] = s[i]
			continue
		}

		if smap[s[i]] != t[i] || tmap[t[i]] != s[i] {
			return false
		}
	}

	return true
}
