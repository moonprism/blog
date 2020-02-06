<template>
    <div id="crop-wrapper" v-loading="loading">
        <Cropper ref="cropper" @ready="ready" :src="this.$props.image.url"/>
        <div class="button-wrapper">
            <el-button @click="cropImage" size="small" icon="el-icon-scissors" plain>Crop</el-button>
        </div>
    </div>
</template>

<script>
    import { Cropper } from "vue-advanced-cropper";

    export default {
        name: "Crop",
        data() {
            return {
                loading: true
            };
        },
        methods: {
            cropImage() {
                const result = this.$refs.cropper.getResult();
                this.$props.image.url = result.canvas.toDataURL('image/jpeg')
                if (this.$props.dialog.afterCropClose) {
                    this.$props.dialog.visible = false
                }
            },
            ready() {
                this.loading = false
            }
        },
        props: {
            image: {
                url: ''
            },
            dialog: {
                visible: false
            },
        },
        components: {
            Cropper
        }
    };
</script>

<style>
    .button-wrapper {
        text-align: center;
        padding-top: 20px;
    }
</style>