package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int
	for _, birdsInDay := range birdsPerDay {
		totalBirds += birdsInDay
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	daysInWeek := 7
	startingIndex := (week - 1) * daysInWeek
	birdsInWeek := birdsPerDay[startingIndex : startingIndex+daysInWeek]
	return TotalBirdCount(birdsInWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx, _ := range birdsPerDay {
		if idx%2 == 0 {
			birdsPerDay[idx] += 1
		}
	}
	return birdsPerDay
}
