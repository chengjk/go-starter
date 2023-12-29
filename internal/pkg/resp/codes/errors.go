package codes

type BizErr struct {
	Code *BizCode
	Err  error
}

func (e BizErr) Error() string {
	return e.Err.Error()
	//return fmt.Sprintf("%s,error:%s", e.Code.String(), e.Err.Error())
}
