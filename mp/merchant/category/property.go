// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package category

// 获取指定分类的所有属性 成功时返回结果的数据结构
//
//  {
//      "id": "1075741879",
//      "name": "品牌",
//      "property_value": [
//          {
//              "id": "200050867",
//              "name": "VIC&"
//          },
//          {
//              "id": "200050868",
//              "name": "Kate&"
//          },
//          {
//              "id": "200050971",
//              "name": "M&"
//          },
//          {
//              "id": "200050972",
//              "name": "Black&"
//          }
//      ]
//  }
type Property struct {
	Id     string `json:"id"`   // 属性id, int64?
	Name   string `json:"name"` // 属性name
	Values []struct {
		Id   string `json:"id"`   // 属性值id, int64?
		Name string `json:"name"` // 属性值name
	} `json:"property_value,omitempty"`
}
