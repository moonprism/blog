import css from 'rollup-plugin-css-only';
import terser from '@rollup/plugin-terser';

export default [
    {
        input: [
            'src/main.js',
        ],
        output: {
            dir: 'dist',
            format: 'es',
            entryFileNames: '[name]-[hash].js',
            // https://github.com/thgh/rollup-plugin-css-only/issues/25
            assetFileNames: 'assets/style-[hash].css',
            // plugins: [terser()],
        },
        plugins: [
            css(),
        ]
    },
];
