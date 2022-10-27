package main

/**
 * @title main
 * @author CH00SE1
 * @date 2022-10-26 15:26:22
 */

import (
	"djwkYbParse/entity"
	"djwkYbParse/service"
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	//r := gin.Default()
	//r.POST("/ybShopInsert", service.JsonRequestBody)
	//r.Run(":8300")
	url := "https://code.nhsa.gov.cn/yp/stdGoodsPublic/getStdGoodsPublicData.html"
	method := "POST"
	text := "goodsCode=&companyNameSc=&registeredProductName=&approvalCode=&_search=false&nd=&rows=" + strconv.Itoa(20) + "&page=" + strconv.Itoa(10) + "&sidx=&sord=asc"
	info, _ := service.RequestMedicineInfo(url, text, method, entity.MedicineResponse1001{})
	medicine1003 := entity.MedicineResponse1003{}
	json.Unmarshal(info, &medicine1003)
	fmt.Println(medicine1003)
}
