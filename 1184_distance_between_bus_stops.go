func distanceBetweenBusStops(distance []int, start int, destination int) int {
	cw_dist, ccw_dist, n := 0, 0, len(distance)
	if start > destination {
		start, destination = destination, start
	}
	for i := start; i < destination; i++ {
		cw_dist += distance[i]
	}
	for i := destination; i < start+n; i++ {
		ccw_dist += distance[i%n]
	}
	if cw_dist < ccw_dist {
		return cw_dist
	}
	return ccw_dist
}