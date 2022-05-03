<template>
  <div>
    <div id="editor" class="editor" />
  </div>
</template>

<script>
import $ from 'jquery'
let editor
let that
// 自定义按钮事件 name 为唯一标识符 tip为提示 click为点击触发事件，icon为图标
const info = { name: 'blueBackground', tip: '蓝色背景', click () { editor.updateValue(`[info]${editor.getSelection()}[/info]`) }, icon: '<img alt src="/static/images/buttons/info.png">' }
const danger = { name: 'dangerBackground', tip: '红色背景', click () { editor.updateValue(`[danger]${editor.getSelection()}[/danger]`) }, icon: '<img alt src="/static/images/buttons/danger.png">' }
const success = { name: 'greenBackground', tip: '绿色背景', click () { editor.updateValue(`[success]${editor.getSelection()}[/success]`) }, icon: '<img alt  src="/static/images/buttons/success.png">' }
const infoBox = { name: 'infoBox', tip: '蓝色标题栏', click () { editor.updateValue(`[infobox title="标题"]${editor.getSelection()}[/infobox]`) }, icon: '<img alt  src="/static/images/buttons/infobox.png">' }
const dangerBox = { name: 'dangerBox', tip: '红色标题栏', click () { editor.updateValue(`[dangerbox title="标题"]${editor.getSelection()}[/dangerbox]`) }, icon: '<img alt  src="/static/images/buttons/dangerbox.png">' }
const successBox = { name: 'successBox', tip: '绿色标题栏', click () { editor.updateValue(`[successbox title="标题"]${editor.getSelection()}[/successbox]`) }, icon: '<img alt  src="/static/images/buttons/successbox.png">' }
const block = { name: 'block', tip: '代码块', click () { editor.updateValue(`[block]${editor.getSelection()}[/block]`) }, icon: '<img alt  src="/static/images/buttons/codeblock.png">' }
const collSpan = { name: 'collSpan', tip: '展开收缩', click () { editor.updateValue(`[collapse title="标题"]${editor.getSelection()}[/collapse]`) }, icon: '<img alt  src="/static/images/buttons/accordion.png">' }
const music = { name: 'music', tip: '网易云音乐', click () { editor.updateValue(`[music autoplay="0"]${editor.getSelection()}[/music]`) }, icon: '<img alt  src="/static/images/buttons/music.png">' }
const bilibili = { name: 'bilibili', tip: '哔哩哔哩', click () { editor.updateValue(`[bilibili cid="" page="1"]${editor.getSelection()}[/bilibili]`) }, icon: '<img alt  src="/static/images/buttons/bilibili.png">' }
const yBtn = { name: 'yBtn', tip: '云盘下载', click () { editor.updateValue(`[ypbtn]${editor.getSelection()}[/ypbtn]`) }, icon: '<img alt  src="/static/images/buttons/ypbtn.png">' }
// import '~vditor/src/assets/scss/index'
export default {
  name: 'PostEdit',
  props: {
    html: { type: String, default: '' }, // 文章渲染的html内容
    md: { type: String, default: '' } // 文章渲染的md内容
  },
  data () {
    that = this
    return {
      // vditor编辑器的设置
      option: {
        height: 740,
        resize: { enable: true }, // 允许拖拽
        tab: '    ', // 设置tab为四个空格
        typewriterMode: true, // 开启打字机模式（编辑位置上下居中）
        placeholder: '随便写点什么吧~', // 默认提示文字
        counter: { enable: true, type: 'text' }, // 开启文本字数统计
        mode: 'ir', // 使用分屏渲染功能
        icon: 'material', // 使用material风格的图标
        toolbar: ['undo', 'redo', '|',
          'bold', 'italic', 'strike', 'quote', 'line', '|',
          'headings', 'list', 'ordered-list', 'outdent', 'indent', '|',
          'upload', 'link', 'code', 'inline-code', 'check', 'table', 'emoji', '|',
          'both', 'edit-mode', 'fullscreen', 'preview', 'outline', 'export', '|',
          info, danger, success, infoBox, dangerBox, successBox, block, collSpan, music, bilibili, yBtn
        ],
        cache: { enable: false }, // 关闭本地缓存
        input () { that.$emit('update:html', editor.getHTML()); that.$emit('update:md', editor.getValue()) }, // 通过调用$emit方法自动更新父属性的值
        upload: { // 配置文件上传
          accept: 'image/*', // 只允许上传图片
          handler (files) { // 我们自己处理文件上传
            // 启动图片上传
            const data = new FormData()
            data.append('file', files[0])
            that.$store.dispatch('admin-file/uploadImage', data).then((data) => {
              editor.focus()
              editor.updateValue(`![](${data.url})`)
            })
          },
          multiple: false // 只允许上传一个文件
        },
        preview: {
          delay: 0 // 我们设置延时解析0秒，立即显示预览效果
        }
      }
    }
  },
  watch: {
    html (content) {
      try {
        // 注意因为文章数据是异步获取的，所以我们需要监听内容变化，注意这里要设置延时要不然会报错
        setTimeout(() => {
          // 只有当编辑器没有内容的时候才赋值
          if (editor.getValue().length === 1 && editor !== null && this.html !== '') {
            editor.setValue(this.md ? this.md : this.html)
          }
        }, 200)
      } catch (e) {
        this.$message.error('加载出现问题，请尝试关闭并重新进入！')
      }
    }
  },
  mounted () {
    // 初始化
    this.init()
    // 监听复制粘贴事件(注意这个要延迟一段时间要不然会无法监听成功)
    setTimeout(() => {
      $('.vditor-reset').on('paste', (e) => {
        const items = (e.clipboardData || e.originalEvent.clipboardData).items
        // 判断图片类型
        if (items.length !== 0 && items[0].type.includes('image')) {
          const file = items[0].getAsFile()
          const data = new FormData()
          data.append('file', file)
          this.$store.dispatch('admin-file/uploadImage', data).then((data) => {
            editor.focus()
            editor.updateValue(`![](${data.url})`)
          })
        }
      })
    }, 500)
  },
  methods: {
    // 设置内容
    setContent (content) {
      editor.setValue(content)
    },
    init () {
      // 因为直接在外部引用会有报错，所以我们只在mounted里面引用
      this.Vditor = require('vditor')
      // 引入css样式
      require('vditor/src/assets/scss/index.scss')
      editor = new this.Vditor('editor', this.option)
    }
  }
}
</script>

<style>
.vditor-tooltipped img{
  width: 20px;
  padding: 2px;
  margin-top: -5px;
}
.vditor-tooltipped img:hover{
  border: 1px solid #ddd;
  background: #eee;
}
</style>
