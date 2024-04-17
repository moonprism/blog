<script lang="ts">
  import { cn } from '$lib/utils.js'
  import { Button } from '@/components/ui/button'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import Separator from '@/components/ui/separator/separator.svelte'
  import { ChevronsUpDown, ArrowDown, ArrowUp, SeparatorHorizontal, Eraser } from 'lucide-svelte'

  let className: string | undefined | null = undefined
  export { className as class }
  export let sort: {
    order: 'asc' | 'desc' | undefined
    toggle: (_: Event) => void
    clear: () => void
    disabled: boolean
  }

  // 要想不出bug还得在初始化时指定
  // sort: addSortBy({
  //   toggleOrder: ["asc", "desc"],
  // ),
  function handleAscSort(e: Event) {
    if (sort.order === 'asc') {
      return
    }
    sort.toggle(e)
  }

  function handleDescSort(e: Event) {
    if (sort.order === 'desc') {
      return
    }
    if (sort.order === undefined) {
      sort.toggle(e)
    }
    sort.toggle(e)
  }

  function handleNopSort(e: Event) {
    sort.clear()
  }
</script>

{#if !sort.disabled}
  <div class={cn('flex items-center', className)}>
    <DropdownMenu.Root>
      <DropdownMenu.Trigger asChild let:builder>
        <Button
          variant="ghost"
          builders={[builder]}
          class="-ml-3 h-8 data-[state=open]:bg-accent"
          size="sm"
        >
          <slot />
          {#if sort.order === 'desc'}
            <ArrowDown class="ml-1 h-3 w-3 text-gray-950" />
          {:else if sort.order === 'asc'}
            <ArrowUp class="ml-1 h-3 w-3 text-gray-950" />
          {:else}
            <ChevronsUpDown class="ml-1 h-3 w-3 text-muted-foreground/80" />
          {/if}
        </Button>
      </DropdownMenu.Trigger>
      <DropdownMenu.Content align="start">
        <DropdownMenu.Item on:click={handleAscSort}>
          <ArrowUp class="mr-2 h-3.5 w-3.5 text-muted-foreground/70" />
          Asc
        </DropdownMenu.Item>
        <DropdownMenu.Item on:click={handleDescSort}>
          <ArrowDown class="mr-2 h-3.5 w-3.5 text-muted-foreground/70" />
          Desc
        </DropdownMenu.Item>
        {#if sort.order !== undefined}
          <Separator />
          <DropdownMenu.Item on:click={handleNopSort}>
            <Eraser class="mr-2 h-4 w-4 text-muted-foreground/70" />
            Clean
          </DropdownMenu.Item>
        {/if}
      </DropdownMenu.Content>
    </DropdownMenu.Root>
  </div>
{:else}
  <slot />
{/if}
