// components/table/index.js

Component({
  options: {
    styleIsolation: "isolated"
  },
  /**
   * 组件的属性列表
   */
  properties: {
    list: {
      type: Array,
      value: []
    },
    headers: {
      type: Array,
      value: []
    },
    hasBorder: {
      type: String,
      value: "no"
    },
    height: { //table的高度
      type: Number || String,
      value: ''
    },
    width: {
      type: Number,
      value: 0
    },
    tdWidth: {
      type: Number,
      value: 35
    }

  },
  externalClasses: ['s-class-header', 's-class-row'],
  lifetimes: {
    attached: function() {
      this.createSelectorQuery().select('.thead > .td').boundingClientRect( (rect) => {
        console.log('rect', rect)
        let tempWidth = rect.width * this.properties.headers.length 
        if (this.properties.width > 0 && this.properties.width < tempWidth){
          this.setData({
            theadWidth: tempWidth,
            scrolWidth: this.properties.width
          })          
        }else{
          this.setData({
            theadWidth: tempWidth,
            scrolWidth: tempWidth
          })
        } 
        let tempHeight = rect.height * this.properties.list.length 
        if (tempHeight < this.properties.height){
          this.setData({
            height: this.properties.height
          })

        }
      }).exec()
      // 在组件实例进入页面节点树时执行
      
    },
    detached: function() {
      // 在组件实例被从页面节点树移除时执行
    }
  },

  /**
   * 组件的初始数据
   */
  data: {
    theadWidth: 0,    
    scrolWidth:0
  },

  /**
   * 组件的方法列表
   */
  methods: {

  }
})