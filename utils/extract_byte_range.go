package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ExtractByteRanges(byteRange string) (int, int, error) {
	var err error
	var start, end int

	byteRange = strings.ReplaceAll(byteRange, "\"", "")
	if !strings.Contains(byteRange, "bytes=") {
		return 0, 0, fmt.Errorf("invalid byte range format")
	}

	parts := strings.Split(byteRange, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid byte range format")
	}

	start, err = extractNumber(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid byte range format")
	}
	end, err = extractNumber(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid byte range format")
	}

	return start, end, nil
}

func extractNumber(str string) (int, error) {
	re := regexp.MustCompile("\\d+")
	matches := re.FindAllString(str, -1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("no number found")
	}
	num, err := strconv.Atoi(matches[0])
	return num, err
}
