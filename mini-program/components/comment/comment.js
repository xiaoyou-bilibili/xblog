// components/comment/comment.js
// 单独引入vuex
import {getComment} from '../../store/modules/posts'
import towxml from '../../lib/towxml/index'
Component({
  properties: {
    postId: { type: String, value: ''}, // 文章id, 这里我们监听这个值的变化
  },
  // 属性监听事件
  observers: {
    'postId':function(id){
      if(id>0){
        // 这里说明获取到了文章id
        getComment(id).then(data=>{
          let comments = data
          // 遍历评论来设置评论内容
          for(let i=0;i<comments.length;i++){
            comments[i].content = towxml(comments[i].content, 'html')
          }
          // 设置评论数据
          this.setData({comments})
        })
      }
    }
  },
  options: {
    // 使用全局属性
    addGlobalClass: true
  },
  data: {
    comments: [] // 文章所有的评论
  },
  methods: {
    replay(e){
      const id=e.currentTarget.dataset.id
      const name=e.currentTarget.dataset.name
      // 这里我们触发一个回调事件，由父组件来进行处理
      var myEventDetail = {id,name} // detail对象，提供给事件监听函数
      var myEventOption = {} // 触发事件的选项
      this.triggerEvent('replay', myEventDetail, myEventOption)
    }
  }
})

