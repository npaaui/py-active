package v1

import (
	"github.com/gin-gonic/gin"
	"active/common"
)

func index(c *gin.Context) {
	common.ReturnData(c, "this is index")
}
