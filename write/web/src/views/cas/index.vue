<template>
    <div>跳转中...</div>
</template>

<script>
    export default {
        mounted() {
            this.$nextTick(() => {
                window.location.href = process.env.VUE_APP_READ_ORIGIN+'login?redirect='
                    + encodeURI(this.$route.params.url)+'&'
                    + 'key=' + process.env.VUE_APP_CAS_KEY
            })
        }
    }
</script>
