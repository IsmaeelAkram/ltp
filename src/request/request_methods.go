package request

type MethodType int

const (
	REQUEST_INVALID    = 0
	REQUEST_DATETIME   = 1
	REQUEST_DATE       = 2
	REQUEST_TIME       = 3
	REQUEST_MONTH      = 4
	REQUEST_DAY        = 5
	REQUEST_YEAR       = 6
	REQUEST_HOUR       = 7
	REQUEST_MINUTE     = 8
	REQUEST_SECOND     = 9
	REQUEST_NANOSECOND = 10
)

func GetMethodType(x string) MethodType {
	switch x {
	case "DATETIME":
		return REQUEST_DATETIME
	case "DATE":
		return REQUEST_DATE
	case "TIME":
		return REQUEST_TIME
	case "MONTH":
		return REQUEST_MONTH
	case "DAY":
		return REQUEST_DAY
	case "YEAR":
		return REQUEST_YEAR
	case "HOUR":
		return REQUEST_HOUR
	case "MINUTE":
		return REQUEST_MINUTE
	case "SECOND":
		return REQUEST_SECOND
	case "NANOSECOND":
		return REQUEST_NANOSECOND
	default:
		return REQUEST_INVALID
	}
}
