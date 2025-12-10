import react from "@vitejs/plugin-react-swc";
import { defineConfig } from "vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  esbuild: {
    drop: process.env.dev ? ["console"] : [],
  },
  build: {
    outDir: "../server/client",
    emptyOutDir: true,
  },
});
