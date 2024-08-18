<script lang="ts">
  import * as Dialog from '@/components/ui/dialog/index.js'

  import Table from '../../attachment/(components)/table.svelte'
  import type { Attachment } from '$src/types/stream'
  import { type Writable } from 'svelte/store'

  export let open:Writable<boolean>
  export let callback:typeof itemClick

  function itemClick(row:Attachment) {
    callback(row)
    $open = false
  }
</script>

<Dialog.Root bind:open={$open}>
  <Dialog.Content class="max-w-screen z-[51] h-screen cursor-pointer overflow-auto md:w-screen pt-9">
    <Table {itemClick} />
  </Dialog.Content>
</Dialog.Root>
