package easy

func LongestCommonPrefix(strs []string) string {

	//find shortest item
	var shortest = strs[0]
	for _, e := range strs {
		if len(e) < len(shortest) {
			shortest = e
		}
	}

	//use entire shortest word as prefix, if this doesnt work, remove 1 letter from it and try again.
	for len(shortest) > 0 {
		ok := true
		for _, st := range strs {
			if st[:len(shortest)] != shortest {
				shortest = shortest[:len(shortest)-1]
				ok = false
			}
		}

		if ok {
			break
		}
	}

	return shortest
}
