// v3版本文章接口

import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'api/v3/admin/posts'

// 获取文章列表
export function getArticles (data) { return request(base + '/articles', data, 'get') }
// 管理员更新文章状态
export function updateArticles (id, data) { return request(base + `/articles/${id}`, data, 'put') }
// 获取文章内容
export function getPostContent (id) { return request(base + `/articles/${id}`, null, 'get') }
// 删除文章内容
export function deletePost (id) { return request(base + `/articles/${id}`, null, 'delete') }
// 新增文章
export function addPost (data) { return request(base + '/articles', data, 'post') }

// 邮件订阅的通知功能
export function noticeUser (id) { return request(base + `/${id}/subscription`, null, 'post') }
// 获取分类目录
export function getCategory (data) { return request(base + '/category', data, 'get') }
// 更新文章分类
export function updateCategory (id, data) { return request(base + `/category/${id}`, data, 'put') }
// 新增文章分类
export function addCategory (data) { return request(base + '/category', data, 'post') }
// 删除文章分类
export function deleteCategory (id) { return request(base + `/category/${id}`, null, 'delete') }

// 获取日记列表
export function getDiary (data) { return request(base + '/diary', data, 'get') }
// 获取日记内容
export function getDiaryContent (id) { return request(base + `/diary/${id}`, null, 'get') }
// 新建日记
export function addDiary (data) { return request(base + '/diary', data, 'post') }

// 获取删除的文章
export function getTrash (data) { return request(base + '/trash', data, 'get') }

// 获取所有文档
export function getDocs (data) { return request(base + '/docs', data, 'get') }
// 获取文档内容
export function getDocsContent (id) { return request(base + `/docs/${id}`, null, 'get') }
// 添加新文档
export function addDocs (data) { return request(base + '/docs', data, 'post') }
