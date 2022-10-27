package entity

import (
	"gorm.io/gorm"
	"time"
)

/**
 * @title medicine1001
 * @author CH00SE1
 * @date 2022-10-27 16:33:46
 */

type MedicineResponse1002 struct {
	Records     int            `json:"records"`
	Total       int            `json:"total"`
	Rows        []Medicine1002 `json:"rows"`
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

type Medicine1002 struct {
	gorm.Model
	PreparationId                    string     `json:"preparationId"`
	PreparationCode                  string     `json:"preparationCode"`
	PreparationType                  string     `json:"preparationType"`
	PreparationPrename               string     `json:"preparationPrename"`
	PreparationMedicinemodel         string     `json:"preparationMedicinemodel"`
	PreparationOutlook               string     `json:"preparationOutlook"`
	PreparationFactor                string     `json:"preparationFactor"`
	PreparationPacknuit              string     `json:"preparationPacknuit"`
	PreparationNuit                  string     `json:"preparationNuit"`
	PreparationMaterialname          string     `json:"preparationMaterialname"`
	PreparationName                  string     `json:"preparationName"`
	PreparationAddress               string     `json:"preparationAddress"`
	PreparationCommissionname        string     `json:"preparationCommissionname"`
	PreparationCommissionaddress     string     `json:"preparationCommissionaddress"`
	PreparationApprovalcode          string     `json:"preparationApprovalcode"`
	PreparationValiditydate          *time.Time `json:"preparationValiditydate"`
	PreparationPermitnumber          string     `json:"preparationPermitnumber"`
	PreparationExestandard           string     `json:"preparationExestandard"`
	PreparationApplicabledisease     string     `json:"preparationApplicabledisease"`
	PreparationDosage                string     `json:"preparationDosage"`
	PreparationChildmedication       string     `json:"preparationChildmedication"`
	PreparationOldatientmedication   string     `json:"preparationOldatientmedication"`
	PreparationContactname           string     `json:"preparationContactname"`
	PreparationContactnumber         string     `json:"preparationContactnumber"`
	PreparationPerdocattachment      string     `json:"preparationPerdocattachment"`
	PreparationApprovaldocattachment string     `json:"preparationApprovaldocattachment"`
	PreparationDesdocattachment      string     `json:"preparationDesdocattachment"`
	AreaName                         string     `json:"areaName"`
	AreaId                           string     `json:"areaId"`
	HosId                            string     `json:"hosId"`
	HosName                          string     `json:"hosName"`
	InitializationState              int        `json:"initializationState"`
	SubmitTime                       time.Time  `json:"submitTime"`
	ReauditUserId                    string     `json:"reauditUserId"`
	ReauditUserName                  string     `json:"reauditUserName"`
	ReauditAddTime                   time.Time  `json:"reauditAddTime"`
	ReauditRemark                    string     `json:"reauditRemark"`
	AuditUserId                      string     `json:"auditUserId"`
	AuditUserName                    string     `json:"auditUserName"`
	AuditAddTime                     time.Time  `json:"auditAddTime"`
	AuditRemark                      string     `json:"auditRemark"`
	AddUserId                        string     `json:"addUserId"`
	AddUserName                      string     `json:"addUserName"`
	AddTime                          time.Time  `json:"addTime"`
	LastUpdateUserId                 string     `json:"lastUpdateUserId"`
	LastUpdateUserName               string     `json:"lastUpdateUserName"`
	LastUpdateTime                   time.Time  `json:"lastUpdateTime"`
	RkFlag                           bool       `json:"rkFlag"`
	RkkFlag                          bool       `json:"rkkFlag"`
	FinalPushStatus                  int        `json:"finalPushStatus"`
	PushVersion                      string     `json:"pushVersion"`
	Message                          string     `json:"message"`
	DataType                         int        `json:"dataType"`
	IsuFlag                          int        `json:"isuFlag"`
	RkTime                           time.Time  `json:"rkTime"`
	PushTime                         time.Time  `json:"pushTime"`
	ProductInsuranceType             int        `json:"productInsuranceType"`
	ProductName                      string     `json:"productName"`
	ProductMedicineModel             string     `json:"productMedicineModel"`
	ProductRemark                    string     `json:"productRemark"`
	ProductCode                      string     `json:"productCode"`
	PayStandard                      string     `json:"payStandard"`
}
