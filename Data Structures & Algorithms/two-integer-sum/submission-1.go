func twoSum(nums []int, target int) []int {
    m1:=make(map[int]int)
	for i, v:= range nums{
		difference:=target-v
		if j,ok:=m1[difference];ok{
			if i<j{
				return []int{i, j}
			}else{
				return []int{j, i}
			}
		}else{
			m1[v]=i
		}
	}
	return []int{}
}
