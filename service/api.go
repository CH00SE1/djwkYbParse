package service

import (
	"djwkYbParse/db"
	"djwkYbParse/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

/**
 * @title api
 * @author CH00SE1
 * @date 2022-10-26 16:12:59
 */

func JsonRequestBody(c *gin.Context) {
	var ybShopResponse entity.YbShopResponse
	err := c.ShouldBindJSON(&ybShopResponse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows := ybShopResponse.Rows
	fmt.Println(len(rows))
	groom, _ := db.Goorm()
	groom.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.ShopYbInfo{})
	for _, row := range rows {
		groom.Create(&row).Callback()
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "操作成功",
	})
}
