<template>
    <div id="tag-select-wrapper">
        <i-tag class="tag" v-for="tag in availableTags" v-bind:key="tag.id" :tag="tag" v-on:click.native="handleClick(tag)"></i-tag>
    </div>
</template>

<script>
    import iTag from "./tag"
    import tagApi from "@/api/tag"

    export default {
        components: {
            iTag
        },
        name: "tagSelect",
        data() {
            return {
                tags: []
            }
        },
        props: {
            selectTags: {}
        },
        computed: {
            availableTags() {
                let selectTagsMap = {}
                this.$props.selectTags.forEach((element) => {
                    selectTagsMap[element.id] = true
                })
                return this.tags.filter(function (element) {
                    return !selectTagsMap[element.id]
                })
            }
        },
        methods: {
            async fetchAndSetList() {
                const res = await tagApi.list()
                this.tags = res.data
            },
            handleClick(tag) {
                this.$emit('click-tag', tag)
            }
        },
        mounted() {
            this.fetchAndSetList()
        }
    }
</script>

<style scoped>
    .tag {
        cursor: pointer;
    }
</style>
