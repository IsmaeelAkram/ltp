package util

import (
	"fmt"
	"time"

	"github.com/IsmaeelAkram/ltp/src/request"
)

func GetTime(method request.MethodType) string {
	now := time.Now()
	switch method {
	case request.REQUEST_DATETIME:
		return fmt.Sprintf("%d-%d-%d %d:%d:%d:%d", now.Month(), now.Day(), now.Year(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
	case request.REQUEST_DATE:
		return fmt.Sprintf("%d-%d-%d", now.Month(), now.Day(), now.Year())
	case request.REQUEST_TIME:
		return fmt.Sprintf("%d:%d:%d:%d", now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
	case request.REQUEST_MONTH:
		return fmt.Sprintf("%d", now.Month())
	case request.REQUEST_DAY:
		return fmt.Sprintf("%d", now.Day())
	case request.REQUEST_YEAR:
		return fmt.Sprintf("%d", now.Year())
	case request.REQUEST_HOUR:
		return fmt.Sprintf("%d", now.Hour())
	case request.REQUEST_MINUTE:
		return fmt.Sprintf("%d", now.Minute())
	case request.REQUEST_SECOND:
		return fmt.Sprintf("%d", now.Second())
	case request.REQUEST_NANOSECOND:
		return fmt.Sprintf("%d", now.Nanosecond())
	default:
		return ""
	}
}
