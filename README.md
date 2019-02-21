[![Build Status](https://travis-ci.org/lvhuat/httpsrv.svg?branch=master)](https://travis-ci.org/lvhuat/httpsrv)
[![codecov](https://codecov.io/gh/lvhuat/httpsrv/branch/master/graph/badge.svg)](https://codecov.io/gh/lvhuat/httpsrv)

# httpsrv


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

# Support Handler Functions

+ func(ginCtx *gin.Context) (interface{}, code.Error)

+ func(ginCtx *gin.Context) (interface{}, error)

+ func(ginCtx *gin.Context) error

+ func(ginCtx *gin.Context) code.Error

+ func(ginCtx *gin.Context)  

+ func(ginCtx *gin.Context) (httpsrv.NoWrapperResponse, int)

+ func(ginCtx *gin.Context) httpsrv.NoWrapperResponse
