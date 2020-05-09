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
            <h2>background</h2>
            <p>
                <ul>
                    <li> <a @click="waterfallDialog.visible = true">select background</a> </li>
                    <li> <a @click="clear">clear</a> </li>
                    <li> 
                        
                        <el-upload
                            action="#"
                            :auto-upload="false"
                            :show-file-list="false"
                            :on-change="selectImageFromLocal">
                            <a>test local image</a>
                        </el-upload>
                    </li>
                </ul>
            </p>
            <h2>pages</h2>
            <p>
                <ul>
                    <li> <a @click="$router.push({name: 'article_edit', params: {id: 1}})">link</a> </li>
                    <li> <a @click="$router.push({name: 'article_edit', params: {id: 2}})">about</a> </li>
                </ul>
            </p>
        </div>
    </el-row>
</template>

<script>
import ImageWaterfall from "@/components/image/waterfall"
import '@/style/markdown.css'
import {getBase64FromRaw} from "@/utils/file"

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
        async save() {
            await this.$store.dispatch('setting/post', this.SettingInfo)
        },
        async selectImage(imageModel) {
            this.SettingInfo.background_image = process.env.VUE_APP_FILE_ORIGIN + imageModel.key
            await this.save()
            this.waterfallDialog.visible = false
        },
        async clear() {
            this.SettingInfo.background_image = '',
            await this.save()
        },
        async selectImageFromLocal (file) {
            const data = await getBase64FromRaw(file.raw)
            this.$store.commit('setting/setBackgroundImage', data)
        }
    }
}
</script>

<style scoped>
</style>