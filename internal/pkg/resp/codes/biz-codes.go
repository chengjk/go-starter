package codes

import (
	"errors"
	"fmt"
	"strconv"
)

var SystemBizCode = make(map[int]*BizCode, 16)
var (
	OK            BizCode = addCode(10000, "OK")
	TooManyReq            = addCode(429, "too many request")
	BadRequest            = addCode(400, "bad request")
	InvalidUser           = addCode(401, "Invalid User.")
	IllegalArgs           = addCode(40001, "Illegal Args.")
	IllegalState          = addCode(40002, "Illegal State.")
	AlreadyExist          = addCode(40003, "Object Already Exist")
	EmptyState            = addCode(40004, "Empty Result")
	NoPermission          = addCode(40100, "No Permission")
	NotExist              = addCode(40400, "Object Not Exist")
	InternalError         = addCode(50000, "Server Internal Error")
	Immutable             = addCode(50001, "Immutable Object, Can Not Update. ")
)

func addCode(code int, name string) BizCode {
	if SystemBizCode[code] != nil {
		panic("found duplicate biz code:" + strconv.Itoa(code))
	}
	newCode := BizCode{
		Code: code,
		Name: name,
	}
	SystemBizCode[code] = &newCode
	return newCode
}

type IBizCode interface {
	Name() string
	Code() int
}
type BizCode struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

func (c BizCode) String() string {
	return fmt.Sprintf("code:%d,name:%s", c.Code, c.Name)
}

func (c BizCode) Error(err error) BizErr {
	return BizErr{
		Err:  err,
		Code: &c,
	}
}
func (c BizCode) ErrorStr(msg string) BizErr {
	return BizErr{
		Err:  errors.New(msg),
		Code: &c,
	}
}

// ErrorStrF formats as biz error
func (c BizCode) ErrorStrF(tpl string, args ...any) BizErr {
	return BizErr{
		Err:  errors.New(fmt.Sprintf(tpl, args)),
		Code: &c,
	}
}
