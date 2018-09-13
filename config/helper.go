package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//{"code":"0","msg":"\u6210\u529f","status":"ok","data":[]}
func Output(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		//"data": gin.H{
		//	"addJob": gin.H{"name": "crontab_worker"},
		//	"runJob": gin.H{"name": "crontab_worker"},
		//}
	})
}
