import { fileURLToPath, URL } from "url";
import { defineConfig,loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import vuetify from "@vuetify/vite-plugin";

// https://vitejs.dev/config/
export default defineConfig( {
  define: {
    'process.env': process.env
  },
  plugins: [
    vue(),
    vuetify({
      autoImport: true,
    }),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:9999',
        changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/api/, '')
      },
    }
  },
  css: {
    preprocessorOptions: {
      scss: { charset: false },
      css: { charset: false },
    },
  },
});


// devServer: {
//   proxy: {
//     "/api": {
//       target: "http://127.0.0.1:9999",
//         secure: false,
//           logLevel: "debug"
//     }
//   }
// }