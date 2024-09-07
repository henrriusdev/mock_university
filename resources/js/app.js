import { createInertiaApp } from '@inertiajs/svelte';
import '../css/app.css';

createInertiaApp({resolve:n=>{const e=import.meta.glob("./Pages/**/*.svelte",{eager:!0});return e[`./Pages/${n}.svelte`]},setup({el:n,App:e,props:t}){new e({target:n,props:t})}});
