// Package ginrest objetive is to have in one place ours
// rest operations, as standart rest IO with norm output
// Author: Rolando Lucio <rolando@compropago.com,rolandolucio@gmail.com>
package ginrest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// Error general errors response
type Error struct {
	Code int
	Msg  string
}

// Error Interface implementations
func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}
func (e Error) errorNumber() int {
	return e.Code
}

// Payload map for multiple inputs
type Payload map[string]interface{}

// IO struct is the base app/json output formatter
// according to our defined schema
type IO struct {
	res *response
	Gin *gin.Context
}

// response base object
// To OPTIMIZE struct size order elements by bits size
type response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
	Request int64  `json:"request"`
	Object  string `json:"object"`
}

// New initializer
func New(path, object string) *IO {
	io := new(IO)
	io.res = new(response)
	io.res.Status = "error"
	io.res.Code = 500
	io.res.Message = http.StatusText(io.res.Code)
	if path != "" {
		io.res.Path = path
	}
	if object == "" {
		io.res.Object = "object.undefined"
	} else {
		io.res.Object = object
	}

	return io
}

// SetGin method implementation
func (s *IO) SetGin(c *gin.Context) *IO {
	s.Gin = c
	return s
}

// Res Rest Response According to proper type
// params: code proper httpStatusCode
// elements map of structs(interfaces) to append to the json output
// msg if want to override http default message for status code
//
// if elements contains a key defined on response struct it will get
// overrided, in case code is override will not change the response code
//
// The method will eval which http reference is set for the output
// like *gin.Context, if not found it will panic
//
// References,inspirations :
// Eduardo Aguilar's Generator
// https://kev.inburke.com/kevin/golang-json-http/
// http://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/
// http://attilaolah.eu/2013/11/29/json-decoding-in-go/
// http://eagain.net/articles/go-dynamic-json/
// http://gregtrowbridge.com/golang-json-serialization-with-interfaces/
func (s *IO) Res(code int, elements Payload, msg string) *IO {
	// Set defaults to code
	s.httpCodes(code)
	if msg != "" {
		s.res.Message = msg
	}
	// struct to map  s.res for equivalent management to elements
	m := structs.Map(s.res)
	f := make(Payload)

	for k, v := range m {
		l := strings.ToLower(k)
		f[l] = v
	}
	if elements != nil {
		for k, v := range elements {
			f[k] = v
		}
	}
	// first enconde to bytes so could manage easily
	j, err := json.Marshal(f)
	if err != nil {
		//trying to avoid inecesary panics
		log.Println("f marshal err, sending 500", err)
		return s.Res(500, nil, "")
	}

	// as we dont know a proper struct unmarshall to struct
	// some http handlers may just need the bytes , rlly? chek it later
	var parsed interface{}
	e := json.Unmarshal(j, &parsed)

	if e != nil {
		//trying to avoid inecesary panics
		log.Println("j unmarshal err, sending 500", e)
		return s.Res(500, nil, "")
	}

	// Pick the proper selector and output
	switch {
	case s.Gin != nil:
		//s.Gin.JSON(code, parsed)
		s.Gin.JSON(code, parsed)
	default:
		panic("No http reference defined")
		//return s
	}
	return s
}

// httpCodes sets defaults http Messages to IANA
// https://golang.org/src/net/http/status.go
// Inspired in Vhuerta JS solution IO middleware
func (s *IO) httpCodes(code int) *IO {
	if code <= 0 {
		code = 500
	}

	if code >= 200 && code < 300 {
		s.res.Status = "success"
	}

	s.res.Code = code
	s.res.Message = http.StatusText(code)
	s.res.Request = int64(time.Now().UnixNano()) / 100000

	return s
}
