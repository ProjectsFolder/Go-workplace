package dto

import (
    "errors"
)

type ResponseInterface interface {
    GetSuccess() interface{}
    GetMessage() string
}

type Response struct {
    Success interface{} `json:"success"`
    Message string `json:"message"`
}

func (r *Response) GetSuccess() interface{} {
    return r.Success
}

func (r *Response) GetMessage() string {
    return r.Message
}

func (r *Response) CheckResponse() error {
    success := r.GetSuccess()
    if success == false || success == float64(0) {
        err := errors.New(r.GetMessage())

        return err
    }

    return nil
}
