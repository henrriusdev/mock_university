<script lang="ts">
	import { enhance } from '$app/forms';
	import { DataTable, resolver } from '$lib';
	import { Button, Spinner, Tooltip } from 'flowbite-svelte';
	import type { MateriaHistorico } from '../../../../app';
	import { ArrowDownToBracketOutline } from 'flowbite-svelte-icons';

	let { data } = $props();
	let historicoMaterias = $derived(data.historico);
	let loadingPlanificacion = $state(false);
</script>

<div class="container mx-auto p-4">
	<h1 class="text-3xl font-bold mb-6">Materias cursando actualmente</h1>

	{#snippet action(row: MateriaHistorico)}
		<form
			method="POST"
			action="?/planificacion"
			use:enhance={({ formData }) => {
				formData.set('materia', row.id);
				loadingPlanificacion = true;
				return resolver(() => (loadingPlanificacion = false));
			}}
		>
			<Button
				color="primary"
				class="p-2! grid place-content-center"
				pill
				type="submit"
				disabled={loadingPlanificacion}
			>
				{#if loadingPlanificacion}
					<Spinner size="4" color="gray" />
				{:else}
					<ArrowDownToBracketOutline class="h-5 w-5" />
				{/if}
			</Button>
			<Tooltip>
				{#if loadingPlanificacion}
					Cargando...
				{:else}
					Descargar planificación
				{/if}
			</Tooltip>
		</form>
	{/snippet}
	<DataTable data={historicoMaterias} actions={action} />
</div>
