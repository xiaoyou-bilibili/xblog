// pages/me/advice/advice.js
import create from '../../../utils/create'
import store from '../../../store/index'
create.Page(store, {
  data: {

  },
  // 用户提交意见反馈
  submitAdvice(e){
    const data=e.detail.value
    console.log(data)
    if (data.content==''){
      wx.showToast({ icon: 'none', title: '意见不能为空' })
    }else{
      // 用户提交意见反馈
      this.store.tools.submitAdvice(data).then(_=>{
        wx.showModal({
          title: '提示',
          content: '意见反馈成功!',
          confirmText: '回到主页',
          showCancel: false,
          success (res) {
            wx.switchTab({url: '/pages/me/home/home'})
          }
        })
      })
    }
  }

})