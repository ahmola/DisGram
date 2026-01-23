package common

import (
	"strconv"
)

func ConvertStringToUint(str string) (uint, error) {
	convID, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(convID), nil
}
