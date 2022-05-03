/**自己封装的一个工具类*/
// 引入toast
import Toast from '../miniprogram_npm/@vant/weapp/toast/toast';
// 引入dialog
import Dialog from '../miniprogram_npm/@vant/weapp/dialog/dialog';

// 设置剪贴板内容
const setClipboard=(title,message,content,ok='复制',cancel='取消')=>{
  wx.showModal({
    title: title,
    content: message,
    confirmText: ok,
    cancelText: cancel,
    success: res=>{
      if (res.confirm) {
        wx.setClipboardData({data: content})
      }
    }
  })
}
//判断邮箱格式是否正确
const isEmailOk=(email)=>{
  var regExp = /\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/;
  return regExp.test(email);
};
// 下载图片
const downloadImage=(url)=>{
  //先下载网络图片
  wx.downloadFile({
    url: url,
    success: function (res) {
      // 保存图片
      wx.saveImageToPhotosAlbum({
        filePath: res.tempFilePath,
        success: (result) =>{
          wx.showModal({
            title: '提示',
            content: '已保存到系统相册',
            showCancel: false,
            success: _ => {
            }
          })
        },
        fail:  (err) =>{wx.showToast({ icon: 'error', title: '保存图片失败' })}
      });
    },
    fail(){
      wx.showToast({ icon: 'none', title: '无法下载二维码' })
    }
  })
}

// 微信登录授权
const wechatGetInfo=()=>{
  return new Promise((resove,reject) => {
    // 获取用户信息
    wx.getSetting({
      success: res => {
        // 判断用户是否授权过
        if (res.authSetting['scope.userInfo']) {
          // 已经授权，可以直接调用 getUserInfo 获取头像昵称，不会弹框
          wx.getUserInfo({success: res =>resove(res),fail: res => reject(res)})
        }else{
          reject('用户未授权')
        }
      },
      fail: res => reject(res)
    })
  })
}

// 数字格式化(前面是数字，后面的是小数点位数)
const numFormat=(num, digits)=>{
  var si = [
    { value: 1, symbol: "" },
    { value: 1E3, symbol: "k" },
    { value: 1E4, symbol: "W" }
  ];
  var rx = /\.0+$|(\.[0-9]*[1-9])0+$/;
  var i;
  for (i = si.length - 1; i > 0; i--) {
    if (num >= si[i].value) {
      break;
    }
  }
  return (num / si[i].value).toFixed(digits).replace(rx, "$1") + si[i].symbol;
}

module.exports = {
  Dialog,
  setClipboard,
  isEmailOk,
  wechatGetInfo,
  downloadImage,
  numFormat
} 