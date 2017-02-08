package mtree

import "../common"

func conventionalWord(data common.Word) ([]byte, []byte) {
	first := rev(data[:32])
	first = append(first, rev(data[32:64])...)
	second := rev(data[64:96])
	second = append(second, rev(data[96:128])...)
	return first, second
}

func rev(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
