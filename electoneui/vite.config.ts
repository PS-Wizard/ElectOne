import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import fs from 'fs';
import { fileURLToPath } from 'url';
import { dirname, resolve } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

export default defineConfig({
    plugins: [sveltekit()],
    server: {
        https: {
            key: fs.readFileSync(resolve(__dirname, '../assets/certificates/localhost+2-key.pem')),
            cert: fs.readFileSync(resolve(__dirname, '../assets/certificates/localhost+2.pem'))
        },
        proxy: {}
    }
});
