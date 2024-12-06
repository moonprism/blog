<svelte:options customElement="mod-img-modal" />

<script>
  export let src = '';

  $: {
    if (src === '') {
      document.body.style.overflow = '';
    } else {
      document.body.style.overflow = 'hidden';
    }
  }

  function close() {
    src = '';
  }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div style="display: {src === '' ? 'none' : 'flex'};" on:click={close}>
  <img {src} alt={src} />
</div>

<style>
  /* 模态框样式 */
  div {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: var(--background);
    justify-content: center;
    align-items: center;
    cursor: zoom-out;
  }

  img {
    max-width: 97%;
    max-height: 97%;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
</style>
