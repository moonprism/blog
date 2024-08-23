<script lang="ts">
  import DotsHorizontal from 'svelte-radix/DotsHorizontal.svelte'
  import type { Article } from '$src/types/stream.js'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import { Button } from '@/components/ui/button'
  import { openForm, tableData } from '../(data)/data'
  import { fet } from '@/helpers/fetch'
  import { alertDialog } from '@/components/blocks/dialog/alert'
  import { goto } from '$app/navigation'
  import { page } from '$app/stores'

  export let row: Article

  function edit() {
    openForm(row)
  }

  function del() {
    fet.delete(`article/${row.id}`).then((res) => {
      if (res.ok) {
        $tableData = $tableData.filter((v) => v.id !== row.id)
      }
    })
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
    <DropdownMenu.Item
      on:click={() => {
        goto(`${$page.url.pathname}/write/${row.id}`)
      }}>编辑内容</DropdownMenu.Item
    >
    <DropdownMenu.Item>预览</DropdownMenu.Item>
    <DropdownMenu.Item
      on:click={() => {
        alertDialog(`删除文章：${row.title}`).then(() => {
          del()
        })
      }}
    >
      删除
      <DropdownMenu.Shortcut>⌘⌫</DropdownMenu.Shortcut>
    </DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>
