package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractByteRanges(byteRange string) (int, int, error) {
	var err error
	var start, end int

	if !strings.HasPrefix(byteRange, "bytes=") {
		return 0, 0, fmt.Errorf("invalid byte range format")
	}

	byteRange = strings.TrimPrefix(byteRange, "bytes=")

	parts := strings.Split(byteRange, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid byte range format")
	}

	start, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid byte range format")
	}
	end, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid byte range format")
	}

	return start, end, nil
}
