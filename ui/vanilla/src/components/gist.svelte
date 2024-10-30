<svelte:options customElement={{ tag: "mod-gist", shadow: "none" }} />

<script>
  import { debounce } from "../utils";
  import Loading from "./loading.svelte";
  import markdown from "moonprism-markdown";

  let isLoading = false;

  export let cdn = "";
  let currentPage = 1;
  let queryText = "";

  /**
   * @type {Array<{id: number, title: string, lang: string, content: string}>}
   */
  let gists = [],
    pageGists = [];

  async function search(query = "") {
    if (queryText === "" || queryText.startsWith("/")) {
      isLoading = false;
      return;
    }
    const response = await fetch(`/gists/search?q=${query}`);
    if (!response.ok) {
      //TODO 查询错误处理
      isLoading = false;
      return;
    }
    gists = await response.json();
    isLoading = false;
  }

  async function fetchGists(page = 1) {
    isLoading = true;
    const response = await fetch(`/gists/search?page=${page}`);
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    pageGists = await response.json();
    isLoading = false;
  }
  fetchGists(1);

  let isHelp = false;
  /**
   * @param {boolean} show
   */
  function help(show = true) {
    isHelp = show;
  }

  /**
   * 执行自定义命令
   * @param {string} command
   * @param {...string} args
   */
  function exec(command = "", ...args) {
    switch (command) {
      case "tags":
        break;
      default:
        help();
    }
  }

  const debounceSearch = debounce(search, 600);

  function loading() {
    isLoading = true;
  }

  $: {
    if (queryText !== "") {
      if (queryText.startsWith("/")) {
        exec(...queryText.slice(1).split(" "));
      } else {
        debounceSearch(queryText);
        loading();
        help(false);
      }
    } else {
      help(false);
    }
  }
</script>

<div class="gist-main">
  <input bind:value={queryText} placeholder="/" />
  <div class="gists">
    {#if isHelp}
      <div class="gist">
        <div class="gist-title">readme<span>.md</span></div>
        <div class="gist-content markdown-body">
          <p>
            一个网络空间中随处可见的输入框，颜色和长度似乎根据设计者的某些喜好精心设定。<br
            />在曾经个人博客繁荣时代它们常常承载着陌生者的思想游荡于网路，如今却很少被使用了。
          </p>
          <h3>Commands</h3>
          <ul>
            <li><code>/p </code>搜索posts</li>
            <li><code>/tags</code>列出全部标签</li>
            <li><code>/dark|light</code>开启日|夜间模式</li>
            <li><code>/ai </code>AI对话</li>
          </ul>
          <div class="admonition ad-caution">
            <p class="admonition-title">TIP</p>
            <p>正在研发中...</p>
          </div>
        </div>
      </div>
    {:else if isLoading}
      <Loading color="#ff1493" />
    {:else if queryText !== ""}
      {#each gists as gist}
        <div class="gist">
          <div class="gist-title">
            {@html gist.title} <span>.{@html gist.lang}</span>
          </div>
          {#if gist.lang.replace(new RegExp(`^<em>|</em>$`, "g"), "") === "md"}
            <div class="gist-content markdown-body">
              {@html markdown(gist.content, { imageCdnUrl: cdn })}
            </div>
          {:else}
            <div class="gist-content">
              <pre><code>{@html gist.content}</code></pre>
            </div>
          {/if}
        </div>
      {/each}
    {:else}
      {#each pageGists as gist}
        <div class="gist">
          <div class="gist-title">
            {@html gist.title} <span>.{@html gist.lang}</span>
          </div>
          <div class="gist-content {gist.lang === 'md' ? 'markdown-body' : ''}">
            {@html gist.content}
          </div>
        </div>
      {/each}
      {#if pageGists.length >= 12}
        <div class="next-btn-container">
          <button
            on:click={() => {
              fetchGists(++currentPage);
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
  .gist-main input::placeholder {
    color: var(--outline);
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
    padding: 13px 17px;
    overflow: auto;
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
    font-style: normal;
  }
</style>
