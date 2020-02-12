<template>
    <el-row :gutter="20" style="margin-top: 15px;text-align: center">
        <el-col class="code-list" :xs="24" :md="12" :lg="12" :xl="10">
            <div class="article-list">
                <el-table
                        :data="codeList"
                        style="width: 100%">
                    <el-table-column
                            prop="description"
                            label="Description"
                            width="400px">
                    </el-table-column>
                    <el-table-column
                            width="100px"
                            label="Language">
                        <template slot-scope="scope">
                            <el-tag
                                    type="success"
                                    disable-transitions>
                                {{scope.row.lang}}
                            </el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column
                            align="right"
                            width="150px">
                        <template slot="header">
                            <el-button
                                    icon="el-icon-plus"
                                    type="success"
                                    size="mini"
                                    @click="handleCreate">Create</el-button>
                        </template>
                        <template slot-scope="scope">
                            <el-button
                                    size="mini"
                                    @click="handleEdit(scope.row)">Edit</el-button>
                            <el-button
                                    size="mini"
                                    type="danger"
                                    @click="handleDelete(scope.row)">Delete</el-button>
                        </template>
                    </el-table-column>
                </el-table>
                <div class="pagination">
                    <el-pagination
                            background
                            @current-change="handlePageChange"
                            :current-page="pagination.page"
                            :page-size="pagination.limit"
                            layout="prev, pager, next"
                            :total="pagination.total">
                    </el-pagination>
                </div>
            </div>
        </el-col>
    </el-row>
</template>

<script>
    import codeApi from '@/api/code'

    export default {
        data() {
            return {
                codeList: [],
                pagination: {
                    limit: 0,
                    page: 0,
                    total: 0,
                }
            }
        },
        created() {
            this.fetchAndSetList(this.$route.params.page || 1)
        },
        beforeRouteUpdate (to, from, next) {
            this.fetchAndSetList(to.params.page || 1)
            next()
        },
        methods: {
            handlePageChange(currentPage) {
                this.$router.push({name:'code_list', params: {page: currentPage}})
            },
            async fetchAndSetList(page) {
                const res = await codeApi.list(page)
                this.codeList = res.data.data
                this.pagination = res.data.pagination
            },
            handleCreate() {
                this.$router.push({name: 'code_create'})
            },
            handleEdit(code) {
                this.$router.push({name: 'code_edit', params: {id: code.id}})
            },
            async handleDelete(code) {
                this.$confirm('[ '+code.description+' ]', 'delete code', {
                    confirmButtonText: 'Delete',
                    cancelButtonText: 'Cancel',
                }).then(async () => {
                    const res = await codeApi.delete(code.id)
                    if (res.data === 'ok') {
                        this.$message({
                            message: "delete success",
                            type: 'success',
                            duration: 3 * 1000
                        })
                        this.fetchAndSetList(this.$route.params.page || 1)
                    }
                })
            }
        }
    }
</script>

<style scoped>
    .pagination {
        margin-top: 50px;
        text-align: center;
    }
    .code-list {
        float: none;
        margin: 0 auto;
    }
</style>