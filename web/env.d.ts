/// <reference types="vite/client" />

declare module '*.vue' {
    import type { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
}

interface ImportMetaEnv {
    readonly VITE_TITLE: 'cxbooks'
    readonly VITE_API_URL: '/api/'
    readonly VITE_UI_URL: '/ui/'
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}