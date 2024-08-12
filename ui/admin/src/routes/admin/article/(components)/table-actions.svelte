<script lang="ts">
  import DotsHorizontal from 'svelte-radix/DotsHorizontal.svelte'
  import type { Article } from '$src/types/stream.js'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import { Button } from '@/components/ui/button'
  import { openForm } from '../(data)/data'

  export let row: Article
  
  function edit() {
    openForm(row)
  }
</script>

<DropdownMenu.Root>
  <DropdownMenu.Trigger asChild let:builder>
    <Button
      variant="ghost"
      builders={[builder]}
      class="flex h-8 w-8 p-0 data-[state=open]:bg-muted"
    >
      <DotsHorizontal class="h-4 w-4" />
      <span class="sr-only">Open Menu</span>
    </Button>
  </DropdownMenu.Trigger>
  <DropdownMenu.Content class="w-[160px]" align="end">
    <DropdownMenu.Item on:click={edit}>更改信息</DropdownMenu.Item>
    <DropdownMenu.Item>编辑内容</DropdownMenu.Item>
    <DropdownMenu.Item>预览</DropdownMenu.Item>
    <DropdownMenu.Item>
      删除
      <DropdownMenu.Shortcut>⌘⌫</DropdownMenu.Shortcut>
    </DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>
