<template>
  <div>
    <el-card class="box-card">
      <!-- 头像修改区域-->
      <div class="edit-avatar">
        <a href="javascript:void(0)" class="fa">
          <img alt :src="userInfo.avatar" class="up-face">
          <img alt :src="userInfo.hang" class="pendant">
        </a>
        <div>
          <el-button @click="uploadImg">
            修改头像
          </el-button>
          <div class="change-info">
            选择一张你喜欢的图片，裁剪后会自动生成264x264大小，上传图片大小不能超过2M。
          </div>
          <!--图片裁剪框 start-->
          <div style="display: none" class="tailoring-container">
            <div class="black-cloth" @click="closeTailor(this)" />
            <div class="tailoring-content">
              <div class="tailoring-content-one">
                <label title="上传图片" for="chooseImg" class="l-btn choose-btn">
                  <input
                    id="chooseImg"
                    type="file"
                    accept="image/jpg,image/jpeg,image/png"
                    name="file"
                    class="hidden"
                    @change="selectImg"
                  >
                  选择图片
                </label>
                <div class="close-tailoring" @click="closeTailor(this)">
                  ×
                </div>
              </div>
              <div class="tailoring-content-two">
                <div class="tailoring-box-parcel">
                  <img id="tailoringImg" :src="url" alt>
                </div>
                <div class="preview-box-parcel">
                  <p>图片预览：</p>
                  <div class="square previewImg" />
                  <div class="circular previewImg" />
                </div>
              </div>
              <div class="tailoring-content-three">
                <button class="l-btn cropper-reset-btn" @click="resetBtn">
                  复位
                </button>
                <button class="l-btn cropper-rotate-btn" @click="rotateBtn">
                  旋转
                </button>
                <button id="sureCut" class="l-btn sureCut" @click="sureCut">
                  确定
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- 个人信息修改 -->
      <el-divider content-position="left">
        个人信息修改
      </el-divider>
      <el-form ref="form" label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="userInfo.username" :disabled="true" />
        </el-form-item>
        <el-form-item label="头像挂件">
          <el-input v-model="userInfo.hang" />
        </el-form-item>
        <el-form-item label="昵称">
          <el-input v-model="userInfo.nickname" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="userInfo.email" />
        </el-form-item>
        <el-form-item label="签名">
          <el-input v-model="userInfo.sign" type="textarea" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="updateInfo">
            保存
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import $ from 'jquery'
import Cropper from 'cropperjs'
import { mapState } from 'vuex'
import admin from '@/components/mixin/admin-seo'
export default {
  layout: 'admin',
  mixins: [admin],
  data () {
    return {
      cropper: '',
      picValue: '',
      url: ''
    }
  },
  computed: {
    ...mapState('user', ['userInfo'])
  },
  mounted () {
    // 获取个人信息
    this.$store.dispatch('user/getUserInfo')
    // 初始化裁剪框
    const img = document.getElementById('tailoringImg')
    // cropper图片裁剪
    this.cropper = new Cropper(img, {
      aspectRatio: 1, // 默认比例
      preview: '.previewImg', // 预览视图
      guides: true, // 裁剪框的虚线(九宫格)
      autoCropArea: 0.8, // 0-1之间的数值，定义自动剪裁区域的大小，默认0.8
      // movable: false, //是否允许移动图片
      dragCrop: true, // 是否允许移除当前的剪裁框，并通过拖动来新建一个剪裁框区域
      movable: true, // 是否允许移动剪裁框
      resizable: true, // 是否允许改变裁剪框的大小
      zoomable: false, // 是否允许缩放图片大小
      mouseWheelZoom: false, // 是否允许通过鼠标滚轮来缩放图片
      touchDragZoom: false, // 是否允许通过触摸移动来缩放图片
      rotatable: true // 是否允许旋转图片
    })
  },
  methods: {
    // 获取图片的地址（不同的浏览器图片地址不同）
    getObjectURL (file) {
      let url = null
      if (window.createObjectURL !== undefined) { // basic
        url = window.createObjectURL(file)
      } else if (window.URL !== undefined) { // mozilla(firefox)
        url = window.URL.createObjectURL(file)
      } else if (window.webkitURL !== undefined) { // webkit or chrome
        url = window.webkitURL.createObjectURL(file)
      }
      return url
    },
    sureCut () { // 点击剪切图片
      // 是否有图片
      if ($('#tailoringImg').attr('src') == null) {
        return false
      } else {
        // 获取cropper里面的图片
        this.cropper.getCroppedCanvas().toBlob((blob) => {
          // 初始化一个formData对象(这个对象可以用于图片上传)
          const fromData = new FormData()
          // 放入我们的图片
          fromData.append('file', blob)
          // 开始上传图片
          this.$store.dispatch('admin-file/uploadImage', fromData).then((data) => {
            // 这里上传成功，获取到图片地址
            this.$store.dispatch('user/updateInfo', { avatar: data.url }).then(() => {
              // 重新获取用户信息
              this.$store.dispatch('user/getUserInfo')
              // 手动关闭弹框
              this.closeTailor()
            }).catch(() => { this.$message.error('更新用户信息失败') })
          })
        }, 'image/jpeg')
      }
    },
    selectImg (e) { // 旋转一张图片
      // 先获取input里面的文件
      const files = e.target.files || e.dataTransfer.files
      // 判断文件是否为空
      if (!files.length) { return }
      // 设置图片的内容和地址
      this.picValue = files[0]
      this.url = this.getObjectURL(this.picValue)
      // 每次替换图片要重新得到新的url
      if (this.cropper) {
        this.cropper.replace(this.url)
      }
    },
    rotateBtn () { // 点击旋转按钮
      this.cropper.rotate(45)
    },
    resetBtn () { // 点击复位按钮
      this.cropper.reset()
    },
    uploadImg () { // 点击修改头像按钮
      $('.tailoring-container').toggle()
    },
    closeTailor () { // 关闭头像框上传
      $('.tailoring-container').toggle()
    },
    updateInfo () { // 更新个人信息
      // 获取需要更新的信息
      const data = {
        avatar: this.userInfo.avatar,
        hang: this.userInfo.hang,
        nickname: this.userInfo.nickname,
        email: this.userInfo.email,
        sign: this.userInfo.sign
      }
      this.$store.dispatch('user/updateInfo', data).then(() => {
        // 提示更新成功
        this.$message.success('更新成功')
        // 重新获取个人信息
        this.$store.dispatch('user/getUserInfo')
      }).catch((msg) => { this.$message.error(msg) })
    }
  }
}

</script>

<style src="@/assets/css/cropper.min.css"/>

<style scoped>
/*头像部分*/
.fa{
  position: relative;
  margin-right: 20px;
}
.fa .up-face{
  border-radius: 50%;
  width: 60px;
  height: 60px;
}
.pendant{
  position: absolute;
  overflow: hidden;
  top: -21px;
  left: -21px;
  width: 102px;
  height: 102px;
  z-index: 2;
}
.edit-avatar{
  display: flex;
}
.change-info{
  color: #4d5259;
  font-size: 14px;
}

/*图片裁剪*/
/*隐藏上传按钮*/
#chooseImg{
  display: none;
}
.l-btn{
  display: inline-block;
  outline: none;
  resize: none;
  padding:5px 10px;
  background: #8C85E6;
  color: #fff;
  border:solid 1px #8C85E6;
  border-radius: 3px;
  font-size: 14px;
}
.l-btn:hover{
  background: #8078e3;
  animation: anniu 1s infinite;
}
.l-btn:active{
  box-shadow: 0 2px 3px rgba(0,0,0,.2) inset;
}
.tailoring-container, .tailoring-container div, .tailoring-container p{
  margin: 0;padding: 0;
  box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
}
.tailoring-container{
  position: fixed;
  width: 100%;
  height: 100%;
  z-index: 1000;
  top: 0;
  left: 0;
}
.tailoring-container .black-cloth{
  position: fixed;
  width: 100%;
  height: 100%;
  background: #111;
  opacity: .9;
  z-index: 1001;
}
.tailoring-container .tailoring-content{
  position: absolute;
  width: 768px;
  height: 560px;
  background: #fff;
  z-index: 1002;
  left: 0;
  top: 0;
  border-radius: 10px;
  box-shadow: 0 0 10px #000;
  padding: 10px;
}

.tailoring-content-one{
  height: 40px;
  width: 100%;
  border-bottom: 1px solid #DDD ;
}
.tailoring-content .choose-btn{
  float: left;
}
.tailoring-content .close-tailoring{
  display: inline-block;
  height: 30px;
  width: 30px;
  border-radius: 100%;
  background: #eee;
  color: #fff;
  font-size: 22px;
  text-align: center;
  line-height: 30px;
  float: right;
  cursor: pointer;
}
.tailoring-content .close-tailoring:hover{
  background: #ccc;
}

.tailoring-content .tailoring-content-two{
  width: 100%;
  height: 460px;
  position: relative;
  padding: 5px 0;
}
.tailoring-content .tailoring-box-parcel{
  width: 520px;
  height: 450px;
  position: absolute;
  left: 0;
  border: solid 1px #ddd;
}
.tailoring-content .preview-box-parcel{
  display: inline-block;
  width: 228px;
  height: 450px;
  position: absolute;
  right: 0;
  padding: 4px 14px;
}
.preview-box-parcel p{
  color: #555;
}
.previewImg{
  width: 200px;
  height: 200px;
  overflow: hidden;
}
.preview-box-parcel .square{
  margin-top: 10px;
  border: solid 1px #ddd;
}
.preview-box-parcel .circular{
  border-radius: 100%;
  margin-top: 10px;
  border: solid 1px #ddd;
}

.tailoring-content .tailoring-content-three{
  width: 100%;
  height: 40px;
  border-top: 1px solid #DDD ;
  padding-top: 10px;
}
.sureCut{
  float: right;
}

</style>
