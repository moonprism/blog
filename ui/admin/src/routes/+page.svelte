<script lang="ts">
  import { goto } from '$app/navigation'
  import { getJwtInfo } from '@/helpers/jwt'

  if (getJwtInfo() != null) {
    goto('admin')
  } else {
    goto('login')
  }
</script>
