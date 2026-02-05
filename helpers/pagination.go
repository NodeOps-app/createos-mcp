package mcputils

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

func DecodeCursor(cursor string) (int, error) {
	if cursor == "" {
		return 0, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, fmt.Errorf("invalid cursor: %w", err)
	}

	offset, err := strconv.Atoi(string(decoded))
	if err != nil {
		return 0, fmt.Errorf("invalid cursor offset: %w", err)
	}

	if offset < 0 {
		return 0, fmt.Errorf("invalid cursor offset: must be non-negative")
	}

	return offset, nil
}

func EncodeCursor(offset int) string {
	if offset <= 0 {
		return ""
	}

	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(offset)))
}
