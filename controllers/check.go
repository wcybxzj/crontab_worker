package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Check api
func Check(c *gin.Context) {
	c.Data(http.StatusOK, "", nil)
}
