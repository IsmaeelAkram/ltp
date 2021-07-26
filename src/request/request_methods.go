package request

type MethodType int

const (
	REQUEST_INVALID     = 0
	REQUEST_DATETIME    = 1
	REQUEST_DATE        = 2
	REQUEST_TIME        = 3
	REQUEST_HOUR        = 4
	REQUEST_MINUTE      = 5
	REQUEST_SECOND      = 6
	REQUEST_MILLISECOND = 7
)

func GetMethodType(x string) MethodType {
	switch x {
	case "DATETIME":
		return REQUEST_DATETIME
	case "DATE":
		return REQUEST_DATE
	case "TIME":
		return REQUEST_TIME
	case "HOUR":
		return REQUEST_HOUR
	case "MINUTE":
		return REQUEST_MINUTE
	case "SECOND":
		return REQUEST_SECOND
	case "MILLISECOND":
		return REQUEST_MILLISECOND
	default:
		return REQUEST_INVALID
	}
}
