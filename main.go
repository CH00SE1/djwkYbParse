package main

/**
 * @title main
 * @author CH00SE1
 * @date 2022-10-26 15:26:22
 */

import (
	"djwkYbParse/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//db.CreateTableGrom(entity.ShopYbInfo{})
	r := gin.Default()
	r.POST("/ybShopInsert", service.JsonRequestBody)
	r.Run(":8300")
}
