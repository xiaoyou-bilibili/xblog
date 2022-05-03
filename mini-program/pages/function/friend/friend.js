// pages/functions/friend/friend.js
import create from '../../../utils/create'
import store from '../../../store/index'
// 引入自己封装好的函数
import {Dialog,setClipboard,isEmailOk} from '../../../utils/tools'
const app = getApp()
create.Page(store, {
  // 声明依赖
  use: ['plugins'],
  // 计算属性
  computed: {
    friends() { return this.plugins.friends }
  },
  data: {
    setting:{},
    animation:true,//是否开启动画效果
    process: true, // 骨架屏是否加载
    link: {}, // 个人友链数据
    addFriend:false,//打开申请弹窗
  },
  onLoad: function (options) {
    // 获取微信小程序设置
    this.setData({setting:app.globalData.setting})
    if(this.data.setting.friend){
      // 获取我的友链
      this.store.plugins.getFriend().then(_=>{this.setData({process:false})})
    }
    // 设置动画效果
    setTimeout(()=>{this.setData({animation: false})}, 1000)
    // 获取友链设置
    this.store.setting.getFriend().then(data=>{this.data.link=data})
  },
  // 获取友链
  addFriend(){
    Dialog.confirm({
      title: '注意',
      message: "1.申请时请先加上本站的连接(〃'▽'〃)\n2.原创博客，非采集站，全站 HTTPS 优先\n* 图标仅支持 png / jpg /gif 等格式，请勿提交 ico 或 分辨率小于 100x100 的图标",
      confirmButtonText: '开始申请',
      cancelButtonText:'查看本站友链',
      closeOnClickOverlay:true
    }).then(() => {
      //申请友链界面
      this.setData({addFriend:true})
    }).catch((e)=>{
      let data="名字:"+this.data.link.name+"\n描述:"+this.data.link.dec+"\n网址:"+this.data.link.link+"\n头像:"+this.data.link.avatar
      // 设置剪贴板内容
      setClipboard('提示','点击复制即可复制友链数据',data)
    })
  },
  //显示友链的详细信息
  showFriend(e){
    const data=e.currentTarget.dataset
    console.log(e)
    // 设置剪贴板内容
    setClipboard(data.name,data.dec,data.src,'复制地址')
  },
  //添加友链数据
  submitFriend(e){
    const data=e.detail.value
    // console.log(data)
    //判端邮箱是否正确
    if (!isEmailOk(data.email)){
      wx.showToast({icon:'none',title: '邮箱格式不正确'})
    }else if(data.name=='' || data.site==''){
      wx.showToast({ icon: 'none', title: '名字和站点地址不能为空' })
    }else{
      // 提交友链数据
      this.store.plugins.submitFriend(data).then(_=>{
        wx.showToast({ icon: 'none', title: '申请成功,审核通过后系统会发送邮件通知你!' ,duration:2000})
        this.setData({addFriend: false})
      }).catch(msg=>{wx.showToast({ icon: 'none', title: msg })})
    }
  },
  // 隐藏友链
  hideModal(){
    this.setData({addFriend:false})  
  }
})