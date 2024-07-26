<script>
    import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
    import {Button} from "$lib/components/ui/button/index.js";
    import {Input} from "$lib/components/ui/input/index.js";
    import * as Card from "$lib/components/ui/card/index.js";
    import DatePicker from "$lib/components/DatePicker.svelte";
    import DateRangePicker from "$lib/components/DateRangePicker.svelte";
    import {CalendarDate} from "@internationalized/date";
    import {Label} from "$lib/components/ui/label/index.js";

    export let notesNumber = 1;
    export let paymentNumber = 1;

    let actual = "#notes";

    let paymentDates = Array.from({length: paymentNumber}, () => new CalendarDate(new Date().getFullYear(), new Date().getMonth(), new Date().getDate()));

    let subjectInscriptionStart = new CalendarDate(2024, 1, 1);
    let subjectInscriptionEnd = new CalendarDate(2024, 4, 1);
    let cycleStart = new CalendarDate(2024, 4, 1);
    let cycleEnd = new CalendarDate(2024, 7, 1);

    let subjectStart = subjectInscriptionStart.toString();
    let subjectEnd = subjectInscriptionEnd.toString();
    let startCycle = cycleStart.toString();
    let endCycle = cycleEnd.toString();

    $: if(subjectInscriptionStart) {
        subjectStart = subjectInscriptionStart.toString();
    }

    $: if(subjectInscriptionEnd) {
        subjectEnd = subjectInscriptionEnd.toString();
    }

    $: if(cycleStart) {
        startCycle = cycleStart.toString();
    }

    $: if(cycleEnd) {
        endCycle = cycleEnd.toString();
    }

    $: notesNumber = notesNumber > 10 ? 10 : notesNumber;
    $: if(paymentNumber) {
        paymentNumber = paymentNumber > 10 ? 10 : paymentNumber;
        paymentDates = Array.from({length: paymentNumber}, () => new CalendarDate(new Date().getFullYear(), new Date().getMonth(), new Date().getDate()));
    }
</script>

<DirectiveLayout title="Configuration page - MockU"
                 description="Set the web system settings like cycle, notes, percentages, and more on this page" keywords="configuration, settings, mocku">
    <div class="mx-auto grid w-full max-w-6xl gap-2">
        <h1 class="text-3xl font-semibold">Configuration</h1>
    </div>
    <div
            class="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]"
    >
        <nav class="grid gap-4 text-sm text-muted-foreground"
             data-x-chunk-container="chunk-container after:right-0">
            <a href="#notes" on:click={() => actual = "#notes"} class="{actual === '#notes' ? 'font-semibold text-primary' : '' }">Notes</a>
            <a href="#cycle" on:click={() => actual = "#cycle"} class="{actual === '#cycle' ? 'font-semibold text-primary' : '' }">Cycle</a>
            <a href="#paids" on:click={() => actual = "#paids"} class="{actual === '#paids' ? 'font-semibold text-primary' : '' }">Payments</a>
        </nav>
        {#if actual === "#notes"}
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
                        <Input placeholder="Notes quantity" name="notes" type="number" min="{1}" max="{10}" bind:value={notesNumber}/>
                    <Button type="submit" class="w-fit m-2 px-4">Save</Button>
                    </form>
                </Card.Content>
            </Card.Root>
            <Card.Root>
                <Card.Header>
                    <Card.Title>Notes percentage</Card.Title>
                    <Card.Description>
                        The percentages of each note of the semester
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <form class="flex flex-col gap-4" method="post" action="/settings/notes/percentages">
                        {#each Array.from({length: notesNumber}) as _, i}
                            <Input placeholder="Note {i + 1}" type="number" name="note-{i + 1}"/>
                        {/each}
                        <Button type="submit" class="w-fit m-2 px-4">Save</Button>
                    </form>
                </Card.Content>
            </Card.Root>
        </div>
        {:else if actual === "#cycle"}
        <div class="grid gap-6" id="cycle">
            <Card.Root>
                <Card.Header>
                    <Card.Title>Subject inscription And Cycle Dates</Card.Title>
                    <Card.Description>
                        Set the start and end date for the subject inscription and the cycle dates for the semester
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <form method="post" action="/settings/dates" class="flex flex-col gap-2">
                        <Label>Subject Inscription</Label>
                        <DateRangePicker bind:startValue={subjectInscriptionStart} bind:endValue={subjectInscriptionEnd} placeholder="Subject Inscription start and end dates"/>
                        <Label>Subject Inscription</Label>
                        <DateRangePicker bind:startValue={cycleStart} bind:endValue={cycleEnd} placeholder="Cycle start and end dates"/>
                        <input type="hidden" name="subjectInscriptionStart" bind:value={subjectStart} />
                        <input type="hidden" name="subjectInscriptionEnd" bind:value={subjectEnd} />
                        <input type="hidden" name="cycleStart" bind:value={startCycle} />
                        <input type="hidden" name="cycleEnd" bind:value={endCycle} />
                        <Button type="submit" class="w-fit m-2 px-4">Save</Button>
                    </form>
                </Card.Content>
            </Card.Root>
        </div>
        {:else if actual === "#paids"}
        <div class="grid gap-6" id="paids">
            <Card.Root>
                <Card.Header>
                    <Card.Title>Payments Amount</Card.Title>
                    <Card.Description>
                        Set the amount of payments for the cycle
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <form method="post" action="/settings/payment">
                        <Input placeholder="Payments amount" name="payments" type="number" min="{1}" max="{10}" bind:value={paymentNumber}/>
                        <Button type="submit" class="w-fit m-2 px-4">Save</Button>
                    </form>
                </Card.Content>
            </Card.Root>
            <Card.Root>
                <Card.Header>
                    <Card.Title>Payment dates</Card.Title>
                    <Card.Description>
                        Set the payment dates for the cycle
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <form method="post" action="/settings/payments" class="flex flex-col gap-4">
                        {#each paymentDates as paymentDate}
                            <DatePicker bind:value={paymentDate} />
                        {/each}
                        <Button type="submit" class="w-fit m-2 px-4">Save</Button>
                    </form>
                </Card.Content>
            </Card.Root>
        </div>
        {/if}
    </div>
</DirectiveLayout>