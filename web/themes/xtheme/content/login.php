<div id="login">
    <el-form ref="ruleForm" class="loginFrom" :model="loginData" :rules="rules">
        <el-form-item prop="userName">
            <el-input
                v-model="loginData.username"
                class="login-inputorbuttom"
                prefix-icon="el-icon-user"
                placeholder="登录名"
            ></el-input>
        </el-form-item>
        <el-form-item prop="password">
            <el-input
                v-model="loginData.password"
                class="login-inputorbuttom"
                prefix-icon="el-icon-lock"
                placeholder="密码"
                type="password"
            ></el-input>
        </el-form-item>
        <el-form-item class="login-item">
            <el-checkbox
                v-model="loginData.remember"
                style="float:left;margin-bottom:15px"
                label="记住我的登录信息"
                name="type"
            ></el-checkbox>
            <el-button
                v-popover:popover
                class="login-inputorbuttom login-bottom"
                type="primary"
                :loading="loginIng"
                @click="loginButton"
            >
                登 录
            </el-button>
            <div class="login-options clearfix">
                <a href="/" class="login-help">
                    返回首页
                </a>
                <a id="login-account-switch" href="javascript:void(0);" @click="forget">
                    忘记密码
                </a>
                <a href="/access/register" class="register">
                    立即注册
                </a>
            </div>
        </el-form-item>
    </el-form>
</div>