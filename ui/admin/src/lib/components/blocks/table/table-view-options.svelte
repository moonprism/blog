<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { Button } from '@/components/ui/button'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import { ListTodo } from 'lucide-svelte'

  export let tableModel: TableViewModel<any>
  const { pluginStates, flatColumns } = tableModel
  const { hiddenColumnIds } = pluginStates.hide

  function handleHide(id: string) {
    hiddenColumnIds.update((ids: string[]) => {
      if (ids.includes(id)) {
        return ids.filter((i) => i !== id)
      }
      return [...ids, id]
    })
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
  </DropdownMenu.Content>
</DropdownMenu.Root>
