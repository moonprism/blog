<script lang="ts">
  import { onMount } from 'svelte'
  import { loadingImageCount } from '../(data)/data'
  import { getRealSrc } from '../../(data)/data'

  export let src: string
  export let alt = ''

  onMount(() => {
    $loadingImageCount++
    return () => {
      if (isDel === false) {
        $loadingImageCount--
      }
    }
  })

  let isDel = false
  function destroy() {
    isDel = true
    $loadingImageCount--
  }
</script>

<img class="w-full" src={getRealSrc(src)} {alt} on:load={destroy} on:error={destroy} />
