// pages/index/home/home.js
import create from '../../../utils/create'
import store from '../../../store/index'
create.Page(store, {
  data: {
    posts: [], // 文章内容
    current: 1, // 当前第几页
    page: 0, // 总页数
    key: '' // 搜索关键词
  },
  onLoad: function (options) {
    console.log(options)
    this,
    // 设置搜索关键词
    this.setData({key:options.key})
    this.getPosts()
  },
  // 获取文章列表
  getPosts(page=1){
    wx.showLoading({title: '加载中'})
    this.store.post.getPostList({page,q: this.data.key}).then(data=>{
      wx.hideLoading()
      // 根据页数，自动更新文章数据
      let posts=[]
      if(data.current===1){
        posts = data.contents
      }else{
        posts = this.data.posts.concat(data.contents)
      }
      // 设置数据
      this.setData({
        current: data.current,
        page: data.total,
        posts
      })
    })
  },
  // 上拉加载事件
  onReachBottom (){
    if(this.data.current>=this.data.page){
      wx.showToast({title: '到底了',icon: 'error'})      
    } else {
      this.getPosts(parseInt(this.data.current)+1)
    }
  },
  // 点击文章卡片显示文章
  postDetail(e) {
    const id=e.currentTarget.dataset.id;
    const img=e.currentTarget.dataset.img;
    wx.navigateTo({url: '/pages/index/post/post?id='+id+'&img='+img})
  },
})