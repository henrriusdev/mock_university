<script>
    import { Check, ChevronsUpDown } from "lucide-svelte";
    import { tick } from "svelte";
    import * as Command from "$lib/components/ui/command/index.js";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { cn } from "$lib/utils.js";

    /** @type {Array<{ value: string; label: string }>} */
    export let options = [];

    /** @type {boolean} */
    export let multiple = false;
    // Usar un arreglo para almacenar los valores seleccionados
    /** @type {string | string[]} */
    export let value = multiple ? [] : "";

    export let placeholder = "Select an item...";

    export let { class: className = "" } = $$props;

    const MAX_DISPLAY_ITEMS = 1;
    let open = false;
    let triggerId = "popover-trigger-id";
    $: selectedValue = multiple
        ? (() => {
            const selectedLabels = options
                .filter((option) => value.includes(option.value))
                .map((option) => option.label);

            if (selectedLabels.length > MAX_DISPLAY_ITEMS) {
                return `${selectedLabels.slice(0, MAX_DISPLAY_ITEMS).join(', ')}...`;
            }
            return selectedLabels.join(', ') || placeholder;
        })()
        : options.find((f) => f.value === value)?.label ?? placeholder;

    function closeAndFocusTrigger() {
        open = false;
        tick().then(() => {
            document.getElementById(triggerId)?.focus();
        });
    }

    /** @param {string} currentValue */
    function handleSelect(currentValue) {
        if (multiple) {
            // Si es múltiple, agregamos o quitamos el valor seleccionado del arreglo
            if (value.includes(currentValue)) {
                // @ts-ignore
                value = value.filter(/** @param {string} val*/ (val) => val !== currentValue);
            } else {
                value = [...value, currentValue];
            }
        } else {
            value = currentValue;
            closeAndFocusTrigger();
        }
    }
</script>

<Popover.Root bind:open>
  <Popover.Trigger asChild let:builder>
    <Button
            builders="{[builder]}"
            variant="outline"
            id="{triggerId}"
            role="combobox"
            aria-expanded="{open}"
            class="w-[200px] justify-between {className}"
    >
      {selectedValue}
      <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
    </Button>
  </Popover.Trigger>
  <Popover.Content class="w-[200px] p-0">
    <Command.Root>
      <Command.Input {placeholder} />
      <Command.Empty>No item found.</Command.Empty>
      <Command.Group>
        {#each options as option}
          <Command.Item
                  value="{option.value}"
                  onSelect="{() => handleSelect(option.value)}"
          >
            <Check
                    class="{cn(
                'mr-2 h-4 w-4',
                multiple
                  ? !value.includes(option.value) && 'text-transparent'
                  : value !== option.value && 'text-transparent'
              )}"
            />
            {option.label}
          </Command.Item>
        {/each}
      </Command.Group>
    </Command.Root>
  </Popover.Content>
</Popover.Root>
