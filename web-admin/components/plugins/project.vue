<template>
  <div>
    <vxe-grid
      ref="grid"
      border
      resizable
      height="780"
      row-id="id"
      :pager-config="{pageSize: 10}"
      :proxy-config="tableProxy"
      :form-config="tableForm"
      :toolbar-config="tableToolbar"
      :columns="tableColumn"
      :edit-config="editConfig"
      @checkbox-change="checkboxChangeEvent"
      @checkbox-all="checkboxChangeEvent"
    >
      <!-- 项目图片 -->
      <template v-slot:img="{ row }">
        <img class="project-img" :src="row.img">
      </template>
      <!-- 操作模板 -->
      <template v-slot:option="{ row }">
        <!-- 编辑保存按钮 -->
        <vxe-button
          v-if="$refs.grid.isActiveByRow(row)"
          icon="fa fa-save"
          status="primary"
          title="保存"
          circle
          @click="saveRowEvent(row)"
        />
        <vxe-button v-else icon="fa fa-edit" title="编辑" circle @click="editRowEvent(row)" />
        <!--删除按钮 onConfirm 点击确认按钮激发-->
        <el-popconfirm title="确定要删除这个项目？" @onConfirm="deleteProject(row._id)">
          <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
        </el-popconfirm>
      </template>
    </vxe-grid>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { pluginRequest } from '@/api/admin/plugins'
// 定义友链状态
const friendReview = [{ label: '是', value: true }, { label: '否', value: false }]
export default {
  layout: 'admin',
  mixins: [admin], // 混入对象用于批量修改文章状态
  data () {
    return {
      // 多选框选中的值
      selects: '',
      // 使用动态代理数据
      tableProxy: {
        ajax: {
          // 任何支持 Promise API 的库都可以对接（fetch、jquery、axios、xe-ajax）
          query: ({ page, form }) => pluginRequest({ url: '/projects', type: 'get', data: { page: page.currentPage, page_size: page.pageSize, search_type: form.search_type, search_key: form.search_key } }),
          // 删除方法被触发
          delete: _ => this.deleteProject(this.selects)
        }
      },
      // 搜索表单设置
      tableForm: {
        items: [
          {
            field: 'search_type',
            title: '搜索类型',
            itemRender: {
              name: '$select',
              options: [
                { label: '项目名字', value: 'name' },
                { label: '项目介绍', value: 'description' },
                { label: '是否置顶', value: 'is_top' }
              ]
            }
          },
          // visibleMethod用来控制控件是否显示
          { field: 'search_key', visibleMethod: e => e.data.search_type !== 'is_top', title: '关键词', itemRender: { name: '$input', attrs: { placeholder: '请输入关键词' } } },
          { field: 'search_key', visibleMethod: e => e.data.search_type === 'is_top', title: '置顶', itemRender: { name: '$select', options: friendReview } },
          { itemRender: { name: '$button', props: { content: '查询', type: 'submit', status: 'primary' } } },
          { itemRender: { name: '$button', props: { content: '重置', type: 'reset' } } }
        ]
      },
      // 工具栏设置
      tableToolbar: {
        buttons: [
          { code: 'insert_actived', name: '添加', icon: 'fa fa-plus', status: 'primary' },
          { code: 'delete', name: '删除选中', icon: 'fa fa-trash', status: 'danger' }
        ],
        zoom: true, // 缩放
        custom: true // 自定义显示内容
      },
      // 开启行编辑功能
      editConfig: {
        // 手动触发
        trigger: 'manual',
        // 单元格编辑模式
        mode: 'row',
        // 点击其他地方不清除激活状态
        autoClear: false,
        // 不显示按钮
        icon: 'none'
      },
      // 表格结构
      tableColumn: [
        { type: 'checkbox', width: 50 },
        { field: 'name', title: '项目名称', editRender: { name: 'input' } },
        { field: 'img', width: 180, title: '项目图片', slots: { default: 'img' }, editRender: { name: 'input' } },
        { field: 'make_time', title: '制作时间', editRender: { name: '$input', props: { type: 'date' } } },
        { field: 'description', title: '项目描述', editRender: { name: 'textarea' } },
        { field: 'video_url', title: '视频地址', editRender: { name: 'input' } },
        { field: 'blog_url', title: '博客地址', editRender: { name: 'input' } },
        { field: 'code_url', title: '代码地址', editRender: { name: 'input' } },
        { field: 'link', title: '轮播图地址', editRender: { name: 'input' } },
        { field: 'is_top', title: '轮播图置顶', formatter: 'formatTop', width: 100, editRender: { name: '$select', options: friendReview } },
        { title: '操作', slots: { default: 'option' } }
      ]
    }
  },
  methods: {
    // 多选框改变时触发的事件
    checkboxChangeEvent (data) {
      const ids = []
      // 把我们选中值的id提取出来
      data.records.map(item => ids.push(item.id))
      this.selects = ids.toString()
    },
    // 点击编辑事件
    editRowEvent (row) {
      // 我们激活编辑框
      this.$refs.grid.setActiveRow(row)
    },
    // 点击保存事件
    saveRowEvent (row) {
      // 点击保存的时候自动触发代理的保存事件
      this.$refs.grid.clearActived().then(() => {
        // 判断是新增还是保存
        if (row._id.includes('row')) {
          this.$store.dispatch('admin-plugins/pluginRequest', { url: '/projects', type: 'post', data: row }).then(_ => this.updateData()).catch((msg) => { this.$message.error(msg); this.updateData(false) })
        } else {
          this.$store.dispatch('admin-plugins/pluginRequest', { url: `/projects/${row._id}`, type: 'put', data: row }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
        }
      })
    },
    // 提示更新成功并刷新数据
    updateData (show = true) {
      // 是否需要提示更新信息
      if (show) { this.$message.success('更新成功') }
      // 调用commitProxy 手动更新数据
      this.$refs.grid.commitProxy('query')
    },
    // 删除项目
    deleteProject (id) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        // 删除评论
        this.$store.dispatch('admin-plugins/pluginRequest', { url: `/projects/${id}`, type: 'delete', data: null }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
      }
    }
  }
}
</script>

<style scoped>
img.project-img {
  width: 100%;
  height: 100px;
  border-radius: 5px;
}

.vxe-button+.vxe-button{
  margin-left: 0!important;
}
</style>
