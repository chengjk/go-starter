package codes

type Code struct {
	Value   int
	Message string
}

func add(va int, msg string) Code {
	return Code{
		va,
		msg,
	}
}

var (
	OK            Code = add(0, "ok")
	TooManyReq         = add(429, "too many request")
	BadRequest    Code = add(400, "bad request")
	InternalError      = add(500, "internal error")
)
