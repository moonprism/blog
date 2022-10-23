<template>
    <el-row :gutter="20" style="margin-top: 15px;text-align: center">
        <el-col class="code-list">
        <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="评论" name="comment">
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
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>time:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.created_time }}</div></el-col>
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
        </el-tab-pane>
        <el-tab-pane label="链接申请" name="link">
            <div class="link-list">
                <el-table
                        :data="links"
                        style="width: 100%">
                    <el-table-column type="expand">
                        <template slot-scope="scope">
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>email:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.email }}</div></el-col>
                            </el-row>
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>link:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.link }}</div></el-col>
                            </el-row>
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>icon:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.avatar }}</div></el-col>
                            </el-row>
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>intro:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.message }}</div></el-col>
                            </el-row>
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>message:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.data }}</div></el-col>
                            </el-row>
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>time:</span></el-col>
                                <el-col :span="16"><div>{{ scope.row.created_time }}</div></el-col>
                            </el-row>
                            <el-row :gutter="20" class="exp">
                                <el-col :span="4"><span>code:</span></el-col>
                                <el-col :span="16"><div>> [![]({{ scope.row.avatar }}){{ scope.row.name }}`{{ scope.row.icon }}*{{ scope.row.message }}*`]({{ scope.row.link }})</div></el-col>
                            </el-row>
                        </template>
                    </el-table-column>
                    <el-table-column
                            prop="id"
                            label="ID">
                    </el-table-column>
                    <el-table-column
                            width="250"
                            prop="email"
                            label="Email">
                    </el-table-column>
                    <el-table-column
                            width="70"
                            prop="icon"
                            label="Emoji">
                    </el-table-column>
                    <el-table-column
                            prop="name"
                            label="Name">
                    </el-table-column>
                    <el-table-column
                            align="right">
                    </el-table-column>
                </el-table>
                <div class="pagination">
                    <el-pagination
                            background
                            @current-change="fetchLinkList"
                            :current-page="linkPagination.page"
                            :page-size="linkPagination.limit"
                            layout="prev, pager, next"
                            :total="linkPagination.total">
                    </el-pagination>
                </div>
            </div>
        </el-tab-pane>
        </el-tabs>
        </el-col>
    </el-row>
</template>

<script>
    import commentApi from '@/api/comment'
    import linkApi from '@/api/link'

    export default {
        data() {
            return {
                activeName: 'comment',
                comments: [],
                pagination: {
                    limit: 0,
                    page: 0,
                    total: 0,
                },
                linkPagination: {
                    limit: 0,
                    page: 0,
                    total: 0,
                },
                links: [],
            }
        },
        created() {
            this.fetchAndSetList(this.$route.params.page || 1)
            this.fetchLinkList(1)
        },
        beforeRouteUpdate (/* to, from, next */) {
            /* this.fetchAndSetList(to.params.page || 1) */
            /* next()                                    */
        },
        methods: {
            // 意料之外的新数据
            async fetchLinkList(page) {
                const res = await linkApi.list(page)
                this.links = res.data.data
                this.linkPagination = res.data.pagination
            },

            articleLink(art_id) {
                return process.env.VUE_APP_READ_ORIGIN+'article/'+art_id+'#com_up';
            },
            handlePageChange(currentPage) {
                /* this.$router.push({name:'comment_list', params: {page: currentPage}}) */
                this.fetchAndSetList(currentPage)
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

