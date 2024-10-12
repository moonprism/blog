<script>
  import { createEventDispatcher } from "svelte";
  // 评论表单提交成功后触发父组件事件刷新列表
  const dispatch = createEventDispatcher();

  export let id = 0;
  // 后台jwt存在localStorage里所以没有使用csrfToken
  // 前台这里感觉也没必要，实现起来要维护session增加复杂度
  // export let csrfToken = '';

  let name = "";
  let email = "";
  let content = "";

  let errors = {
    Name: "",
    Email: "",
    Content: "",
  };

  let startCheck = false;

  let isPostIn = false;

  export let replyCommentId;
  export let replyName;

  async function post() {
    if (isPostIn) return;
    if (!checkFields()) {
      startCheck = true;
      return;
    }
    isPostIn = true;
    const response = await fetch("/api/comment", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        article_id: Number(id),
        reply_comment_id: Number(replyCommentId),
        root_comment_id: 0,
        name,
        email,
        content,
      }),
    });
    const data = await response.json();
    isPostIn = false;

    if (!response.ok) {
      errors = data;
      return;
    }

    startCheck = false;
    email = name = content = "";
    dispatch("post", data);
  }

  function checkFields() {
    errors.Name = name === "" ? "姓名不能为空" : "";
    errors.Email = email === "" ? "邮箱不能为空" : "";
    errors.Content = content === "" ? "没有什么想说的吗？(｡ŏ_ŏ)" : "";
    return errors.Name === "" && errors.Email === "" && errors.Content === "";
  }

  $: {
    name;
    email;
    content;
    if (startCheck) checkFields();
  }

  let isOpenMemeBoard = false;
  function triggerMemeBoard() {
    isOpenMemeBoard = !isOpenMemeBoard;
  }
  const memes = [
    "(｡・`ω´･)",
    "（￣▽￣）",
    "(╹╹)",
    "( ×ω× )",
    "(=・ω・=)",
    "☆～（ゝ。∂)",
    "(｀・ω・´)",
    "(｡ŏ_ŏ)",
    "（¯﹃¯）",
    "≧▽≦y",
    "(⁄ ⁄•⁄ω⁄•⁄ ⁄)",
    ">ㅂ<",
    "╮(￣▽￣)╭",
    "<(｀^´)>",
    "╮(｡>口<｡)╭",
    "♪(^∇^*)",
    "（︶︿︶）",
    "( T﹏T )",
    '("▔□▔)/',
    "(ﾟДﾟ≡ﾟдﾟ)!?",
    "(｡￫‿￩｡)",
    "( >﹏<。)～",
    "(^・ω・^ )",
    "(｡･ω･｡)",
    "_(:3」∠)_",
    "(・ω< )★",
    "(*゜ロ゜)ノ",
    "（＞д＜）",
    "(*´∀`)~♥",
    "(T_T)",
    "∑(O_O；)",
    "Σ(ﾟдﾟ;)",
    "Ciallo～(∠・ω＜)⌒☆",
  ];

  let textareaDom;
</script>

<form>
  <div class="form-fields">
    <div class="form-field">
      <div class="form-field-appear">
        <label for="name">姓名：</label>
        <input
          type="text"
          id="name"
          bind:value={name}
          placeholder="Name / [Name](🔗)"
          required
        />
      </div>
      <div class="form-field-error">{errors.Name ? errors.Name : ""}</div>
    </div>
    <div class="form-field">
      <div class="form-field-appear">
        <label for="email">邮箱：</label>
        <input
          type="text"
          id="email"
          bind:value={email}
          placeholder="✉️"
          required
        />
      </div>
      <div class="form-field-error">{errors.Email ? errors.Email : ""}</div>
    </div>
  </div>
  <div>
    <div class="form-text">
      <textarea bind:value={content} bind:this={textareaDom} />
      {#if isOpenMemeBoard}
        <div class="form-memes">
          {#each memes as meme}
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <!-- svelte-ignore a11y-missing-attribute -->
            <a
              on:click={() => {
                content += meme;
                triggerMemeBoard();
                textareaDom.focus();
              }}>{meme}</a
            >
          {/each}
        </div>
      {/if}
    </div>
    <div style="display: flex;">
      <div class="reply-name">
        {#if replyName !== ""}
          @{replyName}
        {/if}
      </div>
      <div class="form-field-error">{errors.Content ? errors.Content : ""}</div>
    </div>
  </div>
  <div class="form-buttons">
    <button type="button" class="form-button-meme" on:click={triggerMemeBoard}
      ><span>("▔□▔)/</span></button
    >
    <button type="button" class="form-button-post" on:click={post}>
      {#if isPostIn}
        <div class="loading"></div>
      {:else}
        发布
      {/if}
    </button>
  </div>
</form>

<style>
  form {
    display: flex;
    flex-direction: column;
  }
  form > :not(:last-child) {
    margin-bottom: 0.6rem;
  }
  .form-fields {
    display: flex;
    flex-direction: column;
    flex: 1 1 auto;
    padding-left: 2px;
    padding-right: 2px;
  }
  .form-fields > :not(:last-child) {
    margin-bottom: 0.4rem;
  }
  .form-field {
    flex: 1 1 auto;
  }
  .form-field .form-field-appear {
    flex: 1;
    display: flex;
    align-items: center;
  }
  .form-field-error {
    text-align: right;
    font-size: 0.8rem;
    color: hsl(var(--destructive));
  }
  .form-field input {
    flex: 1;
    margin-left: 2px;
  }
  .form-field input,
  form textarea {
    padding: 6px 10px;
    background-color: var(--background);
    border: 1px solid var(--border);
    border-radius: 3px;
    outline: 0;
    transition: border-color 0.3s;
  }
  form textarea:focus,
  .form-field input:focus {
    border-color: var(--outline);
  }
  .form-text {
    position: relative;
    display: flex;
  }
  form .form-text textarea {
    height: 6rem;
    color: var(--foreground);
    font-size: 0.9rem;
    flex: 1;
  }
  .form-text .form-memes {
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background-color: var(--background);
  }
  .form-text .form-memes a {
    font-size: 0.85rem;
    margin: 2px 2px;
    border: 1px solid var(--line);
    cursor: pointer;
    display: inline-block;
    padding: 1px 3px;
  }
  .form-text .form-memes a:hover {
    color: #ff1493;
    border-color: #ff1493;
  }
  @media (min-width: 768px) {
    .form-fields {
      flex-direction: row;
    }
    .form-field:first-child {
      width: 10px;
      margin-right: 12px;
    }
    .form-field:last-child {
      width: 60px;
    }
    .form-fields > :not(:last-child) {
      margin-bottom: 0;
    }
  }
  .form-buttons {
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }
  .reply-name {
    flex: 1 1 auto;
    margin-left: 1px;
    margin-top: 0.5px;
    font-size: 0.8rem;
    color: var(--gray-4);
  }
  .form-buttons > button {
    margin-left: 3px;
    border: 1px solid var(--outline);
    border-radius: 3px;
    cursor: pointer;
    background-color: var(--background);
    color: var(--foreground);
    font-size: 0.9rem;
  }
  .form-button-meme {
    padding: 3px 5px;
  }
  .form-button-meme:hover {
    display: inline-block;
    animation: shake-sm 0.3s infinite;
    border-color: var(--foreground);
  }
  .form-button-meme:hover span {
    display: inline-block;
    animation: shake 0.4s infinite;
  }
  .form-button-post {
    padding: 3px 13px;
    transition: background-color 0.3s ease;
    transition: color 0.3s ease;
  }
  .form-button-post:hover {
    background-color: var(--foreground);
    color: var(--background);
  }
  /* 抖动效果 */
  @keyframes shake {
    0%,
    100% {
      transform: translateY(0);
    }
    25% {
      transform: translateY(-1.5px);
    }
    50% {
      transform: translateY(1.5px);
    }
    75% {
      transform: translateY(-1.5px);
    }
  }
  @keyframes shake-sm {
    0%,
    100% {
      transform: translateX(0);
    }
    25% {
      transform: translateX(-0.5px);
    }
    50% {
      transform: translateX(0.5px);
    }
    75% {
      transform: translateX(-0.5px);
    }
  }
  .loading {
    border: 2px solid var(--background);
    border-top: 2px solid var(--foreground);
    border-radius: 50%;
    width: 0.9rem;
    height: 0.9rem;
    animation: spin 0.6s linear infinite;
  }
  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>
