<svelte:options customElement="mod-box" />

<script>
  export let height = '';
</script>

<div style="height: {height}px"></div>

<style>
@media (max-width: 768px) {
  div {
    display: none;
  }
}
</style>
