// components/post-bottom/post-buttom.js
const app =getApp()
// 引入自己定义的请求函数
import {getPostCollection,updatePostCollection,commitComment} from '../../store/modules/posts'
Component({
  properties: {
    postId: { type: String, value: ''}, // 文章id, 这里我们监听这个值的变化
    replay: { type: Object, value: ()=>{}}, // 回复的信息
  },
  // 属性监听事件
  observers: {
    // 监听文章id变化
    'postId':function(id){
      if(id>0){
        // 这里说明获取到了文章id
        getPostCollection(id,wx.getStorageSync('openid')).then(data=>{
          this.setData({status:data})
        })
      }
    },
    // 监听回复内容变化
    'replay':function(replay){
      if(replay.id!==undefined && replay.id!==null){
        // 设置提示内容并设为焦点
        this.setData({placeholder:`@${replay.name}`,focus:true})
      }
    }
  },
  // 组件所在页面的周期
  pageLifetimes: {
    show() {
      // 设置用户的头像等信息
      if(app.globalData.userInfo!==null){
        this.setData({avatar:app.globalData.userInfo.avatarUrl})
      }
    }
  },
  options: {
    // 使用全局属性
    addGlobalClass: true
  },
  data: {
    isShow: false, // 是否显示顶部内容
    avatar: '', // 个人头像
    status: {}, // 文章收藏和点赞的状态
    placeholder: '随便说点什么吧~', // 输入框提示信息
    focus: false, // 编辑框设为焦点
    content: '' // 输入框内容
  },
  methods: {
    // 显示和影响底部导航栏
    showMenuBox(){
      // console.log(app.globalData.userInfo.avatarUrl)
      this.setData({isShow: !this.data.isShow})
    },
    // 文章收藏或点赞
    updateCollection(e){
      const option=e.currentTarget.dataset.option
      let status = this.data.status
      // 判断哪里需要更新
      if (option=='collect'){
        status.collection = !status.collection
      }else if(option=='good'){
        status.good = !status.good
      }
      console.log('更新文章')
      // 更新文章状态
      updatePostCollection(this.properties.postId,wx.getStorageSync('openid'),status).then(status=>{
        this.setData({status})
      }).catch('')
    },
    // 发送评论
    sendComment(e){
      // 用户评论的内容
      let content = e.detail.value
      // 用户内容不能为空
      if(app.globalData.userInfo===null){
        wx.showToast({title: '请先登录',icon: 'none'})
        return
      }else if(content.content===''){
        wx.showToast({title: '评论内容不能为空',icon: 'none'})
        return
      }
      // 用户其他内容
      const id = this.properties.postId
      content.name = app.globalData.userInfo.nickName
      content.openid = wx.getStorageSync('openid')
      content.avatar = app.globalData.userInfo.avatarUrl
      if(this.properties.replay.id!=undefined){
        content.parent = this.properties.replay.id
      }
      // 清空输入框
      this.setData({content:''})
      // 提交评论
      commitComment(id,content).then(_=>{
        wx.showModal({content: '为避免垃圾评论，评论需要审核后才能显示！',showCancel: false})
      }).catch(_=>wx.showToast({title: '好像出了点问题~',icon:'none'}))
    },
    showQrcode(){
      // 显示赞赏码
      wx.previewImage({
        current: [app.globalData.setting.price],
        urls: [app.globalData.setting.price]
      })
    },
    // 显示海报
    showPoster(){
      var myEventDetail = {} // detail对象，提供给事件监听函数
      var myEventOption = {} // 触发事件的选项
      this.triggerEvent('showPoster', myEventDetail, myEventOption)
    }
  }
})
