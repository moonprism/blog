<script lang="ts">
  export let color: string
</script>

<div class="flex items-center text-muted-foreground">
  <div class="mr-1 h-4 w-4 rounded-sm shadow" style="background-color:{color}"></div>
  {color}
</div>
