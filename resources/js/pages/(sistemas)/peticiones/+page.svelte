<script lang="ts">
	import { DataTable } from '$lib';
	import { Button, Label, Select } from 'flowbite-svelte';
	import { CheckOutline, CloseOutline } from 'flowbite-svelte-icons';
	import { resolver } from '$lib/utilidades/resolver';
	import type { SubmitFunction } from '@sveltejs/kit';
	import type { Peticion } from '../../../app';
	import type { Peticion as PeticionGet } from '$lib/types';
	import { enhance } from '$app/forms';

	const handleSubmit: SubmitFunction = () => {
		return resolver(() => {});
	};

	let { data } = $props();

	let peticiones = $derived(data.peticiones);
	let estado: 'Todas' | 'Pendiente' | 'Rechazada' | 'Aprobada' = $state('Todas');
	let docente: string | undefined = $state('');
	let estudiante: string | undefined = $state('');

	let peticionesFiltradas = $derived.by(() => {
		return peticiones.filter((p) => {
			const matchEstado =
				estado === 'Todas' || p.peticion.estado?.toLowerCase() === estado.toLowerCase();
			const matchDocente = docente === undefined || docente === '' || p.docente.cedula === docente;
			const matchEstudiante =
				estudiante === undefined || estudiante === '' || p.estudiante?.cedula === estudiante;

			return matchEstado && matchDocente && matchEstudiante;
		});
	});

	const buscarValores = (key: string, value: any) => {
		switch (key) {
			case 'docente':
				return value.nombre;
			case 'estudiante':
				return value.nombre;
			case 'materia':
				return value.nombre;
			default:
				return `${(value as Peticion).descripcion}; corte: ${(value as Peticion).campo}, nota: ${(value as Peticion).valor}`;
		}
	};
</script>

<div class="container mx-auto">
	<h2 class="md:text-2xl text-xl font-bold mb-6">Peticiones de cambio de notas</h2>
	{#snippet actions(row: PeticionGet)}
		{#if ['control', 'administrador'].includes(data.rol) && row.peticion.estado === 'Pendiente'}
			<div class="flex gap-2">
				<form method="post" use:enhance={handleSubmit} action="?/aprobar">
					<input type="hidden" name="id" value={row.peticion.id} />
					<Button type="submit" pill color="green" outline class="p-1!" size="sm">
						<CheckOutline class="w-5 h-5" />
					</Button>
				</form>
				<form method="post" use:enhance={handleSubmit} action="?/rechazar">
					<input type="hidden" name="id" value={row.peticion.id} />
					<Button type="submit" pill color="red" outline class="p-1!" size="sm">
						<CloseOutline class="w-5 h-5" />
					</Button>
				</form>
			</div>
		{/if}
	{/snippet}

	<div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-4">
		<div>
			<Label>Estado</Label>
			<Select
				placeholder="Seleccionar"
				bind:value={estado}
				items={[
					{ value: 'Todas', name: 'Todas' },
					{ value: 'Aprobado', name: 'Aprobadas' },
					{ value: 'Pendiente', name: 'Pendientes' },
					{ value: 'Denegado', name: 'Denegadas' }
				]}
			/>
		</div>
		<div>
			<Label>Docente</Label>
			<Select
				placeholder="Seleccionar"
				bind:value={docente}
				items={[
					{ value: '', name: 'Todos' },
					...data.docentes.map((d) => ({ value: d.cedula, name: d.nombre }))
				]}
			/>
		</div>
		<div>
			<Label>Estudiante</Label>
			<Select
				placeholder="Seleccionar"
				bind:value={estudiante}
				items={[
					{ value: '', name: 'Todos' },
					...data.estudiantes.map((d) => ({ value: d.cedula, name: d?.nombre ?? '' }))
				]}
			/>
		</div>
	</div>

	<DataTable data={peticionesFiltradas} {actions} onSearch={buscarValores} />
</div>
