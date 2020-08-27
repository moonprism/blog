<template>
    <div class="login-container">
        <el-form class ="login-form" status-icon>
            <el-form-item prop="name">
                <el-input v-model="loginForm.username" placeholder="Username" prefix-icon="el-icon-user" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item prop="pass">
                <el-input type="password" v-model="loginForm.password" placeholder="Password" prefix-icon="el-icon-lollipop" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" style="width:100%" @click="login()" :loading="loading">Login</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                loginForm: {
                    username: '',
                    password: '',
                },
                loading: false
            }
        },
        methods: {
            async login() {
                this.loading = true
                try {
                    await this.$store.dispatch('auth/login', this.loginForm)
                } catch (err) {
                    this.loading = false
                }
                this.$router.push({path: '/'})
            }
        },
    }
</script>

<style>
    .login-container {
        position: fixed;
        top: 0;
        left: 0;
        height: 100%;
        width: 100%;
        background-image: url(https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/gFjsJhJcSZCsIyBGbLVL.jpg);
        background-repeat: no-repeat;
        /* background-size: cover; */
    }
    .login-form {
        width: 277px;
        margin: 260px auto;
    }
    .login-form input {
        font-family: 'Space Mono';
    }
    .login-form button {
        letter-spacing: .03em;
        font-family: 'Space Mono';
    }
</style>
