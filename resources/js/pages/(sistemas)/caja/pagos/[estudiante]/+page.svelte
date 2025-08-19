<script lang="ts">
	import type { Pago } from '$lib';
	import DataTable from '$lib/componentes/DataTable.svelte';
	import { Button, Modal } from 'flowbite-svelte';
	import { EyeOutline } from 'flowbite-svelte-icons';
	import type { PageData } from './$types';

	let { data } = $props<{ data: PageData }>();
	const { nombre, pagos } = data;

	let selectedPago: Pago | null = $state(null);
	let open = $state(false);

	function mostrarPago(pago: Pago) {
		selectedPago = pago;
		open = true;
	}

	function onSearch() {
		return "N/A";
	}
</script>

<h2 class="text-xl font-bold text-center mt-6 mb-4">
	Pagos de <span class="text-blue-600 underline cursor-pointer">{nombre}</span>
</h2>

{#if pagos.length > 0}
	{#snippet actions(row: Pago)}
		<div class="flex gap-2">
			<Button pill size="xs" color="light" class="p-1!" onclick={() => mostrarPago(row)}>
				<EyeOutline class="w-5 h-5" />
			</Button>
		</div>
	{/snippet}
	<DataTable data={pagos} {actions} onSearch={onSearch} />
{:else}
	<p class="text-center">No hay pagos registrados</p>
{/if}

<Modal bind:open size="sm">
	{#snippet header()}
		<div class="text-lg font-semibold text-gray-900">Información del Pago</div>
	{/snippet}

	{#if selectedPago}
		<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm text-gray-700">
			<div>
				<p><span class="font-medium">Ciclo:</span> {selectedPago.ciclo}</p>
				<p><span class="font-medium">Fecha de pago:</span> {selectedPago.fecha}</p>
				<p><span class="font-medium">Monto:</span> Bs. {selectedPago.monto}</p>
				{#if selectedPago.referencia}
					<p><span class="font-medium">Nro. de referencia:</span> {selectedPago.referencia}</p>
				{/if}
			</div>
			<div>
				<p><span class="font-medium">Método de pago:</span> {selectedPago.metodo}</p>
			</div>
		</div>

		{#if selectedPago.metodo === 'Efectivo' && selectedPago.billetes}
			<div class="mt-4">
				<h3 class="text-sm font-semibold mb-2">Billetes asociados</h3>
				<ul class="list-disc list-inside text-sm">
					{#each selectedPago.billetes as billete}
						<li>{billete.denominacion} $. × {billete.cantidad}</li>
					{/each}
				</ul>
			</div>
		{/if}
	{/if}
</Modal>
