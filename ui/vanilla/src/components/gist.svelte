<svelte:options customElement={{ tag: "mod-gist", shadow: "none" }} />

<script>
  import { debounce } from "../utils";
  import Loading from "./loading.svelte";

  let isLoading = false;

  let page = 1;
  let queryText = "";

  /**
   * @type {Array<{id: number, title: string, lang: string, content: string}>}
   */
  let gists = [],
    pageGists = [];

  async function fetchGists() {
    const response = await fetch(`/gists/search?q=${queryText}&page=${page}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    if (queryText !== "") {
      gists = await response.json();
    } else {
      pageGists = await response.json();
    }
    isLoading = false;
  }

  fetchGists();

  const search = debounce(fetchGists, 750);

  $: {
    if (queryText.startsWith("/")) {
      gists = [{
        title: "readme",
        lang: "md",
        content: "/p, /ai",
      }]
    } else if (queryText !== "" || page !== 1) {
      isLoading = true;
      search();
    }
  }
</script>

<div class="gist-main">
  <input bind:value={queryText} placeholder="/" />
  <div class="gists">
    {#if isLoading}
      <Loading color="#ff1493" />
    {:else if queryText !== ""}
      {#each gists as gist}
        <div class="gist">
          <div class="gist-title">
            {@html gist.title} <span>.{@html gist.lang}</span>
          </div>
          <div class="gist-content">{@html gist.content}</div>
        </div>
      {/each}
    {:else}
      {#each pageGists as gist}
        <div class="gist">
          <div class="gist-title">
            {@html gist.title} <span>.{@html gist.lang}</span>
          </div>
          <div class="gist-content markdown-body">{@html gist.content}</div>
        </div>
      {/each}
      {#if pageGists.length >= 12}
        <div class="next-btn-container">
          <button
            on:click={() => {
              page++;
            }}>下一页</button
          >
        </div>
      {/if}
    {/if}
  </div>
</div>

<style>
  .gist-main {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 15px 10px 20px;
  }
  .gist-main input {
    width: 44.5%;
    box-sizing: border-box;
    border-radius: 3px;
    font-size: 16px;
    padding: 8px 13px;
    border: 1px solid var(--border);
    box-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05);
    background-color: var(--background);
    outline: 0;
  }
  .gist-main input:focus {
    border-color: aqua;
  }
  .gists {
    width: 100%;
  }
  .gist {
    margin: 15px 33px;
  }
  .gist .gist-title {
    margin: 2px;
    font-size: 17px;
  }
  .gist .gist-content {
    border: 1px solid var(--border);
    border-radius: 4px;
    padding: 10px 12px;
  }
  .next-btn-container {
    text-align: center;
    margin-top: 18px;
  }
  .next-btn-container button {
    font-size: 0.9rem;
  }
  :global(em) {
    border-bottom: 2px solid #ff1493;
    border-radius: 2px;
    font-style: normal;
  }
</style>
