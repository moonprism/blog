<script lang="ts">
  import { Plus } from 'lucide-svelte'
  import Check from 'svelte-radix/Check.svelte'
  import * as Command from '@/components/ui/command'
  import * as Popover from '@/components/ui/popover'
  import { Button } from '@/components/ui/button'
  import { cn } from '$lib/utils.js'
  import { Separator } from '@/components/ui/separator'
  import { Badge } from '@/components/ui/badge'
  import { type option } from '$src/types/table'

  export let filterValues: number[] = []
  export let title: string
  export let options: option[]

  let open = false

  function handleSelect(currentValue: number) {
    if (Array.isArray(filterValues) && filterValues.includes(currentValue)) {
      filterValues = filterValues.filter((v) => v !== currentValue)
    } else {
      filterValues = [...(Array.isArray(filterValues) ? filterValues : []), currentValue]
    }
  }
</script>

<Popover.Root bind:open>
  <Popover.Trigger asChild let:builder>
    <Button builders={[builder]} variant="outline" size="sm" class="h-8 border-dashed">
      <Plus class="mr-1 h-4 w-4" />
      {title}
      {#if filterValues.length > 0}
        <Separator orientation="vertical" class="mx-2 h-4" />
        <Badge variant="secondary" class="rounded-sm px-1 font-normal lg:hidden">
          {filterValues.length}
        </Badge>
        <div class="hidden space-x-1 lg:flex">
          {#if filterValues.length > 2}
            <Badge variant="secondary" class="rounded-sm px-1 font-normal">
              {filterValues.length} Selected
            </Badge>
          {:else}
            {#each options as option}
              {#if filterValues.includes(option.id)}
                <Badge variant="secondary" class="rounded-sm px-1 font-normal">
                  {option.label}
                </Badge>
              {/if}
            {/each}
          {/if}
        </div>
      {/if}
    </Button>
  </Popover.Trigger>
  <Popover.Content class="w-[200px] p-0" align="start" side="bottom">
    <Command.Root>
      <Command.Input placeholder={title} />
      <Command.List>
        <Command.Empty>No results found.</Command.Empty>
        <Command.Group>
          {#each options as option}
            {@const Icon = option.icon}
            <Command.Item
              value={option.label}
              onSelect={() => {
                handleSelect(option.id)
              }}
            >
              <div
                class={cn(
                  'mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary',
                  filterValues.includes(option.id)
                    ? 'bg-primary text-primary-foreground'
                    : 'opacity-50 [&_svg]:invisible'
                )}
              >
                <Check className={cn('h-4 w-4')} />
              </div>
              <Icon class="mr-2 h-4 w-4 text-muted-foreground" />
              <span>
                {option.label}
              </span>
            </Command.Item>
          {/each}
        </Command.Group>
        {#if filterValues.length > 0}
          <Command.Separator />
          <Command.Item
            class="justify-center text-center"
            onSelect={() => {
              filterValues = []
            }}
          >
            Clear filters
          </Command.Item>
        {/if}
      </Command.List>
    </Command.Root>
  </Popover.Content>
</Popover.Root>
