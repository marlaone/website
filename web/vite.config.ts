import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    lib: {
      name: "marlaone",
      entry: 'src/index.ts',
      fileName: "marla-elements"
    },
  }
})
