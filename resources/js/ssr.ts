import { createInertiaApp } from '@inertiajs/svelte'
import { resolvePageComponent } from 'laravel-vite-plugin/inertia-helpers'
import createServer from '@inertiajs/svelte/server'

// Import PrimeIcons
import 'primeicons/primeicons.css'

// Import CSS
import '../css/app.css'

createServer((page) =>
  createInertiaApp({
    page,
    resolve: (name: string) => resolvePageComponent(`./pages/${name}.svelte`, import.meta.glob('./pages/**/*.svelte')),
    setup({ App, props }: { App: any, props: any }) {
      return new App({ props })
    },
  })
)
