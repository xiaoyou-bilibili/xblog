import * as post from "./modules/posts"
import * as plugins from "./modules/plugins"
import * as setting from "./modules/settings"
import * as tools from "./modules/tools"
import * as user from "./modules/user"
export default {
  // 存储数据
  data: {
    post: post.state,
    plugins: plugins.state,
    setting: setting.state,
    tools: tools.state,
    user: user.state
  },
  // 存储函数
  post,
  plugins,
  setting,
  tools,
  user,
  //无脑全部更新，组件或页面不需要声明 use
  //updateAll: true,
  debug: true
}