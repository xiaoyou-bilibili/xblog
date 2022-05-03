// components/article/article.js
import {numFormat} from '../../utils/tools'
Component({
  properties: {
    item: { type: Object, value: ()=>{}}
  },
  options: {
    // 使用全局属性
    addGlobalClass: true
  },
  // 监听器监听数据变化自动设置内容
  observers:{
    item(list){
      this.setData({
        view:numFormat(list.view,2),
        good:numFormat(list.good,2),
        comment:numFormat(list.comment,2)
      })
    }
  },
  data: {
    view:0,
    good:0,
    comment:0
  }
})