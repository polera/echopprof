package echopprof

import (
	"github.com/labstack/echo"
	"net/http/pprof"
)

func Wrap(e *echo.Echo) {
	e.Get("/debug/pprof/", fromHandlerFunc(pprof.Index))
	e.Get("/debug/pprof/heap", fromHTTPHandler(pprof.Handler("heap")))
	e.Get("/debug/pprof/goroutine", fromHTTPHandler(pprof.Handler("goroutine")))
	e.Get("/debug/pprof/block", fromHTTPHandler(pprof.Handler("block")))
	e.Get("/debug/pprof/threadcreate", fromHTTPHandler(pprof.Handler("threadcreate")))
	e.Get("/debug/pprof/cmdline", fromHandlerFunc(pprof.Cmdline))
	e.Get("/debug/pprof/profile", fromHandlerFunc(pprof.Profile))
	e.Get("/debug/pprof/symbol", fromHandlerFunc(pprof.Symbol))
}

var Wrapper = Wrap
