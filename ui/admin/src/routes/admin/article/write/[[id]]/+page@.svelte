<script lang="ts">
  import Editor from './(components)/editor.svelte'
  import { appInfo, initAppInfo } from '$src/routes/admin/(data)/data'

  initAppInfo()
</script>

{#if $appInfo.attachmentCDN}
  <Editor />
{/if}
