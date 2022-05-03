<div id="primary" class="content-area">
    <main id="main" class="site-main" role="main">
        <div class="ex-login">
            <?php if (get_err_msg()!=""){
                echo "<div id='login_error'><strong>错误</strong>：".get_err_msg()."<br></div>";
            } else if (get_ok_msg()!=""){
                echo "<div class='success'>".get_ok_msg()."<br></div>";
            } ?>
            <div class="ex-login-title">
                <p><img src="https://cdn.jsdelivr.net/gh/moezx/cdn@3.2.9/img/Sakura/images/none.png"></p>
            </div>
            <form action="" method="post">
                <p><input type="text" name="username" id="username" value="" size="25" placeholder="用户名" required /></p>
                <p><input type="password" name="password" id="password" value="" size="25" placeholder="密码" required /></p>
                <p class="forgetmenot">
                    <label for="rememberme">
                        <input name="remember" type="checkbox" id="remember" value="forever" />
                    </label>
                    记住我
                </p>
                <input class="button login-button" name="submit" type="submit" value="登 入">
            </form>
            <div class="ex-new-account" style="padding: 0;">
                <p><a href="/access/register" target="_blank">注册</a>|<a href="/access/forget" target="_blank">忘记密码</a></p>
            </div>
        </div>
    </main>
</div>