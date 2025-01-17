package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/repo"
	srv "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/service/v1"
)

type Controller interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

type controller struct {
	srv srv.Service
}

func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
