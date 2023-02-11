package router

import (
	"net/http"
	"github.com/ninja-dark/echelon_task/internal/infrastructure/api/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
	hs *handler.Hanlder
}

func NewRouter(hs *handler.Hanlder) *Router {
	r := gin.Default()
	ret := &Router{
		hs: hs,
	}
	r.POST("/api/v1/remote-execution", ret.ExecuteCommand)
	ret.Engine = r
	return ret
}

type Command struct {
	Cmd   string
	Os    string
	Stdin string
}

func (rt *Router) ExecuteCommand(c *gin.Context) {
	var f Command
	err := c.BindJSON(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	d, err := rt.hs.ExecuteCommand(c.Request.Context(), f.Cmd, f.Os, f.Stdin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, d)
}
