package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{})

}
