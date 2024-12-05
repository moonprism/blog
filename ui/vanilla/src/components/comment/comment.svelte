<svelte:options customElement="mod-comment" />

<script>
  /**
   * @type {string} 文章ID
   */
  let id;
  export { id as i };

  import { timeAgoStr } from '@/utils';
  import CommentForm from './form.svelte';
  import Loading from '@/components/loading.svelte';

  let isRequestIn = false;

  /**
   * @typedef {Object} Comment
   * @property {number} id
   * @property {string} name
   * @property {string} email
   * @property {string} content
   * @property {string} link
   * @property {number} article_id
   * @property {number} reply_comment_id
   * @property {number} root_comment_i
   * @property {Comment} sub_comments
   */

  /**
   * @typedef {Object} CmntInfo
   * @property {Comment} data
   * @property {boolean} has_next
   */

  /**
   * @type {CmntInfo}
   */
  let cmntInfo = {
    data: [],
    has_next: false
  };

  let replyCommentId = 0;
  let replyName = '';

  /**
   * @type {Element}
   */
  let formDom;
  let replyLineDom;
  function replySelect(event, artId, name) {
    replyCommentId = artId;
    replyName = name;
    replyLineDom.style.height = '0';
    if (artId === 0) {
      return;
    }

    // 连线
    let replyCmntDom = event.target;
    for (let i = 0; i < 3; i++) {
      if (!replyCmntDom.classList.contains('cmnt-box')) {
        replyCmntDom = replyCmntDom.parentElement;
      } else {
        break;
      }
    }
    const lineHeight = replyCmntDom.offsetTop - formDom.offsetTop - formDom.offsetHeight + 10;
    replyLineDom.style.height = lineHeight + 'px';

    const rect = formDom.getBoundingClientRect();
    // 计算目标元素的绝对位置
    const scrollToY = window.scrollY + rect.top;
    // 滚动到目标位置
    window.scrollTo({ top: scrollToY, behavior: 'smooth' });
  }

  function request(rootId = 0, page = 1) {
    isRequestIn = true;
    fetch(`/comments/${id}?root_id=${rootId}&page=${page}`)
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => {
        if (rootId === 0 && page === 1) {
          cmntInfo = data;
        } else {
          if (rootId === 0) {
            cmntInfo.data = [...cmntInfo.data, ...data.data];
            cmntInfo.has_next = data.has_next;
          } else {
            const cmnt = cmntInfo.data.find((c) => c.id === rootId);
            cmnt.sub_comments = [...cmnt.sub_comments, ...data.data];
            cmnt.has_next = data.has_next;
            cmntInfo = cmntInfo;
          }
        }
        isRequestIn = false;
      })
      .catch((error) => console.error('Fetch error:', error));
  }
  request();

  let rootPage = 1;

  // 评论成功后展示数据&清空预设 data.detail: Comment
  function setAll(data) {
    const rootCmntId = data.detail.root_comment_id;
    if (rootCmntId === 0) {
      data.detail.sub_comments = [];
      cmntInfo.data.unshift(data.detail);
      cmntInfo.data = cmntInfo.data;
    } else {
      const rootCmnt = cmntInfo.data.find((c) => c.id === rootCmntId);
      rootCmnt.sub_comments.push(data.detail);
      const index = cmntInfo.data.indexOf(rootCmnt);
      if (index > -1) {
        const element = cmntInfo.data.splice(index, 1)[0]; // 删除元素
        cmntInfo.data.unshift(element); // 将元素添加到数组开头
      }
      cmntInfo.data = cmntInfo.data;
    }
    replyCommentId = 0;
    replyName = '';
    replyLineDom.style.height = '0';
  }
</script>

<div class="main-comment-form" bind:this={formDom}>
  <CommentForm {id} {replyName} {replyCommentId} on:post={setAll} />
  <div class="reply-line" bind:this={replyLineDom}></div>
</div>

<div class="board">
  {#each cmntInfo.data as c1}
    <div class="cmnt">
      <div class="cmnt-box">
        <div class="cmnt-left{replyCommentId == c1.id ? ' selected' : ''}">
          <div class="cmnt-header">
            <img
              alt="avatar"
              loading="lazy"
              src="https://gravatar.loli.net/avatar/{c1.email}?d=retro"
            />
            <div class="cmnt-header-info">
              <div class="name">{c1.name}</div>
              <div class="ago">{timeAgoStr(c1.created)}</div>
            </div>
            {#if replyCommentId == c1.id}
              <button
                class="reply-close-btn"
                on:click={(e) => {
                  replySelect(e, 0, '');
                }}
              >
                <svg><use href="#icon-close" /></svg>
              </button>
            {/if}
          </div>
          <p>{c1.content}</p>
        </div>
        {#if replyCommentId != c1.id}
          <button class="reply-btn" on:click={(e) => replySelect(e, c1.id, c1.name)}>
            <svg><use href="#icon-reply" /></svg>
          </button>
        {/if}
      </div>
      {#each c1.sub_comments as c2}
        <!--todo 超出400行再抽象出组件-->
        <div class="cmnt">
          <div class="cmnt-box">
            <div class="cmnt-left{replyCommentId == c2.id ? ' selected' : ''}">
              <div class="cmnt-header">
                <img
                  alt="avatar"
                  loading="lazy"
                  src="https://gravatar.loli.net/avatar/{c2.email}?d=retro"
                />
                <div class="cmnt-header-info">
                  <div class="name">
                    {c2.name}
                    {#if c2.reply_comment_id != c1.id}
                      <span>@{c1.sub_comments.find((j) => j.id === c2.reply_comment_id).name}</span>
                    {/if}
                  </div>
                  <div class="ago">{timeAgoStr(c2.created)}</div>
                </div>
                {#if replyCommentId == c2.id}
                  <button
                    class="reply-close-btn"
                    on:click={(e) => {
                      replySelect(e, 0, '');
                    }}
                  >
                    <svg><use href="#icon-close" /></svg>
                  </button>
                {/if}
              </div>
              <p>{c2.content}</p>
            </div>
            {#if replyCommentId != c2.id}
              <button class="reply-btn" on:click={(e) => replySelect(e, c2.id, c2.name)}>
                <svg><use href="#icon-reply" /></svg>
              </button>
            {/if}
          </div>
        </div>
      {/each}
      {#if c1.has_next}
        <div class="load-more">
          <button
            on:click={() => {
              // ts做得到吗
              if (!c1.currentPage) {
                c1.currentPage = 1;
              }
              request(c1.id, ++c1.currentPage);
            }}>more..</button
          >
        </div>
      {/if}
    </div>
  {/each}
  {#if cmntInfo.has_next && !isRequestIn}
    <div class="load-more">
      <button
        on:click={() => {
          request(0, ++rootPage);
        }}>more.</button
      >
    </div>
  {/if}
  {#if isRequestIn}
    <Loading />
  {/if}
</div>

<symbol id="icon-close" viewBox="0 0 100 100">
  <path
    style="fill:black;stroke:red;stroke-width:6;"
    d="M 20,4 3,21 33,50 3,80 20,97 49,67 79,97 95,80 65,50 95,20 80,4 50,34 z"
  />
</symbol>

<symbol id="icon-reply" viewBox="0 0 24 24">
  <path d="M22 2H2v14h2V4h16v12h-8v2h-2v2H8v-4H2v2h4v4h4v-2h2v-2h10V2z" fill="var(--foreground)" />
</symbol>

<style>
  .main-comment-form {
    padding: 40px 50px 10px;
    position: relative;
  }
  .reply-line {
    position: absolute;
    right: 76px;
    width: 1px;
    background: var(--border);
    transition: height 0.3s;
    -webkit-transition: height 0.3s;
  }
  .board {
    margin: 10px 40px 40px;
  }
  .board .cmnt {
    display: flex;
    flex-direction: column;
    padding-top: 10px;
    border-top: 1px solid var(--line);
    margin-left: 10px;
  }
  .board .cmnt .cmnt {
    padding-top: 5px;
    margin-left: 35px;
  }
  .cmnt-box {
    display: flex;
  }
  .cmnt-box:hover > .reply-btn {
    opacity: 1;
  }
  .board .cmnt .cmnt-left {
    flex: 1 1 auto;
    padding: 4px 0 4px 5px;
  }
  .board .cmnt:not(:first-child) {
    margin-top: 7px;
  }
  .cmnt-header {
    display: flex;
    align-items: center;
  }
  .cmnt-header .name {
    margin-left: 7px;
  }
  .cmnt-header .name span {
    font-size: 0.9rem;
    color: var(--gray-4);
  }
  .cmnt-header .ago {
    margin-left: 6px;
    color: var(--gray-4);
    font-size: 0.8rem;
  }
  .cmnt-header img {
    height: 30px;
    width: 30px;
    border-radius: 50%;
  }
  .cmnt-header-info {
    flex: 1 1 auto;
  }
  .cmnt p {
    margin: 10px 10px 0px;
    white-space: break-spaces;
  }
  button {
    border: 0;
    margin: 0;
    padding: 0;
    cursor: pointer;
    background-color: var(--background);
  }
  .cmnt .reply-btn {
    padding: 5px;
    opacity: 0;
    transition: opacity 0.2s ease-in-out;
  }
  .cmnt .reply-close-btn {
    margin-right: 4px;
    align-self: flex-start;
  }
  .reply-btn svg {
    width: 17px;
    height: 17px;
  }
  .reply-close-btn svg {
    width: 15px;
    height: 15px;
  }
  .reply-btn:hover svg,
  .reply-close-btn:hover svg {
    transform: scale(1.2);
  }
  .cmnt-left {
    border: 1px solid var(--background);
  }
  .selected {
    border: 1px solid var(--border);
    border-radius: 2px;
  }
  .load-more {
    text-align: center;
    margin-top: 15px;
    margin-bottom: -20px;
  }
  .load-more button {
    font-size: 0.9rem;
    padding: 0px 2px 1px;
    border-radius: 2px;
  }
  .load-more button:hover {
    background-color: var(--foreground);
    color: var(--background);
  }
  @media (max-width: 768px) {
    .main-comment-form {
      padding: 40px 10px 10px;
    }
    .board {
      margin: 10px 10px 40px;
    }
  }
</style>
