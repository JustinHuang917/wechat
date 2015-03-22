// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package merchant

import (
	"errors"

	"github.com/chanxuehong/wechat/mp/merchant/group"
)

// 增加分组
//  NOTE: 无需指定 Id 字段
func (c *Client) MerchantGroupAdd(_group *group.GroupEx) (groupId int64, err error) {
	if _group == nil {
		err = errors.New("_group == nil")
		return
	}

	_group.Id = 0 // 无需指定 Id 字段

	var request = struct {
		GroupDetail *group.GroupEx `json:"group_detail"`
	}{
		GroupDetail: _group,
	}

	var result struct {
		Error
		GroupId int64 `json:"group_id"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := merchantGroupAddURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		groupId = result.GroupId
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

// 删除分组
func (c *Client) MerchantGroupDelete(groupId int64) (err error) {
	var request = struct {
		GroupId int64 `json:"group_id"`
	}{
		GroupId: groupId,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := merchantGroupDeleteURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
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
		err = &result
		return
	}
}

// 修改分组名称
func (c *Client) MerchantGroupRename(groupId int64, newName string) (err error) {
	if newName == "" {
		return errors.New(`newName == ""`)
	}

	var request = struct {
		GroupId   int64  `json:"group_id"`
		GroupName string `json:"group_name"`
	}{
		GroupId:   groupId,
		GroupName: newName,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := merchantGroupPropertyModifyURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
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
		err = &result
		return
	}
}

// 修改分组商品
func (c *Client) MerchantGroupModifyProduct(modifyRequest *group.GroupModifyProductRequest) (err error) {
	if modifyRequest == nil {
		return errors.New("modifyRequest == nil")
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := merchantGroupProductModifyURL(token)

	if err = c.postJSON(url_, modifyRequest, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
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
		err = &result
		return
	}
}

// 获取所有分组
func (c *Client) MerchantGroupGetAll() (groups []group.Group, err error) {
	var result struct {
		Error
		GroupsDetail []group.Group `json:"groups_detail"`
	}
	result.GroupsDetail = make([]group.Group, 0, 16)

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := merchantGroupGetAllURL(token)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		groups = result.GroupsDetail
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

// 根据分组ID获取分组信息
func (c *Client) MerchantGroupGetById(groupId int64) (_group *group.GroupEx, err error) {
	var request = struct {
		GroupId int64 `json:"group_id"`
	}{
		GroupId: groupId,
	}

	var result struct {
		Error
		GroupDetail group.GroupEx `json:"group_detail"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := merchantGroupGetByIdURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		_group = &result.GroupDetail
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
