<script lang="ts">
  import type { Tag } from '$src/types/stream.js'
  import { openForm, tableData } from '../(data)/data'
  import { fet } from '@/helpers/fetch'
  import * as AlertDialog from '@/components/ui/alert-dialog/index.js'
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

<div class="flex h-8 items-center justify-center gap-1">
  <button on:click={edit} class="rounded-sm border border-black bg-white px-2 text-xs text-black"
    >Edit</button
  >
  <button
    on:click={() => (isDel = true)}
    class="rounded-sm border border-black bg-black px-2 text-xs text-white">Delete</button
  >
</div>
