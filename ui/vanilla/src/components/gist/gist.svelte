<svelte:options customElement={{ tag: "mod-gist", shadow: "none" }} />

<script>
  import { debounce } from "@/utils";
  import Loading from "@/components/loading.svelte";
  import markdown from "moonprism-markdown";

  import { foresee, getUrlIds, setUrlIds, isMdLang } from "./funcs";

  let isLoading = false;

  export let cdn = "";

  let currentPage = 1;
  // 当前输入框文本
  let currentKeyword = "";

  /**
   * @type {Array<{id: number, title: string, lang: string, content: string}>}
   */
  let gists = [], // gists 页面数据
    specialGists = []; // 查询结果/自定义数据

  async function search(keyword = "", source = "") {
    if (currentKeyword === "" || keyword === "") {
      isLoading = false;
      return;
    }
    const response = await fetch(`/gists/search?q=${keyword}&o=${source}`);
    if (!response.ok) {
      //TODO 查询错误处理
      isLoading = false;
      return;
    }
    specialGists = await response.json();
    isLoading = false;
    setUrlIds(specialGists.map((g) => g.id));
  }

  const ids = getUrlIds().join(",");

  async function fetchGists(page = 1) {
    isLoading = true;
    const response = await fetch(`/gists/search?page=${page}&ids=${ids}`);
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    gists = await response.json();
    // 按照ids排序
    if (ids !== "") {
      const idOrder = ids.split(",").map((s) => Number(s));
      gists = gists.sort((a, b) => {
        return idOrder.indexOf(a.id) - idOrder.indexOf(b.id);
      });
    }
    isLoading = false;
  }
  fetchGists(1);

  let showHelpPanel = false;
  /**
   * @param {boolean} show
   */
  function help(show = true) {
    showHelpPanel = show;
  }

  /**
   * 执行自定义命令
   * @param {string} command
   * @param {...string} args
   */
  function exec(command = "", ...args) {
    help(false);
    switch (command) {
      case "tags":
        execTagsCommand();
        break;
      case "p":
        const keyword = args.join(" ").trim();
        if (keyword === "") {
          help();
        } else {
          loading();
          debounceSearch(keyword, "art");
        }
        break;
      default:
        help();
    }
  }

  /**
   * @type {Array<{id: number, color: string, name: string}>}
   */
  let tags = [];
  async function execTagsCommand() {
    if (tags.length !== 0) {
      return;
    }
    isLoading = true;
    const response = await fetch("/api/tag");
    if (!response.ok) {
      throw new Error("Request tags network response was not ok");
    }
    const res = await response.json();
    tags = res.data;
    specialGists = [
      {
        title: "Tags",
        lang: "md",
        content: tags
          .map((t) => {
            return `<a class="art-tag"
                    href="posts/tag/${t.name}"
                    style="background-color: ${t.color};
                    margin-right: 7px;">${t.name}</a>`;
          })
          .join("")
      }
    ];
    isLoading = false;
  }

  const debounceSearch = debounce(search, 600);

  function loading() {
    isLoading = true;
  }

  $: {
    if (currentKeyword !== "") {
      if (currentKeyword.startsWith("/")) {
        exec(...currentKeyword.slice(1).split(" "));
      } else {
        debounceSearch(currentKeyword.trim());
        loading();
        help(false);
      }
    } else {
      setUrlIds([]);
      help(false);
    }
  }
</script>

<div class="gist-main">
  <input bind:value={currentKeyword} placeholder="/" />
  <div class="gists">
    {#if showHelpPanel}
      <div class="gist">
        <div class="gist-title">README<span>.md</span></div>
        <div class="gist-content markdown-body">
          <p>
            一个网络空间中随处可见的输入框，颜色和长度似乎根据设计者的某些喜好精心设定。<br
            />在曾经个人博客繁荣时代它们常常承载着陌生者的思想游荡于网路，如今却很少被使用了。
          </p>
          <h3>Commands</h3>
          <ul>
            <li><code>.</code>检索 gists</li>
            <li><code>/p .</code>检索 posts</li>
            <li><code>/tags</code>列出全部标签</li>
            <li><code>/dark|light</code>开启日|夜间模式</li>
            <li><code>/ai </code></li>
          </ul>
          <div class="admonition ad-caution">
            <p class="admonition-title">TIP</p>
            <p>正在完善中...</p>
          </div>
          <h3>Tips</h3>
          <p>gists 的搜索结果将记录在当前 url 上，刷新渲染</p>
        </div>
      </div>
    {:else if isLoading}
      <Loading color="#ff1493" />
    {:else if currentKeyword !== ""}
      {#each specialGists as gist}
        <div class="gist">
          <div class="gist-title">
            {@html foresee(gist.title)} <span>.{@html foresee(gist.lang)}</span>
          </div>
          {#if isMdLang(gist.lang)}
            <div class="gist-content markdown-body">
              {@html foresee(markdown(gist.content, { imageCdnUrl: cdn }))}
            </div>
          {:else}
            <div class="gist-content">
              <pre><code>{@html foresee(gist.content)}</code></pre>
            </div>
          {/if}
        </div>
      {/each}
    {:else}
      {#each gists as gist}
        <div class="gist">
          <div class="gist-title">
            {@html gist.title} <span>.{@html gist.lang}</span>
          </div>
          <div class="gist-content {gist.lang === 'md' ? 'markdown-body' : ''}">
            {@html gist.content}
          </div>
        </div>
      {/each}
      <div class="next-btn-container">
        {#if ids !== ""}
          <div>[{ids}]</div>
        {:else if gists.length >= 12 && ids === ""}
          <button
            on:click={() => {
              fetchGists(++currentPage);
            }}>下一页</button
          >
        {/if}
      </div>
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
    margin: 3px 2px;
    font-size: 16.5px;
  }
  .gist .gist-title span {
    font-size: 15px;
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
  :global(.em) {
    border-bottom: 2px solid #ff1493;
    font-style: normal;
  }
</style>
