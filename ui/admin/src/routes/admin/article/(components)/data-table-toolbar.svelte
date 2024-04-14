<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import Cross2 from 'svelte-radix/Cross2.svelte'
  import type { Writable } from 'svelte/store'
  import { DataTableFacetedFilter } from './index.js'
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
  import type { Article } from '$src/types/stream.js'
  import TableViewOptions from './table-view-options.svelte'
  import { statuses } from '../(data)/data.js'

  export let tableModel: TableViewModel<Article>

  const { pluginStates } = tableModel
  const {
    filterValue
  }: {
    filterValue: Writable<string>
  } = pluginStates.filter

  const {
    filterValues
  }: {
    filterValues: Writable<{
      status: number[]
    }>
  } = pluginStates.colFilter

  $: showReset = Object.values({ ...$filterValues, $filterValue }).some((v) => v.length > 0)
</script>

<div class="flex items-center justify-between">
  <div class="flex flex-1 items-center space-x-2">
    <Input
      placeholder="Filter tasks..."
      class="h-8 w-[150px] lg:w-[250px]"
      type="search"
      bind:value={$filterValue}
    />

    <DataTableFacetedFilter
      bind:filterValues={$filterValues.status}
      title="Status"
      options={statuses}
    />
    {#if showReset}
      <Button
        on:click={() => {
          $filterValue = ''
          $filterValues.status = []
        }}
        variant="ghost"
        class="h-8 px-2 lg:px-3"
      >
        Reset
        <Cross2 class="ml-2 h-4 w-4" />
      </Button>
    {/if}
  </div>

  <TableViewOptions {tableModel} />
</div>
