import svelte from 'rollup-plugin-svelte';
import resolve from '@rollup/plugin-node-resolve';
import terser from '@rollup/plugin-terser';

export default [
    {
        input: [
            'src/components/code.svelte',
            'src/components/comment.svelte'
        ],
        output: {
            dir: 'dist/mod',
            format: 'es',
            entryFileNames: '[name]-[hash].js',
            // plugins: [terser()],
        },
        plugins: [
            svelte({
                compilerOptions: {
                    customElement: true,
                }
            }),
            resolve({
                browser: true,
                dedupe: ['svelte'],
                exportConditions: ['svelte'],
            }),
        ]
    },
];
