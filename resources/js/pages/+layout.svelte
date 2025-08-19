<script lang="ts">
	import ToastContainer from '$lib/componentes/ToastContainer.svelte';
	import '../app.css';
	let { children } = $props();
</script>

<div class="relative">
	<ToastContainer position="bottom-right" />
	{@render children()}
</div>
