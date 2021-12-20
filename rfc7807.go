package go_rfc7807

import (
	"encoding/json"
	"fmt"
)

type Rfc7807Error struct {
	Code          int                  `json:"code"`
	Type          string               `json:"type"`
	Title         string               `json:"title"`
	Detail        string               `json:"detail,omitempty"`
	Instance      string               `json:"instance,omitempty"`
	InvalidParams []invalidParamDetail `json:"invalid-params,omitempty"`
}

func (e *Rfc7807Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *Rfc7807Error) PutDetail(format string, a ...interface{}) *Rfc7807Error {
	e.Detail = fmt.Sprintf(format, a...)
	return e
}

func (e *Rfc7807Error) PutParam(name, reason string, a ...interface{}) *Rfc7807Error {
	e.InvalidParams = append(e.InvalidParams, invalidParamDetail{
		Name:   name,
		Reason: fmt.Sprintf(reason, a...),
	})
	return e
}

type invalidParamDetail struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}
