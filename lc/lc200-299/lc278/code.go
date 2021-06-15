/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	i := 0
	j := n

	if ok := isBadVersion(i); ok {
		return i
	}

	for j != i+1 {
		if ok := isBadVersion((i + j) / 2); ok {
			j = (i + j) / 2
		} else {
			i = (i + j) / 2
		}
	}

	return j
}