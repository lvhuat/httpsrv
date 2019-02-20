package httpsrv

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroup(t *testing.T) {
	wrapper := New(nil)
	v1 := wrapper.Group("/v1")
	handler := func(*gin.Context) (interface{}, error) {
		return nil, nil
	}
	v1.Get("/user", handler)
	v1.Post("/user", handler)
	v1.Patch("/user", handler)
	v1.Put("/user", handler)
	v1.Options("/user", handler)
	v1.Head("/user", handler)
	v1.Delete("/user", handler)
	v1.Any("/ping", handler)
	v1.HandlePprof()
	v1.HandleStat()

	profile := v1.Group("profile")
	profile.Get("info", handler)

	assert.Equal(t, "/v1", v1.BasePath())
}
