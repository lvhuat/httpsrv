[![Build Status](https://www.travis-ci.org/lvhuat/httpsrv.svg?branch=master)](https://www.travis-ci.org/lvhuat/httpsrv)
[![codecov](https://codecov.io/gh/lvhuat/httpsrv/branch/master/graph/badge.svg)](https://codecov.io/gh/lvhuat/httpsrv)
[![Go Report Card](https://goreportcard.com/badge/github.com/lvhuat/httpsrv)](https://goreportcard.com/report/github.com/lvhuat/httpsrv)
[![GoDoc](https://godoc.org/github.com/lvhuat/httpsrv?status.svg)](https://godoc.org/github.com/lvhuat/httpsrv)
[![Release](https://img.shields.io/github/release-pre/lvhuat/httpsrv.svg?style=flat-square)](https://github.com/lvhuat/httpsrv/releases)

# httpsrv
A easy use wrapper of [Gin](https://github.com/gin-gonic/gin)

# Quick Start
```
wrapper := httpsrv.New(&httpsrv.Option{})
wrapper.HandleStat()
wrapper.HandlePprof()
wrapper.Get("/v1/hello", func(ctx *gin.Context) (interface{}, code.Error) {
    return []string{"Hello", "World"}, nil
})
wrapper.Get("/v1/error", func(ctx *gin.Context) (interface{}, code.Error) {
    return nil, code.NewMcode("ERROR", "test error")
})
wrapper.Run(":8080")
```

```
$ curl 127.0.0.1:8080/v1/hello
data":["Hello","World"],"result":true,"timestamp":1550259190025}
$ curl 127.0.0.1:8080/v1/error
{"mcode":"ERROR","message":"test error","result":false,"timestamp":1550259260240}
```