<template>
    <el-row :gutter="20">
        <el-dialog
            v-if="waterfallDialog.visible"
                :close-on-click-modal="false"
                :visible.sync="waterfallDialog.visible">
            <!--图片瀑布流选择模块-->
            <image-waterfall :size="waterfallSizeConfig" @select-image="selectImage"></image-waterfall>
        </el-dialog>
        <div class="markdown">
            <h2>Pages</h2>
            <p>
                <ul>
                    <li> <a @click="$router.push({name: 'article_edit', params: {id: 1}})">❤️ Link</a></li>
                    <li> <a @click="$router.push({name: 'article_edit', params: {id: 2}})">🍥 About</a> </li>
                </ul>
            </p>

            <h2>Background</h2>
            <p>
                <ul>
                    <li> <a @click="waterfallDialog.visible = true">选择图片</a> </li>
                    <li>
                        <el-upload
                            action="#"
                            :auto-upload="false"
                            :show-file-list="false"
                            :on-change="selectImageFromLocal">
                            <a>试用本地图片</a>
                        </el-upload>
                    </li>
                    <li> <a @click="clear">Clear</a> </li>
                </ul>
            </p>

            <h2>Email</h2>
            <p>
            开启评论邮箱提醒: 
            <el-switch
            v-model="$store.state.setting.setting.notice"
            @change="runEmailSwitch"
            active-color="#13ce66"
            inactive-color="#ff4949">
            </el-switch>
            </p>

            <h2>CSS</h2>
            <p>
                <code-edit ref="codeEditCSS" :code="{lang:'css', content:this.css}" :config="{useVim:false}"></code-edit>
            </p>

            <h2>JS</h2>
            <p>
                <code-edit ref="codeEditJS" :code="{lang:'js', content:this.js}" :config="{useVim:false}"></code-edit>
            </p>
            <br><br><hr>
            <p style="text-align: center; position: relative; top: 15px;">
                <el-button @click="save()" type="primary" size="medium" icon="el-icon-edit" circle></el-button>
            </p>
        </div>
    </el-row>
</template>

<script>
import ImageWaterfall from "@/components/image/waterfall"
import '@/style/markdown.css'
import {getBase64FromRaw} from "@/utils/file"
import codeEdit from '@/components/code-edit'
import { mapGetters } from 'vuex'
import settingApi from '@/api/setting'

export default {
    components: {
        ImageWaterfall,
        codeEdit
    },
    data() {
        return {
            waterfallDialog: {
                visible: false
            },
            settingInfo: {
                background_image: '',
            },
            waterfallSizeConfig: {
                width: 150
            },
            jsCodeEditConf: {
                lang: 'js',
                content: '',
            },
            cssCodeEditConf: {
                lang: 'css',
                content: '',
            },
        }
    },
    computed: {
        ...mapGetters({
            img: 'setting/backgroundImage',
            js: 'setting/globalJS',
            css: 'setting/globalCSS',
            captcha: 'setting/captcha',
            notice: 'setting/notice'
        })
    },
    methods: {
        async sync() {
            this.settingInfo.background_image = this.img
            this.settingInfo.global_js = this.$refs.codeEditJS.getContent()
            this.settingInfo.global_css = this.$refs.codeEditCSS.getContent()
        },
        async save() {
            this.sync()
            await this.$store.dispatch('setting/post', this.settingInfo)
            this.$message({
                message: 'write setting',
                type: 'success',
                duration: 1 * 1000
            })
        },
        async selectImage(imageModel) {
            this.settingInfo.background_image = process.env.VUE_APP_FILE_ORIGIN + imageModel.key
            this.$store.commit('setting/set', this.settingInfo)
            this.waterfallDialog.visible = false
        },
        clear() {
            this.settingInfo.background_image = ''
            this.$store.commit('setting/setBackgroundImage', '')
        },
        async runEmailSwitch(status) {
            let emailData = {"notice": status}
            const res = await settingApi.updateEmail(emailData)
            if (res.data === 'ok') {
                this.sync()
                this.settingInfo.notice = status
                this.$store.commit('setting/set', this.settingInfo)
                this.$message({
                    message: "邮箱提醒: "+(status?"开":"关"),
                    type: 'success',
                    duration: 2 * 1000
                })
                this.$store.commit('setting/init')
            }
        },
        async selectImageFromLocal (file) {
            const data = await getBase64FromRaw(file.raw)
            this.$store.commit('setting/setBackgroundImage', data)
        },
    },
}
</script>

<style scoped>
p {
    font-family: "Space Mono", yuan
}
</style>
