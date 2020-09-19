func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _, book := range bookings {
		res[book[0]-1] += book[2]
		if book[1] < n {
			res[book[1]] -= book[2]
		}
	}
	for i := 1; i < n; i++ {
		res[i] += res[i-1]
	}
	return res
}