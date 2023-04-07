package helper

import "strconv"

const (
	defaultPageNumber      = 0
	defaultLimitPageNumber = 8
)

// this function will return offset value and limit value
func ParseOffsetValue(pageRequest, limitRequest string) (offset int, limit int, page int) {
	page, err := strconv.Atoi(pageRequest)
	if err != nil {
		page = defaultPageNumber
	}

	limit, err = strconv.Atoi(limitRequest)
	if err != nil {
		limit = defaultLimitPageNumber
	}

	offset = (page - 1) * limit

	return
}
