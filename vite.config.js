import { defineConfig } from 'vite';
import laravel from 'laravel-vite-plugin';
import {svelte} from "@sveltejs/vite-plugin-svelte"
import path from "path";

export default defineConfig({plugins: [laravel({input: 'resources/js/app.js',publicDirectory: 'public',buildDirectory: 'build',refresh: ['resources/**'],}),svelte({compilerOptions: {hydratable: true,},}),],build: {manifest: true, outDir: 'public/build',rollupOptions: {input: 'resources/js/app.js',output: {entryFileNames: 'assets/[name].js',chunkFileNames: 'assets/[name].js',assetFileNames: 'assets/[name].[ext]',manualChunks: undefined, },},},resolve: {alias: {$lib: path.resolve(__dirname, 'resources/js/lib'),},}});
