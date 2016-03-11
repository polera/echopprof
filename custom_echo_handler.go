package echopprof

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

type customEchoHandler struct {
	handlerFunc func(echo.Context) error
}

func (ceh *customEchoHandler) Handle(c echo.Context) error {
	return ceh.handlerFunc(c)
}

func fromHTTPHandler(httpHandler http.Handler) *customEchoHandler {
	return &customEchoHandler{
		handlerFunc: standard.WrapHandler(httpHandler),
	}
}

func fromHandlerFunc(serveHTTP func(w http.ResponseWriter, r *http.Request)) *customEchoHandler {
	return &customEchoHandler{
		handlerFunc: standard.WrapHandler(&customHTTPHandler{serveHTTP: serveHTTP}),
	}
}
