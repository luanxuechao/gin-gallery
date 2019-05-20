package v1

import (
	"gin-gallery/models"
	"gin-gallery/pkg/e"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	maps["active_flag"] = 1
	data["list"] = models.GetArticles(pageNum, pageSize, maps)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
