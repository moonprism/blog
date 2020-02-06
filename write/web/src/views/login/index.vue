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
                <el-button type="primary" style="width:100%" @click="login()">Login</el-button>
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
                    password: ''
                },
            }
        },
        methods: {
            async login() {
                await this.$store.dispatch('auth/login', this.loginForm)
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
        background-image: url(https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/cc/bg.jpg)
    }
    .login-form {
        width: 277px;
        margin: 170px auto;
    }
    .login-form input {
        font-family: 'Roboto Mono';
    }
    .login-form button {
        letter-spacing: .03em;
        font-family: 'Roboto Mono';
    }
</style>