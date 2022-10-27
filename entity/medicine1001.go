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

type MedicineResponse1001 struct {
	Records     int            `json:"records"`
	Total       int            `json:"total"`
	Rows        []Medicine1001 `json:"rows"`
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

type Medicine1001 struct {
	gorm.Model
	PiecesId                     string    `json:"piecesId"`
	PiecesCode                   string    `json:"piecesCode"`
	PiecesName                   string    `json:"piecesName"`
	MedicinalMaterialsName       string    `json:"medicinalMaterialsName"`
	ProcessingMethod             string    `json:"processingMethod"`
	EfficacyClassification       string    `json:"efficacyClassification"`
	SubjectSource                string    `json:"subjectSource"`
	MaterialSource               string    `json:"materialSource"`
	PharmaceuticalSite           string    `json:"pharmaceuticalSite"`
	PropertyflavorChanneltropism string    `json:"propertyflavorChanneltropism"`
	FunctionIndications          string    `json:"functionIndications"`
	UsageDosage                  string    `json:"usageDosage"`
	PaymentPolicy                string    `json:"paymentPolicy"`
	PaymentPolicyP               string    `json:"paymentPolicyP"`
	SpecificationName            string    `json:"specificationName"`
	SpecificationPageNumber      string    `json:"specificationPageNumber"`
	SpecificationAttachment      string    `json:"specificationAttachment"`
	AreaName                     string    `json:"areaName"`
	AreaId                       string    `json:"areaId"`
	InitializationState          int       `json:"initializationState"`
	SubmitTime                   time.Time `json:"submitTime"`
	IsUsing                      int       `json:"isUsing"`
	IsUsingRemark                string    `json:"isUsingRemark"`
	ReauditUserId                string    `json:"reauditUserId"`
	ReauditUserName              string    `json:"reauditUserName"`
	ReauditAddTime               time.Time `json:"reauditAddTime"`
	ReauditRemark                *string   `json:"reauditRemark"`
	AuditUserId                  string    `json:"auditUserId"`
	AuditUserName                string    `json:"auditUserName"`
	AuditAddTime                 time.Time `json:"auditAddTime"`
	AuditRemark                  *string   `json:"auditRemark"`
	AddUserId                    string    `json:"addUserId"`
	AddUserName                  string    `json:"addUserName"`
	AddTime                      time.Time `json:"addTime"`
	LastUpdateUserId             string    `json:"lastUpdateUserId"`
	LastUpdateUserName           string    `json:"lastUpdateUserName"`
	LastUpdateTime               time.Time `json:"lastUpdateTime"`
	EfficacyClassificationid     string    `json:"efficacyClassificationid"`
	IsReport                     string    `json:"isReport"`
	RkFlag                       bool      `json:"rkFlag"`
	RkkFlag                      bool      `json:"rkkFlag"`
	SourceAreaName               string    `json:"sourceAreaName"`
	SourceAreaId                 string    `json:"sourceAreaId"`
	PushVersion                  string    `json:"pushVersion"`
	FinalPushStatus              int       `json:"finalPushStatus"`
	Message                      string    `json:"message"`
	IsuFlag                      int       `json:"isuFlag"`
	RkTime                       time.Time `json:"rkTime"`
	PushTime                     time.Time `json:"pushTime"`
}
