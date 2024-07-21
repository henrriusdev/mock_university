<script>
    import Navbar from "$lib/components/Navbar.svelte";
    export let title = 'Mock University - The best Programming University';
    export let description = 'Mock University is the best Programming University in the world. We offer the best programming courses and tutorials.';
    export let careers = [];
</script>
<svelte:head>
    <title>{title}</title>
    {#if description}
        <meta name="description" content="{description}">
    {/if}
</svelte:head>

<!-- The MainLayout.svelte file is a layout file that will be used to wrap all the pages in the application. -->
<div class="container-lg mx-auto">
    <Navbar {careers}/>
    <slot></slot>
</div>