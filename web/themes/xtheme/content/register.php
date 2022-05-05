<div id="signup">
    <el-form ref="logindata" class="loginFrom" status-icon :model="registered" :rules="rules">
        <el-form-item prop="username">
            <el-input
                v-model="registered.username"
                class="login-inputorbuttom"
                prefix-icon="el-icon-user"
                placeholder="登录名"
            ></el-input>
        </el-form-item>
        <el-form-item prop="nickname">
            <el-input
                v-model="registered.nickname"
                class="login-inputorbuttom"
                prefix-icon="el-icon-user-solid"
                placeholder="昵称"
            ></el-input>
        </el-form-item>
        <el-form-item prop="password">
            <el-input
                v-model="registered.password"
                class="login-inputorbuttom"
                prefix-icon="el-icon-lock"
                placeholder="密码"
                type="password"
            ></el-input>
        </el-form-item>
        <el-form-item prop="repeat">
            <el-input
                v-model="registered.repeat"
                class="login-inputorbuttom"
                prefix-icon="el-icon-unlock"
                placeholder="重复密码"
                type="password"
            ></el-input>
        </el-form-item>
        <el-form-item prop="email">
            <el-input
                v-model="registered.email"
                class="login-inputorbuttom"
                prefix-icon="el-icon-message"
                placeholder="邮箱地址"
                type="input"
            ></el-input>
        </el-form-item>
        <el-form-item class="verify" prop="verify">
            <el-input
                v-model="registered.verify"
                type="text"
                placeholder="请输入验证码"
            ></el-input>
            <canvas
                id="code"
                width="70px"
                height="27px"
                style="float: right; border:1px solid #d3d3d3;"
                @click="createCode"
            ></canvas>
        </el-form-item>
        <el-form-item class="login-item">
            <el-button
                v-popover:popover
                class="resignbutton"
                type="primary"
                :loading="loading"
                @click="loginButton()"
            >
                注册
            </el-button>
            <div class="login-options clearfix">
                <a href="/" class="login-help">
                    返回首页
                </a>
                <a id="login-account-switch" href="javascript:void()" @click="forget">
                    忘记密码
                </a>
                <a href="/access/login" class="register">
                    立即登录
                </a>
            </div>
        </el-form-item>
    </el-form>
</div>