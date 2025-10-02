package models

import "strconv"

type NetworkHeader struct {
	HeightStr string `json:"height"`
}

func (n NetworkHeader) Height() uint64 {
	// Convert HeightStr to uint64
	num, err := strconv.ParseUint(n.HeightStr, 10, 64)
	if err != nil {
		return 0
	}
	return num
}
