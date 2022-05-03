<div id="login">
    <h1><a href="javascript:">注册界面</a></h1>
    <?php
    if (get_err_msg()!=""){
        echo "<div id='login_error'><strong>错误</strong>：".get_err_msg()."<br></div>";
    } else if (get_ok_msg()!=""){
        echo "<div class='success'>".get_ok_msg()."<br></div>";
    }
    ?>
    <form name="loginform" id="loginform" action="" method="post">
        <p>
            <label for="username">登录名</label>
            <input type="text" name="username" id="username" class="username" size="20" autocapitalize="off" />
        </p>
        <p>
            <label for="nickname">昵称</label>
            <input type="text" name="nickname" id="nickname" class="input" size="20" autocapitalize="off">
        </p>
        <p>
            <label for="email">邮箱</label>
            <input type="text" name="email" id="email" class="input" size="20" autocapitalize="off">
        </p>
        <div class="user-pass-wrap">
            <label for="password">密码</label>
            <div class="wp-pwd">
                <input type="password" name="password" id="password" class="input password-input" value="" size="20">
            </div>
        </div>
        <div class="user-pass-wrap">
            <label for="confirm">确认密码</label>
            <div class="wp-pwd">
                <input type="password" name="confirm" id="confirm" class="input password-input" value="" size="20">
            </div>
        </div>
        <p class="submit">
            <input type="submit" name="wp-submit" id="wp-submit" class="button button-primary button-large" value="注册">
        </p>
    </form>
    <p id="nav">
        <a href="<?php echo setting_web().'/access/login' ?>">登录</a> |
        <a href="<?php echo setting_web().'/access/forget' ?>">忘记密码？</a>
    </p>
    <p id="backtoblog">
        <a href="<?php echo setting_web() ?>">← 返回主页</a>
    </p>
</div>