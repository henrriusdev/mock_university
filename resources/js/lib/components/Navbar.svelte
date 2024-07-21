<script>

    export let careers = [];

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
                    About <span class="text-bold">MockU</span>
                </Menubar.Item>
                <Menubar.Item>
                    Mission & Vision
                </Menubar.Item>
                <Menubar.Separator/>
                <Menubar.Sub>
                    <Menubar.SubTrigger>More about us</Menubar.SubTrigger>
                    <Menubar.SubContent>
                        <Menubar.Item>Our campuses</Menubar.Item>
                        <Menubar.Separator/>
                        <Menubar.Item>The Student Community</Menubar.Item>
                        <Menubar.Item>The Teacher Community</Menubar.Item>
                        <Menubar.Item>Admissions</Menubar.Item>
                    </Menubar.SubContent>
                </Menubar.Sub>
            </Menubar.Content>
        </Menubar.Menu>
        <Menubar.Menu>
            <Menubar.Trigger>Contact</Menubar.Trigger>
            <Menubar.Content>
                <Menubar.Item inset>WhatsApp</Menubar.Item>
                <Menubar.Separator/>
                <Menubar.Item inset>Facebook</Menubar.Item>
                <Menubar.Item inset>Twitter</Menubar.Item>
                <Menubar.Item inset>Instagram</Menubar.Item>
                <Menubar.Separator/>
                <Menubar.Item inset>Contact page</Menubar.Item>
            </Menubar.Content>
        </Menubar.Menu>
        <Menubar.Menu>
            <Menubar.Trigger class="bg-default">
                <a href="/login" class="text-default-foreground">Login</a>
            </Menubar.Trigger>
        </Menubar.Menu>
    </Menubar.Root>
</div>