<script lang="ts">
  import Button from '@/components/ui/button/button.svelte'
  import { Plus } from 'lucide-svelte'

  export let text: string
  export let onClick = () => {}
</script>

<Button variant="ghost" class="my-2 py-0" on:click={onClick}>
  <Plus class="mr-1 h-4 w-4" />
  {text}
</Button>
