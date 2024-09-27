<script lang="ts">
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import Table from './(components)/table.svelte'
  import { openForm, tableData } from './(data)/data'
  import { Award } from 'lucide-svelte'
  import Button from '@/components/ui/button/button.svelte'
  import { goto } from '$app/navigation'
  import * as Tooltip from '$lib/components/ui/tooltip/index.js'
  import { RefreshCcw } from 'lucide-svelte'
  import { renderMD } from './write/[[id]]/(data)/data'
  import { fet } from '@/helpers/fetch'
  import type { ArticleDetail } from '$src/types/stream'
  import toast from '$lib/helpers/toast'
  import { base } from '$app/paths'

  let syncCount = 0
  let successCount = 0
  let isRequest = false

  async function render(id: number) {
    const getRes = await fet.get(`article/${id}`)
    if (!getRes.ok) {
      return
    }
    const article = <ArticleDetail>getRes.data
    const text = article.content.text
    const html = await renderMD(text, true)
    const putRes = await fet.put(`article/${id}`, { text, html })
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
    const ids = $tableData.map((i) => i.id)
    syncCount = ids.length
    //  Promise.all(...)
    ids.forEach(async (id) => {
      await render(id)
    })
    isRequest = true
  }
</script>

<div class="container mx-auto">
  <div class="flex items-center justify-between">
    <div>
      <DataActionButton text="New Article" onClick={() => openForm()} />
      <Tooltip.Root>
        <Tooltip.Trigger asChild let:builder>
          <Button
            builders={[builder]}
            variant="ghost"
            size="icon"
            on:click={() => {
              goto(`${base}/admin/tag`)
            }}
          >
            <Award class="h-5 w-5"></Award>
          </Button>
        </Tooltip.Trigger>
        <Tooltip.Content>标签管理</Tooltip.Content>
      </Tooltip.Root>
    </div>
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
