<script lang="ts">
  import { cn } from '$lib/utils.js'

  export let colors: string[]
  let className: string | undefined | null = undefined
  export { className as class }
</script>

<div class="flex gap-1">
  {#each colors as color}
    <div style="background-color: {color}" class={cn('h-3 w-3 rounded shadow', className)}></div>
  {/each}
</div>
