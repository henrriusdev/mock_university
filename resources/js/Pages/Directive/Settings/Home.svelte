<script>
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
  import { Button } from "$lib/components/ui/button/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import DatePicker from "$lib/components/DatePicker.svelte";
  import DateRangePicker from "$lib/components/DateRangePicker.svelte";
  import { CalendarDate } from "@internationalized/date";
  import { Label } from "$lib/components/ui/label/index.js";

  export let notesNumber = 1;
  export let paymentNumber = 1;
  export let startRegSubj = undefined;
  export let endRegSubj = undefined;
  export let cycleStart = undefined;
  export let cycleEnd = undefined;
  export let percentages = Array.from({ length: notesNumber }, () => 0);
  if (percentages == null) {
    percentages = Array.from({ length: notesNumber }, () => 0);
  }
  export let paymentDates = Array.from(
    { length: paymentNumber },
    () =>
      new CalendarDate(
        new Date().getFullYear(),
        new Date().getMonth() + 1,
        new Date().getDate(),
      ),
  );
  let actual = "#notes";

  const payments = paymentNumber;

  let paymentDatesParsed =
    paymentDates?.map((date) => {
      let dateStr = date.toString();
      let dateArr = dateStr.split("-");
      dateArr[2] = dateArr[2].split("T")[0];

      return new CalendarDate(
        parseInt(dateArr[0]),
        parseInt(dateArr[1]),
        parseInt(dateArr[2]),
      );
    }) ??
    Array.from(
      { length: paymentNumber },
      () =>
        new CalendarDate(
          new Date().getFullYear(),
          new Date().getMonth() + 1,
          new Date().getDate(),
        ),
    );

  let inputPaymentDates =
    paymentDatesParsed?.map((date) => date.toString()) ?? [];

  startRegSubj = startRegSubj ? new Date(startRegSubj) : new Date();
  endRegSubj = endRegSubj ? new Date(endRegSubj) : new Date();
  cycleStart = cycleStart ? new Date(cycleStart) : new Date();
  cycleEnd = cycleEnd ? new Date(cycleEnd) : new Date();

  let subjectInscriptionStart = startRegSubj
    ? new CalendarDate(
        startRegSubj.getFullYear(),
        startRegSubj.getMonth() + 1,
        startRegSubj.getDate(),
      )
    : new CalendarDate(2024, 3, 1);
  let subjectInscriptionEnd = endRegSubj
    ? new CalendarDate(
        endRegSubj.getFullYear(),
        endRegSubj.getMonth() + 1,
        endRegSubj.getDate(),
      )
    : new CalendarDate(2024, 3, 1);
  let cycleStartDate = cycleStart
    ? new CalendarDate(
        cycleStart.getFullYear(),
        cycleStart.getMonth() + 1,
        cycleStart.getDate(),
      )
    : new CalendarDate(2024, 3, 1);
  let cycleEndDate = cycleEnd
    ? new CalendarDate(
        cycleEnd.getFullYear(),
        cycleEnd.getMonth() + 1,
        cycleEnd.getDate(),
      )
    : new CalendarDate(2024, 3, 1);

  let subjectStart = subjectInscriptionStart.toString();
  let subjectEnd = subjectInscriptionEnd.toString();
  let startCycle = cycleStartDate.toString();
  let endCycle = cycleEndDate.toString();

  $: if (subjectInscriptionStart) {
    subjectStart = subjectInscriptionStart.toString();
  }

  $: if (subjectInscriptionEnd) {
    subjectEnd = subjectInscriptionEnd.toString();
  }

  $: if (cycleStartDate) {
    startCycle = cycleStartDate.toString();
  }

  $: if (cycleEndDate) {
    endCycle = cycleEndDate.toString();
  }

  $: notesNumber = notesNumber > 10 ? 10 : notesNumber;
  $: if (paymentNumber != payments) {
    paymentNumber = paymentNumber > 10 ? 10 : paymentNumber;
  }

  // check if the paymentDatesParsed is updated and then update the inputPaymentDates
  $: {
    inputPaymentDates =
      paymentDatesParsed?.map((date) => date.toString()) ?? [];
  }
</script>

<DirectiveLayout
  title="Configuration page - MockU"
  description="Set the web system settings like cycle, notes, percentages, and more on this page"
  keywords="configuration, settings, mocku"
>
  <div class="mx-auto grid w-full max-w-6xl gap-2">
    <h1 class="text-3xl font-semibold">Configuration</h1>
  </div>
  <div
    class="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]"
  >
    <nav
      class="grid gap-4 text-sm text-muted-foreground"
      data-x-chunk-container="chunk-container after:right-0"
    >
      <a
        href="#notes"
        on:click={() => (actual = "#notes")}
        class={actual === "#notes" ? "font-semibold text-primary" : ""}>Notes</a
      >
      <a
        href="#cycle"
        on:click={() => (actual = "#cycle")}
        class={actual === "#cycle" ? "font-semibold text-primary" : ""}>Cycle</a
      >
      <a
        href="#paids"
        on:click={() => (actual = "#paids")}
        class={actual === "#paids" ? "font-semibold text-primary" : ""}
        >Payments</a
      >
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
              <Input
                placeholder="Notes quantity"
                name="notes"
                type="number"
                min={1}
                max={10}
                bind:value={notesNumber}
              />
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
            <form
              class="flex flex-col gap-4"
              method="post"
              action="/settings/notes/percentages"
            >
              {#each percentages as percentage, i}
                <Input
                  placeholder="Note {i + 1}"
                  type="number"
                  name="note-{i + 1}"
                  min="1"
                  max="100"
                  value={percentage !== 0 ? percentage * 100 : percentage}
                />
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
              Set the start and end date for the subject inscription and the
              cycle dates for the semester
            </Card.Description>
          </Card.Header>
          <Card.Content>
            <form
              method="post"
              action="/settings/dates"
              class="flex flex-col gap-2"
            >
              <Label class="my-4">Subject Inscription</Label>
              <DateRangePicker
                bind:startValue={subjectInscriptionStart}
                bind:endValue={subjectInscriptionEnd}
                placeholder="Subject Inscription start and end dates"
              />
              <Label class="my-4">Cycle Start and End</Label>
              <DateRangePicker
                bind:startValue={cycleStartDate}
                bind:endValue={cycleEndDate}
                placeholder="Cycle start and end dates"
              />
              <input
                type="hidden"
                name="start_registration_subjects"
                bind:value={subjectStart}
              />
              <input
                type="hidden"
                name="end_registration_subjects"
                bind:value={subjectEnd}
              />
              <input type="hidden" name="cycle_start" bind:value={startCycle} />
              <input type="hidden" name="cycle_end" bind:value={endCycle} />
              <Button type="submit" class="w-fit m-2 px-4">Save</Button>
            </form>
          </Card.Content>
        </Card.Root>
        <Card.Root>
          <Card.Header>
            <Card.Title class="text-red-500">Start new cycle</Card.Title>
            <Card.Description class="text-destructive-foreground">
              Start a new cycle for the semester. This will change all the
              configurations for the new cycle to default values. Be careful
              with this action, it <span class="text-red-500"
                >does not have a confirmation</span
              >.
            </Card.Description>
          </Card.Header>
          <Card.Content>
            <form method="post" action="/settings/cycle">
              <Button variant="destructive" type="submit" class="w-fit m-2 px-4"
                >Start</Button
              >
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
              <Input
                placeholder="Payments amount"
                name="payments"
                type="number"
                min={1}
                max={10}
                bind:value={paymentNumber}
              />
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
            <form
              method="post"
              action="/settings/payment/dates"
              class="flex flex-col gap-4"
            >
              {#each paymentDatesParsed as paymentDate, i}
                <DatePicker bind:value={paymentDate} />
              {/each}
              {#each inputPaymentDates as paymentDate, i}
                <input
                  type="hidden"
                  name="payment-{i + 1}"
                  bind:value={paymentDate}
                />
              {/each}
              <Button type="submit" class="w-fit m-2 px-4">Save</Button>
            </form>
          </Card.Content>
        </Card.Root>
      </div>
    {/if}
  </div>
</DirectiveLayout>
