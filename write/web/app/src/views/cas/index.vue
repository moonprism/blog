<template>
    <div>跳转中...</div>
</template>

<script>
    import casApi from "@/api/cas"

    export default {
        async mounted() {
            const res = await casApi.key()
            const cas_key = res.data.cas_key
            this.$nextTick(() => {
                window.location.href = process.env.VUE_APP_READ_ORIGIN+'cas/auth?redirect='
                    + btoa(encodeURIComponent(this.$route.params.url))+'&'
                    + 'key=' + cas_key
            })
        }
    }
</script>
