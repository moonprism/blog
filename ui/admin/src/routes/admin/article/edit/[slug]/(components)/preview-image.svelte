<script>
  import { getRealSrc } from '@/helpers/fetch'

  export let src
  export let alt
</script>

<img src={`${getRealSrc(src)}`} {alt} />
