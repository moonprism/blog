<svelte:options customElement="mod-img-loading" />

<script>
  // to be continued
  export let color = '#0d92f4';
</script>

<div style="background-color: {color};"></div>

<style>
  div {
    width: 30px;
    height: 30px;
    padding: 0;
    margin: 0;
    border-radius: 100%;
    animation: scaleout 1.5s infinite ease-in-out;
  }
  @keyframes scaleout {
    0% {
      -webkit-transform: scale(0);
    }
    100% {
      -webkit-transform: scale(1);
      opacity: 0;
    }
  }
</style>
