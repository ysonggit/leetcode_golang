func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _, book := range bookings {
		for i := book[0]; i <= book[1]; i++ {
			res[i-1] += book[2]
		}
	}
	return res
}