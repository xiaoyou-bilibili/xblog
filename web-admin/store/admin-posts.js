// v3版本api 文章管理

// 和管理员文章相关的vuex管理
import { requestProcess } from '@/utils/request-process-v3'
import {
  updateArticles,
  noticeUser,
  getPostContent,
  getCategory,
  addPost,
  deletePost,
  getDiaryContent,
  addDiary,
  updateCategory,
  addCategory,
  deleteCategory,
  getDocs,
  addDocs,
  getDocsContent
} from '@/api/admin/posts'

import { DataProcess } from '@/utils/tree'
// 全局参数
export const state = () => ({
  // 文章分类(用于文章编辑器)
  category: [{
    value: 0, // 节点值
    label: '', // 节点名字
    children: [{ value: 0, label: '' }]
  }],
  categoryList: [], // 文章分类，用于分类设置
  deleteList: [{ // 删除的文章信息
    id: 0,
    title: '',
    content: '',
    type: ''
  }],
  docList: [],
  docContent: {
    title: '', // 标题
    content: '' // 内容
  }
})

// 属性获取
export const getters = {
  category (state) { return state.category },
  categoryList (state) { return state.categoryList },
  deleteList (state) { return state.deleteList },
  docList (state) { return state.docList },
  docContent (state) { return state.docContent }
}

// 数据修改
export const mutations = {
  // 对获取的分类进行处理转换为element可以识别的
  setCategory (state, data) {
    const categoryS = []
    // 遍历父节点
    // todo 优化计算算法
    Object.values(data).forEach((item) => {
      if (item.parent === 0) {
        const category = {
          value: item.id,
          label: item.title,
          children: []
        }
        // 遍历子节点获取该父节点下的数据
        Object.values(data).forEach((item2) => {
          if (item2.parent === item.id) {
            category.children.push({
              value: item2.id,
              label: item2.title
            })
          }
        })
        categoryS.push(category)
      }
    })
    state.category = categoryS
  },
  // 修改分类和标签设置里面的分类
  setCategoryList (state, data) { state.categoryList = DataProcess(data) },
  // 分类列表添加值
  pushCategoryList (state, data) { state.categoryList.push(data) },
  // 设置文档列表
  setDocList (state, data) { state.docList = DataProcess(data) },
  // 设置文档内容
  setDocContent (state, data) { state.docContent = data },
  // 根节点添加子节点数据
  docListPush (state, data) { state.docList.push(data) }
}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  // 更新文章状态;
  updatePost ({ commit }, data) { return requestProcess(updateArticles, none, none, data.id, data.data) },
  // 获取文章内容
  getPostContent ({ commit }, data) { return requestProcess(getPostContent, none, none, data) },
  // 删除文章内容
  deletePost ({ commit }, data) { return requestProcess(deletePost, none, none, data) },
  // 添加文章
  addPost ({ commit }, data) { return requestProcess(addPost, none, none, data) },
  // 邮件订阅通知功能
  noticeUser ({ commit }, data) { return requestProcess(noticeUser, none, none, data) },
  // 获取文章分类
  getCategory ({ commit }) { return requestProcess(getCategory, res => commit('setCategory', res)) },
  // 获取日记内容
  getDiaryContent ({ commit }, data) { return requestProcess(getDiaryContent, none, none, data) },
  // 新增日记
  addDiary ({ commit }, data) { return requestProcess(addDiary, none, none, data) },
  // 获取文章分类（用于分类和标签设置）
  getCategoryList ({ commit }) { return requestProcess(getCategory, data => commit('setCategoryList', data)) },
  // 更新文章分类（用于分类和标签设置）
  updateCategoryList ({ commit }, data) { return requestProcess(updateCategory, none, none, data.id, data.data) },
  // 添加文章分类（用于分类和标签设置）
  addCategory ({ commit }, data) { return requestProcess(addCategory, none, none, data) },
  // 删除文章分类
  deleteCategory ({ commit }, data) { return requestProcess(deleteCategory, none, none, data) },
  // 获取所有文档
  getDocs ({ commit }, data) { return requestProcess(getDocs, data => commit('setDocList', data), none, data) },
  // 获取文章内容
  getDocsContent ({ commit }, data) { return requestProcess(getDocsContent, data => commit('setDocContent', data), none, data) },
  // 添加新文档
  addDocs ({ commit }, data) { return requestProcess(addDocs, none, none, data) }
}
