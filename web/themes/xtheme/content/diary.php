<div class="diary-card">
    <img alt class="user-icon" src="<?php plugins_diary_avatar(); ?>" />
    <div class="arrow"></div>
    <div class="diary-content">
        <div class="diary-head">
            <div class="diary-head-name">
                <?php plugins_diary_nickname(); ?>
            </div><div class="diary-head-time">
                <?php plugins_diary_date(); ?>
            </div>
        </div>
        <?php if (!get_plugins_diary_encrypt()):?>
        <div class="diary-body">
            <?php plugins_diary_content(); ?>
        </div>
        <?php else: ?>
        <div class="lock-face">
            <div class="lock-text">
                这篇日记已被上锁了，请输入密码解锁
            </div>
            <div class="lock-input">
                <el-input v-model="password" type="password" placeholder="请输入访问密码">
                    <el-button slot="prepend" icon="el-icon-key" ></el-button>
                </el-input>
            </div>
            <el-button class="lock-btn" type="primary" @click="getEncryptContent(<?php echo get_plugins_diary_id() ?>)">
                点击访问
            </el-button>
        </div>
        <?php endif;?>
        <div class="diary-footer" @click="gotoDiary(<?php echo get_plugins_diary_id() ?>)">
            <div class="diary-comment">
                <i class="far fa-comment"></i><?php plugins_diary_comment(); ?>条评论
            </div>
            <div class="diary-like">
                <i class="far fa-thumbs-up"></i><?php plugins_diary_good(); ?>人点赞
            </div>
        </div>
    </div>
</div>

