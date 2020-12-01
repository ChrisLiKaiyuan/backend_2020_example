package models

type SuccessReturn struct {
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
	//cache bool
	//code string
}

type MakeErrorReturn struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
