<template>
    <el-row :gutter="20" style="margin-top: 15px;text-align: center">
        <el-col class="code-list">
            <div class="article-list">
                <el-table
                        :data="comments"
                        style="width: 100%">
                    <el-table-column type="expand">
                        <template slot-scope="scope">
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>email:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.email }}</div></el-col>
                            </el-row>
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>article:</span></el-col>
                                <el-col :span="16"><div><el-link target="_blank" type="primary" :href="articleLink(scope.row.art_id)">{{ scope.row.art_id }}</el-link></div></el-col>
                            </el-row>
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>to:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.to_id }}</div></el-col>
                            </el-row>
                        </template>
                    </el-table-column>
                    <el-table-column
                            prop="id"
                            label="ID">
                    </el-table-column>
                    <el-table-column
                            width="400"
                            prop="text"
                            label="Content">
                    </el-table-column>
                    <el-table-column
                            prop="name"
                            label="Name">
                    </el-table-column>
                    <el-table-column
                            align="right">
                        <template slot-scope="scope">
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
    import commentApi from '@/api/comment'

    export default {
        data() {
            return {
                comments: [],
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
            articleLink(art_id) {
                return process.env.VUE_APP_READ_ORIGIN+'article/id/'+art_id+'#com_up';
            },
            handlePageChange(currentPage) {
                this.$router.push({name:'comment_list', params: {page: currentPage}})
            },
            async fetchAndSetList(page) {
                const res = await commentApi.list(page)
                this.comments = res.data.data
                this.pagination = res.data.pagination
            },
            async handleDelete(comment) {
                this.$confirm('[ '+ comment.name + ': ' + comment.text+' ]', 'delete comment', {
                    confirmButtonText: 'Delete',
                    cancelButtonText: 'Cancel',
                }).then(async () => {
                    const res = await commentApi.delete(comment.id)
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
    .exp {
        margin: 5px;
    }
    .exp div {
        display: block;
    }
    .exp span {
        color: #99a9bf;
    }
</style>

