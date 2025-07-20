import { useTheme, themes } from './ThemeContext';

// PrimeReact components
import { Button } from 'primereact/button';
import { OverlayPanel } from 'primereact/overlaypanel';
import { useRef } from 'react';

export const ThemeSwitcher = () => {
  const { theme, setTheme } = useTheme();
  const op = useRef<OverlayPanel>(null);
  
  const currentTheme = themes.find(t => t.value === theme) || themes[0];
  
  // Toggle between light and dark version of the current color
  const toggleLightDark = () => {
    const currentThemeObj = themes.find(t => t.value === theme);
    if (currentThemeObj) {
      const currentColor = currentThemeObj.value.split('-')[2]; // 'blue', 'indigo', etc.
      const isDark = currentThemeObj.dark;
      
      // Find the opposite theme with the same color
      const newThemeType = isDark ? 'light' : 'dark';
      // Find the matching theme from our predefined list
      const targetTheme = themes.find(t => 
        t.value.includes(`lara-${newThemeType}`) && t.value.includes(currentColor)
      );
      
      if (targetTheme) {
        setTheme(targetTheme.value);
      }
    }
  };
  
  return (
    <div className="flex items-center gap-2">
      <Button 
        icon={currentTheme.dark ? "pi pi-sun" : "pi pi-moon"}
        onClick={toggleLightDark} 
        rounded
        text
        aria-label="Toggle light/dark mode"
        tooltip={currentTheme.dark ? "Switch to light mode" : "Switch to dark mode"}
        tooltipOptions={{ position: 'bottom' }}
      />
      
      <Button 
        label={currentTheme.name}
        icon="pi pi-chevron-down"
        iconPos="right"
        outlined
        onClick={(e) => op.current?.toggle(e)}
        aria-label="Select theme"
        className="p-button-sm"
      />
      
      <OverlayPanel ref={op} className="w-56">
        <div className="p-2">
          {themes.map((themeOption) => (
            <Button
              key={themeOption.value}
              label={themeOption.name}
              onClick={() => {
                setTheme(themeOption.value);
                op.current?.hide();
              }}
              className={`mb-1 w-full justify-content-start ${theme === themeOption.value ? 'p-button-outlined' : 'p-button-text'}`}
              size="small"
            />
          ))}
        </div>
      </OverlayPanel>
    </div>
  );
};

export default ThemeSwitcher;
