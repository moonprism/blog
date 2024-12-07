import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  base: "/blog/docs/",
  title: "blog 使用手册",
  description: "User Guide",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "指南", link: "/guide/getting-started", activeMatch: "/guide/" },
      { text: "示例", link: "https://kicoe.com/" },
    ],
    
    logo: '/favicon.svg',

    sidebar: {
      "/guide/": {
        base: "/guide/",
        items: [
          {
            text: "指南",
            items: [
              { text: "快速开始", link: "getting-started" },
              {
                text: "图床配置",
                base: "/guide/attachment-",
                items: [
                  { text: "Github API", link: "github-api" },
                  { text: "阿里云 OSS", link: "ali-oss" },
                ],
              },
              { text: "部署", link: "deploy" },
            ],
          },
        ],
      },
    },

    socialLinks: [
      { icon: "github", link: "https://github.com/moonprism/blog" },
    ],
  },
});
