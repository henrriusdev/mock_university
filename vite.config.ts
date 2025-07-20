import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import path from "path";
import laravel from "laravel-vite-plugin";
import tailwind from "@tailwindcss/vite";

export default defineConfig({
  plugins: [
    laravel({
      input: ["resources/js/app.tsx", "resources/css/app.css"],
      publicDirectory: "public",
      buildDirectory: "build",
      refresh: ["resources/**"],
    }) as any,
    react(),
    tailwind(),
  ],
  build: {
    manifest: true,
    outDir: "public/build",
    rollupOptions: {
      input: "resources/js/app.tsx",
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
