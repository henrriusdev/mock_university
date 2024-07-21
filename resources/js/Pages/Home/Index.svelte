<script>
    import MainLayout from "$lib/layouts/MainLayout.svelte";

    export let careers = [];
</script>

<MainLayout title="Mock University - The best Programming University" description="Mock University is the best Programming University in the world. We offer the best programming courses and tutorials."
            {careers}>
    <div class="container mx-auto">
        <h1 class="text-2xl font-bold text-center mt-4">Welcome to Mock University</h1>
    </div>
</MainLayout>