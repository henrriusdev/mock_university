<script lang="ts">
	import { DataTable } from '$lib';
	import ToastContainer from '$lib/componentes/ToastContainer.svelte';
	import { Button, ButtonGroup, Input, Label, Modal, TableSearch, Tooltip } from 'flowbite-svelte';
	import { PenOutline, PlusOutline } from 'flowbite-svelte-icons';
	import type { Carrera } from '../../../../app';
	import type { PageData } from './$types';

	// Datos de la página
	let { data } = $props<{
		data: PageData;
	}>();

	// Estado para el modal
	let modalVisible = $state(false);
	let isEditing = $state(false);
	let searchTerm = $state('');
	let estudianteActual: any = $state({});

	$effect(() => {
		if (!modalVisible) {
			estudianteActual = {};
		}
	});

	// Filtrar estudiantes por término de búsqueda
	let carrerasFiltradas = $derived(
		data.carreras.filter(
			(car: Carrera) =>
				car.nombre.toLowerCase().includes(searchTerm.toLowerCase()) ||
				car.id.toString().toLowerCase().includes(searchTerm.toLowerCase())
		)
	);

	// Función para abrir el modal en modo edición
	function editarEstudiante(estudiante: any) {
		estudianteActual = { ...estudiante };
		if (typeof estudianteActual.carrera === 'object') {
			estudianteActual.carrera = estudianteActual.carrera.id;
		}
		isEditing = true;
		modalVisible = true;
	}

	// Función para abrir el modal en modo creación
	function crearEstudiante() {
		estudianteActual = {
			cedula: '',
			nombre: '',
			correo: '',
			activo: true,
			semestre: 1,
			carrera: 1,
			promedio: 0,
			direccion: '',
			fecha_nac: '',
			sexo: 'M'
		};
		isEditing = false;
		modalVisible = true;
	}
</script>

<div class="w-full">
	<div class="flex justify-between items-center mb-6">
		<h1 class="text-2xl font-bold">Carreras</h1>
		<Button color="blue" onclick={crearEstudiante}>
			<PlusOutline class="mr-2 h-5 w-5" />
			Crear carrera
		</Button>
	</div>

	<div class="mb-4">
		<TableSearch bind:inputValue={searchTerm} placeholder="Buscar por nombre, cédula o correo..." />
	</div>

	{#snippet actions(row: Carrera)}
		<div class="flex gap-2">
			<div class="relative">
				<Button size="xs" color="light" onclick={() => editarEstudiante(row)}>
					<PenOutline class="w-4 h-4" />
				</Button>
				<Tooltip placement="top">Editar carrera</Tooltip>
			</div>
		</div>
	{/snippet}
	<DataTable data={carrerasFiltradas} {actions}></DataTable>

	<Modal title={isEditing ? 'Editar Carrera' : 'Nueva Carrera'} bind:open={modalVisible} size="xs">
		<form action={isEditing ? '?/edit' : '?/create'} method="POST">
			{#if isEditing}
				<input type="hidden" name="id" value={estudianteActual.id} />
			{/if}
			<div class="mb-4 w-full">
				<div>
					<Label for="nombre" class="mb-2">Nombre</Label>
					<ButtonGroup class="w-full">
						<Input
							id="nombre"
							name="nombre"
							placeholder="Informática"
							value={estudianteActual.nombre}
							required
						/>
						<Button type="submit">{isEditing ? 'Actualizar' : 'Guardar'}</Button>
					</ButtonGroup>
				</div>
			</div>
		</form>
		{#snippet footer()}
			<ToastContainer />
		{/snippet}
	</Modal>

</div>
