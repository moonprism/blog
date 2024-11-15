<script lang="ts">
  import { cn } from '$lib/utils.js'
  import { getRealSrc } from '$src/routes/admin/(data)/data'

  export let src: string
  export let alt = ''

  let className: string = ''
  export { className as class }
</script>

<img class={cn('max-h-[100px]', className)} src={getRealSrc(src)} {alt} />
