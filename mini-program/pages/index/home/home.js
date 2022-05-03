// pages/index/home/home.js
import create from '../../../utils/create'
import store from '../../../store/index'
const app = getApp()
create.Page(store, {
  use: ['post'],
  data: {
    index: '0', // 当前选中的内容(显示文章或者日记)
    posts: [], // 文章内容
    current: 1, // 当前第几页
    page: 0, // 总页数
    search: false, // 是否显示搜索的遮罩层
    keyList: '', // 关键词列表
    statusBar: app.globalData.CustomBar // 吸顶高度
  },
  onLoad: function (options) {
    this.getPosts()
  },
  // 获取文章列表
  getPosts(page=1){
    wx.showLoading({title: '加载中'})
    this.store.post.getPostList({page}).then(data=>{
      this.processData(data,data.contents)
    })
  },
  // 获取日记列表
  getDiary(page=1){
    wx.showLoading({title: '加载中'})
    this.store.plugins.getDiary({page,filter:1}).then(data=>{
      this.processData(data,data.contents)
    })
  },
  // 处理文章数据(文章和日记的逻辑是一样的)
  processData(data,list){
    wx.hideLoading()
    // 根据页数，自动更新文章数据
    let posts=[]
    if(data.current===1){
      // 页面自动滚动到顶部
      wx.pageScrollTo({scrollTop: 0})
      posts = list
    }else{
      posts = this.data.posts.concat(list)
    }
    // 设置数据
    this.setData({
      current: data.current,
      page: data.total,
      posts
    })
  },
  tabChange(e){ // tab切换事件
    const choose = e.currentTarget.dataset.id
    // 先修改顶部的状态栏
    this.setData({ index : choose})
    // 获取对应的内容
    if(choose==='0'){
      this.getPosts()
    }else{
      this.getDiary()
    }
  },
  // 上拉加载事件
  onReachBottom (){
    if(this.data.current>this.data.page){
      Toast.fail('到底了');
    } else {
      let page = parseInt(this.data.current)+1
      // 判断加载的是日记还是文章
      if(this.data.index==='0'){
        this.getPosts(page)
      }else{
        this.getDiary(page)
      }
    }
  },
  // 用户关键词输入改变事件
  onChange(e){
    const key = e.detail
    // 设置是否显示遮罩层
    this.setData({search: key !== ""})
    // 显示搜索关键词
    this.store.post.getPostList({q:key}).then(data=>{
      this.setData({keyList:data.contents})
    })
  },
  // 用户搜索框搜索事件
  onSearch(e){
    wx.navigateTo({url: '/pages/index/search/search?key=' + e.detail})
  },
  // 关键词点击事件
  searchGo(e){
    wx.navigateTo({url: '/pages/index/post/post?id=' + e.currentTarget.dataset.id})
  },
  // 点击文章卡片显示文章
  postDetail(e) {
    const id=e.currentTarget.dataset.id;
    const img=e.currentTarget.dataset.img;
    wx.navigateTo({url: '/pages/index/post/post?id='+id+'&img='+img})
  },
})