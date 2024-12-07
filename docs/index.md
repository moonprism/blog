---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "kicoe's Blog"
  text: ""
  tagline: 简单的博客系统，使用 Golang + Svelte 开发
  actions:
    - theme: brand
      text: 快速开始
      link: /guide/getting-started
    - theme: alt
      text: 在线演示(后台)
      link: https://moonprism.github.io/blog/demo/login

features:
  - icon: 🚀
    title: 小巧、高速
    details: 仅需少量的内存与磁盘空间。
  - icon: 🔍
    title: 强大的检索功能
    details: 基于 SQLite 的 FTS 模块实现，支持中文分词和拼音的全文搜索。
  - icon: 📝
    title: 跨平台
    details: 或许可以当个本地记事本...
---
