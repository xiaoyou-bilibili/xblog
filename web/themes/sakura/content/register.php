<div id="primary" class="content-area">
    <main id="main" class="site-main" role="main">
        <div class="ex-login">
            <?php
            if (get_err_msg()!=""){
                echo "<div id='login_error'><strong>错误</strong>：".get_err_msg()."<br></div>";
            } else if (get_ok_msg()!=""){
                echo "<div class='success'>".get_ok_msg()."<br></div>";
            }
            ?>
            <div class="ex-login-title">
                <p><img src="https://cdn.jsdelivr.net/gh/moezx/cdn@3.2.9/img/Sakura/images/none.png"></p>
            </div>
            <form action="" method="post">
                <p><input type="text" name="username" id="username" value="" placeholder="用户名" required /></p>
                <p><input type="text" name="nickname" id="nickname" value="" placeholder="昵称" required /></p>
                <p><input type="text" name="email" id="email" value="" placeholder="邮箱" required /></p>
                <p><input type="password" name="password" id="password" value="" placeholder="密码" required /></p>
                <p><input type="password" name="confirm" id="confirm" value="" placeholder="确认密码" required /></p>
                <input class="button login-button" name="submit" type="submit" value="注册">
            </form>
            <div class="ex-new-account" style="padding: 0;">
                <p><a href="/access/login" target="_blank">登录</a>|<a href="/access/forget" target="_blank">忘记密码</a></p>
            </div>
        </div>
    </main>
</div>