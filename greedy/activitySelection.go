package greedy

func SessionSelection(start, end []int) []int {
	var session []int
	j := 0
	session = append(session, 0)
	for i := 1; i < len(start); i++ {
		if start[i] >= end[j] {
			session = append(session, i)
			j = i
		}
	}
	return session
}
