<script lang="ts">
  import Table from './(components)/table.svelte'
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import { openForm, tableData } from './(data)/data'
  import { RefreshCcw } from 'lucide-svelte'
  import { renderMD } from '../article/write/[slug]/(data)/data'
  import { fet } from '@/helpers/fetch'
  import toast from '$lib/helpers/toast'
  import Button from '@/components/ui/button/button.svelte'

  let syncCount = 0
  let successCount = 0
  let isRequest = false

  async function render(id: number, text: string) {
    const html = await renderMD(text)
    const putRes = await fet.put(`gist/${id}`, { html })
    if (!putRes.ok) {
      return
    }
    successCount++
    if (successCount === syncCount) {
      successCount = 0
      isRequest = false
      toast.success(`同步完成 ${syncCount} 条`)
    }
  }

  function sync() {
    // 同步渲染当前展示条目的Markdown
    if (isRequest) {
      return
    }
    const gists = $tableData
    syncCount = gists.length
    //  Promise.all(...)
    gists.forEach(async (gist) => {
      let content = gist.content
      if (gist.lang !== 'md') {
        content = `${'```'}${gist.lang}\n${gist.content}\n${'```'}`
      }
      await render(gist.id, content)
    })
    isRequest = true
  }
</script>

<div class="container mx-auto">
  <div class="flex items-center justify-between">
    <DataActionButton text="New Gist" onClick={() => openForm()} />
    <div class="flex items-center">
      {#if isRequest}
        <div class="text-sm">{successCount}/{syncCount}</div>
      {/if}
      <Button variant="link" size="icon" on:click={sync}>
        <RefreshCcw class="h-4 w-4 {isRequest ? 'animate-spin' : ''}" />
      </Button>
    </div>
  </div>
  <Table />
</div>
