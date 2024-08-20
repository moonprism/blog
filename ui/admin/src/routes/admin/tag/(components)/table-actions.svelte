<script lang="ts">
  import type { Tag } from '$src/types/stream.js'
  import { openForm, tableData } from '../(data)/data'
  import { fet } from '@/helpers/fetch'
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
  <button on:click={edit} class="btn">Edit</button>
  <button
    on:click={() => {
      alertDialog(`删除标签：${row.name}`).then(() => {
        del()
      })
    }}
    class="btn bg-foreground text-background">Delete</button
  >
</div>

<style>
  .btn {
    @apply min-w-14 rounded-sm border border-foreground text-xs hover:font-bold hover:shadow;
  }
</style>
