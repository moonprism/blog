<script lang="ts">
  import * as Dialog from '@/components/ui/dialog/index.js'

  import Table from '$src/routes/admin/attachment/(components)/table.svelte'
  import type { Attachment } from '$src/types/stream'
  import { type Writable } from 'svelte/store'
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import { openForm } from '$src/routes/admin/attachment/(data)/data'

  export let open: Writable<boolean>
  export let callback: typeof itemClick

  function itemClick(row: Attachment) {
    callback(row)
    $open = false
  }

  function add() {
    openForm()
  }
</script>

<Dialog.Root bind:open={$open}>
  <Dialog.Content
    class="max-w-screen z-[51] h-screen cursor-pointer overflow-auto pt-11 md:w-screen"
  >
    <DataActionButton class="absolute left-[20px] m-1 px-3" text="Upload" onClick={add} />
    <Table {itemClick} />
  </Dialog.Content>
</Dialog.Root>
