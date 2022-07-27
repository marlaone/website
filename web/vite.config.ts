import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    minify: 'esbuild',
    lib: {
      name: "marlaone",
      entry: 'src/index.ts',
      fileName: "marla-elements"
    },
  }
})
