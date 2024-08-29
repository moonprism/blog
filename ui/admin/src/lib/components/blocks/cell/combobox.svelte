<script lang="ts">
  import Check from 'svelte-radix/Check.svelte'
  import CaretSort from 'svelte-radix/CaretSort.svelte'
  import { tick } from 'svelte'
  import * as Command from '$lib/components/ui/command/index.js'
  import * as Popover from '$lib/components/ui/popover/index.js'
  import { Button } from '$lib/components/ui/button/index.js'
  import { cn } from '$lib/utils.js'

  export let items: { value: string; label: string }[]
  export let value: string

  let open = false

  $: selectedValue = items.find((f) => f.value === value)?.label ?? 'Select ...'

  function closeAndFocusTrigger(triggerId: string) {
    open = false
    tick().then(() => {
      document.getElementById(triggerId)?.focus()
    })
  }
</script>

<Popover.Root bind:open let:ids>
  <Popover.Trigger asChild let:builder>
    <Button
      builders={[builder]}
      variant="outline"
      role="combobox"
      aria-expanded={open}
      class="w-[200px] justify-between"
    >
      {selectedValue}
      <CaretSort class="ml-2 h-4 w-4 shrink-0 opacity-50" />
    </Button>
  </Popover.Trigger>
  <Popover.Content class="w-[200px] p-0">
    <Command.Root>
      <Command.Input placeholder="Search ..." class="h-9" />
      <Command.Empty>No item found.</Command.Empty>
      <Command.Group>
        {#each items as framework}
          <Command.Item
            class="cursor-pointer"
            value={framework.value}
            onSelect={(currentValue) => {
              value = currentValue
              closeAndFocusTrigger(ids.trigger)
            }}
          >
            <Check class={cn('mr-2 h-4 w-4', value !== framework.value && 'text-transparent')} />
            {framework.label}
          </Command.Item>
        {/each}
      </Command.Group>
    </Command.Root>
  </Popover.Content>
</Popover.Root>
