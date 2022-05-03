// pages/functions/price/price.js
import create from '../../../utils/create'
import store from '../../../store/index'
// 引入自己封装好的函数
import {downloadImage} from '../../../utils/tools'

create.Page(store, {
  data: {
    pay:'alipay', // 支付方式
    payCode:{alipay:'',wechat:''}, // 支付二维码
    headers: [{ text: 'nickname',display:'昵称'}, { text:'donate',display: '赞助额'}, {text:'comment',display:'备注'}],
    row:[]
  },
  onLoad: function (options) {
    // 获取赞助二维码
    this.store.setting.getDonate().then(data=>this.setData({payCode:data}))
    // 获取赞助记录
    this.store.plugins.getDonate().then(data=>this.setData({row:data}))
  },
  //切换支付方式
  changePay(e){this.setData({pay:e.currentTarget.dataset.choose})},
  // 保存二维码
  saveQrcode(e){
    //判断二维码类型
    let url=this.data.payCode.alipay
    if (this.data.pay ==='weixin'){
      url=this.data.payCode.wechat
    }
    downloadImage(url)
  }
})