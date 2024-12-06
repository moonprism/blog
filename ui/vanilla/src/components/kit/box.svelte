<svelte:options customElement="mod-box" />

<script>
  export let h = '';
</script>

<div style="height: {h}px"></div>

<style>
@media (max-width: 768px) {
  div {
    display: none;
  }
}
</style>
