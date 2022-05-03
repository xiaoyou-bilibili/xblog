<template>
  <div class="side-contain" :class="{'sideBarOpen':sideStatus}">
    <sidebar-logo :collapse="sideStatus" />
    <el-menu class="side-menu el-menu-vertical-demo" :collapse="sideStatus">
      <side-menu v-for="(item,index) in (userInfo.identity === 1 ? adminRouter : userRouter)" :key="index" :index-key="index" :item="item" />
    </el-menu>
  </div>
</template>

<script>
import adminRouter from '@/utils/data/amdin-router'
import userRouter from '@/utils/data/user-router'
import SideMenu from '@/components/content/admin/side-bar/side-menu'
import SidebarLogo from '@/components/content/admin/side-bar/logo'
import { mapGetters, mapState } from 'vuex'
export default {
  name: 'SideBarIndex',
  components: { SidebarLogo, SideMenu },
  data () {
    return {
      adminRouter,
      userRouter
    }
  },
  computed: {
    ...mapState('admin', ['sideStatus']),
    ...mapGetters('user', ['userInfo'])
  }
}
</script>

<style>
  /**修复侧边栏收缩还是显示文字的问题 参考：https://github.com/ElemeFE/element/issues/17391**/
  .el-menu--collapse > div > .el-submenu > .el-submenu__title span,
  .el-menu--collapse > div > .el-submenu > .el-submenu__title .el-submenu__icon-arrow {
    display: none;
  }
</style>

<style scoped>
  .sideBarOpen{
    width: 64px !important;
  }
  .side-contain{
    width: 210px;
    position: fixed;
    bottom: 0;
    top: 0;
    left: 0;
  }
  .side-menu{
    height: 100%;
  }
</style>
