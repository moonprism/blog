<svelte:options customElement="mod-gist" />

<script>
  let queryText = "";

  async function search() {
    const response = await fetch("/gists/search?q=" + queryText, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const data = await response.json();
    console.log(data);
  }
</script>

<input bind:value={queryText} /><button on:click={search}>搜索</button>
