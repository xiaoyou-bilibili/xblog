// pages/functions/douban/douban.js
import create from '../../../utils/create'
import store from '../../../store/index'
// 引入自己封装好的函数
import {SetClipboard} from '../../../utils/tools'
const app = getApp()
create.Page(store, {
  data: {
    setting: {},
    option:'book',//当前选中的项目
    records:[],//获取的豆瓣数据
    current:1,//当前第几页
    page: 0,
    process: true, // 显示加载进度条
    statusBar: app.globalData.CustomBar // 吸顶高度
  },
  onLoad: function (options) {
    // 获取微信小程序设置
    this.setData({setting:app.globalData.setting})
    // 获取豆瓣记录
    this.getDouban();
  },
  //切换选项卡
  changeItem(e){
    this.setData({option: e.currentTarget.dataset.option})
    //刷新内容
    this.getDouban()
  },
  // 获取豆瓣内容
  getDouban(id=1){
    if(!this.data.setting.dou_ban){
      return
    }
    // 如果id为1，那么就显示加载骨架屏
    if(id===1){
      // 页面自动滚动到顶部
      wx.pageScrollTo({scrollTop: 0})
      this.setData({process:true})
    }
    this.store.plugins.getDouBan(this.data.option,{page:id}).then(data=>{
      // console.log(data)
      // 根据页数，自动更新动画数据
      let records=[]
      if(data.current===1){
        records = data.contents
      }else{
        records = this.data.records.concat(data.contents)
      }
      // 设置数据
      this.setData({
        current: data.current,
        page: data.total,
        records: records,
        process:false
      })
    })
  },
  // 上拉加载事件
  onReachBottom (){
    if(this.data.current>=this.data.page){
      wx.showToast({
        title: '没有更多数据了',
        icon:'error'
      })
    } else {
      this.getDouban(parseInt(this.data.current)+1)
    }
  },
  // 点击事件
  itemClick(e){
    const url = e.currentTarget.dataset.url
    SetClipboard('提示','因为微信小程序限制，点击复制链接后用浏览器打开即可查看详细信息',url,'复制链接');
  }
})