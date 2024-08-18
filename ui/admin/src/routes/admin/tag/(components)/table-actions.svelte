<script lang="ts">
  import type { Tag } from '$src/types/stream.js'
  import { openForm, tableData } from '../(data)/data'
  import { fet } from '@/helpers/fetch'
  import * as AlertDialog from '@/components/ui/alert-dialog/index.js'
  import Badge from '@/components/ui/badge/badge.svelte'
  import { alertDialog } from '@/components/blocks/dialog/alert'

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

</script>

<div class="flex h-8 items-center justify-center gap-1">
  <button on:click={edit} class="rounded-sm border border-black bg-white px-2 text-xs text-black"
    >Edit</button
  >
  <button
    on:click={() => {
      alertDialog(`删除标签：${row.name}`).then(() => {
        del()
      })
    }}
    class="rounded-sm border border-black bg-black px-2 text-xs text-white">Delete</button
  >
</div>
