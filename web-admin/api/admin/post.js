// 管理员文章管理接口

import request from '~/utils/request'

// 基本路径
const base = process.env.SERVER + 'api/v2/admin_post/'

// 获取文章列表
export function getPostList (data) { return request(base + 'get/post_list', data, 'get') }
// 管理员更新文章状态
export function updatePost (data) { return request(base + 'update/post', data, 'post') }
// 邮件订阅的通知功能
export function noticeUser (data) { return request(base + 'notice/user', data, 'post') }
// 获取日记列表
export function getDiaryList (data) { return request(base + 'get/diary_list', data, 'get') }
// 获取文章内容
export function getPostContent (data) { return request(base + 'get/post_content', data, 'get') }
// 更新文章的状态
export function updatePostContent (data) { return request(base + 'update/post_content', data, 'post') }
// 获取文章内容
export function getDiaryContent (data) { return request(base + 'get/diary_content', data, 'get') }
// 更新文章内容
export function updateDiaryContent (data) { return request(base + 'update/diary_content', data, 'post') }
// 获取分类目录
export function getCategoryList (data) { return request(base + 'get/category_list', data, 'get') }
// 更新分类目录
export function updateCategoryList (data) { return request(base + 'update/category_list', data, 'post') }
// 获取删除的文章列表
export function getDeleteList (data) { return request(base + 'get/delete_list', data, 'get') }
// 更新删除的文章信息
export function updateDeleteList (data) { return request(base + 'update/delete_list', data, 'post') }
