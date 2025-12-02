import ApplicationLogo from '@/Components/ApplicationLogo';
import { Avatar, AvatarFallback, AvatarImage } from '@/Components/ui/avatar';
import { Button } from '@/Components/ui/button';
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/Components/ui/dropdown-menu';
import { Separator } from '@/Components/ui/separator';
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarHeader,
    SidebarInset,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarProvider,
    SidebarSeparator,
    SidebarTrigger,
} from '@/Components/ui/sidebar';
import { Roles } from '@/types/props';
import { Link, usePage } from '@inertiajs/react';
import {
    BookOpen,
    BookOpenCheck,
    BookOpenText,
    BookPlus,
    Cog,
    CreditCard,
    FileText,
    GraduationCap,
    LayoutDashboard,
    LogOut,
    Pencil,
    ScrollText,
    ShieldUser,
    SquareUser,
    TableOfContents,
    User2,
    UserCog,
    UserPen,
} from 'lucide-react';
import { PropsWithChildren, ReactNode } from 'react';

export default function Authenticated({
    header,
    children,
}: PropsWithChildren<{ header?: ReactNode }>) {
    const user = usePage().props.auth.user;
    console.log(user);

    const navigation = [
        {
            label: 'Dashboard',
            href: route('dashboard'),
            active: Boolean(route().current('dashboard')),
            icon: LayoutDashboard,
            roles: [Roles.ADMIN],
        },
        {
            label: 'Professors',
            href: route('professors.index'),
            active: Boolean(route().current('professors.*')),
            icon: UserPen,
            roles: [Roles.ADMIN, Roles.PROFESSOR_BOSS, Roles.CAREER_LEAD],
        },
        {
            label: 'Professor Dashboard',
            href: 'professor-dashboard.index',
            active: Boolean(route().current('professor-dashboard.*')),
            icon: LayoutDashboard,
            roles: [Roles.PROFESSOR, Roles.PROFESSOR_BOSS, Roles.CAREER_LEAD],
        },
        {
            label: 'Careers',
            href: route('careers.index'),
            active: Boolean(route().current('careers.*')),
            icon: GraduationCap,
            roles: [Roles.ADMIN, Roles.CAREER_LEAD],
        },
        {
            label: 'Leader Dashboard',
            href: 'career-dashboard.index',
            active: Boolean(route().current('career-dashboard.*')),
            icon: LayoutDashboard,
            roles: [Roles.CAREER_LEAD],
        },
        {
            label: 'Students',
            href: 'students.index',
            active: Boolean(route().current('students.*')),
            icon: User2,
            roles: [Roles.ADMIN, Roles.CAREER_LEAD],
        },
        {
            label: 'Student Dashboard',
            href: 'student-dashboard.index',
            active: Boolean(route().current('student-dashboard.*')),
            icon: LayoutDashboard,
            roles: [Roles.STUDENT],
        },
        {
            label: 'Registrars',
            href: 'registrars.index',
            active: Boolean(route().current('registrars.*')),
            icon: ShieldUser,
            roles: [Roles.ADMIN],
        },
        {
            label: 'Registrar Dashboard',
            href: 'registrar-dashboard.index',
            active: Boolean(route().current('registrar-dashboard.*')),
            icon: LayoutDashboard,
            roles: [Roles.REGISTRAR],
        },
        {
            label: 'Settings',
            href: 'profile.edit',
            active: Boolean(route().current('profile.*')),
            icon: Cog,
            roles: [
                Roles.ADMIN,
                Roles.CAREER_LEAD,
                Roles.REGISTRAR,
                Roles.CASHIER,
            ],
        },
        {
            label: 'Subjects',
            href: 'subjects.index',
            active: Boolean(route().current('subjects.*')),
            icon: BookOpen,
            roles: [
                Roles.ADMIN,
                Roles.PROFESSOR,
                Roles.PROFESSOR_BOSS,
                Roles.CAREER_LEAD,
            ],
        },
        {
            label: 'Current subjects',
            href: 'current-subjects.index',
            active: Boolean(route().current('current-subjects.*')),
            icon: BookOpenCheck,
            roles: [Roles.STUDENT],
        },
        {
            label: 'Note History',
            href: 'note-history.index',
            active: Boolean(route().current('note-history.*')),
            icon: BookOpenText,
            roles: [Roles.STUDENT],
        },
        {
            label: 'Subject Inscription',
            href: 'subject-inscriptions.index',
            active: Boolean(route().current('subject-inscriptions.*')),
            icon: BookPlus,
            roles: [Roles.STUDENT],
        },
        {
            label: 'Requests',
            href: 'requests.index',
            active: Boolean(route().current('requests.*')),
            icon: FileText,
            roles: [
                Roles.STUDENT,
                Roles.PROFESSOR,
                Roles.CAREER_LEAD,
                Roles.REGISTRAR,
                Roles.ADMIN,
                Roles.PROFESSOR_BOSS,
            ],
        },
        {
            label: 'Audit logs',
            href: 'traces.index',
            active: Boolean(route().current('traces.*')),
            icon: ScrollText,
            roles: [Roles.ADMIN],
        },
        {
            label: 'Cashiers',
            href: 'cashiers.index',
            active: Boolean(route().current('cashiers.*')),
            icon: SquareUser,
            roles: [Roles.ADMIN],
        },
        {
            label: 'Payments',
            href: 'payments.index',
            active: Boolean(route().current('payments.*')),
            icon: BookOpenCheck,
            roles: [Roles.CASHIER],
        },
        {
            label: 'Invoices',
            href: 'invoices.index',
            active: Boolean(route().current('invoices.*')),
            icon: CreditCard,
            roles: [Roles.CASHIER, Roles.REGISTRAR, Roles.ADMIN],
        },
        {
            label: 'Payment reports',
            href: 'payment-reports.index',
            active: Boolean(route().current('payment-reports.*')),
            icon: TableOfContents,
            roles: [Roles.CASHIER],
        },
        {
            label: 'Cashier Dashboard',
            href: 'cashier-dashboard.index',
            active: Boolean(route().current('cashier-dashboard.*')),
            icon: LayoutDashboard,
            roles: [Roles.CASHIER],
        },
        {
            label: 'Reports',
            href: 'reports.index',
            active: Boolean(route().current('reports.*')),
            icon: TableOfContents,
            roles: [Roles.ADMIN],
        },
        {
            label: 'User Management',
            href: 'user-management.index',
            active: Boolean(route().current('user-management.*')),
            icon: UserCog,
            roles: [Roles.ADMIN],
        },
        {
            label: 'My logs',
            href: 'my-logs.index',
            active: Boolean(route().current('my-logs.*')),
            icon: ScrollText,
            roles: [
                Roles.ADMIN,
                Roles.CAREER_LEAD,
                Roles.REGISTRAR,
                Roles.CASHIER,
                Roles.PROFESSOR,
                Roles.PROFESSOR_BOSS,
                Roles.STUDENT,
            ],
        },
        {
            label: 'Roles & Permissions',
            href: 'roles-permissions.index',
            active: Boolean(route().current('roles-permissions.*')),
            icon: Pencil,
            roles: [Roles.ADMIN],
        },
    ];

    const initials =
        user?.name
            ?.split(' ')
            .filter(Boolean)
            .map((segment: string) => segment[0]?.toUpperCase())
            .slice(0, 2)
            .join('') ||
        user?.email?.[0]?.toUpperCase() ||
        '?';

    const hasHeader = Boolean(header);

    return (
        <SidebarProvider defaultOpen className="bg-muted/20">
            <Sidebar collapsible="icon">
                <SidebarHeader className="p-2 border-b border-sidebar-border">
                    <Link
                        href="/"
                        className="flex items-center h-12 gap-2 px-2 rounded-md"
                    >
                        <ApplicationLogo className="w-auto h-8 fill-foreground" />
                    </Link>
                </SidebarHeader>

                <SidebarContent>
                    <SidebarGroup>
                        <SidebarGroupLabel>Navigation</SidebarGroupLabel>
                        <SidebarGroupContent>
                            <SidebarMenu>
                                {navigation
                                    .filter((item) =>
                                        item.roles.some(
                                            (role) => user?.role === role,
                                        ),
                                    )
                                    .map(({ icon: Icon, ...item }) => (
                                        <SidebarMenuItem key={item.label}>
                                            <SidebarMenuButton
                                                asChild
                                                isActive={item.active}
                                                tooltip={item.label}
                                            >
                                                <Link
                                                    href={item.href}
                                                    className="flex items-center w-full gap-2"
                                                >
                                                    <Icon
                                                        className="w-4 h-4"
                                                        aria-hidden="true"
                                                    />
                                                    <span className="truncate">
                                                        {item.label}
                                                    </span>
                                                </Link>
                                            </SidebarMenuButton>
                                        </SidebarMenuItem>
                                    ))}
                            </SidebarMenu>
                        </SidebarGroupContent>
                    </SidebarGroup>
                </SidebarContent>

                <SidebarSeparator />

                <SidebarFooter className="p-2 mt-auto border-t border-sidebar-border">
                    <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                            <Button
                                variant="ghost"
                                className="w-full justify-start gap-3 rounded-md px-2 py-2 group-data-[collapsible=icon]:justify-center"
                            >
                                <Avatar className="w-8 h-8">
                                    {user?.profile_picture && (
                                        <AvatarImage
                                            src={user.profile_picture}
                                            alt={user?.name ?? 'User avatar'}
                                        />
                                    )}
                                    <AvatarFallback>{initials}</AvatarFallback>
                                </Avatar>
                                <div className="flex min-w-0 flex-col text-left group-data-[collapsible=icon]:hidden">
                                    <span className="text-sm font-medium truncate">
                                        {user?.name}
                                    </span>
                                    <span className="text-xs truncate text-muted-foreground">
                                        {user?.email}
                                    </span>
                                </div>
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent
                            align="end"
                            sideOffset={8}
                            className="w-56"
                        >
                            <DropdownMenuLabel className="flex flex-col gap-1">
                                <span className="text-xs text-muted-foreground">
                                    Signed in as
                                </span>
                                <span className="text-sm font-medium text-foreground">
                                    {user?.name}
                                </span>
                                <span className="text-xs text-muted-foreground">
                                    {user?.email}
                                </span>
                            </DropdownMenuLabel>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem asChild>
                                <Link
                                    href={route('profile.edit')}
                                    className="flex items-center gap-2"
                                >
                                    <User2 className="w-4 h-4" />
                                    Profile
                                </Link>
                            </DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem asChild>
                                <Link
                                    href={route('logout')}
                                    method="post"
                                    as="button"
                                    className="flex items-center w-full gap-2"
                                >
                                    <LogOut className="w-4 h-4" />
                                    Log Out
                                </Link>
                            </DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>
                </SidebarFooter>
            </Sidebar>

            <SidebarInset className="bg-muted/20">
                <div className="flex items-center gap-2 px-4 border-b h-14 bg-background sm:px-6 lg:px-8">
                    <SidebarTrigger />
                    {hasHeader && (
                        <>
                            <Separator orientation="vertical" className="h-6" />
                            <div className="flex items-center flex-1 overflow-hidden">
                                <div className="flex-1 overflow-hidden">
                                    {header}
                                </div>
                            </div>
                        </>
                    )}
                </div>
                <div className="flex-1 p-4 sm:p-6 lg:p-8">{children}</div>
            </SidebarInset>
        </SidebarProvider>
    );
}
