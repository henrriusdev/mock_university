<script lang="ts">
	import { Input, Button, Label, Select, Card, Datepicker } from 'flowbite-svelte';
	import { format, parse } from 'date-fns';
	import { toZonedTime } from 'date-fns-tz';
	import { enhance } from '$app/forms';
	import { resolver } from '$lib/utilidades/resolver.js';

	let { data } = $props();

	// Datos de configuración existentes (si los hay)
	const config = data.configuracion;

	let cicloYear = $derived(config?.ciclo?.split('-')[0] ?? new Date().getFullYear());
	let cicloPeriodo = $derived(config?.ciclo?.split('-')[1] ?? '1');
	let num_porcentaje = $state(config?.num_porcentaje ?? 3);
	let num_cuotas = $state(config?.num_cuotas ?? 5);
	let horario = $state({
		from: config?.horario_inicio
			? toZonedTime(config.horario_inicio, 'UTC')
			: toZonedTime(parse('2025-04-15', 'yyyy-MM-dd', new Date()), 'UTC'),
		to: config?.horario_fin
			? toZonedTime(config.horario_fin, 'UTC')
			: toZonedTime(parse('2025-05-15', 'yyyy-MM-dd', new Date()), 'UTC')
	});
	let porcentajes = $state<number[]>(config?.porcentajes ?? [30, 30, 40]);
	let cuotas = $state<Date[]>(
		config?.cuotas
			? [...config.cuotas.map((cuota) => toZonedTime(cuota, 'UTC'))]
			: Array(5).fill(new Date())
	);

	function updatePorcentaje(index: number, value: number) {
		porcentajes[index] = value;
	}

	function getValue(date: Date) {
		if (!date || date.toString() === 'Invalid Date') return '';
		return format(date, 'dd/MM/yyyy');
	}
</script>

<form
	use:enhance={() => {
		return resolver(()=>{})
	}}
	method="post"
	class="flex h-full w-full justify-center items-center py-6"
>
	<Card class="w-full max-w-4xl shadow-sm border border-gray-200 p-4" size="lg">
		<div class="border-b border-gray-200 pb-3 mb-6">
			<h1 class="text-xl font-bold text-primary-700">Configuración General</h1>
			<p class="text-sm text-gray-500 mt-1">Configure los parámetros del sistema académico</p>
		</div>

		<!-- Sección: Configuración Básica -->
		<div class="mb-8">
			<div class="flex items-center gap-2 mb-4">
				<span class="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded-full">1</span>
				<h2 class="text-base font-semibold text-blue-800">Configuración Básica</h2>
			</div>
			
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-6 gap-x-6 gap-y-4">
				<!-- Ciclo Académico -->
				<div class="col-span-2">
					<Label class="font-medium mb-1.5 block">Ciclo Académico</Label>
					<div class="flex gap-2 items-center">
						<Select
							name="cicloYear"
							placeholder="Seleccionar"
							bind:value={cicloYear}
							class="flex-1"
							items={Array.from({ length: 5 }, (_, i) => ({
								value: String(new Date().getFullYear() + i),
								name: String(new Date().getFullYear() + i)
							}))}
						/>
						<span class="font-bold text-gray-700">-</span>
						<Select
							bind:value={cicloPeriodo}
							placeholder="Seleccionar"
							name="cicloPeriodo"
							class="w-24"
							items={[
								{ value: '1', name: '1' },
								{ value: '2', name: '2' },
								{ value: '3', name: '3' }
							]}
						/>
					</div>
					<p class="text-xs text-gray-500 mt-1">Año y periodo académico</p>
				</div>

				<!-- Número de cortes de nota -->
				<div>
					<Label for="num_porcentaje" class="font-medium mb-1.5 block">Cortes de nota</Label>
					<Input
						id="num_porcentaje"
						name="num_porcentaje"
						type="number"
						min={1}
						max={10}
						bind:value={num_porcentaje}
					/>
					<p class="text-xs text-gray-500 mt-1">Cantidad de evaluaciones</p>
				</div>

				<!-- Número de cuotas -->
				<div>
					<Label for="num_cuotas" class="font-medium mb-1.5 block">Número de cuotas</Label>
					<Input 
						type="number" 
						id="num_cuotas" 
						name="num_cuotas" 
						min={1} 
						max={12} 
						bind:value={num_cuotas} 
					/>
					<p class="text-xs text-gray-500 mt-1">Pagos durante el ciclo</p>
				</div>

				<!-- Fechas de horario -->
				<div class="col-span-2">
					<Label class="font-medium mb-1.5 block">Periodo de horarios</Label>
					<Datepicker 
						range={true} 
						bind:rangeFrom={horario.from} 
						bind:rangeTo={horario.to} 
						placeholder="Seleccionar periodo" 
						translationLocale="es-VE" locale="fr-FR"
						dateFormat={{ year: 'numeric', month: '2-digit', day: '2-digit' }}
					/>
					<p class="text-xs text-gray-500 mt-1">Periodo para configurar horarios</p>
					{#if horario.from && horario.to && horario.from.toString() !== 'Invalid Date' && horario.to.toString() !== 'Invalid Date'}
						<input
							type="hidden"
							name="horario_inicio"
							value={horario.from && horario.from.toUTCString()}
						/>
						<input type="hidden" name="horario_fin" value={horario.to && horario.to.toUTCString()} />
					{/if}
				</div>
			</div>
		</div>

		<hr class="my-6 border-gray-200" />

		<!-- Sección: Porcentajes de nota -->
		<div class="mb-8">
			<div class="flex items-center gap-2 mb-4">
				<span class="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded-full">2</span>
				<h2 class="text-base font-semibold text-green-800">Porcentajes de nota</h2>
			</div>
			
			<div class="mb-4">
				<p class="text-xs text-gray-600 mb-3">Entre todos los porcentajes se debe sumar 100%</p>
				
				<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
					{#each Array(num_porcentaje) as _, i}
						<div>
							<Label for={`porcentaje-${i}`} class="font-medium mb-1.5 block">Corte {i + 1}</Label>
							<div class="relative">
								<Input 
									type="number" 
									id={`porcentaje-${i}`}
									min={0}
									max={100}
									inputmode="numeric"
									name={`porcentajes[${i}]`}
									bind:value={porcentajes[i]}
									oninput={(e) => {
										if (e.target instanceof HTMLInputElement) {
											updatePorcentaje(i, +e.target.value);
										}
									}}
									class="pr-8"
								/>
								<span class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500">%</span>
							</div>
						</div>
					{/each}
				</div>
			</div>
		</div>

		<hr class="my-6 border-gray-200" />

		<!-- Sección: Fechas de cuotas -->
		<div class="mb-8">
			<div class="flex items-center gap-2 mb-4">
				<span class="bg-purple-100 text-purple-800 text-xs font-medium px-2.5 py-0.5 rounded-full">3</span>
				<h2 class="text-base font-semibold text-purple-800">Fechas de cuotas</h2>
			</div>
			
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
				{#each Array(num_cuotas) as _, i}
					<div>
						<Label class="font-medium mb-1.5 block">Cuota {i + 1}</Label>
						<Datepicker bind:value={cuotas[i]} placeholder="Seleccione una fecha" translationLocale="es-VE" locale="fr-FR" dateFormat={{ year: 'numeric', month: '2-digit', day: '2-digit' }}/>
						<input type="hidden" name={`cuotas`} value={getValue(cuotas[i])} />
					</div>
				{/each}
			</div>
		</div>

		<div class="border-t border-gray-200 pt-4 mt-6 flex justify-end">
			<Button color="primary" size="md" type="submit" class="px-6 py-2 shadow-sm hover:shadow-md transition-all">
				Guardar Configuración
			</Button>
		</div>
	</Card>
</form>
