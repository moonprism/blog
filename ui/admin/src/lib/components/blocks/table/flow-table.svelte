<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { Render } from 'svelte-headless-table'
  import type { Filter, ViewOption } from '$src/types/table'
  import TableToolbar from './table-toolbar.svelte'
  import { MasonryGrid } from '@egjs/svelte-grid'
  import Loading from '../animation/loading.svelte'
  import { onMount, onDestroy } from 'svelte'
  import { isScrollAtBottom, throttle } from '@/helpers/system'
  import { isRequestIn } from '@/helpers/fetch'

  export let tableModel: TableViewModel<any, any>
  export let viewOption: ViewOption
  export let filters: Filter[] = []

  const { pageRows } = tableModel

  const gap = 10
  const align = 'stretch'
  let maxStretchColumnSize = 400

  let grid: MasonryGrid

  export function reset(size?: number): void {
    if (size) {
      maxStretchColumnSize = size
    }
    grid.updateItems(grid.getItems())
  }

  const { pageSize, hasNextPage } = tableModel.pluginStates.page
  // 先固定20吧，方便计算从服务器取得真实数据需要的参数
  $pageSize = 20

  let th: Element,
    thParent: Element | null = null

  const handleScroll = throttle(() => {
    if (isScrollAtBottom(thParent)) {
      if (!$isRequestIn && $hasNextPage) {
        $pageSize += 20
      }
    }
  }, 200)

  onMount(() => {
    if (th.parentElement && th.parentElement.getAttribute('role') === 'dialog') {
      thParent = th.parentElement
      thParent.addEventListener('scroll', handleScroll)
    } else {
      window.addEventListener('scroll', handleScroll)
    }
  })
  onDestroy(() => {
    if (thParent !== null) {
      thParent.removeEventListener('scroll', handleScroll)
    } else {
      window.removeEventListener('scroll', handleScroll)
    }
  })
</script>

<div class="space-y-3" bind:this={th}>
  <TableToolbar {tableModel} {filters} {viewOption} />

  <!--button on:click={reset}>reset</button-->
  <MasonryGrid bind:this={grid} class="overflow-hidden" {gap} {align} {maxStretchColumnSize}>
    {#each $pageRows as row (row.id)}
      <div class="group cursor-pointer rounded-md border hover:shadow">
        {#each row.cells as cell (cell.id)}
          <Render of={cell.render()} />
        {/each}
      </div>
    {/each}
  </MasonryGrid>

  <div class="flex justify-center text-base leading-9">
    {#if $isRequestIn}
      <Loading></Loading>
    {/if}
  </div>
</div>
