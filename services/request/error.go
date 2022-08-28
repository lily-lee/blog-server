package request

type BizErr struct {
	HttpCode int    `json:"-"`
	ErrCode  int    `json:"-"`
	ErrMsg   string `json:"err_msg"`
}

func (b *BizErr) Error() string {
	return b.ErrMsg
}
