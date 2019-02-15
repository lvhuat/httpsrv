# httpsrv


# 使用方法
```

	wrapper := httpsrv.New(&httpsrv.Option{
		Prefix: "USER",
	})
	//rand.Seed(time.Now().UnixNano())
	wrapper.HandleStat()
	wrapper.HandlePprof()
	wrapper.Get("/v1/hello", func(ctx *gin.Context) (interface{}, code.Error) {
		return []string{"Hello", "World"}, nil
	})

	wrapper.Get("/v1/error", func(ctx *gin.Context) (interface{}, code.Error) {
		return nil, code.NewMcode("ERROR", "test error")
	})

	go func() {
		for {
			time.Sleep(time.Second * 20)
			httpstat.Reset()
		}
	}()

	wrapper.Run(":8080")
```