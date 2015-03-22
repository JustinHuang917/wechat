// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"net/http"
)

// 微信支付消息处理接口
type Agent interface {
	GetAppId() string  // 微信公众号身份的唯一标识
	GetMchId() string  // 商户ID，身份标识
	GetAppKey() string // 商户支付密钥Key

	// 未知类型的消息处理方法
	//  postRawXMLMsg 是 xml 消息体
	ServeUnknownMsg(w http.ResponseWriter, r *http.Request, postRawXMLMsg []byte)

	// Native（原生）支付回调商户后台获取 package
	//  postRawXMLMsg 是原始 xml 消息体
	ServePayPackageRequest(w http.ResponseWriter, r *http.Request, req map[string]string, postRawXMLMsg []byte)

	// 用户在成功完成支付后，微信后台通知（POST）商户服务器（notify_url）支付结果。
	//  postRawXMLMsg 是原始 xml 消息体
	ServeOrderNotification(w http.ResponseWriter, r *http.Request, data map[string]string, postRawXMLMsg []byte)
}
