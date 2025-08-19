<script>
  import { Link } from '@inertiajs/svelte';
  import ThemeSwitcher from '../components/ThemeSwitcher.svelte';
  
  export let user = null;
  export let title = 'Mock University';
  
  let sidebarVisible = false;
  
  // Menu items for the top menubar
  const menuItems = [
    {
      label: 'Home',
      icon: 'pi pi-home',
      url: '/',
    },
    {
      label: 'Courses',
      icon: 'pi pi-book',
      items: [
        {
          label: 'Browse Courses',
          icon: 'pi pi-list',
          url: '/courses',
        },
        {
          label: 'My Courses',
          icon: 'pi pi-user',
          url: '/my-courses',
        },
      ],
    },
    {
      label: 'Calendar',
      icon: 'pi pi-calendar',
      url: '/calendar',
    },
    {
      label: 'Resources',
      icon: 'pi pi-folder',
      items: [
        {
          label: 'Library',
          icon: 'pi pi-book',
          url: '/library',
        },
        {
          label: 'Downloads',
          icon: 'pi pi-download',
          url: '/downloads',
        },
      ],
    },
  ];

  // Sidebar menu items
  const sidebarItems = [
    {
      label: 'Dashboard',
      icon: 'pi pi-home',
      url: '/',
    },
    {
      label: 'Academic',
      icon: 'pi pi-book',
      items: [
        {
          label: 'Courses',
          icon: 'pi pi-list',
          url: '/courses',
        },
        {
          label: 'Grades',
          icon: 'pi pi-chart-bar',
          url: '/grades',
        },
        {
          label: 'Schedule',
          icon: 'pi pi-calendar',
          url: '/schedule',
        },
      ],
    },
    {
      label: 'Administration',
      icon: 'pi pi-cog',
      items: [
        {
          label: 'Profile',
          icon: 'pi pi-user',
          url: '/profile',
        },
        {
          label: 'Settings',
          icon: 'pi pi-sliders-h',
          url: '/settings',
        },
      ],
    },
    {
      label: 'Help & Support',
      icon: 'pi pi-question-circle',
      url: '/support',
    },
  ];
  
  function toggleSidebar() {
    sidebarVisible = !sidebarVisible;
  }
  
  function closeSidebar() {
    sidebarVisible = false;
  }
  
  function handleClickOutside(event) {
    if (sidebarVisible && !event.target.closest('.sidebar')) {
      closeSidebar();
    }
  }
</script>

<svelte:window on:click={handleClickOutside} />

<div class="min-h-screen flex flex-col">
  <!-- Sidebar -->
  {#if sidebarVisible}
    <div class="fixed inset-0 z-40 bg-black bg-opacity-50" on:click={closeSidebar}></div>
  {/if}
  
  <div class="sidebar fixed left-0 top-0 h-full w-64 bg-white dark:bg-gray-800 border-r border-gray-300 dark:border-gray-600 z-50 transform transition-transform duration-300 {sidebarVisible ? 'translate-x-0' : '-translate-x-full'}">
    <div class="p-3">
      <h2 class="text-xl font-bold mb-4">{title}</h2>
      <nav class="w-full">
        {#each sidebarItems as item}
          <div class="mb-2">
            {#if item.items}
              <div class="font-semibold text-gray-700 dark:text-gray-300 px-3 py-2 flex items-center">
                <i class="{item.icon} mr-2"></i>
                {item.label}
              </div>
              {#each item.items as subItem}
                <Link href={subItem.url} class="flex items-center px-6 py-2 text-gray-600 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 rounded">
                  <i class="{subItem.icon} mr-2"></i>
                  {subItem.label}
                </Link>
              {/each}
            {:else}
              <Link href={item.url} class="flex items-center px-3 py-2 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded">
                <i class="{item.icon} mr-2"></i>
                {item.label}
              </Link>
            {/if}
          </div>
        {/each}
      </nav>
    </div>
  </div>

  <!-- Top Navigation -->
  <nav class="bg-white dark:bg-gray-800 border-b border-gray-300 dark:border-gray-600 shadow-md">
    <div class="flex items-center justify-between px-4 py-3">
      <!-- Start: Hamburger menu and logo -->
      <div class="flex items-center">
        <button 
          class="mr-2 p-2 text-gray-600 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded"
          on:click={toggleSidebar}
        >
          <i class="pi pi-bars"></i>
        </button>
        <Link href="/" class="text-xl font-bold no-underline text-gray-900 dark:text-white">
          {title}
        </Link>
      </div>
      
      <!-- Center: Main menu items -->
      <div class="hidden md:flex items-center space-x-6">
        {#each menuItems as item}
          {#if item.items}
            <div class="relative group">
              <button class="flex items-center text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white">
                <i class="{item.icon} mr-1"></i>
                {item.label}
                <i class="pi pi-chevron-down ml-1"></i>
              </button>
              <div class="absolute left-0 mt-2 w-48 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded shadow-lg opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 z-50">
                {#each item.items as subItem}
                  <Link href={subItem.url} class="flex items-center px-4 py-2 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700">
                    <i class="{subItem.icon} mr-2"></i>
                    {subItem.label}
                  </Link>
                {/each}
              </div>
            </div>
          {:else}
            <Link href={item.url} class="flex items-center text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white">
              <i class="{item.icon} mr-1"></i>
              {item.label}
            </Link>
          {/if}
        {/each}
      </div>

      <!-- End: User profile and theme switcher -->
      <div class="flex items-center gap-3">
        <ThemeSwitcher />
        {#if user}
          <div class="flex items-center gap-2">
            <span class="text-gray-700 dark:text-gray-300">{user.name}</span>
            <div class="w-8 h-8 bg-gray-300 dark:bg-gray-600 rounded-full flex items-center justify-center">
              {#if user.avatar}
                <img src={user.avatar} alt={user.name} class="w-8 h-8 rounded-full" />
              {:else}
                <i class="pi pi-user text-gray-600 dark:text-gray-400"></i>
              {/if}
            </div>
          </div>
        {/if}
      </div>
    </div>
  </nav>

  <!-- Main Content -->
  <main class="flex-grow p-4">
    <slot />
  </main>

  <!-- Footer -->
  <footer class="p-4 bg-gray-50 dark:bg-gray-800 text-center">
    <p class="text-sm text-gray-500 dark:text-gray-400">
      © {new Date().getFullYear()} Mock University. All rights reserved.
    </p>
  </footer>
</div>
