<template>
    <el-container>
        <div class="an" v-bind:style="{ backgroundImage: backgroundImageUrl}" ></div>
        <el-aside width="77px">
            <el-menu class="sideMenu" :collapse="true">
                <router-link to="/setting">
                    <el-menu-item index="4">
                        <i class="el-icon-setting"></i>
                    </el-menu-item>
                </router-link>
                <router-link to="/article">
                    <el-menu-item index="1">
                        <i class="el-icon-document"></i>
                    </el-menu-item>
                </router-link>
                <router-link to="/tag">
                    <el-menu-item index="3">
                        <i class="el-icon-collection-tag"></i>
                    </el-menu-item>
                </router-link>
                <router-link to="/code">
                    <el-menu-item index="2">
                        <i class="el-icon-postcard"></i>
                    </el-menu-item>
                </router-link>
                <router-link to="/file">
                    <el-menu-item index="5">
                        <i class="el-icon-folder"></i>
                    </el-menu-item>
                </router-link>
            </el-menu>
        </el-aside>
        <el-container>
            <el-header style="height: 20px" class="header">
                <el-dropdown @command="handleCommand">
                    <el-avatar
                            size="small"
                            :src="headerImgUrl"></el-avatar>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item command="logout">log out</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
            </el-header>
            <el-main>
                <section class="main board">
                    <transition name="fade-transform" mode="out-in">
                        <!--子页面-->
                        <router-view :key="key" />
                    </transition>
                </section>
            </el-main>
        </el-container>
    </el-container>
</template>

<script>
    import { MessageBox } from 'element-ui'

    import '@/style/main.css'

    export default {
        data() {
            return {
                headerImgUrl: 'https://q1.qlogo.cn/g?b=qq&nk=1370808234&s=640'
            }
        },
        computed: {
            key() {
                return this.$router.path + Math.random()
            },
            backgroundImageUrl() {
                    return 'url(' + this.$store.getters['setting/backgroundImage'] + ')'
            }
        },
        methods: {
            logout() {
                MessageBox.confirm('log out of blog admin?', '', {
                    confirmButtonText: 'log out',
                    cancelButtonText: 'Cancel',
                    type: 'warning'
                }).then(() => {
                    this.$store.dispatch('auth/reset').then(() => {
                        location.reload()
                    })
                })
            },
            handleCommand(command) {
                if (command == 'logout') {
                    this.logout()
                }
            }
        },
        created() {
            this.$store.commit('setting/init')
        }
    }
</script>

<style scoped>
    .sideMenu {
        padding-bottom: 10px;
        border-radius: 2px;
    }
    .header {
        text-align: right;
        margin: 10px;
    }
    .header-img {
        width: 35px;
        height: 35px;
        border-radius: 50%;
    }
    .an {
        position: fixed;
        width: 100%;
        height: 100%;
        z-index: -9;
        top: 0;
        left: 0;
        background-repeat: no-repeat;
        background-size: cover;
        background-color: #e5edf2;
    }
    .board {
        width: 766px;
        margin: 3px auto 290px;
        background-color: #fff;
        padding: 34px 47px 30px;
        border-radius: 3px;
        min-height: 650px;
    }
</style>