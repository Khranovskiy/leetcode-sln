func guessNumber(n int) int {
	l, r := 0, n
	for l <= r {
		mid := l + (r-l)/2
		g := guess(mid)
		if g == 0 {
			return mid
		} else if g == 1 { // mid is lower than the picked
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	 -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *           otherwise return 0
 * func guess(num int) int;
 */

