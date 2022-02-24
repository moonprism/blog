<template>
    <div class="upload-wrapper">
        <el-upload
                action="#"
                :auto-upload="autoLoad"
                :show-file-list="false"
                :on-change="handleImageChange">
            <el-image
                    class = "illustration"
                    v-if="$props.image.url"
                    :src="$props.image.url"></el-image>
            <el-button v-else icon="el-icon-plus" size="small" type="primary" plain>File</el-button>
        </el-upload>

        <div v-if="this.$props.image.url" class="image-operating" style="text-align: center">
            <el-button v-if="isShowUpload" @click="$emit('on-upload')" type="primary" size="small" icon="el-icon-upload2" plain>Upload</el-button>
            <el-button @click="handleImageCrop" size="small" icon="el-icon-scissors" plain>Crop</el-button>
            <el-button v-if="isCrop" @click="handleImageReset" size="small" icon="el-icon-refresh-left" type="success" plain>Reset</el-button>

            <el-button v-if="isOrigin" @click="handleImageDelete" size="small" icon="el-icon-delete" type="danger" plain>Delete</el-button>
            <el-button v-if="isShowCancel || !isOrigin" @click="handleImageCancel" size="small" type="info" plain>Cancel</el-button>
        </div>
        <div v-if="isShowSelect" class="select-img">
            <el-button @click="handleImageSelect" icon="el-icon-plus" size="small" type="primary" plain>Select</el-button>
        </div>
        <el-dialog
                v-if="waterfallDialog.visible"
                :close-on-click-modal="false"
                :visible.sync="waterfallDialog.visible">
            <!--图片瀑布流选择模块-->
            <image-waterfall :size="waterfallSizeConfig" @select-image="selectImage"></image-waterfall>
        </el-dialog>
        <el-dialog
                v-if="cropDialog.visible"
                :close-on-click-modal="false"
                :visible.sync="cropDialog.visible">
            <!--图片裁剪模块-->
            <crop :image="this.$props.image" :dialog="cropDialog"></crop>
        </el-dialog>
    </div>
</template>

<script>
    import crop from './crop'
    import {getBase64FromRaw} from "@/utils/file";
    import ImageWaterfall from "@/components/image/waterfall"

    export default {
        name: 'UploadImage',
        data() {
            return {
                cropDialog: {
                    visible: false,
                    afterCropClose: true,
                    originImgUrl: '',
                },
                waterfallDialog: {
                    visible: false
                },
                autoLoad: false,
                uploadImageUrl: '',
                waterfallSizeConfig: {
                    width: 150
                }
            }
        },
        computed: {
            isCrop() {
                return (this.$props.image.url !== this.uploadImageUrl) && this.$props.image.originUrl !== this.$props.image.url
            },
            isOrigin() {
                return this.$props.image.originUrl !== ''
            },
            isShowCancel() {
                return this.$props.config && this.$props.config.showCancel
            },
            isShowUpload() {
                return this.$props.config && this.$props.config.showUpload && this.$props.image.url != this.$props.image.originUrl
            },
            isShowSelect() {
                return this.$props.config && this.$props.config.showSelect
            }
        },
        created() {
            if (this.$props.config && this.$props.config.isUpload) {
                this.uploadImageUrl = this.$props.image.url
            }
        },
        components: {
            crop,
            ImageWaterfall
        },
        props: {
            image: {
                url: '',
                originUrl: '',
                key: '',
            },
            config: {
                isUpload: false,
                showCancel: false,
                showUpload: false,
                showSelect: true,
            }
        },
        methods: {
            async handleImageChange(file) {
                const data = await getBase64FromRaw(file.raw)
                this.$props.image.url = this.uploadImageUrl = data
            },
            handleImageDelete() {
                this.$emit('on-del')
            },
            handleImageCrop() {
                this.cropDialog.visible = true
            },
            handleImageReset() {
                if(this.uploadImageUrl !== '')
                    this.$props.image.url = this.uploadImageUrl
                else
                    this.$props.image.url = this.$props.image.originUrl
            },
            handleImageCancel() {
                this.$emit('on-cancel')
            },
            handleImageUpload() {
                this.$emit('on-upload')
            },
            handleImageSelect() {
                this.waterfallDialog.visible = true
            },
            selectImage(imageModel) {
                this.$props.image.url = this.$props.image.originUrl = process.env.VUE_APP_FILE_ORIGIN + imageModel.key
                this.$props.image.key = imageModel.key
                this.waterfallDialog.visible = false
            },
        }
    }
</script>