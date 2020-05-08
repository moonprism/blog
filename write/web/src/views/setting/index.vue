<template>
    <el-row :gutter="20">
        <div class="an" v-bind:style="{ backgroundImage: url}" ></div>
        <el-dialog
            v-if="waterfallDialog.visible"
                :close-on-click-modal="false"
                :visible.sync="waterfallDialog.visible">
            <!--图片瀑布流选择模块-->
            <image-waterfall :size="waterfallSizeConfig" @select-image="selectImage"></image-waterfall>
        </el-dialog>
        <div class="board markdown">
            <h2>background</h2>
            <p>
                <ul>
                    <li> <a @click="waterfallDialog.visible = true">select background</a> </li>
                    <li> <a @click="clear">clear</a> </li>
                </ul>
            </p>
        </div>
    </el-row>
</template>

<script>
import SettingApi from "@/api/setting"
import ImageWaterfall from "@/components/image/waterfall"
import '@/style/markdown.css'

export default {
    components: {
        ImageWaterfall
    },
    data() {
        return {
            waterfallDialog: {
                visible: false
            },
            SettingInfo: {
                background_image: '',
            },
            waterfallSizeConfig: {
                width: 150
            },
        }
    },
    methods: {
        async fetchAndRenderInfo() {
            const res = await SettingApi.detail()
            if (res.data) {
                this.SettingInfo = res.data
            }
        },
        async save() {
            const res = await SettingApi.update(this.SettingInfo)
            if (res.data === 'ok') {
                this.$message({
                    message: "update success",
                    type: 'success',
                    duration: 3 * 1000
                })
            }
        },
        async selectImage(imageModel) {
            this.SettingInfo.background_image = process.env.VUE_APP_FILE_ORIGIN + imageModel.key
            await this.save()
            this.waterfallDialog.visible = false
        },
        async clear() {
            this.SettingInfo.background_image = '',
            await this.save()
        }
    },
    async mounted() {
        await this.fetchAndRenderInfo()
    },
    computed: {
        url() {
            return 'url(' + this.SettingInfo.background_image + ')'
        }
    }
}
</script>

<style scoped>
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
    margin: 10px auto 290px;
    background-color: #fff;
    padding: 35px 47px 30px;
    border-radius: 3px;
    min-height: 800px;
}
</style>