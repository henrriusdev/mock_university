<script>
    import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";

    import {Button} from "$lib/components/ui/button/index.js";
    import {Checkbox} from "$lib/components/ui/checkbox/index.js";
    import {Input} from "$lib/components/ui/input/index.js";
    import * as Card from "$lib/components/ui/card/index.js";
</script>

<DirectiveLayout title="Settings page - MockU"
                 description="Set the program settings or your user settings in this page">
    <div class="mx-auto grid w-full max-w-6xl gap-2">
        <h1 class="text-3xl font-semibold">Settings</h1>
    </div>
    <div
            class="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]"
    >
        <nav class="grid gap-4 text-sm text-muted-foreground"


             data-x-chunk-container="chunk-container after:right-0">
            <a href="##" class="font-semibold text-primary"> General </a>
            <a href="##">Notes</a>
            <a href="##">Configuration</a>
            <a href="##">Careers</a>
            <a href="##">Career's Leader</a>
            <a href="##">Advanced</a>
        </nav>
        <div class="grid gap-6">
            <Card.Root>
                <Card.Header>
                    <Card.Title>Store Name</Card.Title>
                    <Card.Description>
                        Used to identify your store in the marketplace.
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <form>
                        <Input placeholder="Store Name"/>
                    </form>
                </Card.Content>
                <Card.Footer class="border-t px-6 py-4">
                    <Button>Save</Button>
                </Card.Footer>
            </Card.Root>
            <Card.Root>
                <Card.Header>
                    <Card.Title>Plugins Directory</Card.Title>
                    <Card.Description>
                        The directory within your project, in which your plugins are located.
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <form class="flex flex-col gap-4">
                        <Input placeholder="Project Name" value="/content/plugins"/>
                        <div class="flex items-center space-x-2">
                            <Checkbox id="include" checked={true}/>
                            <label
                                    for="include"
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                            >
                                Allow administrators to change the directory.
                            </label>
                        </div>
                    </form>
                </Card.Content>
                <Card.Footer class="border-t px-6 py-4">
                    <Button>Save</Button>
                </Card.Footer>
            </Card.Root>
        </div>
    </div>
</DirectiveLayout>