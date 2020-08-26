<template>
    <el-container class="edit-main">
        <el-header>
            <div class="edit-header">
                <!--       先这样         -->
                <el-button @click="handleUpload" plain>Show</el-button>
                <el-button @click="handleSelect" plain>Select</el-button>
                <el-button @click="handleWrite" type="primary" plain>Write</el-button>
                <a @click="handleClose" class="close"><i class="el-icon-circle-close"></i></a>
            </div>
        </el-header>
        <el-main style="padding:0">
            <div class="edit-wrapper">
            <textarea
                style="display: none"
                ref="edit"
                v-model="article.content"
                :options="cmOptions">
            </textarea>
            <el-dialog
                class="i-dialog"
                fullscreen
                :v-if="articleVisible"
                :visible.sync="articleVisible"
                width="30%">
                <!--文章显示模块-->
                <article-show :articleShow="articleShow"></article-show>
            </el-dialog>
                <el-dialog
                        v-if="waterfallDialog.visible"
                        :close-on-click-modal="false"
                        :visible.sync="waterfallDialog.visible">
                    <!--图片瀑布流选择模块-->
                    <image-waterfall :size="waterfallSizeConfig" @select-image="selectImage"></image-waterfall>
                </el-dialog>
            </div>
        </el-main>
    </el-container>
</template>

<script>
    import articleApi from "@/api/article"
    import CodeMirror from 'codemirror/lib/codemirror'
    import 'codemirror/lib/codemirror.css'
    import 'codemirror/mode/markdown/markdown'
    import 'codemirror/mode/go/go'
    import 'codemirror/mode/javascript/javascript'
    import 'codemirror/mode/shell/shell'
    import 'codemirror/mode/sql/sql'
    import 'codemirror/mode/css/css'
    import 'codemirror/mode/php/php'
    import 'codemirror/keymap/vim'
    import 'codemirror/addon/dialog/dialog.css'
    import 'codemirror/theme/material.css'

    //import markdown from 'moonprism-markdown'
    process.env.VUE_APP_FILE_ORIGIN

    import ArticleShow from "./components/show"

    import ImageWaterfall from "@/components/image/waterfall"
    import {insertStr} from "@/utils/string"

    export default {
        components: {
            ArticleShow,
            ImageWaterfall
        },
        data() {
            return {
                article: {},
                articleShow: {
                    markdownText: '',
                    cursor: {
                        line: 0,
                        ch: 0
                    },
                    config: {
                    }
                },
                articleEdit: {
                    markdownText: '',
                    cursor: {
                        line: 0,
                        ch: 0
                    }
                },
                waterfallDialog: {
                    visible: false
                },
                waterfallSizeConfig: {
                    width: 150
                },
                cmOptions: {
                    mode:   'text/markdown',
                    keyMap: 'vim',
                    theme:  'material',
                    lineWrapping: true,
                    matchBrackets: true,
                    //tyautofocus: true,
                    extraKeys: {"Alt-A": this.show},
                    // 开启这三个选项 光标对齐正常但lineNumber不好调样式
                    inputStyle: "contenteditable",
                    indentUnit: 4,
                    indentWithTabs: true,
                    showCursorWhenSelecting: true,
                    lineNumbers: true,
                },
                articleVisible: false,
                articleHTML: '',
                editor: {}
            }
        },
        async mounted() {
            CodeMirror.Vim.defineEx('q', undefined, this.quit)
            CodeMirror.Vim.defineEx('wq', undefined, this.saveAndQuit)
            await this.fetchAndSetArticle(this.$route.params.id)
            CodeMirror.commands.save = this.save
            let editor = CodeMirror.fromTextArea(
                this.$refs.edit, this.cmOptions
            )
            editor.on('cursorActivity', this.cursorActivity)
            CodeMirror.Vim.map('jk', '<Esc>', 'insert')
            this.articleEdit.markdownText = this.article.content
            this.editor = editor
            this.editor.focus()
        },
        methods: {
            async fetchAndSetArticle(id) {
                const res = await articleApi.detail(id)
                this.article = res.data
            },
            async save() {
                this.article.content = this.articleEdit.markdownText = this.editor.getValue()
                //this.article.html = markdown(this.article.content, process.env.VUE_APP_FILE_ORIGIN)
                const res = await articleApi.update(this.article.id, this.article)
                if (res.data === 'ok') {
                    this.$message({
                        message: 'write success',
                        type: 'success',
                        duration: 1 * 1000
                    })
                    this.editor.setCursor(this.editor.getCursor())
                    window.scrollBy(0,-300)
                }
            },
            quit() {
                this.$router.go(-1)
            },
            saveAndQuit() {
                this.save()
                this.quit()
            },
            show() {
                this.articleShow = {
                    markdownText: this.articleEdit.markdownText,
                    cursor: this.articleEdit.cursor,
                    config: {
                        imageCDN: process.env.VUE_APP_FILE_ORIGIN
                    }
                }
                this.articleVisible = !this.articleVisible
            },
            cursorActivity(cm) {
                this.articleEdit.markdownText = cm.getValue()
                this.articleEdit.cursor = cm.getCursor()
            },
            handleSelect() {
                this.waterfallDialog.visible = true
            },
            handleUpload() {
                this.show()
            },
            handleWrite() {
                this.save()
            },
            handleClose() {
                if (this.articleEdit.markdownText !== this.article.content) {
                    this.$confirm('是否保存修改？', '', {
                        confirmButtonText: 'Save',
                        cancelButtonText: 'Cancel',
                    }).then(() => {
                        this.saveAndQuit()
                    }).catch(action => {
                        if (action === 'cancel') {
                            this.quit()
                        }
                    })
                } else {
                    this.quit()
                }
            },
            selectImage(imageModel) {
                let imageMarkdown = '![]('+imageModel.key+')'
                let lineArray = this.articleEdit.markdownText.split('\n')
                lineArray[this.articleEdit.cursor.line] = insertStr(lineArray[this.articleEdit.cursor.line], this.articleEdit.cursor.ch, imageMarkdown)
                this.articleEdit.markdownText = lineArray.join('\n')
                this.editor.setValue(this.articleEdit.markdownText)
                this.waterfallDialog.visible = false
            },
        }
    }
</script>

<style>
    .edit-main {
        background-color: #e5edf2;
        margin: -8px;
        min-height: 1000px;
    }
    .edit-header {
        padding: 5px;
        position: fixed;
        right: 10px;
        z-index: 2;
    }
    .edit-header .close {
        font-size: 28px;
        margin: 0 0 0 15px;
        position: relative;
        top: 7px;
        transition: all .2s linear;
    }
    .edit-header .close:hover {
        color: #f56c6c;
        cursor: pointer;
        transition: all .2s linear;
    }
    .cm-s-material {
        font-family: "Space Mono",yuan;
        width: 92%;
        margin: 0 auto 30px;
        font-size: 22px;
        padding: 10px;
    }
    .cm-s-material.CodeMirror {
        line-height: 1.5;
        letter-spacing: 1.3px;
        height: auto;
    }
    .cm-s-material .cm-string {
        color: #00a1d6;
    }
    .cm-s-material .cm-header-2 {
        color: #fe346e;
        font-size: 24px;
    }
    .cm-s-material .cm-header-3 {
        color: #4cbbb9;
        font-size: 23px;
    }
    .cm-s-material .cm-comment {
        color: #fb7b6b;
    }
</style>
