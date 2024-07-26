<script>
    import CalendarIcon from "lucide-svelte/icons/calendar";
    import {
        CalendarDate,
        DateFormatter,
        getLocalTimeZone
    } from "@internationalized/date";
    import { cn } from "$lib/utils.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { RangeCalendar } from "$lib/components/ui/range-calendar/index.js";
    import * as Popover from "$lib/components/ui/popover/index.js";

    const df = new DateFormatter("en-US", {
        dateStyle: "medium"
    });

    export let startValue = new CalendarDate(new Date());
    export let endValue = new CalendarDate(new Date()).add({ days: 50 });
    export let placeholder = "Pick a date";

    let value = { start: startValue, end: endValue };

    $: endValue = value.end;
    $: startValue = value.start;

</script>

<div class="grid gap-2">
    <Popover.Root openFocus>
        <Popover.Trigger asChild let:builder>
            <Button
                    variant="outline"
                    class={cn(
          "w-[300px] justify-start text-left font-normal",
          !value && "text-muted-foreground"
        )}
                    builders={[builder]}
            >
                <CalendarIcon class="mr-2 h-4 w-4" />
                {#if startValue}
                    {#if endValue}
                        {df.format(startValue.toDate(getLocalTimeZone()))} - {df.format(
                        endValue.toDate(getLocalTimeZone())
                    )}
                    {:else}
                        {df.format(startValue.toDate(getLocalTimeZone()))}
                    {/if}
                {:else}
                    {placeholder}
                {/if}
            </Button>
        </Popover.Trigger>
        <Popover.Content class="w-auto p-0" align="start">
            <RangeCalendar
                    bind:value
                    bind:startValue
                    initialFocus
                    numberOfMonths={2}
                    placeholder={value?.start}
            />
        </Popover.Content>
    </Popover.Root>
</div>