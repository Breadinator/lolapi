package errors

var Statuses map[int]string = map[int]string{
	200: "OK",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	415: "Unsupported Media Type",
	429: "Rate Limit Exceeded",
	500: "Internal Server Error",
	503: "Service Unavailable",
}
