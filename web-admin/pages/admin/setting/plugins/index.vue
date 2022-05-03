<template>
  <div>
    <input
      id="chooseImg"
      type="file"
      accept="application/zip"
      name="file"
      class="hidden"
      @change="selectImg"
    >
    <el-card class="box-card">
      <vxe-grid
        ref="grid"
        border
        resizable
        height="780"
        row-id="id"
        :pager-config="{pageSize: 10}"
        :proxy-config="tableProxy"
        :toolbar-config="tableToolbar"
        :columns="tableColumn"
        @toolbar-button-click="toolbarClick"
      >
        <template v-slot:author="{ row }">
          <a :href="row.site" target="_blank">{{ row.author }}</a>
        </template>
        <!-- 操作模板 -->
        <template v-slot:option="{row}">
          <vxe-button v-if="row.version !== row.new_version" icon="fa fa-refresh" title="更新" circle @click="update(row.download_url)" />
          <vxe-button icon="fa fa-cog" title="设置" circle @click="pluginsSetting(`/admin/setting/plugins/${row.id}`, row.name)" />
          <!--删除文章按钮 onConfirm 点击确认按钮激发-->
          <el-popconfirm title="确定要删除这个插件？" @onConfirm="deletePlugin(row.id)">
            <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
          </el-popconfirm>
        </template>
      </vxe-grid>
    </el-card>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { getAllPlugins } from '@/api/admin/plugins'
import $ from 'jquery'
export default {
  layout: 'admin',
  mixins: [admin], // 混入对象用于批量修改文章状态
  data () {
    return {
      // 使用动态代理数据
      tableProxy: {
        ajax: {
          // 任何支持 Promise API 的库都可以对接（fetch、jquery、axios、xe-ajax）
          query: () => getAllPlugins(),
          // 删除方法被触发
          delete: _ => this.updateComment(this.selects, null, true)
        },
        // 代理结果的一些属性
        props: {
          result: 'data'
        }
      },
      // 工具栏设置
      tableToolbar: {
        buttons: [
          { code: 'add', name: '安装插件', icon: 'fa  fa-cloud-upload', status: 'success' },
          { code: 'reload', name: '重载插件', icon: 'fa  fa-refresh', status: 'warning' }
        ],
        zoom: true, // 缩放
        custom: true // 自定义显示内容
      },
      // 表格结构
      tableColumn: [
        { field: 'name', title: '插件名字' },
        { field: 'description', title: '插件描述' },
        { field: 'author', title: '开发作者', slots: { default: 'author' } },
        { field: 'version', title: '插件版本' },
        { title: '操作', slots: { default: 'option' } }
      ]
    }
  },
  methods: {
    // 工具栏按钮点击事件
    toolbarClick ({ code, button }) {
      // 根据我们的code来判断不同操作
      switch (code) {
        case 'add':
          this.uploadImage()
          break
        case 'reload':
          this.reloadPlugins()
          break
      }
    },
    selectImg (e) {
      // 先获取input里面的文件
      const files = e.target.files || e.dataTransfer.files
      // 判断文件是否为空
      if (!files.length) { return }
      // 设置图片的内容和地址
      this.picValue = files[0]
      this.url = this.getObjectURL(this.picValue)
      // 上传图片
      if (this.picValue !== null && this.picValue !== undefined) {
        // 初始化一个formData对象(这个对象可以用于图片上传)
        const fromData = new FormData()
        // 放入我们的图片
        fromData.append('file', this.picValue)
        // 开始上传图片
        this.$store.dispatch('admin-plugins/uploadPlugins', fromData).then((data) => {
          this.$message.success('安装成功')
          // 调用commitProxy 手动更新数据
          this.$refs.grid.commitProxy('query')
        }).catch(msg => this.$message.error(msg))
      }
    },
    uploadImage () {
      $('#chooseImg').click()
    },
    // 获取图片的地址（不同的浏览器图片地址不同）
    getObjectURL (file) {
      let url = null
      if (window.createObjectURL !== undefined) { // basic
        url = window.createObjectURL(file)
      } else if (window.URL !== undefined) { // mozilla(firefox)
        url = window.URL.createObjectURL(file)
      } else if (window.webkitURL !== undefined) { // webkit or chrome
        url = window.webkitURL.createObjectURL(file)
      }
      return url
    },
    // 删除插件
    deletePlugin (id) {
      this.$store.dispatch('admin-plugins/deletePlugins', id).then(() => {
        this.$message.success('删除成功')
        this.$refs.grid.commitProxy('query')
      }).catch(msg => this.$message.error(msg))
    },
    // 重新加载插件
    reloadPlugins () {
      this.$store.dispatch('admin-plugins/reloadPlugins').then(() => {
        this.$message.success('重载成功')
        this.$refs.grid.commitProxy('query')
      }).catch(msg => this.$message.error(msg))
    },
    // 进入插件设置界面
    pluginsSetting (path, name) {
      // 跳转路由
      this.$router.push(path)
      // 添加标签
      this.$store.dispatch('admin/addTag', { name: `${name}-设置`, path })
    },
    update (url) {
      const loading = this.$loading({
        lock: true,
        text: '正在更新',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      this.$store.dispatch('admin-plugins/downloadPlugins', { url })
        .then(() => { this.$message.success('更新成功！'); this.$refs.grid.commitProxy('query') })
        .catch((msg) => { this.$message.error(msg) })
        .finally(_ => loading.close())
    }
  }
}
</script>

<style scoped>
#chooseImg{
  display: none;
}
.vxe-button+.vxe-button{
  margin-left: 0!important;
}
</style>
