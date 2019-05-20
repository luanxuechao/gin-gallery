package v1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"gin-gallery/pkg/e"
	"gin-gallery/pkg/util"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, handler, _ := c.Request.FormFile("file")
	fileName := util.GetRandomString(20) + path.Ext(handler.Filename)
	// log.Println(c.Request.RemoteAddr, c.Request.URL, c.Request.Host, c.Request.RequestURI)
	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", handler.Filename))
	out, err := os.Create("static/uploadfile/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["imageUri"] = "http://" + c.Request.Host + "/app/v1/images/" + fileName
	// c.String(http.StatusCreated, "upload successful")
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetImage(c *gin.Context) {
	fileName := c.Param("fileName")
	log.Println(fileName)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "image/jpeg")
	c.File("static/uploadfile/" + fileName)
}
