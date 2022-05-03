// 和管理员相关的vuex
import Cookie from 'js-cookie'
// 全局参数
export const state = () => ({
  // 侧边栏状态(false 为开启 true为关闭)
  sideStatus: false,
  tags: [{ // 当前任务栏所有的任务
    active: true, // 是否激活
    title: '主页', // 任务栏标题
    close: false, // 是否可以关闭
    path: '/admin' // 当前标签的路径
  }]
})

// 属性获取
export const getters = {
  sideStatus (state, data) { return state.sideStatus },
  tags (state, data) { return state.tags }
}

// 数据修改
export const mutations = {
  changeSideStatue (state) { state.sideStatus = !state.sideStatus },
  // 添加新的标签
  addNewTag (state, data) {
    // 查找标签是否存在
    if (state.tags.find(item => item.path === data.path) === undefined) {
      // 所有的标签全部为未激活状态
      state.tags.map(item => (item.active = false))
      // 数组添加新的标签
      state.tags.push(data)
    }
  },
  // 更据path来激活对应的标签
  changeTag (state, path) {
    // map遍历数组，更据路径来切换激活状态
    // 这里加了一个=>(),()表示函数体返回对象字面变量
    state.tags.map(item => (item.active = path === item.path))
  },
  // 删除某一个标签
  delTag (state, index) { state.tags.splice(index, 1) },
  // 这里我们利用数组的过滤来删除我们不需要的袁术
  closeOther (state, index) {
    state.tags = state.tags.filter((value, index1, array) => (index1 === index || !value.close))
    // 设置最后一个标签状态为激活
    state.tags[state.tags.length - 1].active = true
  },
  closeAll (state) {
    state.tags = state.tags.filter((value, index1, array) => (!value.close))
    state.tags[0].active = true
  },
  // 保存当前的tag
  saveStatus (state) {
    Cookie.set('tags', state.tags)
  },
  // 恢复当前的tag
  restore (state) {
    if (Cookie.get('tags') !== undefined) {
      state.tags = JSON.parse(Cookie.get('tags'))
    }
  },
  // 文章编辑器切换id
  editChangeID (state, data) {
    // 查找标签是否存在
    if (state.tags.find(item => item.title === data.title) === undefined) {
      // 所有的标签全部为未激活状态
      state.tags.map(item => (item.active = false))
      // 数组添加新的标签
      state.tags.push(data)
    } else {
      // 主动修改path
      state.tags.map(item => (item.path = item.title === data.title ? data.path : item.path))
    }
  },
  // 切换当前标签栏为普通用户的主页
  switchNormal (state) { state.tags = [{ active: true, title: '主页', close: false, path: '/admin/user/normal' }] },
  // 删除当前激活的标签
  delActiveTag (state) { state.tags = state.tags.filter(item => !item.active) }
}

// 函数调用
export const actions = {
  changeSideBar ({ commit }) { commit('changeSideStatue') },
  addTag ({ commit }, data) { commit('addNewTag', { active: true, title: data.name, close: true, path: data.path }) },
  changeTag ({ commit }, data) { commit('changeTag', data) },
  // 关闭标签
  closeTag ({ commit, state }, data) {
    return new Promise((resolve, reject) => {
      // 为了避免commit里面修改数据的种种问题，所以直接在action中操作
      // 这里是因为不能在action中操作数据，所以我们使用副本
      const tags = state.tags.slice()
      if (tags[data].active) {
        // 判断当前下标是否为最后一个
        let index = tags.length - 1
        if (data === tags.length - 1) {
          index = tags.length - 2
        }
        resolve(tags[index].path)
      } else {
        reject(data)
      }
      commit('delTag', data)
    })
  },
  // 关闭其他标签
  closeOther ({ commit }, data) { commit('closeOther', data) },
  // 关闭所有标签
  closeAll ({ commit }) { commit('closeAll') },
  // 保存当前tag
  saveStatus ({ commit }) { commit('saveStatus') },
  // 保存所有tag
  restoreStatus ({ commit }) { commit('restore') },
  // 编辑器切换id操作(文章编辑器，日记编辑器)
  editChangeId ({ commit }, data) { commit('editChangeID', { active: true, title: data.title, close: true, path: data.path }) },
  // 关闭当前激活的标签
  deleteActiveTag ({ commit }) { commit('delActiveTag') }
}
