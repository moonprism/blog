<script lang="ts">
  import type { Attachment } from '$src/types/stream.js'
  import { openForm } from '../(data)/data'
  import { Keyboard } from 'svelte-radix'

  export let row: Attachment

  function edit() {
    openForm(row)
  }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
<div
  on:click={edit}
  class="absolute left-0 top-0 flex h-full w-full items-end justify-center opacity-0 group-hover:opacity-100"
>
  <Keyboard class="h-5"></Keyboard>
</div>
