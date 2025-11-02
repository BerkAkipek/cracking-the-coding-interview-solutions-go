package drawline

/*
Draw Line: A monochrome screen is stored as a single array of bytes, allowing eight consecutive
pixels to be stored in one byte. The screen has width w, where w is divisible by 8 (that is, no byte will
be split across rows). The height of the screen, of course, can be derived from the length of the array
and the width. Implement a function that draws a horizontal line from ( xl, y) to ( x2, y).
The method signature should look something like:
drawline(byte[] screen, int width, int xl, int x2, int y)


mask = ((0xFF >> start_bit) & (0xFF << (7 - end_bit)))


start_mask := 0xFF >> (x1 % 8)
end_mask   := ^(0xFF >> ((x2 % 8) + 1))
*/

func DrawLine(screen []byte, width, x1, x2, y int) {
	bytesPerRow := width / 8
	rowStart := y * bytesPerRow

	startByte := x1 >> 3
	endByte := x2 >> 3

	startOffset := x1 & 7
	endOffset := x2 & 7

	if startByte == endByte {
		// Both points in same byte
		mask := byte((0xFF >> startOffset) & ^(0xFF >> (endOffset + 1)))
		screen[rowStart+startByte] |= mask
		return
	}

	// Left partial byte: fill from startOffset to end of byte
	startMask := byte(0xFF >> startOffset)

	// Right partial byte: fill from start of byte up to endOffset
	endMask := ^byte(0xFF >> (endOffset + 1))

	// Apply masks
	screen[rowStart+startByte] |= startMask
	screen[rowStart+endByte] |= endMask

	// Fill all full bytes between
	for i := startByte + 1; i < endByte; i++ {
		screen[rowStart+i] = 0xFF
	}
}
