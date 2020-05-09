<template>
    <el-row :gutter="20">
        <el-col>
            <article-form :article="createArticleModel" @update-data="returnList"></article-form>
        </el-col>
    </el-row>
</template>

<script>
    import articleForm from './components/form'

    export default {
        data() {
            return {
                createDialogVisible: false,
                createArticleModel: {
                    status: 1,
                    tags: []
                }
            }
        },
        components: {
            articleForm
        },
        methods: {
            returnList() {
                this.$router.push({name: 'article'})
            }
        }
    }
</script>