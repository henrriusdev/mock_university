<script>
  import { Checkbox } from "$lib/components/ui/checkbox";

  /** @type {import('svelte/store').Writable<boolean>} */
  export let checked;
</script>

<Checkbox bind:checked="{$checked}" />
