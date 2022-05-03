// pages/index/post/post.js
import Poster from '../../../lib/poster/poster';
import create from '../../../utils/create'
import store from '../../../store/index'

import {wechatCode} from '../../../config/api'
// html解析组件
import towxml from '../../../lib/towxml/index';
// 自己定义的工具类
import {setClipboard} from '../../../utils/tools'
// 获取web端连接
import {WEB} from '../../../config/api'
const app = getApp()
create.Page(store, {
  data: {
    setting: {}, // 微信小程序设置
    post: {}, // 文章内容
    id: 0, // 文章id
    image: "",
    trans: 0, // 顶部导航栏透明度
    content: {}, // 文章解析后的内容
    CustomBar: app.globalData.CustomBar, // 顶部导航栏高度
    replayInfo: {}, // 用户点击回复的信息
    password: '', // 加密文章用户输入的密码
    isShowPosterModal: false,//是否展示海报弹窗
    posterImageUrl: "",//海报地址
  },
  onLoad: function (options) {
    // 获取微信小程序设置
    this.setData({setting:app.globalData.setting})
    // 判断用户是否是扫描小程序码进来的
    const scene = decodeURIComponent(options.scene)
    if (scene!=undefined && scene.split('&') >= 1){
      let scenes = scene.split('&')
      if (scenes.length >= 1){
        options.id = scenes[0]
      }
    }
    // 设置搜索关键词
    this.setData({image:options.img,id:options.id})
    // console.log(this.data.id)
    // 获取文章内容
    this.getPost(options.id)
  },
  //页面滚动监听事件(用于实现渐变效果)
  onPageScroll: function (e) {
    // 页面滚动时执行
    if (e.scrollTop < 200) {
      this.setData({trans: (e.scrollTop / 100) / 2})
    } else {
      this.setData({trans: 1})
    }
  },
  // 获取文章内容
  getPost(id){
    wx.showLoading({title: '玩命加载中'})
    this.store.post.getPostContent(id).then(post=>{
      wx.hideLoading()
      // 判断有没有图片传递过来
      let image = this.data.image
      if(image===undefined){image = post.image}
      // 解析文章
      if(!post.encrypt){this.parsePost(post.content)}
      // 设置文章内容
      this.setData({post,image})
    }).catch(_=>{wx.showToast({title: '无法获取内容!',icon: 'error'})})
  },
  // 点击阅读原文
  readOriginal(){
    setClipboard('提示','点击复制链接在浏览器打开即可!',`${WEB}/archives/${this.data.id}`,'复制连接')
  },
  // 访问有密码的文章
  getEncrypt(e){
    let password = this.data.password
    this.store.post.getEncryption(this.data.id,{password}).then(post=>{
      // 解析内容 ['post.encrypt'] 可以修改对象里的某一个值
      this.setData({['post.encrypt']:false})
      this.parsePost(post.content)
    }).catch(msg=>wx.showToast({title: msg,icon:'error'}))
  },
  // 输入框改变的时候自动设置密码
  onChange(e){this.setData({password:e.detail})},
  // 解析文章
  parsePost(contents){
    let content=towxml(contents, 'html', {
      base: '',// 相对资源的base路径
      theme: 'light',	// 主题，默认`light`
      events: {// 为元素绑定的事件方法
        tap: (e) => {
          // 链接点击事件
          if (e.currentTarget.dataset.data.attr.class.indexOf('h2w__a') !== -1) {
            let href = e.currentTarget.dataset.data.attr.href
            setClipboard('提示','微信小程序限制，点击复制链接在浏览器打开即可!',href,'复制链接')
          }
          // 图片点击事件，点击自动放大
          if (e.currentTarget.dataset.data.attr.class.indexOf('h2w__img') !== -1){
            wx.previewImage({
              current: e.currentTarget.dataset.data.attr.src,
              urls: [e.currentTarget.dataset.data.attr.src]
            })
          }
        }
      }
    })
    this.setData({content})
  },
  // 用户点击回复事件
  replay(e){
    // 设置replayInfo，子组件进行监听
    this.setData({replayInfo:e.detail})
    // console.log(info)
  },
  // 隐藏海报
  hideModal(){this.setData({isShowPosterModal: false})},
  // 生成海报成功的回调函数
  onPosterSuccess(e) {
    const { detail } = e;
    this.setData({
      posterImageUrl: detail,
      isShowPosterModal: true
    })
  },
  // 生成海报失败的回调函数
  onPosterFail(err) {
    console.info(err)
  },
  // 显示海报
  showPoster(){
    // 判断是否已经生成过海报，如果是就显示
    if (this.data.posterImageUrl !== "") {
      this.setData({isShowPosterModal: true})
      return;
    }
    // 海报的配置信息
    var posterConfig = {
      width: 750,
      height: 1200,
      backgroundColor: '#fff',
      debug: false
    }
    // 定义外框
    var blocks = [
      {
        width: 690,
        height: 808,
        x: 30,
        y: 183,
        borderWidth: 2,
        borderColor: '#0081ff',
        borderRadius: 20,
      },
      {
        width: 634,
        height: 74,
        x: 59,
        y: 680,
        backgroundColor: '#fff',
        opacity: 0.5,
        zIndex: 100,
      }
    ]
    // 定义要显示的文字
    var texts = [];
    texts = [
      {
        x: 113,
        y: 61,
        baseLine: 'middle',
        text: '小游',
        fontSize: 32,
        color: '#8d8d8d',
        width: 570,
        lineNum: 1
      },
      {
        x: 32,
        y: 113,
        baseLine: 'top',
        text: '发现一篇很有意思的文章',
        fontSize: 38,
        color: '#080808',
      },
      {
        x: 59,
        y: 770,
        baseLine: 'middle',
        text: this.data.post.title,
        fontSize: 38,
        color: '#080808',
        marginLeft: 30,
        width: 570,
        lineNum: 2,
        lineHeight: 50
      },
      {
        x: 59,
        y: 875,
        baseLine: 'middle',
        // 这我们使用正则替换掉所有的html标签
        text: this.data.post.content.replace(/<[^>]+>/g,""),
        fontSize: 28,
        color: '#929292',
        width: 560,
        lineNum: 2,
        lineHeight: 50
      },
      {
        x: 315,
        y: 1100,
        baseLine: 'top',
        text: '长按识别小程序码,立即阅读',
        fontSize: 28,
        color: '#929292',
      }
    ];
    // 定义图片位置
    const code = wechatCode(this.data.id)
    var images = [
      {
        width: 62,
        height: 62,
        x: 32,
        y: 30,
        borderRadius: 62,
        url: app.globalData.userInfo.avatarUrl, //用户头像
      },
      {
        width: 634,
        height: 475,
        x: 59,
        y: 210,
        url: this.data.image,//海报主图
      },
      {
        width: 220,
        height: 220,
        x: 70,
        y: 1000,
        url: code,//二维码的图
      }
    ];
    posterConfig.blocks = blocks;//海报内图片的外框
    posterConfig.texts = texts; //海报的文字
    posterConfig.images = images;
    console.log(images)
    this.setData({ posterConfig: posterConfig }, () => {
      Poster.create(true);    //生成海报图片
    });
  },
  // 保存海报图片
  savePosterImage () {
    console.log(this.data.posterImageUrl)
    wx.saveImageToPhotosAlbum({
      filePath: this.data.posterImageUrl,
      success: (result) =>{
        wx.showModal({
          title: '提示',
          content: '二维码海报已存入手机相册，赶快分享到朋友圈吧',
          showCancel: false,
          success: _ => {
            this.setData({
              isShowPosterModal: false,
              isShow: false
            })
          }
        })
      },
      fail:  (err) =>{
        console.log(err);
        if (err.errMsg === "saveImageToPhotosAlbum:fail auth deny") {
          console.log("再次发起授权");
          wx.showModal({
            title: '用户未授权',
            content: '如需保存海报图片到相册，需获取授权.是否在授权管理中选中“保存到相册”?',
            showCancel: true,
            success: (res)=> {
              if (res.confirm) {
                console.log('用户点击确定')
                wx.openSetting({
                  success: function success(res) {
                    console.log('打开设置', res.authSetting);
                    wx.openSetting({
                      success(settingdata) {
                        console.log(settingdata)
                        if (settingdata.authSetting['scope.writePhotosAlbum']) {
                          console.log('获取保存到相册权限成功');
                        } else {
                          console.log('获取保存到相册权限失败');
                        }
                      }
                    })
                  }
                });
              }
            }
          })
        }
      }
    });
  },
})