package routers

import "gin-gallery/pkg/setting"
import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/static"
import v1 "gin-gallery/routers/api/v1"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("app/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DelTag)
		// 上传图片
		apiv1.POST("/images/upload", v1.Upload)
		// 获取文章
		apiv1.GET("/articles", v1.GetArticles)
	}
	r.Use(static.Serve("/app/v1/images", static.LocalFile("./static/uploadfile", true)))
	return r
}
