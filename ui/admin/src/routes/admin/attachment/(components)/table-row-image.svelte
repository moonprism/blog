<script lang="ts">
  import { loadingImageCount } from '../(data)/data'
  import { getRealSrc } from '../../(data)/data'

  export let src: string
  export let alt = ''

  $loadingImageCount++
</script>

<img
  class="w-full"
  src={getRealSrc(src)}
  {alt}
  on:load={() => {
    $loadingImageCount--
  }}
  on:error={() => {
    $loadingImageCount--
  }}
/>
