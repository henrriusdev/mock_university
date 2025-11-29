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
    Sheet,
    SheetClose,
    SheetContent,
    SheetTrigger,
} from '@/Components/ui/sheet';
import { cn } from '@/lib/utils';
import { Link, usePage } from '@inertiajs/react';
import { LogOut, Menu, User2 } from 'lucide-react';
import { PropsWithChildren, ReactNode } from 'react';

export default function Authenticated({
    header,
    children,
}: PropsWithChildren<{ header?: ReactNode }>) {
    const user = usePage().props.auth.user;

    const navigation = [
        {
            label: 'Dashboard',
            href: route('dashboard'),
            active: Boolean(route().current('dashboard')),
        },
    ];

    const initials =
        (user?.name
            ?.split(' ')
            .filter(Boolean)
            .map((segment: string) => segment[0]?.toUpperCase())
            .slice(0, 2)
            .join('') ||
            user?.email?.[0]?.toUpperCase() ||
            '?');

    return (
        <div className="min-h-screen bg-background">
            <nav className="border-b bg-card">
                <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                    <div className="flex h-16 items-center justify-between">
                        <div className="flex items-center gap-6">
                            <Link href="/" className="flex items-center">
                                <ApplicationLogo className="block h-8 w-auto fill-foreground" />
                            </Link>

                            <div className="hidden items-center gap-1 md:flex">
                                {navigation.map((item) => (
                                    <Link
                                        key={item.label}
                                        href={item.href}
                                        className={cn(
                                            'rounded-md px-3 py-2 text-sm font-medium transition-colors',
                                            item.active
                                                ? 'bg-muted text-foreground'
                                                : 'text-muted-foreground hover:text-foreground',
                                        )}
                                    >
                                        {item.label}
                                    </Link>
                                ))}
                            </div>
                        </div>

                        <div className="flex items-center gap-2">
                            <DropdownMenu>
                                <DropdownMenuTrigger asChild>
                                    <Button
                                        variant="ghost"
                                        className="hidden items-center gap-2 md:flex"
                                    >
                                        <Avatar className="h-8 w-8">
                                            {user?.profile_picture && (
                                                <AvatarImage
                                                    src={user.profile_picture}
                                                    alt={user.name ?? 'User avatar'}
                                                />
                                            )}
                                            <AvatarFallback>
                                                {initials}
                                            </AvatarFallback>
                                        </Avatar>
                                        <span className="text-sm font-medium">
                                            {user?.name}
                                        </span>
                                    </Button>
                                </DropdownMenuTrigger>
                                <DropdownMenuContent align="end" className="w-56">
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
                                        <Link href={route('profile.edit')} className="flex items-center gap-2">
                                            <User2 className="h-4 w-4" />
                                            Profile
                                        </Link>
                                    </DropdownMenuItem>
                                    <DropdownMenuSeparator />
                                    <DropdownMenuItem asChild>
                                        <Link
                                            href={route('logout')}
                                            method="post"
                                            as="button"
                                            className="flex w-full items-center gap-2"
                                        >
                                            <LogOut className="h-4 w-4" />
                                            Log Out
                                        </Link>
                                    </DropdownMenuItem>
                                </DropdownMenuContent>
                            </DropdownMenu>

                            <Sheet>
                                <SheetTrigger asChild>
                                    <Button
                                        variant="ghost"
                                        size="icon"
                                        className="md:hidden"
                                    >
                                        <Menu className="h-5 w-5" />
                                        <span className="sr-only">Open navigation</span>
                                    </Button>
                                </SheetTrigger>
                                <SheetContent side="left" className="w-80 px-0">
                                    <div className="flex flex-col gap-6 p-6">
                                        <div className="flex items-center gap-3">
                                            <Avatar className="h-10 w-10">
                                                {user?.profile_picture && (
                                                    <AvatarImage
                                                        src={user.profile_picture}
                                                        alt={user?.name ?? 'User avatar'}
                                                    />
                                                )}
                                                <AvatarFallback className="text-base">
                                                    {initials}
                                                </AvatarFallback>
                                            </Avatar>
                                            <div className="space-y-0.5">
                                                <p className="text-sm font-semibold text-foreground">
                                                    {user?.name}
                                                </p>
                                                <p className="text-xs text-muted-foreground">
                                                    {user?.email}
                                                </p>
                                            </div>
                                        </div>

                                        <div className="space-y-2">
                                            <p className="text-xs font-semibold uppercase text-muted-foreground">
                                                Navigation
                                            </p>
                                            <Separator />
                                            <div className="flex flex-col gap-1">
                                                {navigation.map((item) => (
                                                    <SheetClose asChild key={item.label}>
                                                        <Link
                                                            href={item.href}
                                                            className={cn(
                                                                'rounded-md px-3 py-2 text-sm font-medium transition-colors',
                                                                item.active
                                                                    ? 'bg-muted text-foreground'
                                                                    : 'text-muted-foreground hover:bg-muted hover:text-foreground',
                                                            )}
                                                        >
                                                            {item.label}
                                                        </Link>
                                                    </SheetClose>
                                                ))}
                                            </div>
                                        </div>

                                        <div className="space-y-2">
                                            <p className="text-xs font-semibold uppercase text-muted-foreground">
                                                Account
                                            </p>
                                            <Separator />
                                            <div className="flex flex-col gap-1">
                                                <SheetClose asChild>
                                                    <Link
                                                        href={route('profile.edit')}
                                                        className="rounded-md px-3 py-2 text-sm text-muted-foreground transition hover:bg-muted hover:text-foreground"
                                                    >
                                                        Profile
                                                    </Link>
                                                </SheetClose>
                                                <SheetClose asChild>
                                                    <Link
                                                        href={route('logout')}
                                                        method="post"
                                                        as="button"
                                                        className="rounded-md px-3 py-2 text-left text-sm text-muted-foreground transition hover:bg-muted hover:text-foreground"
                                                    >
                                                        Log Out
                                                    </Link>
                                                </SheetClose>
                                            </div>
                                        </div>
                                    </div>
                                </SheetContent>
                            </Sheet>
                        </div>
                    </div>
                </div>
            </nav>

            {header && (
                <header className="bg-card shadow">
                    <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
                        {header}
                    </div>
                </header>
            )}

            <main className="bg-muted/20">
                {children}
            </main>
        </div>
    );
}
