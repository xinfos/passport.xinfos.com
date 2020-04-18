package errs

//Errs - customization error
type Errs struct {
	ErrCode        int
	ErrMsg         string
	InternalErrMsg string
}

//NewErrs - return customization error strcut
func NewErrs(code int, errMsg ...string) *Errs {
	return &Errs{
		ErrCode: code,
		ErrMsg:  ErrorMsg[code],
	}
}
