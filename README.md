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