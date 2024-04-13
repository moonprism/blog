<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import Cross2 from 'svelte-radix/Cross2.svelte'
  import type { Writable } from 'svelte/store'
  import { statuses } from '$src/types/table.js'
  import { DataTableFacetedFilter } from './index.js'
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
  import type { Article } from '$src/types/stream.js'

  export let tableModel: TableViewModel<Article>
  export let data: Article[]

  const counts = data.reduce<{
    status: { [index: string]: number }
  }>(
    (acc, { status }) => {
      acc.status[status] = (acc.status[status] || 0) + 1
      return acc
    },
    {
      status: {}
    }
  )

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
      status: string[]
    }>
  } = pluginStates.colFilter

  $: showReset = Object.values({ ...$filterValues, $filterValue }).some((v) => v.length > 0)
</script>

<pre>$filterValues = {JSON.stringify($filterValues, null, 2)}</pre>

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
      counts={counts.status}
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
</div>
