package entity

import (
	"gorm.io/gorm"
)

/**
 * @title medicine1001
 * @author CH00SE1
 * @date 2022-10-27 16:33:46
 */

type MedicineResponse1003 struct {
	Records     int            `json:"records"`
	Total       int            `json:"total"`
	Rows        []Medicine1003 `json:"rows"`
	Page        int            `json:"page"`
	Count       int            `json:"count"`
	FirstResult int            `json:"firstResult"`
	MaxResults  int            `json:"maxResults"`
	Success     bool           `json:"success"`
	Result      interface{}    `json:"result"`
	Conditions  struct {
	} `json:"conditions"`
	Msg       interface{} `json:"msg"`
	Form      interface{} `json:"form"`
	Code      int         `json:"code"`
	OperCount int         `json:"operCount"`
	Sord      string      `json:"sord"`
	Sidx      string      `json:"sidx"`
	Orderby   interface{} `json:"orderby"`
}

type Medicine1003 struct {
	gorm.Model
	MaterialName            string `json:"materialName"`
	CompanyNameSc           string `json:"companyNameSc"`
	RegisteredProductName   string `json:"registeredProductName"`
	Unit                    string `json:"unit"`
	ApprovalCode            string `json:"approvalCode"`
	RegisteredOutlook       string `json:"registeredOutlook"`
	RegisteredMedicinemodel string `json:"registeredMedicinemodel"`
	GoodsStandardCode       string `json:"goodsStandardCode"`
	GoodsCode               string `json:"goodsCode"`
	MinUnit                 string `json:"minUnit"`
	Factor                  int    `json:"factor"`
	GoodsName               string `json:"goodsName"`
	ProductRemark           string `json:"productRemark,omitempty"`
	ProductName             string `json:"productName,omitempty"`
}
