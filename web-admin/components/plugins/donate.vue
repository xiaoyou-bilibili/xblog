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
        <el-popconfirm title="确定要删除这条赞助？" @onConfirm="deleteDonate(row.id)">
          <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
        </el-popconfirm>
      </template>
    </vxe-grid>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { pluginRequest } from '@/api/admin/plugins'

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
          query: ({ page, form }) => pluginRequest({ url: '/sponsors', type: 'get', data: { page: page.currentPage, page_size: page.pageSize, search_type: form.search_type, search_key: form.search_key } }),
          // 删除方法被触发
          delete: _ => this.deleteDonate(this.selects)
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
                { label: '赞助昵称', value: 'nickname' },
                { label: '赞助金额', value: 'amount' },
                { label: '赞助留言', value: 'comment' }
              ]
            }
          },
          // visibleMethod用来控制控件是否显示
          { field: 'search_key', title: '关键词', itemRender: { name: '$input', attrs: { placeholder: '请输入关键词' } } },
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
        { field: 'nickname', title: '赞助昵称', editRender: { name: 'input' } },
        { field: 'amount', title: '赞助金额', editRender: { name: 'input' } },
        { field: 'comment', title: '留言', editRender: { name: 'input' } },
        { field: 'donate_time', title: '赞助时间', editRender: { name: '$input', props: { type: 'date' } } },
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
        // 设置更新的值
        const data = {
          nickname: row.nickname,
          comment: row.comment,
          amount: row.amount,
          donate_time: row.donate_time
        }
        // 判断是新增还是保存
        if (row.id.includes('row')) {
          this.$store.dispatch('admin-plugins/pluginRequest', { url: '/sponsors', type: 'post', data }).then(_ => this.updateData()).catch((msg) => { this.$message.error(msg); this.updateData(false) })
        } else {
          this.$store.dispatch('admin-plugins/pluginRequest', { url: `/sponsors/${row.id}`, type: 'put', data }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
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
    // 删除赞助
    deleteDonate (id) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        // 删除赞助
        this.$store.dispatch('admin-plugins/pluginRequest', { url: `/sponsors/${id}`, type: 'delete', data: null }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
      }
    }
  }
}
</script>

<style scoped>

.vxe-button+.vxe-button{
  margin-left: 0!important;
}
</style>
