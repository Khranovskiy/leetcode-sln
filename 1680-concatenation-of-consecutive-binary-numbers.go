package main

func main() {
	concatenatedBinary(4)

}

func concatenatedBinary(n int) int {
	var (
		modulo int64 = 1000000007
		res    int64 = 0
		bits   int64 = 1
		next   int64 = 2
	)
	for i := int64(1); i <= int64(n); i++ {
		if i == next {
			bits += 1
			next = next * 2
		}
		res = ((res<<bits)%modulo + i) % modulo
	}
	return int(res)
}
