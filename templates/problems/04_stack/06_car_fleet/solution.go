package car_fleet

import "sort"

// solveCarFleet counts car fleets by sorting by position and using a stack of arrival times. O(n log n) time, O(n) space.
func solveCarFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}
	type car struct {
		pos   int
		speed int
	}
	cars := make([]car, n)
	for i := 0; i < n; i++ {
		cars[i] = car{position[i], speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})
	fleets := 0
	var lastTime float64
	for _, c := range cars {
		time := float64(target-c.pos) / float64(c.speed)
		if time > lastTime {
			fleets++
			lastTime = time
		}
	}
	return fleets
}
