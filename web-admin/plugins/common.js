// 这里注册一些常用的组件
import Vue from 'vue'

// 使用v-lazy
import VueLazyLoad from 'vue-lazyload'

// 使用全局变量
import global from '@/utils/const.js'

// 注册工具路由
import tools from '@/utils/tools'

// 配置table
import 'xe-utils'
import VXETable from 'vxe-table'
import 'vxe-table/lib/style.css'

// 配置图标选择组件
import iconPicker from 'e-icon-picker'
import 'e-icon-picker/dist/symbol.js' // 基本彩色图标库
import 'e-icon-picker/dist/index.css' // 基本样式，包含基本图标
import 'font-awesome/css/font-awesome.min.css' // font-awesome 图标库
import 'element-ui/lib/theme-chalk/icon.css' // element-ui 图标库

Vue.use(iconPicker, { FontAwesome: true, ElementUI: true, eIcon: true, eIconSymbol: true })

Vue.use(global)
Vue.use(tools)
// v-lazy配置
Vue.use(VueLazyLoad, {
  preLoad: 1.3,
  error: '',
  loading: '',
  attempt: 2
})
// table
Vue.use(VXETable)
// 挂载弹框
Vue.prototype.$XModal = VXETable.modal
// 表格初始化
VXETable.setup({
  grid: {
    // 自定义动态代理的字段信息
    proxyConfig: {
      // 代理表单搜索
      form: true,
      // 代理结果的一些属性
      props: {
        result: 'data.contents',
        total: 'data.total_num'
      }
    }
  },
  // 分页配置
  pager: {
    perfect: true,
    pageSize: 10,
    pagerCount: 7,
    pageSizes: [5, 10, 15, 20, 50],
    layouts: ['PrevJump', 'PrevPage', 'Jump', 'PageCount', 'NextPage', 'NextJump', 'Sizes', 'Total']
  }
})
// 使用全局格式化函数
VXETable.formats.mixin({
  // 格式化文章状态
  formatPostType ({ cellValue }) {
    switch (cellValue) {
      case 'publish':
        return '公开'
      case 'private':
        return '私有'
      case 'encrypt':
        return '加密'
    }
    return cellValue
  },
  // 格式化文章类型
  formatPostKinds ({ cellValue }) {
    switch (cellValue) {
      case 'post':
        return '文章'
      case 'diary':
        return '日记'
      case 'doc':
        return '文档'
    }
    return cellValue
  },
  // 格式化文章状态
  formatPostStatus ({ cellValue }) { return cellValue ? '草稿' : '已发布' },
  // 格式化审核内容
  formatReview ({ cellValue }) { return cellValue === 1 ? '已通过' : '已拒绝' },
  // 格式化用户状态
  formatUserStatus ({ cellValue }) { return cellValue === 1 ? '已激活' : '未激活' },
  // 格式化用户订阅
  formatUserSubscription ({ cellValue }) { return cellValue ? '已订阅' : '未订阅' },
  // 格式化置顶信息
  formatTop ({ cellValue }) { return cellValue ? '是' : '否' },
  // 格式化权限信息
  formatIdentity ({ cellValue }) { return cellValue === 1 ? '管理员' : '普通用户' }
})
