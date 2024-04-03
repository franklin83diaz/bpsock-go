package utils

// SplitData splits the data into chunks of dmtu size
func SplitData(data []byte, dmtu int) [][]byte {
	var listData [][]byte

	for i := 0; i < len(data); i += dmtu {
		end := i + dmtu
		if end > len(data) {
			end = len(data)
		}

		chunk := data[i:end]
		listData = append(listData, chunk)
	}

	return listData
}
