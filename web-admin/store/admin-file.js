// 上传相关的vuex管理

// 全局参数
import { requestProcess } from '@/utils/request-process-v3'
import { uploadImage } from '@/api/admin/file'

export const state = () => ({

})

// 属性获取
export const getters = {

}

// 数据修改
export const mutations = {

}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  uploadImage ({ commit }, data) { return requestProcess(uploadImage, none, none, data) }
}
