<script>
  import {CalendarIcon} from "lucide-svelte";
  import { DateFormatter, getLocalTimeZone } from "@internationalized/date";
  import { cn } from "$lib/utils.js";
  import { Button } from "$lib/components/ui/button/index.js";
  import { Calendar, CalendarSelect } from "$lib/components/ui/calendar/index.js";
  import * as Popover from "$lib/components/ui/popover/index.js";

  const df = new DateFormatter("en-US", {
    dateStyle: "long",
  });

  /** @type {import('@internationalized/date').DateValue | undefined} */
  export let value = undefined;
</script>

<Popover.Root>
  <Popover.Trigger asChild let:builder>
    <Button
      variant="outline"
      class="{cn(
        'w-[280px] justify-start text-left font-normal',
        !value && 'text-muted-foreground'
      )}"
      builders="{[builder]}"
    >
      <CalendarIcon class="mr-2 h-4 w-4" />
      {value ? df.format(value.toDate(getLocalTimeZone())) : "Pick a date"}
    </Button>
  </Popover.Trigger>
  <Popover.Content class="w-auto p-0">
    <CalendarSelect bind:value initialFocus  />
  </Popover.Content>
</Popover.Root>
