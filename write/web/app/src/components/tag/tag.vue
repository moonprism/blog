<template>
    <el-tag
            class="tag"
            :key="tag.id"
            :color="tag.color"
            effect="dark"
    >
        {{tag.name}}
    </el-tag>
</template>

<script>
    export default {
        name: "iTag",
        props: {
            tag: {}
        }
    }
</script>

<style scoped>
    .tag {
        margin: 5px 10px;
        border: none;
    }
    .tag :hover {
        cursor: pointer;
    }
</style>