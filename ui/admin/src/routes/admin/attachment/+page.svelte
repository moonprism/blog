<script lang="ts">
  import Table from './(components)/table.svelte'
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import { initTableData, loadingImageCount, openForm } from './(data)/data'
  import { appInfo } from '../(data)/data'
  import Button from '@/components/ui/button/button.svelte'
  import { RefreshCcw } from 'lucide-svelte'
  import { fet, isRequestIn } from '@/helpers/fetch'
  import toast from '$lib/helpers/toast'
  import type { SearchParams } from '$src/types/table'
  import { onMount } from 'svelte'

  async function syncFromGithubApi() {
    await fet.post('attachment/sync', {})
    toast.success(`同步完成`)
    const searchParams = {
      page_size: 20,
      filter_text: '',
      filter_values: <{ [index: string]: number[] }>{}
    }
    initTableData(<SearchParams>searchParams, false)
  }

  onMount(() => {
    $loadingImageCount = 0
  })
</script>

<div class="container mx-auto">
  {#if !$appInfo.isGithubImages}
    <DataActionButton text="Upload" onClick={() => openForm()} />
  {:else}
    <Button variant="ghost" class="my-2 py-0" on:click={syncFromGithubApi}>
      <RefreshCcw class="mr-2 h-4 w-4 {$isRequestIn ? 'animate-spin' : ''}" /> Sync from github
    </Button>
  {/if}
  <Table />
</div>
