<script>
    import {Menubar, Trigger, Content, Menu, Item, Separator, Sub, SubContent, SubTrigger} from '$lib/components/ui/menubar/index.js';
    import LightSwitch from "$lib/components/LightSwitch.svelte";
    import {ChevronRight} from "lucide-svelte";
    export let careers = [];

</script>

<div class="flex w-full justify-between items-center py-1 px-4">
    <h1 class="text-2xl font-bold">
        <a href="/" class="">Mock University</a>
    </h1>
    <Menubar>
        <Menu>
            <Trigger>Careers</Trigger>
            <Content>
                {#if careers.length === 0}
                    <Item>No careers available</Item>
                {:else}
                    {#each careers as career}
                        <Item>{career.name}</Item>
                    {/each}
                {/if}
            </Content>
        </Menu>
        <Menu>
            <Trigger>About</Trigger>
            <Content>
                <Item>
                    About <span class="text-bold">MockU</span>
                </Item>
                <Item>
                    Mission & Vision
                </Item>
                <Separator/>
                <Sub>
                    <SubTrigger>More about us</SubTrigger>
                    <SubContent>
                        <Item>Our campuses</Item>
                        <Separator/>
                        <Item>The Student Community</Item>
                        <Item>The Teacher Community</Item>
                        <Item>Admissions</Item>
                    </SubContent>
                </Sub>
            </Content>
        </Menu>
        <Menu>
            <Trigger>Contact</Trigger>
            <Content>
                <Item inset>WhatsApp</Item>
                <Separator/>
                <Item inset>Facebook</Item>
                <Item inset>Twitter</Item>
                <Item inset>Instagram</Item>
                <Separator/>
                <Item inset>Contact page</Item>
            </Content>
        </Menu>
        <Menu>
            <Trigger class="bg-default">
                <a href="/login" class="text-white flex justify-center gap-x-2 items-center">
                    Login <ChevronRight class="w-4 h-4"/>
                </a>
            </Trigger>
        </Menu>
        <LightSwitch/>
    </Menubar>
</div>