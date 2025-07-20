import { createContext, useContext, useEffect, useState, ReactNode } from 'react';

// Define all available themes
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

export const themes: { name: string; value: Theme; dark: boolean }[] = [
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

type ThemeContextType = {
  theme: Theme;
  setTheme: (theme: Theme) => void;
};

const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

export const useTheme = () => {
  const context = useContext(ThemeContext);
  if (!context) {
    throw new Error('useTheme must be used within a ThemeProvider');
  }
  return context;
};

interface ThemeProviderProps {
  children: ReactNode;
}

export const ThemeProvider = ({ children }: ThemeProviderProps) => {
  const [theme, setTheme] = useState<Theme>('lara-light-blue');

  // Load theme from localStorage on mount
  useEffect(() => {
    const savedTheme = localStorage.getItem('theme') as Theme;
    if (savedTheme && themes.some(t => t.value === savedTheme)) {
      setTheme(savedTheme);
    }
  }, []);

  // Save theme to localStorage when it changes
  useEffect(() => {
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
  }, [theme]);

  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};

export default ThemeContext;
