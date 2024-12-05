<svelte:options customElement="mod-pager" />

<script>
  export let count;
  export let currentPage = 1;
  export let pageSize = 10;

  // 总页数
  const pages = Math.ceil(count / pageSize);
  // 当前页数
  const page = Number(currentPage);
  /**
   * @type {Number[]}
   */
  let pageList = [];
  for (let i = 1; i <= pages; i++) {
    if (page > 3 && i > 1 && i < page - 1) {
      if (i === page - 2) {
        // 如果当前页数与第一页距离较大，插入省略号
        pageList.push(0);
      }
      if (page >= pages - 1 && i === pages - 2) {
        pageList.push(i);
      }
      continue;
    } else if (page < pages - 1 && i > page + 1 && i < pages) {
      if (page <= 2 && i === 3) {
        pageList.push(i);
      }
      if (i === pages - 1) {
        // 如果当前页数与最后一页距离较大，插入省略号
        pageList.push(0);
      }
      continue;
    }
    pageList.push(i);
  }
</script>

{#if pages > 1}
  <div>
    {#each pageList as i}
      {#if i === page}
        <span>{i}</span>
      {:else if i === 0}
        ...
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
