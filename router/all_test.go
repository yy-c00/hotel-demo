package router

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var e = echo.New()

type RequestConfig struct {
	Name    string
	Params  string
	Body    interface{}
	Method  string
	Handler echo.HandlerFunc
	Code int
}

func ConfigTest(configs []RequestConfig, t *testing.T) {
	for _, v := range configs {
		t.Run(v.Name, func(st *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					st.Errorf("Recover (%v)", r)
				}
			}()

			b, err := json.Marshal(v.Body)
			if err != nil && v.Body != nil {
				st.Errorf("Ocurrs an unexpected error (%v)", err)
			}

			request := httptest.NewRequest(v.Method, "/?"+v.Params, bytes.NewBuffer(b))
			response := httptest.NewRecorder()

			request.Header.Add("Content-Type", "application/json")

			context := e.NewContext(request, response)

			err = v.Handler(context) //Aquí hay error

			if err != nil {
				st.Errorf("Ocurrs an unexpected error (%v)", err)
			}

			if response.Code != v.Code {
				b, _ = ioutil.ReadAll(response.Body)

				st.Errorf("Expected: %v Got: %v (%s)", v.Code, response.Code, b)
			}
		})
	}
}
