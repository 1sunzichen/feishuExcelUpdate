package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST("/md5file", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		log.Println(file.Filename)
		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		ctx.SaveUploadedFile(file, dst)

		ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	server.Run(":8080")
}
