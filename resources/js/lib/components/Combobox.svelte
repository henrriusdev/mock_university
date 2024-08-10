<script>
  import {Check,ChevronsUpDown} from "lucide-svelte";
  import { tick } from "svelte";
  import * as Command from "$lib/components/ui/command/index.js";
  import * as Popover from "$lib/components/ui/popover/index.js";
  import { Button } from "$lib/components/ui/button/index.js";
  import { cn } from "$lib/utils.js";
 
  /** @type {Array<{ value: string; label: string }>} */
  export let options = [];
 
  let open = false;
  export let value = "";

  export let placeholder = "Select an item...";
 
  $: selectedValue =
    options.find((f) => f.value === value)?.label ?? placeholder;
 
  // We want to refocus the trigger button when the user selects
  // an item from the list so users can continue navigating the
  // rest of the form with the keyboard.
  /** @param {string} triggerId */
  function closeAndFocusTrigger(triggerId) {
    open = false;
    tick().then(() => {
      document.getElementById(triggerId)?.focus();
    });
  }
</script>
 
<Popover.Root bind:open let:ids>
  <Popover.Trigger asChild let:builder>
    <Button
      builders={[builder]}
      variant="outline"
      role="combobox"
      aria-expanded={open}
      class="w-[200px] justify-between"
    >
      {selectedValue}
      <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
    </Button>
  </Popover.Trigger>
  <Popover.Content class="w-[200px] p-0">
    <Command.Root>
      <Command.Input placeholder="Search framework..." />
      <Command.Empty>No framework found.</Command.Empty>
      <Command.Group>
        {#each options as option}
          <Command.Item
            value={option.value}
            onSelect={
            (currentValue) => {
              value = currentValue;
              closeAndFocusTrigger(ids.trigger);
            }}
          >
            <Check
              class={cn(
                "mr-2 h-4 w-4",
                value !== option.value && "text-transparent"
              )}
            />
            {option.label}
          </Command.Item>
        {/each}
      </Command.Group>
    </Command.Root>
  </Popover.Content>
</Popover.Root>