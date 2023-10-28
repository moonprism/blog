<template>
    <div class="image-list-wrapper">
    <div class="picker">
        <el-date-picker
           v-model="mon"
           type="month"
           @change="createList()"
           placeholder="选择月">
        </el-date-picker>
    </div>
        <Waterfall :list="list" :gutter="15" :width="waterfallItemWidth" :phoneCol="2" ref="waterfall">
            <template slot="item" slot-scope="props">
                <el-card :body-style="{padding:'0px'}">
                    <div class="card" v-bind:class="{selected: props.data.selected}" @click="handleClick(props.data)">
                        <div v-if="props.data.selected" class="selected-icon"><i class="el-icon-success"></i></div>
                        <img :src="props.data.src" alt="" @load="$refs.waterfall.refresh()">
                        <div class="info">
                            <div class="bottom-time"><time class="time">{{props.data.time}}</time></div>
                        </div>
                    </div>
                </el-card>
            </template>
        </Waterfall>
        <div class="fetchMore">
            <el-button v-if="!isOver" size="small" @click="fetchList" plain>More..</el-button>
        </div>
    </div>
</template>
<script>
    import Waterfall from "vue-waterfall-plugin"
    import fileApi from "@/api/file"
    import {dateFormat} from "@/utils/time"

    // emit [select-image]
    export default {
        data () {
            return {
                mon: '',
                list: [],
                selectedItem: {},
                startFetchLimit: 25,
                fetchLimit: 12,
                isOver: false
            }
        },
        created() {
            this.createList()
        },
        methods: {
            async createList() {
                let maxDate = new Date()
                if (this.mon) {
                    maxDate = new Date(this.mon)
                }
                if (maxDate.getMonth() == 11) {
                    maxDate = new Date(maxDate.getFullYear() + 1, 0, 1);
                } else {
                    maxDate = new Date(maxDate.getFullYear(), maxDate.getMonth() + 1, 1);
                }
                this.maxTime = dateFormat(maxDate, 'yyyy-MM-dd hh:mm:ss')
                this.list = []
                this.fetchList()
            },
            async fetchList() {
                if (this.isOver)
                    return
                let limit = this.fetchLimit
                if (this.list.length === 0) {
                    limit = this.startFetchLimit
                }
                const res = await fileApi.imageList({max_time: (this.maxTime), size:limit})
                let list = res.data.data
                if (list.length < this.fetchLimit) {
                    this.isOver = true
                    this.$message('fetch end')
                }
                list.forEach(item => {
                    item.src = process.env.VUE_APP_FILE_ORIGIN + item.key
                    item.time = dateFormat(item.created_time, 'yyyy-MM-dd hh:mm:ss')
                    this.maxTime = item.time
                    this.list.push(item)
                })
            },
            handleClick(selectItem) {
                this.$emit('select-image', selectItem)
            }
        },
        components: {
            Waterfall
        },
        props: {
            size: {
                width: 240
            }
        },
        computed: {
            waterfallItemWidth: function () {
                return this.$props.size ? this.$props.size.width : 240
            }
        }
    }
</script>

<style scoped>
    .selected{
        /*opacity: 0.6;*/
    }
    .selected-icon {
        font-size: 22px;
        color: aqua;
        position: absolute;
        left: 7px;
        bottom: 2px;
        border-radius: 50%;
    }
    .card :hover {
        cursor: pointer;
    }
    .bottom-time {
        text-align: right;
        padding: 4px 1px;
    }
    .time {
        font-size: 13px;
        color: #999;
        font-family: 'Roboto Mono';
    }
    .info {
        padding: 0 10px;
    }
    img {
        width: 100%;
    }
    .fetchMore {
        text-align: center;
        margin-top: 50px;
    }
    .picker {
        margin-left: 15px;
    }
</style>
