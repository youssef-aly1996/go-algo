package examples

/*
* This algorithim depends on four main factors:
t: type of each seat wether its an asile or window or middle
g: group of single seat type
r: row of each group of seats
c: colums of each group of seats

* Sorting is applied on the criteria of filling seats with respect to seat type asile or window or middle
*  This algorithim is O(N logn + m)
*/

import (
	"fmt"
	"sort"
)
 
type seat struct {
	r, c, g, t int
}
 
func (s seat) Less(rhs seat) bool {
	t1 := []int{s.t, s.r, s.g, s.c}
	t2 := []int{rhs.t, rhs.r, rhs.g, rhs.c}
	for i := range t1 {
		if t1[i] < t2[i] {
			return true
		} else if t1[i] > t2[i] {
			return false
		}
	}
	return false
}
 
func main() {
	var n int
	fmt.Scanf("%d", &n)
	seatGroups := make([]struct{ x, y int }, n)
	for i := range seatGroups {
		fmt.Scanf("%d %d", &seatGroups[i].y, &seatGroups[i].x)
	}
 
	seats := make([]seat, 0, n)
	for i, sg := range seatGroups {
		x, y := sg.x, sg.y
		for r := 1; r <= x; r++ {
			for c := 1; c <= y; c++ {
				curType := -1
				if (i == 0 && c == 1) || (i == n-1 && c == y) {
					curType = 2
				} else if c == 1 || c == y {
					curType = 1
				} else {
					curType = 3
				}
				seats = append(seats, seat{r, c, i, curType})
			}
		}
 
	}
 
	var m int
	fmt.Scanf("%d", &m)
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].Less(seats[j])
	})
	for i := 0; i < m; i++ {
		s := seats[i]
		fmt.Printf("%d, %d, %d, %d\n", s.r, s.c, s.g, s.t)
	}
}
 