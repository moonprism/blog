<template>
    <el-form label-position="top" label-width="80px" class="updateForm code-form">
    <el-form-item label="language">
        <el-select v-model="$props.code.lang" @change="changeLang" placeholder="status">
            <el-option
                    v-for="(value, key) in languageList"
                    :key="key"
                    :label="value"
                    :value="key">
            </el-option>
        </el-select>
    </el-form-item>
        <el-form-item label="tags">
            <el-tag
                    :key="tag"
                    v-for="tag in tags"
                    closable
                    :disable-transitions="false"
                    @close="handleTagClose(tag)">
                {{tag}}
            </el-tag>
            <el-input
                    class="input-new-tag"
                    v-if="tagInputVisible"
                    v-model="tagInputValue"
                    ref="saveTagInput"
                    size="small"
                    @keyup.enter.native="handleTagInputConfirm"
                    @blur="handleTagInputConfirm"
            >
            </el-input>
            <el-button v-else class="button-new-tag" size="small" @click="showTagInput">+ New Tag</el-button>
        </el-form-item>
    <el-form-item label="description">
        <el-input v-model="$props.code.description" style="width:70%;font-size: 17px;"></el-input>
    </el-form-item>
    <el-form-item label="code">
        <!-- code mirror 编辑模块 -->
        <code-edit ref="codeEdit" :code="$props.code" :config="{useVim:true}"></code-edit>
    </el-form-item>
        <li><a target="_blank" rel="noopener" style="color:#ae81ff;cursor:pointer;font-size:15px" @click="preview">preview</a></li>
        <div class="button-wrapper">
            <el-button @click="update" icon="el-icon-upload2" type="primary">Update</el-button>
        </div>
    </el-form>
</template>
<script>
    import codeEdit from '@/components/code-edit'
    import codeApi from '@/api/code'

    export default {
        components: {
            codeEdit
        },
        data() {
            return {
                tagInputVisible: false,
                tagInputValue: '',
                languageList: {
                    md: 'markdown',
                    js: 'javascript',
                    php: 'php',
                    go: 'go',
                    sh: 'shell',
                    css: 'css',
                    sql: 'sql'
                },
            }
        },
        props: {
            code: {}
        },
        computed: {
            tags() {
                if (!this.$props.code.tags) {
                    return []
                }
                return this.$props.code.tags.split(',')
            }
        },
        mounted() {
            this.$props.code.lang = 'md'
        },
        methods: {
            preview() {
                window.open(this.previewLink(), '_blank');
            },
            previewLink() {
                this.$props.code.content = this.$refs.codeEdit.getContent()
                return process.env.VUE_APP_READ_ORIGIN+
                    'preview/code?lang='+ this.$props.code.lang+
                    '&content='+btoa(encodeURIComponent(this.$props.code.content))+
                    '&description='+encodeURIComponent(this.$props.code.description)+
                    '&tags='+encodeURIComponent(this.$props.code.tags);
            },
            handleTagClose(tag) {
                let tags = this.tags
                tags.splice(tags.indexOf(tag), 1);
                this.$props.code.tags = tags.join(',')
            },
            showTagInput() {
                this.tagInputVisible = true;
                this.$nextTick(() => {
                    this.$refs.saveTagInput.$refs.input.focus();
                });
            },
            handleTagInputConfirm() {
                let inputValue = this.tagInputValue;
                if (inputValue) {
                    if (this.$props.code.tags != '') {
                        this.$props.code.tags += (',' + inputValue)
                    } else {
                        this.$props.code.tags = inputValue
                    }
                }
                this.tagInputVisible = false;
                this.tagInputValue = '';
            },
            changeLang() {
                this.$refs.codeEdit.updateLang()
            },
            async update() {
                this.$props.code.content = this.$refs.codeEdit.getContent()
                if (this.$props.code.id) {
                    // update
                    await codeApi.update(this.$props.code.id, this.$props.code)
                    this.successMessage('update success')
                } else {
                    // create
                    await codeApi.create(this.$props.code)
                    this.successMessage('create success')
                }
            },
            successMessage(text) {
                this.$message({
                    message: text,
                    type: 'success',
                    duration: 3 * 1000
                })
                this.$emit('update-data', this.$props.code)
            },
        }
    }
</script>

<style>
    .code-form .el-tag + .el-tag {
        margin-left: 10px;
    }
    .code-form .button-new-tag {
        margin-left: 10px;
        height: 32px;
        line-height: 30px;
        padding-top: 0;
        padding-bottom: 0;
    }
    .code-form .input-new-tag {
        width: 90px;
        margin-left: 10px;
        vertical-align: bottom;
    }
    .code-form .button-wrapper {
        text-align: center;
    }
    .code-form .el-textarea__inner {
        font-family: 'Space Mono', "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif !important;
    }
    .el-form-item__label {
        font-size: 16px;
    }
</style>
