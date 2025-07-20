import { createRoot } from 'react-dom/client';
import { createInertiaApp } from '@inertiajs/react';
import { resolvePageComponent } from 'laravel-vite-plugin/inertia-helpers';
import { PrimeReactProvider } from 'primereact/api';

// Import PrimeReact styles
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';

// Import CSS
import '../css/app.css';

// Import ThemeProvider
import { ThemeProvider } from './components/ThemeContext';



createInertiaApp({
  resolve: (name) => resolvePageComponent(`./pages/${name}.tsx`, import.meta.glob('./pages/**/*.tsx')),
  setup({ el, App, props }) {
    const root = createRoot(el);

    root.render(
      <ThemeProvider>
        <PrimeReactProvider>
          <App {...props} />
        </PrimeReactProvider>
      </ThemeProvider>
    );
  },
});