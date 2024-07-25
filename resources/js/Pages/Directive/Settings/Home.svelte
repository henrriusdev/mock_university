<script>
    import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";

    import {Button} from "$lib/components/ui/button/index.js";
    import {Checkbox} from "$lib/components/ui/checkbox/index.js";
    import {Input} from "$lib/components/ui/input/index.js";
    import * as Card from "$lib/components/ui/card/index.js";

    let actual = "#notes";
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
            <a href="#notes" on:click={() => actual = "#notes"} class="{actual === '#notes' ? 'font-semibold text-primary' : '' }">Notes</a>
            <a href="#config" on:click={() => actual = "#config"} class="{actual === '#config' ? 'font-semibold text-primary' : '' }">Configuration</a>
            <a href="#careers" on:click={() => actual = "#careers"} class="{actual === '#careers' ? 'font-semibold text-primary' : '' }">Careers</a>
        </nav>
        <div class="grid gap-6" id="notes">
            <Card.Root>
                <Card.Header>
                    <Card.Title>Notes Quantity</Card.Title>
                    <Card.Description>
                        Used for the length of notes in one semester
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <form method="post" action="/settings/notes">
                        <Input placeholder="Notes quantity" name="notes" type="number"/>
                    </form>
                </Card.Content>
                <Card.Footer class="border-t px-6 py-4">
                    <Button>Save</Button>
                </Card.Footer>
            </Card.Root>
            <Card.Root>
                <Card.Header>
                    <Card.Title>Notes percentage</Card.Title>
                    <Card.Description>
                        The percentages of each note of the semester
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <form class="flex flex-col gap-4">
                        <Input placeholder="Note 1"/>
                        <Input placeholder="Note 2"/>
                        <Input placeholder="Note 3"/>
                        <Input placeholder="Note 4"/>
                    </form>
                </Card.Content>
                <Card.Footer class="border-t px-6 py-4">
                    <Button>Save</Button>
                </Card.Footer>
            </Card.Root>
        </div>
    </div>
</DirectiveLayout>