<script>
  import { isRequestIn } from '@/helpers/fetch'
  import { BookUp, LoaderCircle } from 'lucide-svelte'
</script>

{#if $isRequestIn}
  <LoaderCircle class="h-4 w-4 animate-spin" />
{:else}
  <BookUp class="h-5 w-5"></BookUp>
{/if}
