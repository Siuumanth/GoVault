package multipart

import "upload/shared"

func calculateTotalParts(fileSize int64, partSize int64) int64 {
	return int64((fileSize + shared.ChunkSizeBytes - 1) / partSize)
}
