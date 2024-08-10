<script >
  import { onMount, onDestroy, tick } from 'svelte';
  import {imask as action} from '@imask/svelte';
  import IMask from 'imask';
	import { cn } from "$lib/utils.js";

  /** @type {HTMLInputElement} */
  let input;

  /** @type {string} */
  export let value;

  /** @type {import('imask').InputMask<Object> | undefined} */
  let maskRef;

  /** @type {Object} */
  let imask;

  /** @type {string} */
  let unmask;

  /** @type {any} */
  let attrs;


  $: {
    ({ imask, unmask, ...attrs } = $$props);

    if (maskRef) {
      value = value.toString();
      writeValue(value);
      attrs.value = maskRef.value;
      tick().then(() => value = getValue());
    }
  }

  /**
   * Get the current value from the mask.
   * @returns {string} The current value.
   */
  function getValue() {
    if (unmask === 'typed') return maskRef ?  maskRef.typedValue.toString() : '';
    if (unmask) return maskRef ? maskRef.unmaskedValue : '';
    return maskRef ? maskRef.value : '';
  }

  /**
   * Set the value in the mask.
   * @param {string} v - The value to set.
   */
  function setValue(v) {
    v = v == null ? '' : v;
    if (unmask === 'typed' && maskRef) maskRef.typedValue = Number(v);
    else if (unmask && maskRef) maskRef.unmaskedValue = v;
    else if (maskRef) maskRef.value = v;
  }

  /**
   * Write the value to the mask if it is different.
   * @param {string} v - The value to write.
   */
  function writeValue(v) {
    if (getValue() !== v ||
      (typeof v !== 'string' && value === '') && maskRef && !maskRef.el.isActive
    ) {
      setValue(v);
    }
  }

  onMount(() => {
    maskRef = IMask(input, imask);
    value = value.toString();
    setValue(value);
  });

  onDestroy(() => {
    if (maskRef) maskRef.destroy();
    maskRef = undefined;
  });

  /**
   * Event handler for the accept event.
   * @param {any} event - The event object.
   */
  function accept({ detail: mask }) {
    value = getValue();
  }
</script>
<input type="text"
  bind:this={input}
  use:action={maskRef}
  {...attrs}
  on:accept={accept}
  on:accept
  on:complete
  class={cn(
		"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
		attrs
	)}
  value={value}
>
