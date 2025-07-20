import { ReactNode, useState } from 'react';
import { Link } from '@inertiajs/react';
import { ThemeSwitcher } from '../components/ThemeSwitcher';

// PrimeReact components
import { Menubar } from 'primereact/menubar';
import { MenuItem } from 'primereact/menuitem';
import { Avatar } from 'primereact/avatar';
import { Sidebar } from 'primereact/sidebar';
import { PanelMenu } from 'primereact/panelmenu';
import { Button } from 'primereact/button';

interface MainLayoutProps {
  children: ReactNode;
  title?: string;
  user?: {
    name: string;
    email: string;
    avatar?: string;
  };
}

export default function MainLayout({ children, title = 'Mock University', user }: MainLayoutProps) {
  const [sidebarVisible, setSidebarVisible] = useState(false);

  // Menu items for the top menubar
  const menuItems: MenuItem[] = [
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
  const sidebarItems: MenuItem[] = [
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

  // Start template for menubar (hamburger menu and logo)
  const start = (
    <div className="flex items-center">
      <Button 
        icon="pi pi-bars" 
        onClick={() => setSidebarVisible(true)} 
        className="mr-2"
        text
      />
      <Link href="/" className="text-xl font-bold no-underline">
        {title}
      </Link>
    </div>
  );

  // End template for menubar (user profile and theme switcher)
  const end = (
    <div className="flex items-center gap-3">
      <ThemeSwitcher />
      {user && (
        <div className="flex items-center gap-2">
          <span>{user.name}</span>
          <Avatar 
            icon={!user.avatar ? "pi pi-user" : undefined} 
            image={user.avatar} 
            shape="circle" 
            className="cursor-pointer"
          />
        </div>
      )}
    </div>
  );

  return (
    <div className="min-h-screen flex flex-col">
      {/* Sidebar */}
      <Sidebar 
        visible={sidebarVisible} 
        onHide={() => setSidebarVisible(false)}
        className="p-sidebar-md"
      >
        <div className="p-3">
          <h2 className="text-xl font-bold mb-4">{title}</h2>
          <PanelMenu model={sidebarItems} className="w-full" />
        </div>
      </Sidebar>

      {/* Top Navigation */}
      <Menubar 
        model={menuItems} 
        start={start} 
        end={end}
        className="border-none shadow-md"
      />

      {/* Main Content */}
      <main className="flex-grow p-4">
        {children}
      </main>

      {/* Footer */}
      <footer className="p-4 bg-gray-50 dark:bg-gray-800 text-center">
        <p className="text-sm text-gray-500 dark:text-gray-400">
          © {new Date().getFullYear()} Mock University. All rights reserved.
        </p>
      </footer>
    </div>
  );
}
