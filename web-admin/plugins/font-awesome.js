// 使用fontawesome图标
// https://github.com/FortAwesome/vue-fontawesome
import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

// 为了方便起见，这里直接导入所有的图标
library.add(fas, fab, far)
Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.config.productionTip = false
