import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { resolve } from 'path';
import path from 'path';


const fullReloadAlways = {
    // Force a full reload of the page on every change. This is needed if working on django templates
    handleHotUpdate({ server }) {
        // Sleep for a second to ensure the backend has finished rebuilding
        setTimeout(() => {
            // Notify the client to reload the entire page
            server.ws.send({ type: "full-reload" });
        }, 1500);
    },
};

export default defineConfig({
    plugins: [
       vue(), // Vue plugin for Vite
       fullReloadAlways,
    ],
    resolve: {
        alias: {
            'vue': path.resolve(__dirname, 'node_modules/vue/dist/vue.esm-bundler.js'),
        },
    },
    server: {
        host: 'localhost', // Allow local access
        port: 8081,        // Port for Vite dev server
        strictPort: true,  // Ensure the port doesn't change
        watch: {
            usePolling: true
        },
        // hmr: {
        //     overlay: true,  // Show error overlay on changes
        //     client: {
        //         // Notify the client to reload the entire page when the backend changes
        //         reconnect: true, // Ensure the HMR client reconnects when the backend file changes
        //         onError: (err) => {
        //             // Log errors and trigger a full reload
        //             console.error('HMR Error: ', err);
        //             location.reload();
        //         },
        //     },
        // },
        proxy: {
            // Proxy requests for Vite's assets (node_modules, resources, @vite)
            'resources': {
                target: 'http://vite:8081', // Vite server URL
                changeOrigin: true,
                secure: false,
            },
            'node_modules': {
                target: 'http://vite:8081', // Vite server URL
                changeOrigin: true,
                secure: false,
            },
            '@vite': {
                target: 'http://vite:8081', // Vite server URL
                changeOrigin: true,
                secure: false,
            },
            '@id': {
                target: 'http://vite:8081', // Vite server URL
                changeOrigin: true,
                secure: false,
            },
            // All other requests (main HTML and Go routes) should go to the Go backend
            '^/(?!resources|node_modules|@vite|@id)': {
                target: 'http://frontend:8080', // Go server URL
                changeOrigin: true,
                secure: false,
            },
        },
    },
    build: {
        rollupOptions: {
            input: {
                app: resolve(__dirname, 'resources/js/app.js'),
                styles: resolve(__dirname, 'resources/css/app.css'),
            },
        },
        outDir: resolve(__dirname, 'public/assets'), // Output directory for production files
        assetsDir: '', // Place assets directly in the output directory
        emptyOutDir: true, // Clear output directory before build
    },
});