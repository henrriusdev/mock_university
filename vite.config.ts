import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import path from "path";
import laravel from "laravel-vite-plugin";
import tailwind from "@tailwindcss/vite";

export default defineConfig({
  plugins: [
    laravel({
      input: ["resources/js/app.ts", "resources/css/app.css"],
      ssr: "resources/js/ssr.ts",
      publicDirectory: "public",
      buildDirectory: "build",
      refresh: ["resources/**"],
    }) as any,
    svelte(),
    tailwind(),
  ],
  build: {
    manifest: true,
    outDir: "public/build",
    rollupOptions: {
      input: "resources/js/app.ts",
      output: {
        entryFileNames: "assets/[name].js",
        chunkFileNames: "assets/[name].js",
        assetFileNames: "assets/[name].[ext]",
        manualChunks: undefined,
      },
    },
  },
  resolve: { 
    alias: { 
      "@": path.resolve(__dirname, "resources/js"),
      "$lib": path.resolve(__dirname, "resources/js/lib") 
    } 
  },
});
