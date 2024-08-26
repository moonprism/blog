<script>
  import Badge from '@/components/ui/badge/badge.svelte'

  export let text = ''
</script>

<div class="min-w-12">
  <Badge variant="secondary" class="mx-1">{text}</Badge>
</div>
