<svelte:options customElement="mod-pager" />

<script>
  export let count;
  export let currentPage = 1;

  // 暂时写死分页大小吧
  const pages = Math.ceil(count / 10);
</script>

{#if pages > 1}
  <div>
    {#each Array(pages)
      .fill(0)
      .map((_, i) => i + 1) as i}
      {#if i === Number(currentPage)}
        <span>{i}</span>
      {:else}
        <a href="?page={i}">{i}</a>
      {/if}
    {/each}
  </div>
{/if}

<style>
  div {
    margin-bottom: 1.5rem;
  }
  span {
    margin: 0 5px;
    font-weight: 700;
    color: var(--primary);
  }
  a {
    padding: 0 5px;
    color: var(--foreground);
    border-radius: 50%;
    transition: color 0.5s;
    text-decoration: none;
  }
  a:hover {
    text-decoration: underline;
  }
</style>
