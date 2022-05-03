<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>网站设置</span>
      </div>
      <el-tabs v-model="active" type="card" @tab-click="change">
        <el-tab-pane label="站点设置" name="01">
          <setting-site />
        </el-tab-pane>
        <el-tab-pane label="壁纸设置" name="03">
          <setting-background />
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script>
import SettingSite from '@/components/content/admin/content/setting/site'
import SettingBackground from '@/components/content/admin/content/setting/background'
import admin from '@/components/mixin/admin-seo'

export default {
  components: {
    SettingBackground,
    SettingSite
  },
  layout: 'admin',
  mixins: [admin],
  data () {
    return {
      active: '01'
    }
  },
  mounted () {
    // 首先获取网站设置
    this.change({ name: '01' })
  },
  methods: {
    // 当导航栏改变的时候自动获取设置信息
    change (item) {
      switch (item.name) {
        case '01':
          this.$store.dispatch('admin-settings/getSiteOption')
          break
        case '03':
          this.$store.dispatch('admin-settings/getBackgroundOption')
          break
        case '04':
          this.$store.dispatch('admin-settings/getSpiderOption')
          break
        case '05':
          break
        case '06':
          // this.$refs.sponsors.updateData(false)
          break
      }
    }
  }
}
</script>

<style scoped>
</style>
