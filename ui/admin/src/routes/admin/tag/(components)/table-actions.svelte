<script lang="ts">
  import DotsHorizontal from 'svelte-radix/DotsHorizontal.svelte'
  import type { Tag } from '$src/types/stream.js'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import { Button } from '@/components/ui/button'
  import { openForm, tableData } from '../(data)/data'
  import { fet } from '@/helpers/fetch'
  import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js'
  import Badge from '@/components/ui/badge/badge.svelte'

  export let row: Tag

  function edit() {
    openForm(row)
  }

  function del() {
    fet.delete(`tag/${row.id}`).then((res) => {
      if (res.ok) {
        $tableData = $tableData.filter((v) => v.id !== row.id)
      }
    })
  }

  let isDel = false
</script>

<AlertDialog.Root bind:open={isDel}>
  <AlertDialog.Content>
    <AlertDialog.Header>
      <AlertDialog.Title>Are you absolutely sure?</AlertDialog.Title>
      <AlertDialog.Description>
        删除标签: <Badge>{row.name}</Badge>
      </AlertDialog.Description>
    </AlertDialog.Header>
    <AlertDialog.Footer>
      <AlertDialog.Cancel>再想想</AlertDialog.Cancel>
      <AlertDialog.Action on:click={del}>确认</AlertDialog.Action>
    </AlertDialog.Footer>
  </AlertDialog.Content>
</AlertDialog.Root>

<DropdownMenu.Root>
  <DropdownMenu.Trigger asChild let:builder>
    <Button
      variant="ghost"
      builders={[builder]}
      class="flex h-8 w-8 p-0 data-[state=open]:bg-muted"
    >
      <DotsHorizontal class="h-4 w-4" />
    </Button>
  </DropdownMenu.Trigger>
  <DropdownMenu.Content class="w-[100px]" align="end">
    <DropdownMenu.Item on:click={edit}>编辑</DropdownMenu.Item>
    <DropdownMenu.Item on:click={() => (isDel = true)}>
      删除
      <DropdownMenu.Shortcut>⌘⌫</DropdownMenu.Shortcut>
    </DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>
