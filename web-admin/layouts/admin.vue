<!--管理员界面的布局文件-->
<template>
  <client-only>
    <div>
      <side-bar-index />
      <div class="main-container" :class="{'container-open':sideStatus}">
        <div class="main-head">
          <navi-bar />
          <tags-view />
        </div>
        <div class="main-bg" />
        <nuxt class="layout-main" />
      </div>
    </div>
  </client-only>
</template>

<script>
import SideBarIndex from '@/components/content/admin/side-bar/index'
import NaviBar from '@/components/content/admin/head/navi-bar'
import TagsView from '@/components/content/admin/head/tags-view'
import { mapState } from 'vuex'
export default {
  name: 'AdminPage',
  components: { TagsView, NaviBar, SideBarIndex },
  computed: {
    ...mapState('admin', ['sideStatus'])
  },
  mounted () {
    // 恢复状态
    this.$store.dispatch('admin/restoreStatus')
    // 页面刷新时及时保存tag信息
    window.addEventListener('beforeunload', () => this.$store.dispatch('admin/saveStatus'))
  }
}
</script>

<style>
/*图片地址为空时自动隐藏img*/
img[src=""],img:not([src]){
  opacity:0;
}
/*清除输入框选中黑色*/
input{outline:none}
/*输入框界面*/
.el-input.el-input--prefix {
  margin-top: 10px;
}
</style>

<style scoped>
.container-open{
  margin-left: 64px!important;
}
.main-container{
  margin-left: 210px;
  height: 100%;
}
.main-bg{
  position: fixed;
  bottom: 0;
  top: 0;
  left: 0;
  right: 0;
  z-index: -1;
  background: #F5F6FA;
}
.layout-main{
  padding: 15px;
}
</style>
