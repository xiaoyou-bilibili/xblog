<div id="login">
    <h1><a href="javascript">登录界面</a></h1>
    <?php
    if (get_err_msg()!=""){
        echo "<div id='login_error'><strong>错误</strong>：".get_err_msg()."<br></div>";
    } else if (get_ok_msg()!=""){
        echo "<div class='success'>".get_ok_msg()."<br></div>";
    }
    ?>
    <form name="loginform" id="loginform" action="" method="post">
        <p>
            <label for="username">用户名或电子邮箱地址</label>
            <input type="text" name="username" id="username" class="input" value="" size="20" autocapitalize="off">
        </p>

        <div class="user-pass-wrap">
            <label for="password">密码</label>
            <div class="wp-pwd">
                <input type="password" name="password" id="password" class="input password-input" value="" size="20">
            </div>
        </div>
        <p class="forgetmenot"><input name="remember" type="checkbox" id="remember" value="forever"> <label for="remember">记住我</label></p>
        <p class="submit">
            <input type="submit" name="wp-submit" id="wp-submit" class="button button-primary button-large" value="登录">
        </p>
    </form>
    <p id="nav">
        <a href="<?php echo setting_web().'/access/register' ?>">注册</a> |
        <a href="<?php echo setting_web().'/access/forget' ?>">忘记密码？</a>
    </p>
    <p id="backtoblog">
        <a href="<?php echo setting_web() ?>">← 返回主页</a>
    </p>
</div>