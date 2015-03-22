// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// 获取用户信息(需scope为 snsapi_userinfo)
//  NOTE:
//  1. Client 需要指定 OAuth2Config, OAuth2Token
//  2. lang 可能的取值是 zh_CN, zh_TW, en; 如果留空 "" 则默认为 zh_CN.
func (c *Client) UserInfo(lang string) (info *UserInfo, err error) {
	switch lang {
	case "":
		lang = Language_zh_CN
	case Language_zh_CN, Language_zh_TW, Language_en:
		// do nothing
	default:
		err = fmt.Errorf("lang 必须是 \"\",%s,%s,%s 之一",
			Language_zh_CN, Language_zh_TW, Language_en)
		return
	}

	if c.OAuth2Config == nil {
		err = errors.New("没有提供 OAuth2Config")
		return
	}
	if c.OAuth2Token == nil {
		err = errors.New("没有提供 OAuth2Token")
		return
	}

	// 如果过期则自动刷新 access token
	if c.accessTokenExpired() {
		if _, err = c.TokenRefresh(); err != nil {
			return
		}
	}

	if len(c.AccessToken) == 0 {
		err = errors.New("没有有效的 AccessToken")
		return
	}
	if len(c.OpenId) == 0 {
		err = errors.New("没有有效的 OpenId")
		return
	}

	resp, err := c.httpClient().Get(userInfoURL(c.AccessToken, c.OpenId, lang))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", resp.Status)
		return
	}

	var result struct {
		Error
		UserInfo
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	info = &result.UserInfo
	return
}
