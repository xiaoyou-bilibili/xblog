<template>
  <div>
    <el-row>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="side-card">
          <div slot="header" class="clearfix">
            <span>可选择(可直接拖动到对应区域)</span>
          </div>
          <side-draggable ref="unused" :list="sideOptions.unused" />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="side-card">
          <div slot="header" class="clearfix">
            <span>左侧边栏</span>
          </div>
          <side-draggable ref="left" :list="sideOptions.left" />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <el-card class="side-card">
          <div slot="header" class="clearfix">
            <span>右侧边栏</span>
          </div>
          <side-draggable ref="right" :list="sideOptions.right" />
        </el-card>
      </el-col>
    </el-row>
    <el-button style="margin-top: 10px" type="primary" @click="saveOption">
      保存设置
    </el-button>
  </div>
</template>

<script>

import SideDraggable from '@/components/content/admin/content/side/side-draggable'
import { mapGetters } from 'vuex'
export default {
  name: 'SideSideTool',
  components: { SideDraggable },
  computed: {
    // 把getter混入到computed对象中,我们可以直接调用vex里面的内容即可
    ...mapGetters('admin-settings', ['sideOptions'])
  },
  methods: {
    // 保存侧边栏设置
    saveOption () {
      // 设置更新的数据
      const data = {
        left: this.$refs.left.lists,
        right: this.$refs.right.lists
      }
      // 更新记录
      this.$store.dispatch('admin-settings/updateSideOption', data).then(() => {
        this.$message.success('保存成功!')
        this.$store.dispatch('admin-settings/getSideOption')
      }).catch(msg => this.$message.error(msg))
    }
  }
}
</script>

<style>
.side-card .el-card__body {
  padding: 0!important;
}
</style>

<style scoped>
  .side-card {
    margin: 0 5px;
  }
</style>
