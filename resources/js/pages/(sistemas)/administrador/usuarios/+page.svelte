<script lang="ts">
	import { enhance } from '$app/forms';
	import { cedulaMask, DataTable } from '$lib';
	import StatusToggleModal from '$lib/componentes/StatusToggleModal.svelte';
	import ToastContainer from '$lib/componentes/ToastContainer.svelte';
	import { resolver } from '$lib/utilidades/resolver';
	import { imask } from '@imask/svelte';
	import type { SubmitFunction } from '@sveltejs/kit';
	import { Button, Input, Label, Modal, Select, Spinner, TableSearch, Tooltip } from 'flowbite-svelte';
	import { CheckOutline, PenOutline, PlusOutline, UsersOutline } from 'flowbite-svelte-icons';
	import type { Usuario } from '../../../../app';
	import type { ActionData, PageData } from './$types';

	// Datos de la página
	let { data }: { data: PageData; form: ActionData } = $props<{
		data: PageData;
		form: ActionData;
	}>();

	// Estado para el modal
	let statusModalOpen = $state(false);
	let selectedUsuarioForStatus: Usuario | null = $state(null);
	let statusModalTitle = $state('');
	let statusModalMessage = $state('');
	let modalVisible = $state(false);
	let descargando = $state(false);
	let cargando = $state(false);
	let isEditing = $state(false);
	let searchTerm = $state('');
	let formEl: HTMLFormElement | undefined = $state();
	// Estado para el modal de confirmación de eliminación
	let deleteModalOpen = $state(false);
	let selectedUsuarioForDelete: Usuario | null = $state(null);
	let usuarioActual: Partial<{
		id: number;
		cedula: string;
		nombre: string;
		correo: string;
		rol: { id: number; nombre: string };
	}> = $state({
		cedula: '',
		nombre: '',
		correo: '',
		rol: { id: 0, nombre: '' }
	});

	$effect(() => {
		if (!modalVisible) {
			usuarioActual = {};
		}
	});

	$effect(() => {
		if (data.usuarios) {
			usuariosFiltrados =
				data?.usuarios.filter(
					(est) =>
						est?.cedula?.toLowerCase().includes(searchTerm.toLowerCase()) ||
						est?.correo?.toLowerCase().includes(searchTerm.toLowerCase()) ||
						est?.nombre?.toLowerCase().includes(searchTerm.toLowerCase())
				) ?? [];
		}
	});

	let usuarios: Usuario[] = $state(data.usuarios);
	let usuariosFiltrados = $derived(
		usuarios.filter(
			(est) =>
				est?.cedula?.toLowerCase().includes(searchTerm.toLowerCase()) ||
				est?.correo?.toLowerCase().includes(searchTerm.toLowerCase()) ||
				est?.nombre?.toLowerCase().includes(searchTerm.toLowerCase())
		) ?? []
	);

	// Función para abrir el modal en modo edición
	function editarUsuario(usuario: any) {
		usuarioActual = { ...usuario };
		isEditing = true;
		modalVisible = true;
	}

	// Función para abrir el modal de eliminación
	function confirmarEliminarUsuario(usuario: Usuario) {
		selectedUsuarioForStatus = usuario;
		statusModalTitle = usuario.activo ? "Inactivar Usuario" : "Activar Usuario";
		statusModalMessage = usuario.activo ? 
			`¿Estás seguro de que deseas inactivar al usuario ${usuario.nombre}?` : 
			`¿Estás seguro de que deseas activar al usuario ${usuario.nombre}?`;
		statusModalOpen = true;
	}

	// Función para abrir el modal en modo creación
	function crearUsuario() {
		isEditing = false;
		modalVisible = true;
	}

	const handleSubmit: SubmitFunction = ({ formElement }) => {
		return resolver(() => {
			formElement.reset();
			modalVisible = false;
		});
	};

	const downloadSubmit: SubmitFunction = () => {
		if (descargando) {
			return;
		}

		descargando = true;
		return resolver(() => (descargando = false));
	};

	const onSearch = (key: keyof Usuario, row: Usuario) => {
		if (key.toLowerCase() === 'rol') {
			return row.nombre[0].toUpperCase() + row.nombre.slice(1);
		}

		return '';
	};

	const handleSuccess = () => {
		statusModalOpen = false;
		usuarios = data.usuarios;
	};
</script>

<div class="w-full">
	<div class="flex justify-between items-center mb-6">
		<h1 class="text-2xl font-bold">Usuarios</h1>
		<div class="flex flex-col md:flex-row justify-center space-x-4">
			<Button color="blue" onclick={crearUsuario}>
				<PlusOutline class="mr-2 h-5 w-5" />
				Registrar
			</Button>
			<form method="POST" action="?/descargar" use:enhance={downloadSubmit} class="space-y-6">
				<Button type="submit" color="alternative">
					{#if !descargando}
						Descargar excel para carga masiva
					{:else}
						<Spinner color="primary" class="h-5 w-5" />
					{/if}
				</Button>
			</form>
			<form
				method="POST"
				action="?/cargar"
				use:enhance={() => {
					return resolver(() => (cargando = false));
				}}
				enctype="multipart/form-data"
				class="space-y-6 flex items-center gap-4"
				id="uploadForm"
			>
				<input
					id="uploadInput"
					name="archivo"
					type="file"
					accept=".xls,.xlsx,.xlsm"
					class="hidden"
					onchange={() => {
						const exportarForm = document.getElementById('uploadForm');
						if (exportarForm instanceof HTMLFormElement) {
							exportarForm.requestSubmit();
						}
					}}
				/>
				<Button
					type="button"
					color="emerald"
					onclick={() => document.getElementById('uploadInput')?.click()}
					class="flex items-center space-x-3"
				>
					{#if cargando}
						<Spinner color="primary" class="h-5 w-5" />
					{:else}
						<PlusOutline class="mr-2 h-5 w-5" />
					{/if}
					Cargar archivo Excel
				</Button>
			</form>
		</div>
	</div>

	<div class="mb-4">
		<TableSearch bind:inputValue={searchTerm} placeholder="Buscar por nombre, cédula o correo..." />
	</div>

	{#snippet actions(row: Usuario)}
		<div class="flex gap-2">
			<div class="relative">
				<Button size="xs" color="light" onclick={() => editarUsuario(row)}>
					<PenOutline class="w-4 h-4" />
				</Button>
				<Tooltip placement="top">Editar usuario</Tooltip>
			</div>
			<div class="relative">
				<Button size="xs" color={row.activo ? "red" : "green"} onclick={() => confirmarEliminarUsuario(row)}>
					{#if row.activo}
						<UsersOutline class="w-5 h-5" />
					{:else}
						<CheckOutline class="w-5 h-5" />
					{/if}
				</Button>
				<Tooltip placement="top">{row.activo ? "Inactivar usuario" : "Activar usuario"}</Tooltip>
			</div>
		</div>
	{/snippet}
	<DataTable data={usuariosFiltrados} {actions} {onSearch}></DataTable>

	<Modal title={isEditing ? 'Editar Usuario' : 'Nuevo Usuario'} bind:open={modalVisible} size="md">
		<form
			action={isEditing ? '?/edit' : '?/create'}
			method="POST"
			use:enhance={handleSubmit}
			bind:this={formEl}
		>
			{#if isEditing}
				<input type="hidden" name="id" value={usuarioActual!.id} />
			{/if}
			<div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-4">
				<div class="md:col-span-2">
					<Label for="cedula" class="mb-2">Cédula</Label>
					<input
						id="cedula"
						name="cedula"
						placeholder="Ingrese la cédula"
						value={usuarioActual?.cedula || ''}
						oninput={(e) => {
							if (usuarioActual) usuarioActual.cedula = e.currentTarget.value;
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
						value={usuarioActual!.nombre}
						required
						oninput={(e: Event) => {
							const target = e.currentTarget as HTMLInputElement;
							if (usuarioActual && target) usuarioActual.nombre = target.value.replace(/\d+/g, '');
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
						value={usuarioActual?.correo}
						required
					/>
				</div>
				<div class="md:col-span-2">
					<Label for="rol" class="mb-2">Rol</Label>
					<Select
						id="rol"
						name="rol"
						value={usuarioActual?.rol?.id ?? 3}
						required
						placeholder="Seleccionar"
						items={[
							{ name: 'Personal de caja', value: 2 },
							{ name: 'Personal de control de estudio', value: 3 }
						]}
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
					<Button type="submit" color="primary" onclick={() => formEl?.requestSubmit()}
						>{isEditing ? 'Actualizar' : 'Guardar'}</Button
					>
				</div>
				<ToastContainer />
			</div>
		{/snippet}
	</Modal>

	<StatusToggleModal
		bind:open={statusModalOpen}
		title={statusModalTitle}
		message={statusModalMessage}
		action="?/toggleStatus"
		formData={{
			cedula: selectedUsuarioForStatus?.cedula || ''
		}}
		isActivating={selectedUsuarioForStatus ? !selectedUsuarioForStatus.activo : false}
		onSuccess={handleSuccess}
	/>
</div>
