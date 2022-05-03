<template>
  <div>
    <el-card class="option-card" shadow="hover">
      <el-tabs v-model="activeName" type="card">
        <el-tab-pane label="通用设置" name="normal">
          <el-form ref="form" class="col-md-4" label-width="80px">
            <div class="tool-img-item">
              <h4>
                邮件订阅
                <el-switch
                  v-model="userInfo.subscription"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  @change="emailSubscription"
                />
              </h4>
              <div>打开后博主有文章更新时会提醒你</div>
            </div>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import admin from '@/components/mixin/admin-seo'

export default {
  name: 'Normal',
  layout: 'admin',
  mixins: [admin],
  data () {
    return {
      activeName: 'normal'
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  methods: {
    // 邮件订阅功能
    emailSubscription (subscription) {
      // 提交订阅
      this.$store.dispatch('user/updateInfo', { subscription: subscription.toString() }).then(() => {
        this.$message.success('修改成功')
      }).catch((msg) => {
        this.$message.error(msg)
        // 重新获取用户信息
        this.$store.dispatch('user/getUserInfo')
      })
    }
  }
}
</script>

<style scoped lang="scss">
/*侧边栏设置的相关设置*/
.tool-img-item{
  display: flex;
  flex-direction: column;
  .item-input{
    display: flex;
    flex-direction: row;
    margin: 5px 0 0 0;
  }
  .item-input .el-input{
    margin-right: 5px;
  }
  h4{
    font-size: 1em;
    font-weight: 800;
  }
}
</style>
