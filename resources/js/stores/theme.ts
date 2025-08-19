import { writable } from 'svelte/store';

// Define theme types
export type Theme = 
  | 'lara-light-blue' 
  | 'lara-dark-blue' 
  | 'lara-light-indigo' 
  | 'lara-dark-indigo' 
  | 'lara-light-purple' 
  | 'lara-dark-purple' 
  | 'lara-light-teal' 
  | 'lara-dark-teal' 
  | 'lara-light-green' 
  | 'lara-dark-green';

export interface ThemeOption {
  name: string;
  value: Theme;
  dark: boolean;
}

// Define all available themes
export const themes: ThemeOption[] = [
  { name: 'Lara Light Blue', value: 'lara-light-blue', dark: false },
  { name: 'Lara Dark Blue', value: 'lara-dark-blue', dark: true },
  { name: 'Lara Light Indigo', value: 'lara-light-indigo', dark: false },
  { name: 'Lara Dark Indigo', value: 'lara-dark-indigo', dark: true },
  { name: 'Lara Light Purple', value: 'lara-light-purple', dark: false },
  { name: 'Lara Dark Purple', value: 'lara-dark-purple', dark: true },
  { name: 'Lara Light Teal', value: 'lara-light-teal', dark: false },
  { name: 'Lara Dark Teal', value: 'lara-dark-teal', dark: true },
  { name: 'Lara Light Green', value: 'lara-light-green', dark: false },
  { name: 'Lara Dark Green', value: 'lara-dark-green', dark: true },
];

// Get initial theme from localStorage or default
function getInitialTheme(): Theme {
  if (typeof window !== 'undefined') {
    const savedTheme = localStorage.getItem('theme') as Theme;
    if (savedTheme && themes.some(t => t.value === savedTheme)) {
      return savedTheme;
    }
  }
  return 'lara-light-blue';
}

// Create the theme store
export const currentTheme = writable(getInitialTheme());

// Subscribe to theme changes and apply them
if (typeof window !== 'undefined') {
  currentTheme.subscribe((theme: Theme) => {
    localStorage.setItem('theme', theme);
    
    // Apply theme class to body for global styling
    document.body.classList.remove(...themes.map(t => `theme-${t.value}`));
    document.body.classList.add(`theme-${theme}`);
    
    // Set data-theme attribute for potential CSS selectors
    document.documentElement.setAttribute('data-theme', theme);
    
    // Dynamically load the selected theme CSS
    const linkId = 'prime-theme-css';
    let link = document.getElementById(linkId) as HTMLLinkElement;
    
    if (!link) {
      link = document.createElement('link');
      link.id = linkId;
      link.rel = 'stylesheet';
      document.head.appendChild(link);
    }
    
    link.href = `https://cdn.jsdelivr.net/npm/primereact@10.9.6/resources/themes/${theme}/theme.css`;
  });
}

export function getCurrentTheme(): ThemeOption {
  return themes.find(t => t.value === getInitialTheme()) || themes[0];
}

export function toggleLightDark(currentThemeValue: Theme): void {
  const currentThemeObj = themes.find(t => t.value === currentThemeValue);
  if (currentThemeObj) {
    const currentColor = currentThemeObj.value.split('-')[2]; // 'blue', 'indigo', etc.
    const isDark = currentThemeObj.dark;
    
    // Find the opposite theme with the same color
    const newThemeType = isDark ? 'light' : 'dark';
    const targetTheme = themes.find(t => 
      t.value.includes(`lara-${newThemeType}`) && t.value.includes(currentColor)
    );
    
    if (targetTheme) {
      currentTheme.set(targetTheme.value);
    }
  }
}
