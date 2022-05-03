// 管理员后台接口返回模型
export interface AdminPostList{
  id: number, // 文章id
  title: string, // 文章标题
  content: string, // 文章内容
  date: string, // 发布时间
  good: number, // 点赞数
  view: number, // 浏览量
  comment: number, // 评论数
  status: string, // 文章状态
  category: Array<String>, // 文章分类
  tags: Array<string>, // 文章标签
  is_draft: Boolean // 是否是草稿
}

// 列表数据
export interface List<T> {
  total: 0, // 文章页数
  current: 0, // 当前页数
  contents: Array<T>
}
