<script>
  import {
    Menubar,
    MenubarTrigger,
    MenubarContent,
    MenubarMenu,
    MenubarItem,
    MenubarSeparator,
    MenubarSub,
    MenubarSubContent,
    MenubarSubTrigger,
  } from "$lib/components/ui/menubar/index.js";
  import LightSwitch from "$lib/components/LightSwitch.svelte";
  import { ChevronRight } from "lucide-svelte";
  export let careers = [];
</script>

<div class="flex w-full justify-between items-center py-1 px-4">
  <h1 class="text-2xl font-bold">
    <a href="/" class="">Mock University</a>
  </h1>
  <Menubar>
    <MenubarMenu>
      <MenubarTrigger>Careers</MenubarTrigger>
      <MenubarContent>
        {#if careers.length === 0}
          <MenubarItem>No careers available</MenubarItem>
        {:else}
          {#each careers as career}
            <MenubarItem>{career.name}</MenubarItem>
          {/each}
        {/if}
      </MenubarContent>
    </MenubarMenu>
    <MenubarMenu>
      <MenubarTrigger>About</MenubarTrigger>
      <MenubarContent>
        <MenubarItem>
          About <span class="text-bold">MockU</span>
        </MenubarItem>
        <MenubarItem>Mission & Vision</MenubarItem>
        <MenubarSeparator />
        <MenubarSub>
          <MenubarSubTrigger>More about us</MenubarSubTrigger>
          <MenubarSubContent>
            <MenubarItem>Our campuses</MenubarItem>
            <MenubarSeparator />
            <MenubarItem>The Student Community</MenubarItem>
            <MenubarItem>The Teacher Community</MenubarItem>
            <MenubarItem>Admissions</MenubarItem>
          </MenubarSubContent>
        </MenubarSub>
      </MenubarContent>
    </MenubarMenu>
    <MenubarMenu>
      <MenubarTrigger>Contact</MenubarTrigger>
      <MenubarContent>
        <MenubarItem inset>WhatsApp</MenubarItem>
        <MenubarSeparator />
        <MenubarItem inset>Facebook</MenubarItem>
        <MenubarItem inset>Twitter</MenubarItem>
        <MenubarItem inset>Instagram</MenubarItem>
        <MenubarSeparator />
        <MenubarItem inset>Contact page</MenubarItem>
      </MenubarContent>
    </MenubarMenu>
    <MenubarMenu>
      <MenubarTrigger class="bg-default">
        <a
          href="/login"
          class="text-white flex justify-center gap-x-2 items-center"
        >
          Login <ChevronRight class="w-4 h-4" />
        </a>
      </MenubarTrigger>
    </MenubarMenu>
    <LightSwitch />
  </Menubar>
</div>
