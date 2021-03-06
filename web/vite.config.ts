import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';

export default defineConfig({
  plugins: [solidPlugin()],
  server: {
    port: 3000,
  },
  build: {
    minify: 'esbuild',
    lib: {
      name: "marlaone",
      entry: 'src/index.ts',
      fileName: "marla-elements"
    },
  },
});
