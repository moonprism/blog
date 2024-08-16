<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { Render } from 'svelte-headless-table'
  import type { filter } from '$src/types/table'
  import TableToolbar from './table-toolbar.svelte'
  import { MasonryGrid } from '@egjs/svelte-grid'
  import Loading from '../animate/loading.svelte'
  import { onMount, onDestroy } from 'svelte'

  export let tableModel: TableViewModel<any, any>
  export let filters: filter[] = []

  const { pageRows } = tableModel

  let size = 10

  const gap = 10
  const align = 'stretch'
  const maxStretchColumnSize = 400

  let grid: MasonryGrid

  export function reset(): void {
    grid.updateItems(grid.getItems())
  }
  
  let isLoading = false

  let atBottom = false
  function handleScroll(): void {
    const scrollTop = window.scrollY
    const scrollHeight = document.documentElement.scrollHeight
    const clientHeight = window.innerHeight
    atBottom = scrollTop + clientHeight >= scrollHeight - 5 // 5px tolerance
    console.log(atBottom)
    if (atBottom) {
      isLoading = true
      size += 20
      tableModel.pluginStates.page.pageSize.set(size)
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
  <TableToolbar {tableModel} {filters} />

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
