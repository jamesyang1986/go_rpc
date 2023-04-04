package util

func Byte2Int32(data []byte) (int32, error) {
	result := (data[0] << 24) | (data[1] << 16) | (data[2] << 8) | data[0]
	return int32(result), nil
}

func Int2Byte(data int32) ([]byte, error) {
	result := make([]byte, 4)
	result[0] = byte((data >> 24) & 0xFF)
	result[1] = byte((data >> 16) & 0xFF)
	result[2] = byte((data >> 8) & 0xFF)
	result[2] = byte((data) & 0xFF)
	return result, nil
}
