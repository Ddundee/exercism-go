package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for i := 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
		
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0
	start := (week - 1) * 7
	end := start + 7

	if len(birdsPerDay) < start {
		return 0
	}
	if len(birdsPerDay) < end {
		end = len(birdsPerDay)
	}

	for i := start; i < end; i++ {
		total += birdsPerDay[i]
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
