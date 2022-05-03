<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>侧边栏和顶部导航栏设置</span>
      </div>
      <el-tabs v-model="active" type="card" @tab-click="change">
        <el-tab-pane label="顶部导航栏设置" name="01">
          <side-head />
        </el-tab-pane>
        <el-tab-pane label="侧边栏设置" name="02">
          <side-side-tool />
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import SideSideTool from '@/components/content/admin/content/side/side-tool'
import SideHead from '@/components/content/admin/content/side/head'

export default {
  components: { SideSideTool, SideHead },
  layout: 'admin',
  mixins: [admin],
  data () {
    return {
      active: '01'
    }
  },
  mounted () {
    this.change({ name: '01' })
  },
  methods: {
    change (item) {
      switch (item.name) {
        case '01':
          this.$store.dispatch('admin-settings/getNavOption')
          break
        case '02':
          this.$store.dispatch('admin-settings/getSideOption')
          break
      }
    }
  }
}
</script>

<style scoped>

</style>
