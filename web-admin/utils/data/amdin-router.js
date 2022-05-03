// 管理员界面的路由配置
export default [
  {
    path: '/admin',
    icon: 'home',
    text: '首页'
  },
  {
    path: '/admin/post',
    icon: 'book-open',
    text: '文章管理',
    child: [
      {
        path: 'post-list',
        icon: 'list',
        text: '文章列表'
      },
      {
        path: 'diary-list',
        icon: 'calendar-alt',
        text: '日记列表'
      },
      {
        path: 'post-edit',
        icon: 'edit',
        text: '文章编辑器'
      },
      {
        path: 'diary-edit',
        icon: 'pen',
        text: '日记编辑器'
      },
      {
        path: 'tag-setting',
        icon: 'tag',
        text: '分类和标签设置'
      },
      {
        path: 'recycle',
        icon: 'recycle',
        text: '文章回收站'
      }
    ]
  },
  {
    path: '/admin/doc',
    icon: 'file-alt',
    text: '文档管理'
  },
  {
    path: '/admin/comment',
    icon: 'comment',
    text: '评论管理',
    child: [
      {
        path: 'comment-list',
        icon: 'comments',
        text: '评论列表'
      }
    ]
  },
  {
    path: '/admin/user',
    icon: 'user',
    text: '用户管理',
    child: [
      {
        path: 'user-list',
        icon: 'users',
        text: '用户列表'
      }
    ]
  },
  {
    path: '/admin/setting',
    icon: 'sliders-h',
    text: '设置管理',
    child: [
      {
        path: 'web-setting',
        icon: 'cogs',
        text: '网站设置'
      },
      {
        path: 'plugins',
        icon: 'plug',
        text: '插件设置'
      },
      {
        path: 'theme',
        icon: 'palette',
        text: '主题设置'
      },
      {
        path: 'side',
        icon: 'tools',
        text: '导航栏和侧边设置'
      },
      {
        path: 'wechat',
        icon: 'mobile-alt',
        text: '其他客户端设置'
      }
    ]
  }
]
