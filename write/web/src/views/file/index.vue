<template>
    <div class="file-wrapper">
        <div v-if="!isEdit" class="file-list">
            <div class="file-list-banner">
                <el-upload
                        action="#"
                        :auto-upload="false"
                        :show-file-list="false"
                        :on-change="handleImageUpload">
                    <el-button icon="el-icon-upload" type="primary" plain>upload</el-button>
                </el-upload>
            </div>
            <image-waterfall @select-image="selectImage" v-if="hackReset == true"></image-waterfall>
        </div>
        <div v-if="isEdit" class="file-edit">
            <!--图片上传模块-->
            <div class="upload-image-wrapper">
                <upload-img :image="image" :config="uploadConfig" @on-del="deleteFile" @on-cancel="uploadImgCancel" @on-upload="uploadImg"></upload-img>
            </div>
        </div>
    </div>
</template>

<script>
    import UploadImg from '@/components/image/upload'
    import ImageWaterfall from "@/components/image/waterfall"
    import {getBase64FromRaw, dataURItoBlob} from "@/utils/file"
    import FileApi from  '@/api/file'

    export default {
        data() {
            return {
                isEdit: false,
                image: {
                    url: '',
                    originUrl: ''
                },
                uploadConfig: {
                    showCancel: true,
                    isUpload: false,
                    showUpload: true,
                    showSelect: false,
                },
                imageModel: {},
                hackReset: true
            }
        },
        components: {
            ImageWaterfall,
            UploadImg
        },
        methods: {
            selectImage(imageModel) {
                this.image.url = this.image.originUrl = process.env.VUE_APP_FILE_ORIGIN + imageModel.key
                this.isEdit = true
                this.imageModel = imageModel
                this.uploadConfig.isUpload = false
            },
            async handleImageUpload(file) {
                this.imageModel = {}
                const data = await getBase64FromRaw(file.raw)
                this.image.originUrl = ''
                this.image.url = data
                this.uploadConfig.isUpload = true
                this.isEdit = true
            },
            async deleteFile() {
                if (this.imageModel.id) {
                    // delete file
                    this.$confirm("delete "+this.imageModel.key+'?<br><img class="conf-img" src="'+this.image.originUrl+'"/>', 'Delete File', {
                        confirmButtonText: '确定',
                        cancelButtonText: '取消',
                        dangerouslyUseHTMLString: true,
                    }).then(async () => {
                        // 调用删除图片接口
                        let res = await FileApi.deleteImage(this.imageModel.id)
                        if (res.data === "ok") {
                            this.successMessage('删除')
                        }
                    })
                }
                this.isEdit = false
            },
            uploadImgCancel() {
                this.isEdit = false
            },
            async uploadImg() {
                // upload file
                this.$confirm('upload image to server?<br><img class="conf-img" src="'+this.image.url+'"/>', 'upload File', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    dangerouslyUseHTMLString: true,
                }).then(async () => {
                    const blob = dataURItoBlob(this.image.url)
                    const fd = new FormData()
                    fd.append("image", blob)
                    let res = await FileApi.uploadImage(fd)
                    if (res.data.key) {
                        this.successMessage('上传')
                    }
                })
                this.isEdit = false
            },
            successMessage(msg) {
                this.$message({
                    type: 'success',
                    message: msg + '成功!'
                })
                this.hackReset = false
                this.$nextTick(() => {
                    this.hackReset = true
                })
            }
        }
    }
</script>

<style>
    .conf-img {
        width: 100%;
    }
    .file-list-banner {
        text-align: center;
        margin-bottom: 10px;
    }
    .file-edit {
        text-align: center;
        padding-top: 40px;
    }
    .file-edit .image-operating{
        margin-top: 10px;
    }
    .operation {
        text-align: right;
    }
    .upload-image-wrapper {
        width: 70%;
        margin: 0 auto;
    }
</style>