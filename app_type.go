package alipay

import (
	"net/url"
)

type AppAuthInfoType struct {
	TargetID string `json:"target_id"`
}

func (app AppAuthInfoType) getInfoUrl() (value url.Values) {
	p := url.Values{}
	p.Add("target_id", app.TargetID)
	return p
}

type AppAuthOauthTokenRequest struct {
	Code string `json:"code"`
}

func (auto AppAuthOauthTokenRequest) APIName() string {
	return "alipay.system.oauth.token"
}

// 返回参数列表
func (auto AppAuthOauthTokenRequest) Params() map[string]string {
	return map[string]string{
		"code": auto.Code,
	}
}

// 返回扩展 JSON 参数的字段名称
func (auto AppAuthOauthTokenRequest) ExtJSONParamName() string {
	return "grant_type"
}

// 返回扩展 JSON 参数的字段值
func (auto AppAuthOauthTokenRequest) ExtJSONParamValue() string {
	return "authorization_code"
}

type AppAuthOauthTokenResponse struct {
	UserID       string `json:"user_id"`
	AlipayUserID string `json:"alipay_user_id"`
	AccessToken  string `json:"access_token"`
}

func (auto AppAuthOauthTokenResponse) APIName() string {
	return "alipay.user.info.share"
}

// 返回参数列表
func (auto AppAuthOauthTokenResponse) Params() map[string]string {
	return map[string]string{
		"auth_token": auto.AccessToken,
	}
}

// 返回扩展 JSON 参数的字段名称
func (auto AppAuthOauthTokenResponse) ExtJSONParamName() string {
	return ""
}

// 返回扩展 JSON 参数的字段值
func (auto AppAuthOauthTokenResponse) ExtJSONParamValue() string {
	return ""
}

type AppUserInfo struct {
	Code               string `json:"code"`
	UserID             string `json:"user_id"`
	Avatar             string `json:"avatar"`
	Province           string `json:"province"`
	City               string `json:"city"`
	NickName           string `json:"nick_name"`
	UserName           string `json:"user_name"` // 真实姓名
	IsStudentCertified string `json:"is_student_certified"`
	UserType           string `json:"user_type"`
	UserStatus         string `json:"user_status"`
	IsCertified        string `json:"is_certified"`
	Gender             string `json:"gender"`
}

func (a AppUserInfo) IsOk() bool {
	return a.Code == "10000"
}
