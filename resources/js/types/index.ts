// User interface
export interface User {
  name: string;
  email: string;
  avatar?: string;
}

// Course interface
export interface Course {
  id: number;
  code: string;
  name: string;
  instructor: string;
  schedule: string;
}

// Event interface
export interface Event {
  id: number;
  title: string;
  date: string;
}

// Menu item interface
export interface MenuItem {
  label: string;
  icon: string;
  url?: string;
  items?: MenuItem[];
}

// Dashboard props
export interface DashboardProps {
  user: User;
}

// Layout props
export interface LayoutProps {
  user?: User;
  title?: string;
}
