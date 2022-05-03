<template>
  <div>
    <el-row :gutter="20">
      <!--文章编辑器-->
      <el-col :xs="24" :sm="15" :md="20">
        <!--使用sync的目的是为了让子组件改变父主键的值-->
        <post-edit
          :html.sync="postContent.html"
          :md.sync="postContent.md"
        />
      </el-col>
      <!--功能选择区-->
      <el-col :xs="24" :sm="9" :md="4">
        <el-card class="option-card" shadow="hover">
          <div slot="header" style="background: #996600;" class="card-head">
            <font-awesome-icon icon="upload" />
            <span>日记发布</span>
            <span v-show="saveStatus" style="font-size: 12px"><i class="el-icon-circle-check" />已自动保存</span>
          </div>
          <div class="text-label">
            <font-awesome-icon icon="eye" />
            日记可见性
          </div>
          <div>
            <el-radio-group v-model="postContent.status" text-color="#CC9999">
              <el-radio label="publish">
                公开
              </el-radio>
              <el-radio label="encrypt">
                加密
              </el-radio>
              <el-radio label="private">
                私密
              </el-radio>
            </el-radio-group>
            <el-input v-if="postContent.status==='encrypt'" v-model="postContent.password" placeholder="请输入密码" show-password />
          </div>
          <el-divider />
          <div v-if="postContent.is_draft">
            <el-popconfirm title="删除后不可恢复，确定删除？" @onConfirm="post('delete')">
              <el-button slot="reference" type="danger" icon="el-icon-delete">
                删除
              </el-button>
            </el-popconfirm>
            <el-button type="success" icon="el-icon-s-promotion" @click="post('update',true, true)">
              发布
            </el-button>
          </div>
          <div v-else>
            <el-popconfirm title="删除后不可恢复，确定删除？" @onConfirm="post('delete')">
              <el-button slot="reference" type="danger" icon="el-icon-delete">
                删除
              </el-button>
            </el-popconfirm>
            <el-button type="primary" icon="el-icon-tickets" @click="post('update',true, false, true)">
              草稿
            </el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import _ from 'lodash'
import admin from '@/components/mixin/admin-seo'
import PostEdit from '~/components/content/admin/content/post-edit'
export default {
  components: { PostEdit },
  layout: 'admin',
  mixins: [admin],
  data () {
    return {
      radio: '公开',
      password: '',
      // 文章内容
      postContent: {
        post_id: 0, // 日记
        html: '', // 文章标题
        status: 'publish', // 日记状态
        password: '', // 日记密码
        md: '', // 文章markdown格式
        is_draft: true // 日记是否为草稿
      },
      saveStatus: false, // 当前日记是否保存
      flag: true // 这个flag用于跳过watch监听，避免自动更新死循环
    }
  },
  watch: {
    postContent: {
      deep: true,
      handler () {
        if (this.flag) {
          // 设置状态未保存
          this.saveStatus = false
          // 自动保存
          this.debouncedAutoSave()
        } else {
          this.flag = true
        }
      }
    }
  },
  created () {
    // 创建一个防反跳函数
    this.debouncedAutoSave = _.debounce(() => { this.post(this.postContent.post_id ? 'update' : 'publish') }, 1000)
  },
  mounted () {
    // 自动获取编辑器的所有数据
    const id = this.$route.params.id
    if (id !== undefined && id !== null) {
      this.$store.dispatch('admin-posts/getDiaryContent', id).then((res) => {
        // 因为我们不能直接在主界面修改文章内容，所以我们需要单独赋值
        this.postContent = res
      }).catch(() => { })
    }
  },
  methods: {
    // 更新文章 show表示是否需要提示,默认不提示 publish表示是否为发布文章，默认只更新 draft表示是否为草稿
    post (option, show = false, publish = false, draft = false) {
      // 为了避免修改对象造成自动更新，我们这里使用拷贝
      const content = Object.assign({}, this.postContent)
      // 判断参数是否为空
      if (content.html && content.md && content.status) {
        // 判断是发布还是草稿
        if (publish) { content.is_draft = false }
        if (draft) { content.is_draft = true }
        // 判断用户操作，不同操作对应不同内容
        switch (option) {
          // 这里说明是发布文章
          case 'publish':
            // 我们这里发布一下文章
            this.$store.dispatch('admin-posts/addDiary', content).then((data) => {
              this.postContent.post_id = data.post_id
              // 路由替换
              window.history.pushState(null, '后台管理系统', `/admin/post/diary-edit/${data.post_id}`)
              // 更新标签
              this.$store.dispatch('admin/editChangeId', { title: '日记编辑器', path: `/admin/post/diary-edit/${data.post_id}` })
              this.saveStatus = true
            })
            break
          // 更新文章操作
          case 'update':
            // 如果是更新需要转换一下草稿
            content.is_draft = content.is_draft ? 'true' : 'false'
            this.$store.dispatch('admin-posts/updatePost', { id: content.post_id, data: content }).then((data) => {
              this.saveStatus = true
              // 判断是发布还是草稿
              if (publish) {
                this.flag = false
                this.postContent.is_draft = false
                this.$message.success('发布成功')
              } else if (draft) {
                this.flag = false
                this.postContent.is_draft = true
                this.$message.success('保存草稿成功')
              }
            })
            break
          case 'delete':
            // 删除这个文章,同时关闭当前标签页
            this.$store.dispatch('admin-posts/deletePost', content.post_id).then(() => {
              // 删除标签页
              this.$store.dispatch('admin/deleteActiveTag')
              // 跳转到主页路由
              this.$router.replace('/admin')
            })
        }
      } else if (show) {
        // 这里我们如果是自动保存草稿的就不需要自动弹框
        this.$message.error('必须要有内容和标题!')
      }
    }
  }
}
</script>

<style>
/*设置card的边距为0，同时设置分割线的间距*/
.option-card .el-card__header{
  padding: 0!important;
}
.option-card .el-divider--horizontal{
  margin: 10px 0!important;
}
/*消除多选*/
.el-radio-group .el-radio{
  margin-right: 14px!important;
}
</style>

<style scoped>
.post-edit{
  margin-bottom: 10px;
}
.card-head{
  padding: 18px 20px;
  color: #ffffff;
  font-size: 16px;
}
.option-card{
  margin-bottom: 10px;
}
.text-label{
  margin-bottom: 5px;
}
</style>
