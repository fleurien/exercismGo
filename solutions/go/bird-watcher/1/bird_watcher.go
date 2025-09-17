package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0 
	for i := 0; i < len(birdsPerDay); i++{
		total = total + birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekStart := (week - 1) * 7
	weekEnd := weekStart + 6
	newSlice := birdsPerDay[weekStart:weekEnd + 1]
	total := 0
	for i := 0; i < 7; i++{
		total = total + newSlice[i]
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	birdsPerDay[0] = birdsPerDay[0] + 1
	for i:= 2; i < len(birdsPerDay); i += 2{
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}
