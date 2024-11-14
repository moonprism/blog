<script lang="ts">
  import { getRealSrc } from "../../(data)/data"

  export let src: string
  export let alt = ''
</script>

<img class="w-full" src={getRealSrc(src)} {alt} />
