package xor

import "strings"

const XOR_KEY_111 = 0xAC

func GetXORKey121(fileName string) int {
	fileName = strings.ToLower(fileName)
	var ind int
	for i := 0; i < len(fileName); i++ {
		ind += int(fileName[i])
	}
	return ind & 0xff
}
