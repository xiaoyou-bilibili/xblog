<template>
  <div>
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
      >
        <template v-slot:author="{ row }">
          <a :href="row.url" target="_blank">{{ row.author }}</a>
        </template>
        <!-- 操作模板 -->
        <template v-slot:option="{row}">
          <vxe-button v-if="row.version !== row.new_version" icon="fa fa-refresh" title="更新" circle @click="update(row.download_url)" />
          <vxe-button icon="fa fa-cog" title="设置" circle @click="pluginsSetting(`/admin/setting/theme/${row.dir}`, row.name)" />
          <vxe-button v-show="!row.enable" icon="fa fa-check" title="启用" circle @click="chooseTheme(row.dir)" />
          <!--          &lt;!&ndash;删除文章按钮 onConfirm 点击确认按钮激发&ndash;&gt;-->
          <!--          <el-popconfirm title="确定要删除这个插件？" @onConfirm="deletePlugin(row.id)">-->
          <!--            <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />-->
          <!--          </el-popconfirm>-->
        </template>
      </vxe-grid>
    </el-card>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { getThemes } from '@/api/admin/theme'
export default {
  layout: 'admin',
  mixins: [admin], // 混入对象用于批量修改文章状态
  data () {
    return {
      // 使用动态代理数据
      tableProxy: {
        ajax: {
          // 任何支持 Promise API 的库都可以对接（fetch、jquery、axios、xe-ajax）
          query: () => getThemes(),
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
        zoom: true, // 缩放
        custom: true // 自定义显示内容
      },
      // 表格结构
      tableColumn: [
        { field: 'name', title: '主题名字' },
        { field: 'description', title: '主题描述' },
        { field: 'author', title: '开发作者', slots: { default: 'author' } },
        { field: 'version', title: '主题版本' },
        { title: '操作', slots: { default: 'option' } }
      ]
    }
  },
  methods: {
    chooseTheme (path) {
      this.$store.dispatch('admin-theme/chooseTheme', path)
        .then(() => {
          this.$message.success('切换主题成功')
          this.$refs.grid.commitProxy('query')
        }).catch(msg => this.$message.error(msg))
    },
    // 进入主题设置界面
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
      this.$store.dispatch('admin-theme/downloadTheme', { url }).then(() => { this.$message.success('更新成功！'); loading.close(); this.$refs.grid.commitProxy('query') }).catch((msg) => { this.$message.error(msg); loading.close() })
    }
  }
}
</script>

<style scoped>

</style>
