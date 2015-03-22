// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package category

// 获取指定分类的所有子分类 成功时返回结果的数据结构
//
//  {
//      "errcode": 0,
//      "errmsg": "success",
//      "cate_list": [
//          {
//              "id": "537074292",
//              "name": "数码相机"
//          },
//          {
//              "id": "537074293",
//              "name": "家用摄像机"
//          },
//          {
//              "id": "537074298",
//              "name": "单反相机"
//          },
//          {
//              "id": "537082412",
//              "name": "镜头"
//          },
//          {
//              "id": "537102001",
//              "name": "单电/微单"
//          }
//      ]
//  }
type Category struct {
	Id   string `json:"id"`   // 分类id, int64?
	Name string `json:"name"` // 分类name
}
