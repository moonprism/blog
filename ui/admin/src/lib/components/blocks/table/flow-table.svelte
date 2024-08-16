<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { Render } from 'svelte-headless-table'
  import type { filter } from '$src/types/table'
  import TableToolbar from './table-toolbar.svelte'
  import Button from '@/components/ui/button/button.svelte'
  import { MasonryGrid } from '@egjs/svelte-grid'

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

  <div class="text-center">
    <Button
      on:click={() => {
        size += 20
        tableModel.pluginStates.page.pageSize.set(size)
      }}>loading</Button
    >
  </div>
</div>
