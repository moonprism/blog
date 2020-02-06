<template>
    <el-row :gutter="20">
        <el-col :span="13" :offset="4">
            <code-form :code="code" @update-data="returnList"></code-form>
        </el-col>
    </el-row>
</template>

<script>
    import codeForm from './components/form'
    import codeApi from '@/api/code'

    export default {
        data() {
            return {
                code: {}
            }
        },
        components: {
            codeForm
        },
        mounted() {
            this.fetchAndSetCode(this.$route.params.id)
        },
        methods: {
            async fetchAndSetCode(id) {
                const res = await codeApi.detail(id)
                this.code = res.data
            },
            returnList() {
                this.$router.push({name: 'code'})
            }
        }
    }
</script>