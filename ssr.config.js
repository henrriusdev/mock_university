import { defineConfig } from 'vite';
import {svelte} from "@sveltejs/vite-plugin-svelte"
import laravel from "laravel-vite-plugin";
import path from "path";
export default defineConfig({plugins: [laravel({input: ['resources/js/app.js', 'resources/css/app.css'],ssr: 'resources/js/ssr.js', publicDirectory: 'public',buildDirectory: 'bootstrap',refresh: ['resources/**'],}),svelte(),],build: {ssr: true, outDir: 'bootstrap',rollupOptions: {input: 'resources/js/ssr.js',output: {entryFileNames: 'assets/[name].js',chunkFileNames: 'assets/[name].js',assetFileNames: 'assets/[name][extname]',manualChunks: undefined, },},},resolve: {alias: {$lib: path.resolve(__dirname, 'resources/js/lib'),},}})