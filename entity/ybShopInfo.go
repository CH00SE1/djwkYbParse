package entity

import (
	"time"
)

/**
 * @title ybShopInfo
 * @author CH00SE1
 * @date 2022-10-26 15:16:47
 */

// 返回参数
type YbShopResponse struct {
	Records     int          `json:"records"`
	Total       int          `json:"total"`
	Rows        []ShopYbInfo `json:"rows"`
	Page        int          `json:"page"`
	Count       int          `json:"count"`
	FirstResult int          `json:"firstResult"`
	MaxResults  int          `json:"maxResults"`
	Success     bool         `json:"success"`
	Result      string       `json:"result"`
	Conditions  struct {
		LsydName       string `json:"lsydName"`
		AuthStatusList []int  `json:"authStatusList"`
		LsydCodeNew    string `json:"lsydCodeNew"`
		GreaterThan0   string `json:"greaterThan0"`
	} `json:"conditions"`
	Msg       string `json:"msg"`
	Form      string `json:"form"`
	Code      int    `json:"code"`
	OperCount int    `json:"operCount"`
	Sord      string `json:"sord"`
	Sidx      string `json:"sidx"`
	Orderby   string `json:"orderby"`
}

type ShopYbInfo struct {
	CreatedAt                 time.Time
	MaName                    string     `json:"maName"`
	MaOrgName                 string     `json:"maOrgName"`
	Ybjbmc                    string     `json:"ybjbmc"`
	IsSetting                 int        `json:"isSetting"`
	NewAreaBusineseLicense    string     `json:"newAreaBusineseLicense"`
	LsydId                    string     `json:"lsydId"`
	LsydCode                  *string    `json:"lsydCode"`
	LsydName                  string     `json:"lsydName"`
	LsydBuslicense            string     `json:"lsydBuslicense"`
	LsydOrgcode               string     `json:"lsydOrgcode"`
	LsydBusinessstatus        *string    `json:"lsydBusinessstatus"`
	LsydResidence             *string    `json:"lsydResidence"`
	LsydLegrepresentative     *string    `json:"lsydLegrepresentative"`
	LsydIdcardtype            *string    `json:"lsydIdcardtype"`
	LsydCertificatenumber     *string    `json:"lsydCertificatenumber"`
	LsydRegisteredcapital     *string    `json:"lsydRegisteredcapital"`
	LsydBuslicensenumber      *string    `json:"lsydBuslicensenumber"`
	LsydDateissue             *time.Time `json:"lsydDateissue"`
	LsydValiduntil            *time.Time `json:"lsydValiduntil"`
	LsydBusinessscope         *string    `json:"lsydBusinessscope"`
	LsydBusinesstype          *string    `json:"lsydBusinesstype"`
	LsydCeo                   *string    `json:"lsydCeo"`
	LsydContactinformation    *string    `json:"lsydContactinformation"`
	LsydYbname                *string    `json:"lsydYbname"`
	LsydYbphone               *string    `json:"lsydYbphone"`
	LsydEmail                 *string    `json:"lsydEmail"`
	LsydEledocattachment      *string    `json:"lsydEledocattachment"`
	LsydOrgcodedocattachment  string     `json:"lsydOrgcodedocattachment"`
	LsydBuslicensenumtachment *string    `json:"lsydBuslicensenumtachment"`
	AreaName                  string     `json:"areaName"`
	AreaId                    *int       `json:"areaId"`
	LsydStartdate             string     `json:"lsydStartdate"`
	LsydEnddate               string     `json:"lsydEnddate"`
	LsydProtocolstatus        string     `json:"lsydProtocolstatus"`
	LsydServicearea           string     `json:"lsydServicearea"`
	LsydBack                  *string    `json:"lsydBack"`
	LsydBankaccount           *string    `json:"lsydBankaccount"`
	LsydBankDeposit           *string    `json:"lsydBankDeposit"`
	LsydCreditrating          string     `json:"lsydCreditrating"`
	LsydViolation             string     `json:"lsydViolation"`
	LsydAgreementtachment     string     `json:"lsydAgreementtachment"`
	InitializationState       int        `json:"initializationState"`
	SubmitTime                *time.Time `json:"submitTime"`
	TcqcsreauditUserId        string     `json:"tcqcsreauditUserId"`
	TcqcsreauditUserName      string     `json:"tcqcsreauditUserName"`
	TcqcsreauditAddTime       string     `json:"tcqcsreauditAddTime"`
	TcqcsreauditRemark        string     `json:"tcqcsreauditRemark"`
	TcqfsauditUserId          string     `json:"tcqfsauditUserId"`
	TcqfsauditUserName        string     `json:"tcqfsauditUserName"`
	TcqfsauditAddTime         string     `json:"tcqfsauditAddTime"`
	TcqfsauditRemark          string     `json:"tcqfsauditRemark"`
	ShengcsreauditUserId      *string    `json:"shengcsreauditUserId"`
	ShengcsreauditUserName    *string    `json:"shengcsreauditUserName"`
	ShengcsreauditAddTime     *time.Time `json:"shengcsreauditAddTime"`
	ShengcsreauditRemark      string     `json:"shengcsreauditRemark"`
	ShengfsauditUserId        string     `json:"shengfsauditUserId"`
	ShengfsauditUserName      string     `json:"shengfsauditUserName"`
	ShengfsauditAddTime       string     `json:"shengfsauditAddTime"`
	ShengfsauditRemark        string     `json:"shengfsauditRemark"`
	GuocsreauditUserId        string     `json:"guocsreauditUserId"`
	GuocsreauditUserName      string     `json:"guocsreauditUserName"`
	GuocsreauditAddTime       string     `json:"guocsreauditAddTime"`
	GuocsreauditRemark        string     `json:"guocsreauditRemark"`
	GuofsauditUserId          *string    `json:"guofsauditUserId"`
	GuofsauditUserName        *string    `json:"guofsauditUserName"`
	GuofsauditAddTime         *time.Time `json:"guofsauditAddTime"`
	GuofsauditRemark          *string    `json:"guofsauditRemark"`
	LevelStatus               string     `json:"levelStatus"`
	AddUserId                 string     `json:"addUserId"`
	AddUserName               string     `json:"addUserName"`
	AddTime                   *time.Time `json:"addTime"`
	LastUpdateUserId          string     `json:"lastUpdateUserId"`
	LastUpdateUserName        string     `json:"lastUpdateUserName"`
	LastUpdateTime            time.Time  `json:"lastUpdateTime"`
	UserName                  string     `json:"userName"`
	GrantTime                 string     `json:"grantTime"`
	Password                  string     `json:"password"`
	MaxlsydCode               string     `json:"maxlsydCode"`
	HasCount                  string     `json:"hasCount"`
	AuthorizeAttachment       string     `json:"authorizeAttachment"`
	IsMaintainAgreement       string     `json:"isMaintainAgreement"`
	AgreementCount            string     `json:"agreementCount"`
	AuthorizeState            string     `json:"authorizeState"`
	IsDesignatedDrugstore     string     `json:"isDesignatedDrugstore"`
	AuthorizeDetailId         string     `json:"authorizeDetailId"`
	AuthStatus                string     `json:"authStatus"`
	TcqOrgId                  string     `json:"tcqOrgId"`
	TcqOrgName                string     `json:"tcqOrgName"`
	AgreeId                   string     `json:"agreeId"`
	AgreeNum                  string     `json:"agreeNum"`
	ReportStatus              int        `json:"reportStatus"`
	ReportUserId              string     `json:"reportUserId"`
	ReportUserName            string     `json:"reportUserName"`
	ReportTime                string     `json:"reportTime"`
	QueryAreaId               *string    `json:"queryAreaId"`
	RkFlag                    int        `json:"rkFlag"`
	LockStatus                int        `json:"lockStatus"`
	ActualAreaCode            *string    `json:"actualAreaCode"`
	ActualAreaName            *string    `json:"actualAreaName"`
	ActualArea                *string    `json:"actualArea"`
	LatitudeAndLongitude      *string    `json:"latitudeAndLongitude"`
	DredgeStatus              int        `json:"dredgeStatus"`
	ConfirmFlag               int        `json:"confirmFlag"`
	LastConfirmTime           *time.Time `json:"lastConfirmTime"`
	NeedConfirm               int        `json:"needConfirm"`
	DataType                  string     `json:"dataType"`
	RkkFlag                   int        `json:"rkkFlag"`
	Id                        string     `json:"id"`
	IsAudit                   int        `json:"isAudit"`
	AgreeNeedConfirmCount     string     `json:"agreeNeedConfirmCount"`
	YsNum                     string     `json:"ysNum"`
}
