<script lang="ts">
  import MixerHorizontal from 'svelte-radix/MixerHorizontal.svelte'
  import type { TableViewModel } from 'svelte-headless-table'
  import { Button } from '@/components/ui/button'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import type { Article } from '$src/types/stream'

  export let tableModel: TableViewModel<Article>
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

  const hidableCols = ['id', 'image', 'summary', 'created', 'updated']
</script>

<DropdownMenu.Root>
  <DropdownMenu.Trigger asChild let:builder>
    <Button variant="outline" size="sm" class="ml-auto hidden h-8 lg:flex" builders={[builder]}>
      <MixerHorizontal class="mr-2 h-4 w-4" />
      View
    </Button>
  </DropdownMenu.Trigger>
  <DropdownMenu.Content>
    <DropdownMenu.Label>Toggle columns</DropdownMenu.Label>
    <DropdownMenu.Separator />
    {#each flatColumns as col}
      {#if hidableCols.includes(col.id)}
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
