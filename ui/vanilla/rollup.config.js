import svelte from 'rollup-plugin-svelte';
import resolve from '@rollup/plugin-node-resolve';
import copy from 'rollup-plugin-copy';
import terser from '@rollup/plugin-terser';
import del from 'rollup-plugin-delete';
import postcss from 'rollup-plugin-postcss';
import cssnano from 'cssnano';

import path from 'path';
import { fileURLToPath } from 'url';
import alias from '@rollup/plugin-alias';
import crypto from 'crypto';
const hash = crypto.randomBytes(4).toString('hex');
const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default [
  {
    input: 'src/main.js',
    output: {
      dir: 'dist',
      format: 'es',
      entryFileNames: `[name]-${hash}.js`,
      plugins: [terser()]
    },
    plugins: [
      del({ targets: 'dist/*' }),
      postcss({
        extract: `main-${hash}.css`,
        plugins: [cssnano()]
      }),
      copy({
        targets: [
          {
            src: 'templates/*.html',
            dest: 'dist/tmpl',
            transform: (contents, filename) => contents.toString().replaceAll('{hash}', hash)
          },
          {
            src: 'src/favicon.svg',
            dest: 'dist'
          }
        ]
      })
    ]
  },
  {
    input: [
      'src/components/kit.svelte',
      'src/components/pager.svelte',
      'src/components/comment/comment.svelte',
      'src/components/gist/gist.svelte'
    ],
    output: {
      dir: 'dist/mod',
      format: 'es',
      chunkFileNames: `index-${hash}.js`,
      entryFileNames: `[name]-${hash}.js`,
      plugins: [terser()]
    },
    plugins: [
      svelte({
        compilerOptions: {
          customElement: true
        }
      }),
      resolve({
        browser: true,
        dedupe: ['svelte'],
        exportConditions: ['svelte']
      }),
      alias({
        entries: [{ find: '@', replacement: `${__dirname}/src/` }]
      })
    ]
  }
];
