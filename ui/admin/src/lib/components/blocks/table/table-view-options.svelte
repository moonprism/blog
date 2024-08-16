<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { Button } from '@/components/ui/button'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import { ListTodo } from 'lucide-svelte'
  import type { ViewOption } from '$src/types/table'

  export let tableModel: TableViewModel<any>
  export let viewOption: ViewOption
  const selected = viewOption.selected

  const { pluginStates, flatColumns } = tableModel
  const { hiddenColumnIds } = pluginStates.hide

  function handleHide(id: string) {
    hiddenColumnIds.update((ids: string[]) => {
      if (ids.includes(id)) {
        return ids.filter((i) => i !== id)
      }
      return [...ids, id]
    })
    $selected = $hiddenColumnIds
  }

  if (viewOption.type === 'hideColumn') {
    if (viewOption.options.length !== 0) {
      // todo 自定义过滤字段
    }
    hiddenColumnIds.update(() => $selected)
  }

  const noHidableCols = ['id', 'actions']
</script>

<DropdownMenu.Root>
  <DropdownMenu.Trigger asChild let:builder>
    <Button variant="outline" size="sm" class="ml-auto hidden h-8 lg:flex" builders={[builder]}>
      <ListTodo class="mr-1 h-4 w-4" />
      View
    </Button>
  </DropdownMenu.Trigger>
  <DropdownMenu.Content>
    {#if viewOption.type === 'hideColumn'}
      {#each flatColumns as col}
        {#if !noHidableCols.includes(col.id)}
          <DropdownMenu.CheckboxItem
            checked={!$hiddenColumnIds.includes(col.id)}
            on:click={() => handleHide(col.id)}
          >
            {col.header}
          </DropdownMenu.CheckboxItem>
        {/if}
      {/each}
    {:else}
      {#each viewOption.options as o}
        <DropdownMenu.CheckboxItem
          checked={$selected.includes(o)}
          on:click={() => {
            $selected = [o]
          }}
        >
          {o}
        </DropdownMenu.CheckboxItem>
      {/each}
    {/if}
  </DropdownMenu.Content>
</DropdownMenu.Root>
