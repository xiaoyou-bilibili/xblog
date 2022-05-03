// pages/functions/bilibili/bilibili.js
import create from '../../../utils/create'
import store from '../../../store/index'
// 引入自己封装好的函数
import {Toast,setClipboard} from '../../../utils/tools'
const app = getApp()
create.Page(store, {
  data: {
    setting:{},
    // 已追多少
    total: 0,
    // 当前第几页
    current: 1,
    // 总计多少页
    page: 0,
    // 我的追番
    animations: [],
    // 是否加载完成
    process: true,
    // 是否显示番剧详情
    show: false,
    // 番剧详情内容
    detail: {}
  },
  onLoad: function (options) {
    // 获取微信小程序设置
    this.setData({setting:app.globalData.setting})
    this.getAnimation(1);
  },
  // 获取动漫内容
  getAnimation(page){
    if(!this.data.setting.animation){
      return
    }
    this.store.plugins.getAnimation({page}).then(data=>{
      // 根据页数，自动更新动画数据
      let animations=[]
      if(data.current===1){
        animations = data.contents
      }else{
        animations = this.data.animations.concat(data.contents)
      }
      // 设置数据
      this.setData({
        total: data.num,
        current: data.current,
        page: data.total,
        animations: animations,
        process:false
      })
    })
  },
  // 上拉加载事件
  onReachBottom (){
    if(this.data.current>this.data.page){
      Toast.fail('到底了');
    } else {
      this.getAnimation(parseInt(this.data.current)+1)
    }
  },
  //显示我的动画数据
  showAnimation(e){
    // 获取动画数据
    const data=e.currentTarget.dataset
    // 设置内容
    this.setData({
      show:true,
      detail: data
    })
  },
  //隐藏番剧详情
  hideModal(){
    this.setData({show:false})
  },
  openAnimation(e){
    const url = e.currentTarget.dataset.url
    this.setData({show:false})
    // 设置剪贴板内容
    setClipboard('提示','因为微信小程序限制，点击复制播放链接后用浏览器打开即可观看',url,'复制链接')
  },
})


