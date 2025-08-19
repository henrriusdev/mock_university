<script>
  import { currentTheme, themes, toggleLightDark } from '../stores/theme';
  
  let overlayVisible = false;
  let overlayPanel;
  
  $: currentThemeObj = themes.find(t => t.value === $currentTheme) || themes[0];
  
  function handleToggleLightDark() {
    toggleLightDark($currentTheme);
  }
  
  function selectTheme(themeValue) {
    currentTheme.set(themeValue);
    overlayVisible = false;
  }
  
  function toggleOverlay() {
    overlayVisible = !overlayVisible;
  }
  
  function handleClickOutside(event) {
    if (overlayPanel && !overlayPanel.contains(event.target)) {
      overlayVisible = false;
    }
  }
</script>

<svelte:window on:click={handleClickOutside} />

<div class="flex items-center gap-2">
  <button 
    class="p-button p-component p-button-rounded p-button-text"
    on:click={handleToggleLightDark}
    aria-label="Toggle light/dark mode"
    title={currentThemeObj.dark ? "Switch to light mode" : "Switch to dark mode"}
  >
    <span class="p-button-icon pi {currentThemeObj.dark ? 'pi-sun' : 'pi-moon'}"></span>
  </button>
  
  <button 
    class="p-button p-component p-button-outlined p-button-sm"
    on:click={toggleOverlay}
    aria-label="Select theme"
  >
    <span class="p-button-label">{currentThemeObj.name}</span>
    <span class="p-button-icon p-button-icon-right pi pi-chevron-down"></span>
  </button>
  
  {#if overlayVisible}
    <div 
      bind:this={overlayPanel}
      class="p-overlaypanel p-component w-56 absolute z-50 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded shadow-lg mt-2"
      style="top: 100%; right: 0;"
    >
      <div class="p-2">
        {#each themes as themeOption}
          <button
            class="p-button p-component w-full justify-start mb-1 {$currentTheme === themeOption.value ? 'p-button-outlined' : 'p-button-text'}"
            on:click={() => selectTheme(themeOption.value)}
          >
            <span class="p-button-label">{themeOption.name}</span>
          </button>
        {/each}
      </div>
    </div>
  {/if}
</div>

<style>
  .p-button {
    display: inline-flex;
    cursor: pointer;
    user-select: none;
    align-items: center;
    vertical-align: bottom;
    text-align: center;
    overflow: hidden;
    position: relative;
    border: 1px solid;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    transition: background-color 0.2s, color 0.2s, border-color 0.2s, box-shadow 0.2s;
    border-radius: 6px;
  }
  
  .p-button-rounded {
    border-radius: 50%;
    padding: 0.5rem;
  }
  
  .p-button-text {
    background: transparent;
    border-color: transparent;
    color: #6366f1;
  }
  
  .p-button-outlined {
    background: transparent;
    color: #6366f1;
    border: 1px solid #6366f1;
  }
  
  .p-button-sm {
    font-size: 0.875rem;
    padding: 0.375rem 0.75rem;
  }
  
  .p-button:hover {
    background: rgba(99, 102, 241, 0.04);
  }
  
  .p-button-icon {
    margin-right: 0.5rem;
  }
  
  .p-button-icon-right {
    margin-left: 0.5rem;
    margin-right: 0;
  }
  
  .justify-start {
    justify-content: flex-start;
  }
</style>
