<script lang="ts">
	import { cedulaMask, DataTable, telefono, StatusToggleModal } from '$lib';
	import { imask } from '@imask/svelte';
	import { Button, Input, Label, Modal, Select, TableSearch, Tooltip } from 'flowbite-svelte';
	import ToastContainer from '$lib/componentes/ToastContainer.svelte';
	import { PenOutline, PlusOutline, CheckOutline, CloseOutline } from 'flowbite-svelte-icons';
	import type { Coordinador } from '../../../../app';
	import { enhance } from '$app/forms';
	import type { SubmitFunction } from '@sveltejs/kit';
	import { resolver } from '$lib/utilidades/resolver';

	// Datos de la página
	let { data } = $props();

	// Estado para el modal
	let modalVisible = $state(false);
	let isEditing = $state(false);
	let searchTerm = $state('');
	let formEl: HTMLFormElement | undefined = $state();
	// Estado para el modal de confirmación de activación/desactivación
	let statusModalOpen = $state(false);
	let statusModalTitle = $state('');
	let statusModalMessage = $state('');
	let coordinadorActual = $state({
		cedula: '',
		nombre: '',
		correo: '',
		telefono: '',
		carrera: 0,
		usuario: 0,
		activo: true,
		// Internal tracking properties
		id_coordinador: 0
	});
	let selectedCoordinadorForStatus: Coordinador | null = $state(null);

	$effect(() => {
		if (!modalVisible) {
			coordinadorActual = {
				cedula: '',
				nombre: '',
				correo: '',
				telefono: '',
				carrera: 0,
				usuario: 0,
				activo: true,
				id_coordinador: 0
			};
		}
	});

	let coordinadores: Coordinador[] = $state(data.coordinadores);
	let coordinadoresFiltrados = $derived(
		coordinadores.filter(
			(est) =>
				est?.cedula?.includes(searchTerm) ||
				est?.telefono?.toLowerCase().includes(searchTerm.toLowerCase()) ||
				est?.correo?.toLowerCase().includes(searchTerm.toLowerCase()) ||
				est?.nombre?.toLowerCase().includes(searchTerm.toLowerCase())
		) ?? []
	);

	// Función para abrir el modal en modo edición
	function editarCoordinador(coordinador: any) {
		coordinadorActual = {
			cedula: coordinador.cedula || '',
			nombre: coordinador.nombre || '',
			correo: coordinador.correo || '',
			telefono: coordinador.telefono || '',
			carrera: data.carreras.find((car: any) => car.nombre === coordinador.carrera)?.id || 0,
			usuario: coordinador.usuario || 0,
			activo: coordinador.activo || false,
			id_coordinador: coordinador.id_coordinador || 0
		};
		isEditing = true;
		modalVisible = true;
	}

	// Función para abrir el modal en modo creación
	function crearCoordinador() {
		isEditing = false;
		modalVisible = true;
	}

	// Función para abrir el modal de activación/desactivación
	function confirmarCambiarEstadoCoordinador(coordinador: Coordinador) {
		selectedCoordinadorForStatus = coordinador;
		statusModalTitle = coordinador.activo ? "Inactivar Coordinador" : "Activar Coordinador";
		statusModalMessage = coordinador.activo ? 
			`¿Estás seguro de que deseas inactivar al coordinador ${coordinador.nombre}?` : 
			`¿Estás seguro de que deseas activar al coordinador ${coordinador.nombre}?`;
		statusModalOpen = true;
	}

	const handleSubmit: SubmitFunction = () => {
		return resolver(() => {
			modalVisible = false;
		});
	};

	const handleSuccess = () => {
		// Refresh data after successful toggle
		coordinadores = data.coordinadores;
	};
</script>

<div class="w-full">
	<div class="flex justify-between items-center mb-6">
		<h1 class="text-2xl font-bold">Coordinadores</h1>
		<Button color="blue" onclick={crearCoordinador}>
			<PlusOutline class="mr-2 h-5 w-5" />
			Registrar
		</Button>
	</div>

	<div class="mb-4">
		<TableSearch bind:inputValue={searchTerm} placeholder="Buscar por nombre, cédula o correo..." />
	</div>

	{#snippet actions(row: Coordinador)}
		<div class="flex gap-2">
			<div class="relative">
				<Button size="xs" color="light" onclick={() => editarCoordinador(row)}>
					<PenOutline class="w-4 h-4" />
				</Button>
				<Tooltip placement="top">Editar coordinador</Tooltip>
			</div>
			<div class="relative">
				<Button size="xs" color={row.activo ? "red" : "green"} onclick={() => confirmarCambiarEstadoCoordinador(row)}>
					{#if row.activo}
						<CloseOutline class="w-4 h-4" />
					{:else}
						<CheckOutline class="w-4 h-4" />
					{/if}
				</Button>
				<Tooltip placement="top">{row.activo ? 'Inactivar coordinador' : 'Activar coordinador'}</Tooltip>
			</div>
		</div>
	{/snippet}
	<DataTable data={coordinadoresFiltrados} {actions}></DataTable>

	<Modal
		title={isEditing ? 'Editar Coordinador' : 'Nuevo Coordinador'}
		bind:open={modalVisible}
		size="lg"
	>
		<form
			action={isEditing ? '?/edit' : '?/create'}
			use:enhance={handleSubmit}
			method="POST"
			bind:this={formEl}
		>
			{#if isEditing}
				<input type="hidden" name="id_coordinador" value={coordinadorActual.id_coordinador} />
				<input type="hidden" name="id" value={coordinadorActual.usuario} />
			{/if}
			<div class="grid grid-cols-1 md:grid-cols-6 gap-4 mb-4">
				<div class="md:col-span-2">
					<Label for="cedula" class="mb-2">Cédula</Label>
					<input
						id="cedula"
						name="cedula"
						placeholder="Ingrese la cédula"
						value={coordinadorActual?.cedula || ''}
						oninput={(e) => {
							if (coordinadorActual) coordinadorActual.cedula = e.currentTarget.value;
						}}
						required
						use:imask={cedulaMask as any}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500"
					/>
				</div>
				<div class="md:col-span-2">
					<Label for="nombre" class="mb-2">Nombre Completo</Label>
					<Input
						id="nombre"
						name="nombre"
						placeholder="Ingrese el nombre completo"
						value={coordinadorActual.nombre}
						required
						oninput={(e: Event) => {
							const target = e.currentTarget as HTMLInputElement;
							coordinadorActual.nombre = target.value.replace(/\d+/g, '');
						}}
					/>
				</div>
				<div class="md:col-span-2">
					<Label for="correo" class="mb-2">Correo Electrónico</Label>
					<Input
						id="correo"
						name="correo"
						type="email"
						placeholder="correo@ejemplo.com"
						value={coordinadorActual.correo}
						oninput={(e: Event) => {
							const target = e.currentTarget as HTMLInputElement;
							coordinadorActual.correo = target.value;
						}}
						required
					/>
				</div>
				<div class="md:col-span-3">
					<Label for="carrera" class="mb-2">Carrera</Label>
					<Select
						id="carrera"
						name="carrera_id"
						value={coordinadorActual.carrera}
						required
						items={data.carreras.map((carrera: any) => ({
							value: carrera.id,
							name: carrera.nombre
						}))}
						placeholder="Seleccionar"
					/>
				</div>
				<div class="md:col-span-3">
					<Label for="telefono" class="mb-2">Teléfono</Label>
					<input
						id="telefono"
						name="telefono"
						placeholder="Ingrese el teléfono"
						value={coordinadorActual.telefono}
						oninput={(e: Event) => {
							const target = e.currentTarget as HTMLInputElement;
							coordinadorActual.telefono = target.value;
						}}
						required
						use:imask={telefono}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500"
					/>
				</div>
			</div>
		</form>
		{#snippet footer()}
			<div class="flex justify-between items-center w-full">
				<div>
					<Button type="button" color="alternative" onclick={() => (modalVisible = false)}
						>Cancelar</Button
					>
					<Button type="submit" color="primary" onclick={() => formEl?.requestSubmit()}>
						{isEditing ? 'Actualizar' : 'Guardar'}
					</Button>
				</div>
				<ToastContainer />
			</div>
		{/snippet}
	</Modal>

	<!-- Modal de confirmación de activación/desactivación -->
	<StatusToggleModal
		bind:open={statusModalOpen}
		title={statusModalTitle}
		message={statusModalMessage}
		action="?/toggleStatus"
		formData={{
			cedula: selectedCoordinadorForStatus?.cedula || ''
		}}
		isActivating={selectedCoordinadorForStatus ? !selectedCoordinadorForStatus.activo : false}
		onSuccess={handleSuccess}
	/>
</div>
