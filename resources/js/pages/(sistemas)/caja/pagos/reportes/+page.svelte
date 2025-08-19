<script lang="ts">
	import { Button, Card, Heading, Label, Select, Datepicker } from 'flowbite-svelte';
	import { FileLinesOutline } from 'flowbite-svelte-icons';
	import { enhance } from '$app/forms';
	import { resolver } from '$lib/utilidades/resolver';

	let reportType = $state('dia');
	let paymentSelection = $state('todos');
	let reportDate: Date | { from?: Date; to?: Date } | undefined = $state(new Date());

	$effect(() => {
		if (reportType === 'dia') {
			reportDate = new Date();
		} else if (reportType === 'fechas') {
			reportDate = { from: new Date(), to: new Date() };
		} else {
			reportDate = undefined;
		}
	});
</script>

<div class="w-full flex justify-center mt-8">
	<Card class="p-6 w-full max-w-xl">
		<Heading tag="h2" class="text-2xl font-bold text-blue-600 mb-6">Reportes de pagos</Heading>

		<form
			method="POST"
			action="?/generarReporte"
			use:enhance={({ formData }) => {
				formData.append('tipo', reportType);
				formData.append('filtro', paymentSelection);
				if (reportType === 'dia') {
					formData.append('fecha', (reportDate as Date)?.toLocaleDateString('en-CA'));
				}
				if (reportType === 'fechas' || reportType === 'monto') {
					formData.append('fi', (reportDate as any).from?.toLocaleDateString('en-CA'));
					formData.append('ff', (reportDate as any).to?.toLocaleDateString('en-CA'));
				}
				return resolver(() => {});
			}}
			class="space-y-6"
		>
			<div>
				<Label for="report-type" class="mb-2">Seleccione su tipo de reporte</Label>
				<Select
					placeholder="Seleccionar"
					id="report-type"
					bind:value={reportType}
					items={[
						{ value: 'dia', name: 'Día' },
						{ value: 'fechas', name: 'Entre fechas especificas' },
						{ value: 'monto', name: 'Montos totales' }
					]}
				/>
			</div>

			<div>
				<Label for="payment-selection" class="mb-2">
					¿Cuáles pagos desea que estén en el reporte?
				</Label>
				<Select
					placeholder="Seleccionar"
					id="payment-selection"
					bind:value={paymentSelection}
					items={[
						{ value: 'todos', name: 'Todos' },
						{ value: 'transferencia', name: 'Transferencias' },
						{ value: 'efectivo', name: 'Efectivo' },
						{ value: 'punto', name: 'Punto de venta' }
					]}
				/>
			</div>

			<div>
				<Label for="report-date" class="mb-2">Fecha de pagos</Label>
				{#if reportType === 'dia'}
					<Datepicker
						bind:value={reportDate}
						range={reportType !== 'dia'}
						availableTo={new Date()}
						placeholder="Seleccione una fecha"
						translationLocale="es-VE"
						locale="fr-FR"
						dateFormat={{ year: 'numeric', month: '2-digit', day: '2-digit' }}
					/>
				{:else}
					<Datepicker
						bind:rangeFrom={reportDate.from}
						bind:rangeTo={reportDate.to}
						range
						availableTo={new Date()}
						placeholder="Seleccione un rango de fechas"
						translationLocale="es-VE"
						locale="fr-FR"
						dateFormat={{ year: 'numeric', month: '2-digit', day: '2-digit' }}
					/>
				{/if}
			</div>

			<Button color="primary" size="lg" class="w-full" disabled={!reportDate} type="submit">
				<FileLinesOutline class="mr-2 h-5 w-5" />
				Generar reporte
			</Button>
		</form>
	</Card>
</div>
