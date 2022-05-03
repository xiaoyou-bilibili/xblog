<template>
  <div>
    <vxe-grid
      ref="grid"
      border
      resizable
      height="700"
      row-id="id"
      :pager-config="{pageSize: 10}"
      :proxy-config="tableProxy"
      :toolbar-config="tableToolbar"
      :columns="tableColumn"
      :edit-config="editConfig"
      @checkbox-change="checkboxChangeEvent"
      @checkbox-all="checkboxChangeEvent"
    >
      <!-- 选择分类 -->
      <template v-slot:toolbar_buttons>
        <!-- 分类选择 -->
        <el-select v-model="choose" placeholder="请选择" @change="getItems">
          <el-option
            v-for="item in navigation"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
        <vxe-button status="primary" icon="fa fa-plus" @click="changeNavigation(true)">
          添加分类
        </vxe-button>
        <el-popconfirm title="确定要删除这条分类？" @onConfirm="deleteNavigation">
          <vxe-button slot="reference" status="danger" icon="fa fa-trash">
            删除分类
          </vxe-button>
        </el-popconfirm>
        <vxe-button status="info" icon="fa fa-edit" @click="changeNavigation(false)">
          修改分类
        </vxe-button>
        <vxe-button status="success" icon="fa fa-link" @click="addLink">
          添加网站
        </vxe-button>
        <el-popconfirm title="确定要删除？" @onConfirm="deleteDonate(selects)">
          <vxe-button slot="reference" status="danger" icon="fa fa-trash">
            删除选中网站
          </vxe-button>
        </el-popconfirm>
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
        <el-popconfirm title="确定要删除这条链接？" @onConfirm="deleteDonate(row.id)">
          <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
        </el-popconfirm>
      </template>
    </vxe-grid>
    <!--添加新分类-->
    <el-dialog
      :close-on-click-modal="false"
      :title="addOption?'添加分类':'修改分类'"
      :visible.sync="addCategoryDialog"
      width="350px"
    >
      <el-form ref="form" label-width="80px">
        <el-form-item label="分类名字">
          <el-input v-model="category.name" />
        </el-form-item>
        <el-form-item label="分类颜色">
          <div style="display: flex">
            <el-input v-model="category.value" />
            <el-color-picker v-model="category.value" style="margin-left: 3px" />
          </div>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addCategoryDialog = false">取 消</el-button>
        <el-button type="primary" @click="addNavigationCategory">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { pluginRequest } from '@/api/admin/plugins'
let navigationID = -1
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
          query: ({ page, form }) => pluginRequest({ url: `/navigation/${navigationID}/links`, type: 'get', data: { page: page.currentPage, page_size: page.pageSize } }),
          // 删除方法被触发
          delete: _ => this.deleteDonate(this.selects)
        }
      },
      // 工具栏设置
      tableToolbar: {
        // 使用自定义工具栏
        slots: {
          buttons: 'toolbar_buttons'
        },
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
        { field: 'id', title: '导航id', width: 100 },
        { field: 'name', title: '链接名', editRender: { name: 'input' } },
        { field: 'url', title: '链接地址', editRender: { name: 'input' } },
        { title: '操作', slots: { default: 'option' } }
      ],
      // 选择器选择
      choose: '',
      // 添加网站dialog
      addCategoryDialog: false,
      // 添加网站的内容
      category: {
        name: '',
        value: '',
        parent: 0
      },
      // 如果为true就是添加，否则就是删除
      addOption: true,
      // 所有导航值
      navigation: []
    }
  },
  mounted () {
    this.updateNavigation(false, false)
  },
  methods: {
    // 更新导航(update 表示更新链接数据, show 显示提示信息)
    updateNavigation (update = false, show = true) {
      this.addCategoryDialog = false
      this.$store.dispatch('admin-plugins/pluginRequest', { url: '/navigation', type: 'get' })
        .then((data) => { this.navigation = data })
      if (show) { this.$message.success('操作成功') }
      if (update) {
        this.choose = ''
        navigationID = 0
        this.updateData(false)
      }
    },
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
          name: row.name,
          value: row.url,
          parent: navigationID
        }
        // 判断是新增还是保存(如果是新增，那么id就是字符串)
        if (isNaN(row.id)) {
          this.$store.dispatch('admin-plugins/pluginRequest', { url: '/navigation', type: 'post', data })
            .then(_ => this.updateData())
            .catch((msg) => { this.$message.error(msg); this.updateData(false) })
        } else {
          this.$store.dispatch('admin-plugins/pluginRequest', { url: '/navigation/' + row.id, type: 'put', data })
            .then(_ => this.updateData())
            .catch(msg => this.$message.error(msg))
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
    // 删除链接
    deleteDonate (id) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        // 删除导航
        this.$store.dispatch('admin-plugins/pluginRequest', { url: '/navigation/' + id, type: 'delete' })
          .then(() => { this.updateData() })
          .catch(msg => this.$message.error(msg))
      }
    },
    // 获取导航下的网址
    getItems (id) {
      // 设置id并更新数据
      navigationID = id
      this.updateData(false)
    },
    // 修改分类
    changeNavigation (add) {
      // 判断是添加还是更新操作
      if (add) {
        this.addOption = true
        this.category.name = ''
        this.category.value = ''
      } else {
        // 检查是否选择分类
        if (this.checkCategory()) { return }
        // 如果是更新，我们从navigation中找出需要修改的值
        const data = this.navigation
        data.forEach((item) => {
          if (item.id === this.choose) {
            this.category.name = item.name
            this.category.value = item.color
          }
        })
        this.addOption = false
      }
      this.addCategoryDialog = true
    },
    // 添加分类
    addNavigationCategory () {
      const category = this.category
      // 参数验证
      if (category.name === '' || category.value === '') {
        this.$message.warning('请填写内容和颜色')
        return
      }
      if (this.addOption) {
        this.$store.dispatch('admin-plugins/pluginRequest', { url: '/navigation', type: 'post', data: category })
          .then(() => { this.updateNavigation(true) })
          .catch(msg => this.$message.error(msg))
      } else {
        this.$store.dispatch('admin-plugins/pluginRequest', { url: '/navigation/' + navigationID, type: 'put', data: category })
          .then(() => { this.updateNavigation(false) })
          .catch(msg => this.$message.error(msg))
      }
    },
    // 检查分类
    checkCategory () {
      if (navigationID === 0) {
        this.$message.warning('请先选择分类')
        return true
      }
      return false
    },
    // 删除分类
    deleteNavigation () {
      // 检查是否选择分类
      if (this.checkCategory()) { return }
      // 删除分类
      this.$store.dispatch('admin-plugins/pluginRequest', { url: '/navigation/' + navigationID, type: 'delete' })
        .then(() => { this.updateNavigation(true) })
        .catch(msg => this.$message.error(msg))
    },
    // 添加网址
    addLink () {
      // 检查是否选择分类
      if (this.checkCategory()) { return }
      // 手动添加临时数据并激活
      this.$refs.grid.insert().then(({ row }) => this.$refs.grid.setActiveRow(row))
    }
  }
}
</script>

<style scoped>

.vxe-button+.vxe-button{
  margin-left: 0!important;
}
</style>
