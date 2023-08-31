package main

func ReverseBits(oct byte) byte {
	var res byte

	for i := 0; i <= 8; i++ {
		res = res<<1 + oct&1
		oct = oct >> 1
	}

	return res
}
