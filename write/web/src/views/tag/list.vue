<template>
    <el-row :gutter="20" style="margin-top: 20px">
        <el-col :span="13" :offset="4">
            <div class="tag-list">
                <el-table
                        :data="tagList"
                        style="width: 100%">
                    <el-table-column
                            width="120px"
                            label="Tag">
                        <template slot-scope="scope">
                            <i-tag :tag="scope.row"></i-tag>
                        </template>
                    </el-table-column>
                    <el-table-column
                            width="220px"
                            label="Name">
                        <template slot-scope="scope">
                            <el-input v-model="scope.row.name"></el-input>
                        </template>
                    </el-table-column>
                    <el-table-column
                            width="80px"
                            label="Color">
                        <template slot-scope="scope">
                            <el-color-picker v-model="scope.row.color"></el-color-picker>
                        </template>
                    </el-table-column>
                    <el-table-column
                            width="100px"
                            label="Reference">
                        <template slot-scope="scope">
                            <div style="text-align: center">{{scope.row.count}}</div>
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
                                    v-if="scope.row.id"
                                    size="mini"
                                    @click="handleUpdate(scope.row)">Update</el-button>
                            <el-button
                                    v-if="scope.row.id"
                                    size="mini"
                                    type="danger"
                                    @click="handleDelete(scope.row)">Delete</el-button>
                            <el-button
                                    v-else
                                    size="mini"
                                    type="primary"
                                    @click="handleInsert(scope.row)">Insert</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </el-col>
    </el-row>
</template>

<script>
    import tagApi from "@/api/tag"
    import iTag from "@/components/tag/tag"

    export default {
        components: {
            iTag
        },
        data() {
            return {
                tagList: []
            }
        },
        methods: {
            async fetchAndSetTagList() {
                const res = await tagApi.list()
                this.tagList = res.data
            },
            handleCreate() {
                this.tagList.unshift({
                    id: 0,
                    name: '',
                    color: '',
                })
            },
            async handleUpdate(tag) {
                const res = await tagApi.update(tag.id, tag)
                if (res.data === 'ok') {
                    this.$message({
                        message: "update success",
                        type: 'success',
                        duration: 3 * 1000
                    })
                }
            },
            async handleDelete(tag) {
                this.$confirm('<'+tag.name+'>', 'delete tag', {
                    confirmButtonText: 'Delete',
                    cancelButtonText: 'Cancel',
                }).then(async () => {
                    const res = await tagApi.delete(tag.id)
                    if (res.data === 'ok') {
                        this.$message({
                            message: "delete success",
                            type: 'success',
                            duration: 3 * 1000
                        })
                        this.fetchAndSetTagList()
                    }
                })
            },
            async handleInsert(tag) {
                const res = await tagApi.create(tag)
                if (res.data === 'ok') {
                    this.$message({
                        message: "create success",
                        type: 'success',
                        duration: 3 * 1000
                    })
                    this.fetchAndSetTagList()
                }
            }
        },
        mounted() {
            this.fetchAndSetTagList()
        }
    }
</script>
