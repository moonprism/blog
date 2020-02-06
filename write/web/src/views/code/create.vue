<template>
    <el-row :gutter="20">
        <el-col :span="13" :offset="4">
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
