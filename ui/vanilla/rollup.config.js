import css from "rollup-plugin-css-only";
import svelte from "rollup-plugin-svelte";
import resolve from "@rollup/plugin-node-resolve";
import copy from "rollup-plugin-copy";
import terser from "@rollup/plugin-terser";
import del from "rollup-plugin-delete";

import crypto from "crypto";
var hash = crypto.randomBytes(4).toString("hex");

export default [
  {
    input: "src/main.js",
    output: {
      dir: "dist",
      format: "es",
      entryFileNames: `[name]-${hash}.js`,
      // https://github.com/thgh/rollup-plugin-css-only/issues/25
      assetFileNames: `assets/style-${hash}.css`,
      // plugins: [terser()],
    },
    plugins: [
      del({ targets: "dist/*" }),
      css(),
      copy({
        targets: [
          {
            src: "templates/*.html",
            dest: "dist",
            transform: (contents, filename) =>
              contents.toString().replaceAll("{hash}", hash),
          },
        ],
      }),
    ],
  },
  {
    input: [
      "src/components/core.svelte",
      "src/components/pager.svelte",
      "src/components/comment.svelte",
    ],
    output: {
      dir: "dist/mod",
      format: "es",
      chunkFileNames: `index-${hash}.js`,
      entryFileNames: `[name]-${hash}.js`,
      // plugins: [terser()],
    },
    plugins: [
      svelte({
        compilerOptions: {
          customElement: true,
        },
      }),
      resolve({
        browser: true,
        dedupe: ["svelte"],
        exportConditions: ["svelte"],
      }),
    ],
  },
];
