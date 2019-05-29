//@author Devansh Gupta
package vad

import (
	"net/url"
	"regexp"
)

//RequestValidator, validates a request
type RequestValidator struct {
	Params          url.Values
	RequriredFields []string
	ParamPatterns   map[string]string
	state           uint8
}

//AddFieldPattern, this adds a matching pattern against  a field values
func (r *RequestValidator) AddFieldPattern(field, pattern string) {
	if r.ParamPatterns != nil {

	} else {
		r.ParamPatterns = make(map[string]string)
	}
	r.ParamPatterns[field] = pattern
}

//ValidateAgainstPattern, this validates a field values against respective their values
func (r *RequestValidator) ValidateAgainstPattern() (bool, error) {
	for k, v := range r.ParamPatterns {
		for _, x := range r.Params[k] {
			if ok, err := regexp.MatchString(v, x); !ok {
				r.state = 0
				return ok, err
			}

		}
	}
	r.state = 1
	return true, nil
}

//HaveRequiredParams, this returns  nil if request contains all the requried params
//else InvalidInput
func (r *RequestValidator) HaveRequiredParams() error {
	for _, v := range r.RequriredFields {
		if _, ok := r.Params[v]; !ok {
			r.state = 0
			return InvalidInput{"required parameter not required"}
		}
	}
	r.state = 1
	return nil
}

//IsValid, returns true input is correct so far
func (r RequestValidator) IsValid() bool {
	return r.state == valid
}

const (
	valid = uint8(1)
	//Pattern for unsigned integer
	Pattern_UNINT  = "^\\d+$"
	Pattern_Email  = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	Pattern_UFloat = "^\\d+\\.?\\d*$"
	Pattern_Mobile = "^[0-9]{10}$"
)

type InvalidInput struct {
	Msg string
}

func (x InvalidInput) Error() string {
	return x.Msg
}
