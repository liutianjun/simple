package simple

type JsonResult struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func Json(code int, message string, data interface{}, success bool) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   success,
	}
}

func JsonData(data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
	}
}

func JsonPageData(results interface{}, page *Paging) *JsonResult {
	return JsonData(&PageResult{
		Results: results,
		Page:    page,
	})
}

func JsonSuccess() *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      nil,
		Success:   true,
	}
}

func JsonError(err *CodeError) *JsonResult {
	return &JsonResult{
		ErrorCode: err.Code,
		Message:   err.Message,
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorMsg(message string) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorCode(code int, message string) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

func JsonErrorData(code int, message string, data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   false,
	}
}

type RspBuilder struct {
	Data map[string]interface{}
}

func NewEmptyRspBuilder() *RspBuilder {
	return &RspBuilder{Data: make(map[string]interface{})}
}

func NewRspBuilder(obj interface{}) *RspBuilder {
	return NewRspBuilderExcludes(obj)
}

func NewRspBuilderExcludes(obj interface{}, excludes ...string) *RspBuilder {
	return &RspBuilder{Data: StructToMap(obj, excludes...)}
}

func (this *RspBuilder) Put(key string, value interface{}) *RspBuilder {
	this.Data[key] = value
	return this
}

func (this *RspBuilder) Build() map[string]interface{} {
	return this.Data
}

func (this *RspBuilder) JsonResult() *JsonResult {
	return JsonData(this.Data)
}
