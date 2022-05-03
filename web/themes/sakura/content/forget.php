<div id="primary" class="content-area">
    <main id="main" class="site-main" role="main">
        <div class="ex-login">
            <?php if (get_err_msg()!=""){
                echo "<div id='login_error'><strong>错误</strong>：".get_err_msg()."<br></div>";
            } else if (get_ok_msg()!=""){
                echo "<div class='success'>".get_ok_msg()."<br></div>";
            } ?>
            <div id='login_error'>将发送一封找回密码的邮件给你，请点击链接重置密码!</div>
            <form action="" method="post">
                <p><input type="text" name="email" id="email" value="" size="25" placeholder="用户邮箱" required /></p>
                <input class="button login-button" name="submit" type="submit" value="发送邮件">
            </form>
            <div class="ex-new-account" style="padding: 0;">
                <p><a href="/access/register">注册</a>|<a href="/access/login">登录</a></p>
            </div>
        </div>
    </main>
</div>