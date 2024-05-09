package p0007

const MAX int32 = 0x7FFFFFFF
const MIN int32 = -MAX ^ 1

func Reverse(x int) int {
	x32 := int32(x)
	var y int32
	for x32 != 0 {
		aux := x32 % 10
		if (y > 0 && y > (MAX-aux)/10) || (y < 0 && y < (MIN-aux)/10) {
			return 0
		}
		y = 10*y + aux
		x32 /= 10
	}
	return int(y)
}
