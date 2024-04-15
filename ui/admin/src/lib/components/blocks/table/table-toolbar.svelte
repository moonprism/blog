<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import type { Writable } from 'svelte/store'
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
  import { BookmarkX, Delete, ListRestart, Undo, Undo2 } from 'lucide-svelte';
  

  import type { Article } from '$src/types/stream.js'
  import TableViewOptions from './table-view-options.svelte'
  import type { filter } from '$src/types/table.js'
  import TableFilterOption from './table-filter-option.svelte'
  import { capitalizeFirstLetter } from '@/helpers/string'

  export let tableModel: TableViewModel<Article>

  export let filters: filter[]

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
      //status: number[]
      //id: number[]
      [index: string]: number[]
    }>
  } = pluginStates.colFilter

  $: showReset = Object.values({ ...$filterValues, $filterValue }).some((v) => v.length > 0)
</script>

<div class="flex items-center justify-between">
  <div class="flex flex-1 items-center space-x-2">
    <Input
      placeholder="What is it you desire."
      class="h-8 w-[150px] lg:w-[250px]"
      type="search"
      bind:value={$filterValue}
    />

    {#each filters as filter}
      <TableFilterOption
        bind:filterValues={$filterValues[filter.name]}
        title={capitalizeFirstLetter(filter.name)}
        options={filter.options}
      />
    {/each}
    {#if showReset}
      <Button
        on:click={() => {
          $filterValue = ''
          filters.forEach((filter) => ($filterValues[filter.name] = []))
        }}
        variant="ghost"
        class="h-8 px-2 lg:px-3"
      >
        Reset
        <Undo class="h-4 w-4" />
      </Button>
    {/if}
  </div>

  <TableViewOptions {tableModel} />
</div>
