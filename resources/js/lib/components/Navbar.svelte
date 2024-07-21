<script>
    import * as Menubar from "$lib/components/ui/menubar";
    export let careers = [];

    let bookmarks = false;
    let fullUrls = true;

    const profileRadioValue = "benoit";
</script>

<div class="flex w-full justify-between items-center py-1 px-4">
    <h1 class="text-2xl font-bold">
        <a href="/" class="">Mock University</a>
    </h1>
    <Menubar.Root>
        <Menubar.Menu>
            <Menubar.Trigger>Careers</Menubar.Trigger>
            <Menubar.Content>
                {#if careers.length === 0}
                    <Menubar.Item>No careers available</Menubar.Item>
                {:else}
                    {#each careers as career}
                        <Menubar.Item>{career.name}</Menubar.Item>
                    {/each}
                {/if}
            </Menubar.Content>
        </Menubar.Menu>
        <Menubar.Menu>
            <Menubar.Trigger>About</Menubar.Trigger>
            <Menubar.Content>
                <Menubar.Item>
                    Undo
                    <Menubar.Shortcut>⌘Z</Menubar.Shortcut>
                </Menubar.Item>
                <Menubar.Item>
                    Redo
                    <Menubar.Shortcut>⇧⌘Z</Menubar.Shortcut>
                </Menubar.Item>
                <Menubar.Separator/>
                <Menubar.Sub>
                    <Menubar.SubTrigger>Find</Menubar.SubTrigger>
                    <Menubar.SubContent>
                        <Menubar.Item>Search the web</Menubar.Item>
                        <Menubar.Separator/>
                        <Menubar.Item>Find...</Menubar.Item>
                        <Menubar.Item>Find Next</Menubar.Item>
                        <Menubar.Item>Find Previous</Menubar.Item>
                    </Menubar.SubContent>
                </Menubar.Sub>
                <Menubar.Separator/>
                <Menubar.Item>Cut</Menubar.Item>
                <Menubar.Item>Copy</Menubar.Item>
                <Menubar.Item>Paste</Menubar.Item>
            </Menubar.Content>
        </Menubar.Menu>
        <Menubar.Menu>
            <Menubar.Trigger>Services</Menubar.Trigger>
            <Menubar.Content>
                <Menubar.CheckboxItem bind:checked={bookmarks}
                >Always Show Bookmarks Bar
                </Menubar.CheckboxItem
                >
                <Menubar.CheckboxItem bind:checked={fullUrls}>
                    Always Show Full URLs
                </Menubar.CheckboxItem>
                <Menubar.Separator/>
                <Menubar.Item inset>
                    Reload
                    <Menubar.Shortcut>⌘R</Menubar.Shortcut>
                </Menubar.Item>
                <Menubar.Item inset>
                    Force Reload
                    <Menubar.Shortcut>⇧⌘R</Menubar.Shortcut>
                </Menubar.Item>
                <Menubar.Separator/>
                <Menubar.Item inset>Toggle Fullscreen</Menubar.Item>
                <Menubar.Separator/>
                <Menubar.Item inset>Hide Sidebar</Menubar.Item>
            </Menubar.Content>
        </Menubar.Menu>
        <Menubar.Menu>
            <Menubar.Trigger>Contact</Menubar.Trigger>
            <Menubar.Content>
                <Menubar.RadioGroup value={profileRadioValue}>
                    <Menubar.RadioItem value="andy">Andy</Menubar.RadioItem>
                    <Menubar.RadioItem value="benoit">Benoit</Menubar.RadioItem>
                    <Menubar.RadioItem value="Luis">Luis</Menubar.RadioItem>
                </Menubar.RadioGroup>
                <Menubar.Separator/>
                <Menubar.Item inset>Edit...</Menubar.Item>
                <Menubar.Separator/>
                <Menubar.Item inset>Add Profile...</Menubar.Item>
            </Menubar.Content>
        </Menubar.Menu>
    </Menubar.Root>
</div>