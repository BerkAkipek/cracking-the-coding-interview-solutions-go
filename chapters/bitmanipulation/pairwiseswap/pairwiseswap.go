package pairwiseswap

/*
Pairwise Swap: Write a program to swap odd and even bits in an integer with as few instructions as
possible (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, and so on).
*/

func Pairwise(n uint64) uint64 {
	const evenMask uint64 = 0x5555555555555555 // 0101...
	const oddMask uint64 = 0xAAAAAAAAAAAAAAAA  // 1010...
	even := n & evenMask
	odd := n & oddMask
	return (even << 1) | (odd >> 1)
}
