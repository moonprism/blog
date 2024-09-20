<script lang="ts">
  import { goto } from '$app/navigation'
  import Button from '@/components/ui/button/button.svelte'
  import { removeJwt } from '@/helpers/jwt'
</script>

<div class="container mx-auto mt-4">
  <div class="my-2">我从来没有觉得写代码开心过</div>
  <div class="text-sm text-secondary-foreground">
    <!--todo-->
    <p>上次登陆时间：2024 年 9 月 17 日</p>
    <p>Token 时效：72小时(自动刷新)</p>
    <Button size="sm" class="mt-1 h-7 rounded-sm" on:click={() => {
      removeJwt()
      goto('/login')
    }}>退出登陆</Button>
  </div>
</div>
