package alipay

import (
	"crypto"
	"errors"
)

func (this AliPay) GetAuthInfo(autoInfoType AppAuthInfoType) (string, error) {
	urlValue := autoInfoType.getInfoUrl()
	urlValue.Add("apiname", "com.alipay.account.auth")
	urlValue.Add("method", "alipay.open.auth.sdk.code.get")
	urlValue.Add("app_id", this.appId)
	urlValue.Add("app_name", "mc")
	urlValue.Add("biz_type", "openservice")
	urlValue.Add("pid", this.partnerId)
	urlValue.Add("product_id", "APP_FAST_LOGIN")
	urlValue.Add("scope", "kuaijie")
	urlValue.Add("auth_type", "AUTHACCOUNT")
	urlValue.Add("sign_type", "RSA2")
	sign, err := signWithPKCS1v15(urlValue, this.privateKey, crypto.SHA256)
	if err != nil {
		return "", err
	}
	urlValue.Add("sign", sign)
	return urlValue.Encode(), nil
}

// 第一步获取用户授权令牌
func (this AliPay) GetAutoOauthToken(oauthTokenRequest AppAuthOauthTokenRequest) (resp AppAuthOauthTokenResponse, err error) {
	type RespMap struct {
		AlipaySystemOauthTokenResponse AppAuthOauthTokenResponse `json:"alipay_system_oauth_token_response"`
		ErrorResponse ErrorResponseClass `json:"error_response"`
	}
	var respMap RespMap
	err = this.DoRequest("POST", oauthTokenRequest, &respMap)
	if err != nil || respMap.ErrorResponse.Code != "" {
		return
	}
	resp = respMap.AlipaySystemOauthTokenResponse
	return
}

// 第二部通过令牌查询用户信息
func (this AliPay) GetUserInfo(oauthToken AppAuthOauthTokenResponse) (resp AppUserInfo, err error) {
	type RespMap struct {
		AlipayUserInfoShareResponse AppUserInfo `json:"alipay_user_info_share_response"`
	}
	var respMap RespMap
	err = this.DoRequest("POST", oauthToken, &respMap)
	if respMap.AlipayUserInfoShareResponse.Code != "10000" {
		return resp, errors.New("授权失败")
	}
	resp = respMap.AlipayUserInfoShareResponse
	return
}


type ErrorResponseClass struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}