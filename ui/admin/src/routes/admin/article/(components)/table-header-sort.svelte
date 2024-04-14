<script lang="ts">
  import { cn } from '$lib/utils.js'
  import { Button } from '@/components/ui/button'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import Separator from '@/components/ui/separator/separator.svelte'
  import { ChevronsUpDown, ArrowDown, ArrowUp, Eraser } from 'lucide-svelte'

  let className: string | undefined | null = undefined
  export { className as class }
  export let props: {
    select: never
    sort: {
      order: 'desc' | 'asc' | undefined
      toggle: (_: Event) => void
      clear: () => void
      disabled: boolean
    }
    filter: never
  }

  function handleAscSort(e: Event) {
    if (props.sort.order === 'asc') {
      return
    }
    props.sort.toggle(e)
  }

  function handleDescSort(e: Event) {
    if (props.sort.order === 'desc') {
      return
    }
    if (props.sort.order === undefined) {
      // We can only toggle, so we toggle from undefined to 'asc' first
      props.sort.toggle(e)
    }
    props.sort.toggle(e) // Then we toggle from 'asc' to 'desc'
  }

  function handleNopSort(e: Event) {
    props.sort.clear()
  }
</script>

{#if !props.sort.disabled}
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
          {#if props.sort.order === 'desc'}
            <ArrowDown class="ml-1 h-3 w-3" />
          {:else if props.sort.order === 'asc'}
            <ArrowUp class="ml-1 h-3 w-3" />
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
        {#if props.sort.order !== undefined}
          <Separator />
          <DropdownMenu.Item on:click={handleNopSort}>
            <Eraser class="mr-2 h-3 w-3 text-muted-foreground/70" />
            Clean
          </DropdownMenu.Item>
        {/if}
      </DropdownMenu.Content>
    </DropdownMenu.Root>
  </div>
{:else}
  <slot />
{/if}
