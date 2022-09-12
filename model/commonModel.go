package model

type FlowData struct {
	CtrlError
	Request  interface{}            `json:"request"`
	Response interface{}            `json:"response"`
	Data     map[string]interface{} `json:"data"`
}

type CtrlError struct {
	CtrlCode string `json:"ctlcode"`
	ServError
}

type ServError struct {
	ServCode string `json:"servcode"`
	FuncCode string `json:"funccode"`
	Err      error  `json:"err"`
	Msg      string `json:"msg"`
}

type ExternalResponse struct {
	ExternalErrorResponse
	Data interface{} `json:"data"`
}

type ExternalErrorResponse struct {
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

func (resp FlowData) ErrorCode() string {
	return resp.CtrlCode + "-" + resp.ServCode + "-" + resp.FuncCode
}

type Test struct {
	Num int `json:"num"`
}
