<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { Render } from 'svelte-headless-table'
  import type { Filter, ViewOption } from '$src/types/table'
  import TableToolbar from './table-toolbar.svelte'
  import { MasonryGrid } from '@egjs/svelte-grid'
  import Loading from '../animate/loading.svelte'
  import { onMount, onDestroy } from 'svelte'
  import { scrollAtBottom } from '@/helpers/system'

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
  $pageSize = 20

  let isLoading = false
  function handleScroll() {
    if (scrollAtBottom()) {
      isLoading = true
      setTimeout(() => {
        $pageSize += 10
        // todo 平滑滚动& 防抖
        console.log($pageSize)
      }, 1000)
      if (!$hasNextPage) {
        isLoading = false
      }
    }
  }

  onMount(() => {
    window.addEventListener('scroll', handleScroll)
  })
  onDestroy(() => {
    window.removeEventListener('scroll', handleScroll)
  })
</script>

<div class="space-y-3">
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
    {#if isLoading}
      <Loading></Loading>
    {/if}
  </div>
</div>
