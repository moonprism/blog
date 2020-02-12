<template>
    <el-row :gutter="20">
        <el-col class="code-edit" :xs="24" :md="12" :lg="12" :xl="12">
            <code-form :code="createCodeModel" @update-data="returnList"></code-form>
        </el-col>
    </el-row>
</template>

<script>
    import codeForm from './components/form'

    export default {
        data() {
            return {
                createCodeModel: {
                    lang: 'markdown',
                    tags: ''
                }
            }
        },
        components: {
            codeForm
        },
        methods: {
            returnList() {
                this.$router.push({name: 'code'})
            }
        }
    }
</script>

<style scoped>
    .code-edit {
        float: none;
        margin: 0 auto;
    }
</style>