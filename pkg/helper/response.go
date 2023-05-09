package helper

type Response struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func NewResponse() Response {
	return Response{}
}

func (r *Response) Success(status string, data interface{}) Response {
	return Response{
		Status: status,
		Data:   data,
	}
}

func (r *Response) Error(status string, message string, error interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  error,
	}
}

var ResponseJSON = Response{}
