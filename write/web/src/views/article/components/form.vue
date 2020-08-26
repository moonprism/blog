<template>
    <el-form label-position="right" label-width="90px" class="updateForm">
        <el-form-item label="status:">
            <el-select v-model="$props.article.status" placeholder="status">
                <el-option
                        v-for="item in statusList"
                        :key="item.status"
                        :label="item.text"
                        :value="item.status">
                </el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="title:">
            <el-input v-model="$props.article.title" style="width:300px"></el-input>
        </el-form-item>
        <el-form-item label="tags:">
            <el-tag
                    class="tag"
                    v-for="tag in $props.article.tags"
                    :key="tag.id"
                    :color="tag.color"
                    :disable-transitions="false"
                    effect="dark"
                    closable
                    @close="handleTagClose(tag)"
            >
                {{tag.name}}
            </el-tag>
            <el-button size="small" @click="tagSelectVisible = true">+ Add Tag</el-button>
        </el-form-item>
        <el-form-item label="image:">
            <!--图片上传模块-->
            <upload-img :image="image" :config="uploadConfig" @on-del="uploadImgDel" @on-cancel="uploadImgCancel"></upload-img>
        </el-form-item>
        <el-form-item label="summary:">
            <el-input type="textarea" v-model="$props.article.summary" :rows="5"></el-input>
        </el-form-item>
        <div class="button-wrapper">
            <el-button @click="update" icon="el-icon-upload2" type="primary">Update</el-button>
        </div>
        <el-dialog
                title="select tag"
                v-if="tagSelectVisible"
                :close-on-click-modal="false"
                :visible.sync="tagSelectVisible">
            <!--标签选择模块-->
            <tag-select @click-tag="handleTagClick" :selectTags="$props.article.tags"></tag-select>
        </el-dialog>
    </el-form>
</template>

<script>
    import uploadImg from '@/components/image/upload'
    import {dataURItoBlob} from '@/utils/file'
    import fileApi from  '@/api/file'
    import articleApi from '@/api/article'
    import tagSelect from '@/components/tag/select'

    export default {
        name: 'ArticleForm',
        components: {
            uploadImg,
            tagSelect
        },
        data() {
            return {
                image: {
                    url: '',
                    originUrl: '',
                    key: '',
                },
                uploadConfig: {
                    showSelect: true,
                },
                tagSelectVisible: false
            }
        },
        computed: {
            statusList() {
                return this.$store.getters['article/statusList']
            }
        },
        props: {
            article: {
                tags: []
            }
        },
        mounted() {
            if (this.$props.article.image)
                this.image.url = this.image.originUrl = process.env.VUE_APP_FILE_ORIGIN + this.$props.article.image
        },
        methods: {
            async update() {
                if (this.image.key) {
                    this.$props.article.image = this.image.key
                } else if (this.image.url != this.image.originUrl) {
                    // 文件上传 or 文件弃用
                    if (this.image.url) {
                        const blob = dataURItoBlob(this.image.url)
                        const fd = new FormData()
                        fd.append("image", blob)
                        const res = await fileApi.uploadImage(fd)
                        this.$props.article.image = res.data.key
                    } else {
                        this.$props.article.image = ""
                    }
                }
                this.requestApi()
            },
            async requestApi() {
                if (this.$props.article.id) {
                    // update
                    await articleApi.update(this.$props.article.id, this.$props.article)
                    this.successMessage('update success')
                } else {
                    // create
                    await articleApi.create(this.$props.article)
                    this.successMessage('create success')
                }
            },
            successMessage(text) {
                this.$message({
                    message: text,
                    type: 'success',
                    duration: 3 * 1000
                })
                this.$emit('update-data', this.$props.article)
            },
            uploadImgDel() {
                this.image.url = ''
            },
            uploadImgCancel() {
                this.image.url = ''
            },
            handleTagClick(tag) {
                this.$props.article.tags.unshift(tag)
                this.tagSelectVisible = false
            },
            handleTagClose(tag) {
                this.$props.article.tags = this.$props.article.tags.filter((element) => {
                    return tag.id !== element.id
                })
            }
        }
    }
</script>

<style scoped>
    .illustration {
        width: 100%;
        margin-top: 12px;
    }
    .button-wrapper {
        text-align: center;
    }
    .tag {
        border: none;
        margin: 0 5px;
    }
    .el-form-item__label {
        font-size: 16px;
    }
</style>
