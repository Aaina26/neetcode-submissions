func isAnagram(s string, t string) bool {
	m1:=make(map[rune]int)
	m2:=make(map[rune]int)

	for _, r:= range s{
		m1[r]++
	}
	for _, r:= range t{
		m2[r]++
	}
	if len(m1)==len(m2){
		for key, _:= range m1{
			if m1[key]!=m2[key]{
				return false
			}
		}
	}else{
		return false
	}
	return true
}
