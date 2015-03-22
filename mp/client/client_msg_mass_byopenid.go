// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"

	"github.com/chanxuehong/wechat/mp/message/active/massbyopenid"
)

// 根据用户列表群发文本消息.
func (c *Client) MsgMassSendTextByOpenId(msg *massbyopenid.Text) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发图片消息.
func (c *Client) MsgMassSendImageByOpenId(msg *massbyopenid.Image) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发语音消息.
func (c *Client) MsgMassSendVoiceByOpenId(msg *massbyopenid.Voice) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发视频消息.
func (c *Client) MsgMassSendVideoByOpenId(msg *massbyopenid.Video) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发图文消息.
func (c *Client) MsgMassSendNewsByOpenId(msg *massbyopenid.News) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return c.msgMassSendByOpenId(msg)
}

func (c *Client) msgMassSendByOpenId(msg interface{}) (msgid int64, err error) {
	var result struct {
		Error
		MsgId int64 `json:"msg_id"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := messageMassSendByOpenIdURL(token)

	if err = c.postJSON(url_, msg, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		msgid = result.MsgId
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return
	}
}
