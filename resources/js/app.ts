import { createInertiaApp } from '@inertiajs/svelte'
import { resolvePageComponent } from 'laravel-vite-plugin/inertia-helpers'

// Import PrimeIcons
import 'primeicons/primeicons.css'

// Import CSS
import '../css/app.css'

createInertiaApp({
  resolve: (name: string) => resolvePageComponent(`./pages/${name}.svelte`, import.meta.glob('./pages/**/*.svelte')),
  setup({ el, App, props }: { el: Element, App: any, props: any }) {
    new App({ target: el, props })
  },
})
