<script>
  import * as Avatar from "$lib/components/ui/avatar/index.js";

  /** @type {string} */
  export let src = "";

  export let { alt } = $$props;

  /** @type {string} */
  export let name = "";

  const fallback = name
    .split(" ")
    .map((word) => word.charAt(0))
    .join("")
    .toUpperCase();
</script>

<div class="flex items-center justify-start gap-x-1">
  <Avatar.Root>
    <Avatar.Image {src} {alt} />
    <Avatar.Fallback>{fallback}</Avatar.Fallback>
  </Avatar.Root>
  <span class="ml-2">{name}</span>
</div>
