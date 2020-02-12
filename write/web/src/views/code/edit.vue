<template>
    <el-row :gutter="20">
        <el-col class="code-edit" :xs="24" :md="12" :lg="12" :xl="12">
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

<style scoped>
    .code-edit {
        float: none;
        margin: 0 auto;
    }
</style>