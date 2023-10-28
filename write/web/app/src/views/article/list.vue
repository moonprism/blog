<template>
    <el-row :gutter="20" style="margin-top: 10px">
        <el-col class="article-l">
            <div class="article-list">
                <el-table
                        :data="articleList"
                        style="width: 100%">
                    <el-table-column type="expand">
                        <template slot-scope="scope">
                            <!--update模块-->
                            <article-form :article="scope.row" @update-data="loadData"></article-form>
                        </template>
                    </el-table-column>

                    <el-table-column
                            prop="title"
                            label="Title"
                            width="300px">
                    </el-table-column>
                    <el-table-column
                            width="120px"
                            label="Status">
                        <template slot-scope="scope">
                            <el-tag
                                    :type="statusInfo(scope.row.status)['type']"
                                    disable-transitions>{{statusInfo(scope.row.status)['text']}}</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column
                            width="130px"
                            label="Updated">
                        <template slot-scope="scope">
                            <i class="el-icon-time"></i>
                            <span style="margin-left: 7px">{{ dateFormat(scope.row.updated_time) }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column
                            align="right">
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
    import articleApi from "@/api/article"
    import {dateFormat} from "@/utils/time"
    import articleForm from './components/form'

    export default {
        data() {
            return {
                articleList: null,
                pagination: {
                    limit: 0,
                    page: 0,
                    total: 0,
                },
            }
        },
        components: {
            articleForm
        },
        beforeRouteUpdate (to, from, next) {
            this.fetchAndSetArticleList(to.params.page || 1)
            next()
        },
        computed: {
            statusList() {
              return this.$store.getters['article/statusList']
            }
        },
        created() {
            this.loadData()
        },
        methods: {
            handlePageChange(currentPage) {
                this.$router.push({name:'article_list', params: {page: currentPage}})
            },
            async fetchAndSetArticleList(page) {
                const res = await articleApi.list(page)
                this.articleList = res.data.data
                this.pagination = res.data.pagination
            },
            dateFormat(data) {
                return dateFormat(data, 'yyyy-MM-dd')
            },
            statusInfo(status) {
                for (let i=0; i < this.statusList.length; i++) {
                    if (this.statusList[i].status == status) {
                        return this.statusList[i]
                    }
                }
                // 未知状态为draft
                return this.statusList[0]
            },
            loadData() {
                this.fetchAndSetArticleList(this.$route.params.page || 1)
            },
            handleCreate() {
                this.$router.push({name: 'article_create'})
            },
            async handleDelete(article) {
                this.$confirm('《'+article.title+'》', 'delete article', {
                    confirmButtonText: 'Delete',
                    cancelButtonText: 'Cancel',
                }).then(async () => {
                    const res = await articleApi.delete(article.id)
                    if (res.data === 'ok') {
                        this.$message({
                            message: "delete success",
                            type: 'success',
                            duration: 3 * 1000
                        })
                        this.loadData()
                    }
                })
            },
            handleEdit(article) {
                this.$router.push({name: 'article_edit', params: {id: article.id}})
            }
        }
    }
</script>

<style scoped>
    .pagination {
        text-align: right;
        padding: 30px 10px;
        margin-right: 25px;
    }
    .updateForm {
        
    }
    .article-l {
        float: none;
        margin: 0 auto;
    }
</style>
