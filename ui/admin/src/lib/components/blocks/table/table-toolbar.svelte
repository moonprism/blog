<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { writable, type Writable } from 'svelte/store'
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
  import { Undo } from 'lucide-svelte'

  import TableViewOptions from './table-tool-view.svelte'
  import type { Filter, ViewOption } from '$src/types/table.js'
  import TableFilterOption from './table-tool-filter.svelte'
  import { capitalizeFirstLetter } from '@/helpers/string'

  export let tableModel: TableViewModel<any>
  export let viewOption: ViewOption
  export let filters: Filter[]
  export let showHeadSparkles: Writable<boolean> | null = null

  const { pluginStates } = tableModel

  const filterText = pluginStates.filter.filterValue

  let filterValues: Writable<{ [index: string]: any[] }>

  if (pluginStates.colFilter !== undefined) {
    filterValues = pluginStates.colFilter.filterValues
  }

  $: {
    // fix 隐藏列后该列过滤失效
    filters.forEach((e) => {
      if (!(e.name in $filterValues)) {
        $filterValues[e.name] = []
      }
    })
  }

  function onInputFoucus(i: boolean) {
    if (showHeadSparkles !== null) {
      $showHeadSparkles = i
    }
  }

  $: showReset = Object.values({ ...$filterValues, $filterText }).some((v) => v.length > 0)
</script>

<div class="flex items-center justify-between">
  <div class="flex flex-1 items-center space-x-2">
    <Input
      placeholder="What is it you desire."
      class="h-8 w-[150px] lg:w-[250px]"
      type="search"
      bind:value={$filterText}
      on:focus={() => onInputFoucus(true)}
      on:focusout={() => onInputFoucus(false)}
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
          $filterText = ''
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

  <TableViewOptions {tableModel} {viewOption} />
</div>
