<script lang="ts">
  import type { Attachment } from '$src/types/stream.js'
  import { Keyboard } from 'svelte-radix'
  import { openForm } from '../(data)/data'

  export let row: Attachment
  export let click: ((row:Attachment) => void) | null = null

  function cl() {
    if (click === null) {
      openForm(row)
    } else {
      click(row)
    }
  }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
<div
  on:click={cl}
  class="absolute left-0 top-0 flex h-full w-full items-end justify-center opacity-0 group-hover:opacity-100"
>
  {#if click === null}
    <Keyboard class="h-5"></Keyboard>
  {:else}
    <div class="mb-1 flex h-3 w-3 items-center justify-center rounded-full bg-foreground shadow">
      <div class="relative top-[1px] font-mono text-[.6rem] text-background font-bold">A</div>
    </div>
  {/if}
</div>
