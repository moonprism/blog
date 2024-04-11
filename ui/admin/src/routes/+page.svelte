<script lang="ts">
  import { goto } from "$app/navigation"
  import { getJwtInfo } from "@/helpers/fetch"

  if (getJwtInfo() != null) {
    goto('admin')
  } else {
    goto('login')
  }
</script>