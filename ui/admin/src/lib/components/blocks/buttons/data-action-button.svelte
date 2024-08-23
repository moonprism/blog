<script lang="ts">
  import Button from '@/components/ui/button/button.svelte'
  import { Plus } from 'lucide-svelte'
  import { cn } from '$lib/utils.js'

  let className: string = ''
  export { className as class }

  export let text: string
  export let onClick = () => {}
</script>

<Button variant="ghost" class={cn('my-2 py-0', className)} on:click={onClick}>
  <Plus class="mr-1 h-4 w-4" />
  {text}
</Button>
