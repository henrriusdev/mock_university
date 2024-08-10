<script>
  import MaskInput from "svelte-input-mask/MaskInput.svelte";
  import { cn } from "$lib/utils.js";
  let className = undefined;
  /** @type {string | undefined} */
  export let value = undefined;
  export { className as class };
  /** @type {boolean} */
  export let readonly = false;
  /** @type {string | undefined} */
  export let mask = undefined;
  /** @type {string | undefined} */
  export let maskChar = "_";
  /** @type {boolean} */
  export let alwaysShowMask = true;
</script>

<MaskInput
  {alwaysShowMask}
  {maskChar}
  {mask}
  {readonly}
  class="{cn(
    'flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50',
    className
  )}"
  {...$$restProps}
  bind:value
/>
