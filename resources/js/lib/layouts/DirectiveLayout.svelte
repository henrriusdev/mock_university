<script>
    import CircleUser from "lucide-svelte/icons/circle-user";
    import Menu from "lucide-svelte/icons/menu";
    import Package2 from "lucide-svelte/icons/package-2";
    import Search from "lucide-svelte/icons/search";
    import {Button} from "$lib/components/ui/button/index.js";
    import {Input} from "$lib/components/ui/input/index.js";
    import LightSwitch from "$lib/components/LightSwitch.svelte";
    import {ModeWatcher} from "mode-watcher";
    import {
        DropdownMenu,
        DropdownMenuContent,
        DropdownMenuItem,
        DropdownMenuLabel,
        DropdownMenuSeparator,
        DropdownMenuTrigger
    } from "$lib/components/ui/dropdown-menu/index.js";
    import {Sheet, SheetContent, SheetTrigger} from "$lib/components/ui/sheet/index.js";

    export let title;
    export let description;
    export let keywords;

    const routes = [
        {name: "Dashboard", href: "/directive",},
        {name: "Students", href: "/directive/students",},
        {name: "Subjects", href: "/directive/subjects",},
        {name: "Professors", href: "/directive/professors",},
        {name: "Careers", href: "/directive.careers",},
        {name: "Payments", href: "/directive/payments",},
        {name: "Requests", href: "/directive/requests",},
    ];
</script>
<svelte:head>
    {#if title}
        <title>{title}</title>
    {/if}
    {#if description}
        <meta name="description" content={description}/>
    {/if}
    {#if keywords}
        <meta name="keywords" content={keywords}/>
    {/if}
</svelte:head>
<div class="flex min-h-screen w-full flex-col">
    <ModeWatcher/>
    <header class="sticky top-0 flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6">
        <nav
                class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6"
        >
            <LightSwitch/>
            {#each routes as {name, href} (href)}
                <a href={href} class="text-foreground transition-colors hover:text-foreground">
                    {name}
                </a>
            {/each}
        </nav>
        <Sheet>
            <SheetTrigger asChild let:builder>
                <Button
                        variant="outline"
                        size="icon"
                        class="shrink-0 md:hidden"
                        builders={[builder]}
                >
                    <Menu class="h-5 w-5"/>
                    <span class="sr-only">Toggle navigation menu</span>
                </Button>
            </SheetTrigger>
            <SheetContent side="left">
                <nav class="grid gap-6 text-lg font-medium">
                    <a href="##" class="flex items-center gap-2 text-lg font-semibold">
                        <Package2 class="h-6 w-6"/>
                        <span class="sr-only">Acme Inc</span>
                    </a>
                    {#each routes as {name, href} (href)}
                        <a href={href} class="text-foreground transition-colors hover:text-foreground">
                            {name}
                        </a>
                    {/each}
                </nav>
            </SheetContent>
        </Sheet>
        <div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
            <form class="ml-auto flex-1 sm:flex-initial">
                <div class="relative">
                    <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground"/>
                    <Input
                            type="search"
                            placeholder="Search students..."
                            class="pl-8 sm:w-[250px] md:w-[150px] lg:w-[250px]"
                    />
                </div>
            </form>
            <DropdownMenu>
                <DropdownMenuTrigger asChild let:builder>
                    <Button
                            builders={[builder]}
                            variant="secondary"
                            size="icon"
                            class="rounded-full"
                    >
                        <CircleUser class="h-5 w-5"/>
                        <span class="sr-only">Toggle user menu</span>
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                    <DropdownMenuItem inset>
                        <a href="/settings">
                            Settings
                        </a>
                    </DropdownMenuItem>
                    <DropdownMenuSeparator/>
                    <DropdownMenuLabel>My Account</DropdownMenuLabel>
                    <DropdownMenuItem>Logout</DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    </header>
    <main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
        <slot/>
    </main>
</div>
