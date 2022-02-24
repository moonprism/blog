<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
  @font-face {
    font-family: 'yuan';
    src: url(https://cdn.jsdelivr.net/gh/moonprism/cdn/font/yuan.ttf);
  }
  @import url('https://fonts.googleapis.com/css2?family=Space+Mono&display=swap');
  body {
    font-family: "Space Mono","Helvetica Neue",Helvetica,yuan,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
  }
  textarea, input {
      font-family: "Space Mono",yuan;
  }
</style>
