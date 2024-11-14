<script lang="ts">
  import { getRealSrc } from "../../(data)/data"

  export let src: string
  export let alt = ''
</script>

<img class="max-h-[100px]" src={getRealSrc(src)} {alt} />
