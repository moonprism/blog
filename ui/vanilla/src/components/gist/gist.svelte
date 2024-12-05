<svelte:options customElement={{ tag: 'mod-gist', shadow: 'none' }} />

<script>
  import { debounce } from '@/utils';
  import Loading from '@/components/loading.svelte';
  import markdown from 'moonprism-markdown';

  import { foresee, getUrlIds, setUrlIds, isMdLang, toggleTheme } from './funcs';

  let isLoading = false;

  export let cdn = '';

  let currentPage = 1;
  // 当前输入框文本
  let currentKeyword = '';

  /**
   * @type {Array<{id: number, title: string, lang: string, content: string}>}
   */
  let gists = [], // gists 页面数据
    specialGists = []; // 查询结果/自定义数据

  async function search(keyword = '', source = '') {
    if (currentKeyword === '' || keyword === '') {
      isLoading = false;
      return;
    }
    const response = await fetch(`/gists/search?keyword=${keyword}&source=${source}`);
    if (!response.ok) {
      //TODO 查询错误处理
      isLoading = false;
      return;
    }
    specialGists = await response.json();
    const currentSk = currentKeyword.startsWith('/p ')
      ? currentKeyword.slice('/p '.length)
      : currentKeyword;
    if (currentSk != keyword) {
      return;
    }
    isLoading = false;
    if (source === 'art') {
      specialGists.map((g) => {
        g.title = `<a style="color:var(--primary);font-size:1.25rem" href="/post/${g.id}">${g.title}</a>`;
        return g;
      });
    } else {
      setUrlIds(specialGists.map((g) => g.id));
    }
  }

  const ids = getUrlIds().join(',');

  async function fetchGists(page = 1) {
    window.scrollTo(0, 0);
    isLoading = true;
    const response = await fetch(`/gists/search?page=${page}&ids=${ids}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    gists = await response.json();
    // 按照ids排序
    if (ids !== '') {
      const idOrder = ids.split(',').map((s) => Number(s));
      gists = gists.sort((a, b) => {
        return idOrder.indexOf(a.id) - idOrder.indexOf(b.id);
      });
    }
    isLoading = false;
  }
  fetchGists(1);

  let showHelpPanel = false;

  /**
   * 执行自定义命令
   * @param {string} command
   * @param {...string} args
   */
  function exec(command = '', ...args) {
    showHelpPanel = false;
    switch (command) {
      case 'tags':
        execTagsCommand();
        break;
      case 'p':
        const keyword = args.join(' ').trim();
        if (keyword === '') {
          showHelpPanel = true;
        } else {
          isLoading = true;
          debounceSearch(keyword, 'art');
        }
        break;
      case 'dark':
        toggleTheme(true);
        currentKeyword = '';
        break;
      case 'light':
        toggleTheme(false);
        currentKeyword = '';
        break;
      default:
        showHelpPanel = true;
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
    const response = await fetch('/api/tag');
    if (!response.ok) {
      throw new Error('Request tags network response was not ok');
    }
    const res = await response.json();
    tags = res.data;
    specialGists = [
      {
        title: 'Tags',
        lang: 'md',
        content: tags
          .map((t) => {
            return `<mod-tag i="${t.id}" color="${t.color}">${t.name}</mod-tag> `;
          })
          .join('')
      }
    ];
    isLoading = false;
  }

  const debounceSearch = debounce(search, 600);

  // 输入为空格时不作处理
  let lastKeyword = '';

  function handleKeywordInput() {
    if (currentKeyword !== '') {
      if (lastKeyword.trim() === currentKeyword.trim()) {
        lastKeyword = currentKeyword;
        return;
      }
      if (currentKeyword.startsWith('/')) {
        exec(...currentKeyword.slice(1).split(' '));
      } else {
        debounceSearch(currentKeyword.trim());
        isLoading = true;
        showHelpPanel = false;
      }
    } else {
      setUrlIds([]);
      showHelpPanel = false;
      specialGists = [];
    }
    lastKeyword = currentKeyword;
  }

  let copyColor = 'var(--foreground)';
</script>

<div class="gist-main" style="--copy-color: {copyColor};">
  <input bind:value={currentKeyword} on:input={handleKeywordInput} placeholder="/" />
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
          </ul>
          <h3>Tips</h3>
          <ul>
            <li>gists 的搜索结果将记录在当前 url 上，刷新渲染</li>
          </ul>
        </div>
      </div>
    {:else if isLoading}
      <Loading color="#ff1493" />
    {:else if currentKeyword !== ''}
      {#each specialGists as gist}
        <div class="gist">
          <div class="gist-title">
            {#if isMdLang(gist.lang)}
              <svg><use href="#icon-md" /></svg>
            {:else}
              <svg><use href="#icon-terminal" /></svg>
            {/if}
            {@html foresee(gist.title)} <span>.{@html foresee(gist.lang)}</span>
          </div>
          {#if isMdLang(gist.lang)}
            <div class="gist-content markdown-body">
              {#each gist.content.split('\n== 🌟 ==\n') as p}
                {@html foresee(markdown(p, { imageCdnUrl: cdn }))}
              {/each}
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
            {#if gist.lang === 'md'}
              <svg><use href="#icon-md" /></svg>
            {:else}
              <svg><use href="#icon-terminal" /></svg>
            {/if}
            {@html gist.title} <span>.{@html gist.lang}</span>
          </div>
          {#if gist.lang === 'md'}
            <div class="gist-content markdown-body">
              {@html gist.content}
            </div>
          {:else}
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <div
              class="gist-content"
              on:mouseenter={() => {
                copyColor = 'var(--foreground)';
              }}
            >
              <div class="copy">
                <button
                  on:click={(event) => {
                    let el = event.target.parentElement;
                    while (!el.classList.contains('gist-content')) {
                      el = el.parentElement;
                    }
                    navigator.clipboard.writeText(el.innerText);
                    copyColor = 'rgb(34 197 94)';
                  }}><svg><use href="#icon-copy" /></svg></button
                >
              </div>
              {@html gist.content}
            </div>
          {/if}
        </div>
      {/each}
      <div class="next-btn-container">
        {#if ids !== ''}
          <div>[{ids}]</div>
        {:else if gists.length >= 12 && ids === ''}
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

<symbol id="icon-terminal" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
  <rect x="3" y="5" width="17" height="14" rx="2" stroke="var(--icon-sh-color)" stroke-width="2" />
  <path
    d="M7 10L9 12L7 14"
    stroke="var(--icon-sh-color)"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
  <path d="M12 14H16" stroke="var(--icon-sh-color)" stroke-width="2" stroke-linecap="round" />
</symbol>

<!-- License: CC Attribution. Made by Amir Baqian: https://dribbble.com/amirbaqian -->
<symbol id="icon-md" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
  <path
    fill-rule="evenodd"
    clip-rule="evenodd"
    d="M14.9703 3.3437C13.0166 2.88543 10.9834 2.88543 9.02975 3.3437C6.20842 4.00549 4.0055 6.20841 3.3437 9.02975C2.88543 10.9834 2.88543 13.0166 3.3437 14.9703C4.0055 17.7916 6.20842 19.9945 9.02975 20.6563C10.9834 21.1146 13.0166 21.1146 14.9703 20.6563C17.7916 19.9945 19.9945 17.7916 20.6563 14.9703C21.1146 13.0166 21.1146 10.9834 20.6563 9.02975C19.9945 6.20842 17.7916 4.00549 14.9703 3.3437ZM8.55377 9.12812C8.55377 8.8109 8.81093 8.55374 9.12815 8.55374H12.9573C13.2745 8.55374 13.5317 8.8109 13.5317 9.12812C13.5317 9.44533 13.2745 9.70249 12.9573 9.70249H9.12815C8.81093 9.70249 8.55377 9.44533 8.55377 9.12812ZM8.55377 12C8.55377 11.6828 8.81093 11.4256 9.12815 11.4256H14.8719C15.1891 11.4256 15.4462 11.6828 15.4462 12C15.4462 12.3172 15.1891 12.5743 14.8719 12.5743H9.12815C8.81093 12.5743 8.55377 12.3172 8.55377 12ZM8.55377 14.8718C8.55377 14.5546 8.81093 14.2975 9.12815 14.2975H12C12.3172 14.2975 12.5744 14.5546 12.5744 14.8718C12.5744 15.189 12.3172 15.4462 12 15.4462H9.12815C8.81093 15.4462 8.55377 15.189 8.55377 14.8718Z"
    fill="var(--icon-md-color)"
  />
</symbol>

<!-- License: PD. Made by Sargam Icons: https://github.com/planetabhi/sargam-icons -->
<symbol
  id="icon-copy"
  width="18px"
  height="18px"
  viewBox="0 0 24 24"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    d="M18.6 9h-7.2A2.4 2.4 0 0 0 9 11.4v7.2a2.4 2.4 0 0 0 2.4 2.4h7.2a2.4 2.4 0 0 0 2.4-2.4v-7.2A2.4 2.4 0 0 0 18.6 9Z"
    fill="var(--copy-color)"
    fill-opacity=".16"
    stroke="var(--copy-color)"
    stroke-width="1.5"
    stroke-miterlimit="10"
  /><path
    d="M6 15h-.6C4.07 15 3 13.93 3 12.6V5.4C3 4.07 4.07 3 5.4 3h7.2C13.93 3 15 4.07 15 5.4V6"
    stroke="var(--copy-color)"
    stroke-width="1.5"
    stroke-miterlimit="10"
    stroke-linecap="round"
  />
</symbol>

<style>
  .gist-main {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 15px 10px 20px;
    --icon-sh-color: #33363f;
    --icon-md-color: #363853;
  }
  :global(.dark .gist-main) {
    --icon-sh-color: var(--foreground) !important;
    --icon-md-color: var(--foreground) !important;
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
    color: var(--foreground);
    outline: 0;
  }
  @media (max-width: 768px) {
    .gist-main input {
      width: 70%;
    }
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
    margin: 15px 0;
  }
  .gist .gist-title {
    line-height: 1.5rem;
    margin: 3px 2px;
    font-size: 16.5px;
    display: flex;
    align-items: center;
  }
  .gist svg {
    margin-right: 3px;
    width: 20px;
    height: 20px;
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
  .gist-content .copy {
    position: relative;
  }
  .gist-content .copy button {
    position: absolute;
    right: -16px;
    top: -7px;
    opacity: 0;
    transition: opacity 0.2s ease-in-out;
    background-color: unset;
  }
  .gist-content:hover .copy button {
    opacity: 1;
  }
  .next-btn-container {
    text-align: center;
    margin-top: 18px;
  }
  .next-btn-container button {
    font-size: 0.85rem;
  }
  .next-btn-container button:hover {
    text-decoration: underline;
  }
  :global(.em) {
    border-bottom: 2px solid #ff1493;
    font-style: normal;
  }
</style>
