<script lang="ts">
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';
	import { DataTable, ConfirmDeleteModal } from '$lib/componentes';
	import { resolver } from '$lib/utilidades/resolver';
	import type { SubmitFunction } from '@sveltejs/kit';
	import { Button, Input, Label, Modal, MultiSelect, Select, Tooltip } from 'flowbite-svelte';
	import ToastContainer from '$lib/componentes/ToastContainer.svelte';
	import {
		CalendarMonthOutline,
		InfoCircleOutline,
		PenOutline,
		PlusOutline,
		TrashBinOutline,
		UsersGroupOutline
	} from 'flowbite-svelte-icons';
	import type { Asignacion, Horario, Materia } from '../../../app';
	import type { PageData } from './$types';
	import type { AsignacionReq } from '$lib/types';

	interface Form {
		id: string;
		nombre: string;
		prelacion: (string | number)[];
		unidad_credito: number;
		hp: number;
		ht: number;
		semestre: number;
		id_carrera: string;
		ciclo: string;
		maximo: number;
		asignaciones: AsignacionReq[];
	}

	const dias = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado'];

	let form: Form = $state({
		id: '',
		nombre: '',
		prelacion: [],
		unidad_credito: 0,
		hp: 0,
		ht: 0,
		semestre: 1,
		id_carrera: '',
		ciclo: '',
		maximo: 0,
		asignaciones: []
	});

	let formEl: HTMLFormElement;
	let asignacionesVisualizacion: Asignacion[] = $state([]);
	let { data }: { data: PageData } = $props();
	let showModal = $state(false);
	let showAsignaciones = $state(false);
	let prelacion = $state(false);
	let editMode = $state(false);
	let filtroCarrera: number | null = $state(data.rol === 'coordinador' ? data.carrera_id : null);
	let filtroDocente: number | null = $state(null);
	let filtroSemestre = $state('');
	let searchTerm = $state('');
	let showAyuda = $state(false);
	// Estado para el modal de confirmación de eliminación
	let deleteModalOpen = $state(false);
	let selectedMateriaForDelete: Materia | null = $state(null);

	// Opciones de prelación calculadas dinámicamente
	let opcionesPrelacion: { value: string; name: string }[] = $state([]);

	const materiasFiltradas = $derived(
		data.materias.filter((mat) => {
			const matchCarrera = !filtroCarrera || mat.id_carrera === filtroCarrera;
			// Since we don't have a single docente anymore, we check if any asignacion has the selected docente
			const matchDocente = !filtroDocente || mat.asignaciones?.some(asig => asig.profesor?.id === filtroDocente);
			const matchSemestre = !filtroSemestre || mat.semestre === +filtroSemestre;

			const matchSearch = (() => {
				if (!searchTerm.trim()) return true;
				try {
					const fixedExpr = searchTerm
						.replace(
							/([a-zA-Z_]+)\s*==?\s*([a-zA-Z_]+)/g,
							(_, field, value) => `${field} == "${value}"`
						)
						.replace(/<>/g, '!=')
						.replace(/=/g, '==');

					return Function('mat', `with (mat) { return ${fixedExpr} }`)(mat);
				} catch {
					return [mat.nombre, mat.id].some((v) =>
						v.toLowerCase().includes(searchTerm.toLowerCase())
					);
				}
			})();
			return matchCarrera && matchDocente && matchSemestre && matchSearch;
		})
	);

	function openModal(materia: Materia | null = null) {
		if (materia) {
			form = {
				...materia,
				id_carrera: materia.id_carrera.toString(),
				asignaciones: materia.asignaciones?.map(asig => ({
					id: asig.id,
					nombre: asig.nombre,
					profesor_id: asig.profesor?.id,
					horarios: asig.horarios
				})) || [],
				prelacion: materia.prelacion ? materia.prelacion.split(',').map((p) => p.trim()) : []
			};
			editMode = true;
		} else {
			form = {
				id: '',
				nombre: '',
				prelacion: [],
				unidad_credito: 0,
				hp: 0,
				ht: 0,
				semestre: 1,
				id_carrera: '',
				ciclo: '',
				maximo: 30,
				asignaciones: []
			};
			editMode = false;
		}
		showModal = true;
	}

	function addAsignacion() {
		form.asignaciones.push({
			nombre: '',
			profesor_id: undefined,
			horarios: []
		});
	}

	function removeAsignacion(index: number) {
		form.asignaciones.splice(index, 1);
	}

	function addHorarioToAsignacion(asignacionIndex: number) {
		form.asignaciones[asignacionIndex].horarios.push({
			dia: 'Lunes',
			hora_inicio: '08:00',
			hora_fin: '10:00'
		});
	}

	function removeHorarioFromAsignacion(asignacionIndex: number, horarioIndex: number) {
		form.asignaciones[asignacionIndex].horarios.splice(horarioIndex, 1);
	}

	function openAsignaciones(row: Materia) {
		asignacionesVisualizacion = row.asignaciones || [];
		showAsignaciones = true;
	}

	// Función para abrir el modal de eliminación
	function confirmarEliminarMateria(materia: Materia) {
		selectedMateriaForDelete = materia;
		deleteModalOpen = true;
	}

	function searchName(key: 'asignaciones' | 'id_carrera', value: number | Asignacion[]): string {
		if (key === 'asignaciones' && Array.isArray(value)) {
			return value.map(asig => asig.profesor?.nombre || 'Sin profesor').join(', ');
		}

		if (key === 'id_carrera') {
			return data.carreras.find((c) => c.id === value)?.nombre || '';
		}

		return '';
	}

	const handleSubmit: SubmitFunction = () => {
		return resolver(() => (showModal = false));
	};

	// Update prelacion options whenever form.semestre, form.id_carrera, or showModal changes
	$effect(() => {
		// Only populate options when modal is open
		if (!showModal) return;
		
		if (form.semestre === 1) {
			opcionesPrelacion = [];
		} else {
			opcionesPrelacion = data.materias
				.filter(
					(m) => m.id_carrera == form.id_carrera && m.semestre < form.semestre && m.id !== form.id
				)
				.map((m) => ({ value: m.id, name: m.nombre }));
		}
	});
</script>

<!-- Tabla de materias -->
<div class="flex justify-between items-center mb-4">
	<h1 class="text-2xl font-bold">Materias</h1>
	{#if ['coordinador', 'administrador'].includes(data.rol.toLowerCase())}
		<Button onclick={() => openModal()} class="btn btn-primary">
			<PlusOutline class="h-5 w-5 mr-4" />
			Crear Materia
		</Button>
	{/if}
</div>

<div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-4">
	<div>
		<Label>Carrera</Label>
		<Select
			placeholder="Seleccionar"
			bind:value={filtroCarrera}
			items={[
				{ value: '', name: 'Todas' },
				...data.carreras.map((c) => ({ value: c.id, name: c.nombre }))
			]}
		/>
	</div>
	<div>
		<Label>Docente</Label>
		<Select
			placeholder="Seleccionar"
			bind:value={filtroDocente}
			items={[
				{ value: '', name: 'Todos' },
				...data.docentes.map((d) => ({ value: d.id, name: d.nombre }))
			]}
		/>
	</div>
	<div>
		<Label>Semestre</Label>
		<Select
			placeholder="Seleccionar"
			bind:value={filtroSemestre}
			items={[
				{ value: '', name: 'Todos' },
				...Array(10)
					.fill(null)
					.map((_, i) => ({ value: `${i + 1}`, name: `${i + 1}°` }))
			]}
		/>
	</div>
	<div class="col-span-2">
		<Label>Búsqueda o condición</Label>
		<Input bind:value={searchTerm} placeholder="Nombre, código o expresión...">
			{#snippet right()}
				<Button
					color="primary"
					size="xs"
					title="Ayuda"
					class="p-2!"
					onclick={() => (showAyuda = true)}
				>
					<InfoCircleOutline class="w-6 h-6" />
				</Button>
			{/snippet}
		</Input>
	</div>
</div>

{#snippet action(row: Materia)}
	<div class="flex justify-between items-center">
		{#if ['coordinador', 'administrador'].includes(data.rol.toLowerCase())}
			<div class="relative">
				<Button pill class="p-1.5!" size="xs" color="light" onclick={() => openModal(row)}>
					<PenOutline class="w-5 h-5" />
				</Button>
				<Tooltip placement="top">Editar materia</Tooltip>
			</div>
			<!--<div class="relative">
				<Button
					pill
					class="p-1.5!"
					size="xs"
					color="red"
					onclick={() => confirmarEliminarMateria(row)}
				>
					<TrashBinOutline class="w-5 h-5" />
				</Button>
				<Tooltip placement="top">Eliminar materia</Tooltip>
			</div>-->
		{/if}
		<div class="relative">
			<Button pill class="p-1.5!" size="xs" color="alternative" onclick={() => openAsignaciones(row)}>
				<UsersGroupOutline class="w-5 h-5" />
			</Button>
			<Tooltip placement="top">Ver secciones</Tooltip>
		</div>
	</div>
{/snippet}

<DataTable data={materiasFiltradas} actions={action} onSearch={searchName} />

<!-- Modal de creación/edición -->
<Modal bind:open={showModal} size="lg">
	{#snippet header()}
		<div>{editMode ? 'Editar Materia' : 'Crear Materia'}</div>
	{/snippet}
	<form
		method="POST"
		use:enhance={handleSubmit}
		bind:this={formEl}
		action={editMode ? `?/edit` : '?/create'}
	>
		<div class="grid grid-cols-4 gap-4">
			<div class="col-span-2">
				<Label for="id" class="mb-2">Código</Label>
				<Input id="id" name="id" bind:value={form.id} class="input" required readonly={editMode} />
			</div>
			<div class="col-span-2">
				<Label for="nombre" class="mb-2">Nombre</Label>
				<Input name="nombre" bind:value={form.nombre} class="input" required />
			</div>
			<div>
				<Label for="unidad_credito" class="mb-2">Unidades de crédito</Label>
				<Input
					id="unidad_credito"
					name="unidad_credito"
					type="number"
					bind:value={form.unidad_credito}
					class="input"
					required
				/>
			</div>
			<div>
				<Label for="hp" class="mb-2">Horas prácticas</Label>
				<Input id="hp" name="hp" type="number" bind:value={form.hp} class="input" required />
			</div>
			<div>
				<Label for="ht" class="mb-2">Horas Teóricas</Label>
				<Input id="ht" name="ht" type="number" bind:value={form.ht} class="input" required />
			</div>
			<div>
				<Label for="semestre" class="mb-2">Semestre</Label>
				<Input
					id="semestre"
					name="semestre"
					type="number"
					bind:value={form.semestre}
					class="input"
					required
				/>
			</div>
			<div class="col-span-3">
				<Label for="id_carrera" class="mb-2">Carrera</Label>
				<Select
					id="id_carrera"
					name="id_carrera"
					bind:value={form.id_carrera}
					class="select"
					required
					items={data.carreras.map((c) => ({ value: c.id, name: c.nombre }))}
					placeholder="Seleccione"
				></Select>
			</div>
			<div>
				<Label for="maximo" class="mb-2">Máximo de estudiantes</Label>
				<Input id="maximo" name="maximo" type="number" bind:value={form.maximo} class="input" />
			</div>
			<div class="col-span-4">
				<Label for="prelacion" class="mb-2">Prelación</Label>
				<MultiSelect
					id="prelacion"
					dropdownClass="z-[99999]!"
					name="prelacion"
					bind:value={form.prelacion}
					items={opcionesPrelacion}
					placeholder="Seleccione las prelaciones"
				/>
			</div>
		</div>

		<!-- Sección de asignaciones -->
		<div class="mt-6">
			<div class="flex justify-between items-center mb-4">
				<h3 class="text-lg font-medium">Asignaciones de Secciones</h3>
				<Button onclick={addAsignacion} color="blue" size="sm">
					<PlusOutline class="w-4 h-4 mr-2" />
					Agregar Sección
				</Button>
			</div>
			
			{#each form.asignaciones as asignacion, i}
				<div class="border rounded-lg p-4 mb-4 bg-gray-50">
					<div class="flex justify-between items-start mb-3">
						<h4 class="font-medium">Sección {i + 1}</h4>
						<Button onclick={() => removeAsignacion(i)} color="red" size="xs">
							<TrashBinOutline class="w-4 h-4" />
						</Button>
					</div>
					
					<div class="grid grid-cols-2 gap-4 mb-4">
						<div>
							<Label for={`nombre_${i}`} class="mb-2">Nombre de la sección</Label>
							<Input
								id={`nombre_${i}`}
								bind:value={asignacion.nombre}
								placeholder="ej: MA, MB, MC"
								class="input"
							/>
						</div>
						<div>
							<Label for={`profesor_${i}`} class="mb-2">Profesor</Label>
							<Select
								id={`profesor_${i}`}
								bind:value={asignacion.profesor_id}
								items={data.docentes.map((d) => ({ value: d.id, name: d.nombre }))}
								placeholder="Seleccione profesor"
								class="select"
							/>
						</div>
					</div>
					
					<!-- Horarios para esta asignación -->
					<div>
						<div class="flex justify-between items-center mb-2">
							<Label class="text-sm">Horarios</Label>
							<Button onclick={() => addHorarioToAsignacion(i)} color="green" size="xs">
								<PlusOutline class="w-3 h-3 mr-1" />
								Agregar Horario
							</Button>
						</div>
						
						{#each asignacion.horarios as horario, j}
							<div class="grid grid-cols-4 gap-2 mb-2 items-end">
								<div>
									<Label for={`dia_${i}_${j}`} class="text-xs">Día</Label>
									<Select
										id={`dia_${i}_${j}`}
										bind:value={horario.dia}
										items={dias.map(d => ({ value: d, name: d }))}
										class="select"
									/>
								</div>
								<div>
									<Label for={`inicio_${i}_${j}`} class="text-xs">Hora inicio</Label>
									<Input
										id={`inicio_${i}_${j}`}
										type="time"
										bind:value={horario.hora_inicio}
										class="input"
									/>
								</div>
								<div>
									<Label for={`fin_${i}_${j}`} class="text-xs">Hora fin</Label>
									<Input
										id={`fin_${i}_${j}`}
										type="time"
										bind:value={horario.hora_fin}
										class="input"
									/>
								</div>
								<div>
									<Button onclick={() => removeHorarioFromAsignacion(i, j)} color="red" size="xs">
										<TrashBinOutline class="w-3 h-3" />
									</Button>
								</div>
							</div>
						{/each}
					</div>
				</div>
			{/each}
		</div>

		<!-- Hidden inputs -->
		<input type="hidden" name="asignaciones" value={JSON.stringify(form.asignaciones)} />
		<input type="hidden" name="prelacion" value={form.prelacion.join(',')} />
	</form>
	{#snippet footer()}
		<div class="flex justify-between items-center w-full">
			<div>
				<Button type="button" color="alternative" onclick={() => (showModal = false)}
					>Cancelar
				</Button>
				<Button type="submit" color="primary" onclick={() => formEl.requestSubmit()}
					>{editMode ? 'Actualizar' : 'Guardar'}</Button
				>
			</div>
			<ToastContainer />
		</div>
	{/snippet}
</Modal>

<!-- Modal de visualización de asignaciones -->
<Modal bind:open={showAsignaciones} size="lg">
	{#snippet header()}
		<div>Secciones y Horarios</div>
	{/snippet}
	<div class="space-y-4">
		{#each asignacionesVisualizacion as asignacion}
			<div class="border rounded-lg p-4">
				<div class="flex justify-between items-start mb-3">
					<div>
						<h3 class="font-semibold text-lg">Sección: {asignacion.nombre}</h3>
						{#if asignacion.profesor}
							<p class="text-gray-600">Profesor: {asignacion.profesor.nombre}</p>
						{:else}
							<p class="text-gray-500 italic">Sin profesor asignado</p>
						{/if}
					</div>
				</div>
				
				{#if asignacion.horarios && asignacion.horarios.length > 0}
					<div>
						<h4 class="font-medium mb-2">Horarios:</h4>
						<div class="space-y-2">
							{#each asignacion.horarios as horario}
								<div class="bg-gray-50 p-2 rounded">
									<span class="font-medium">{horario.dia}:</span>
									{horario.hora_inicio} - {horario.hora_fin}
								</div>
							{/each}
						</div>
					</div>
				{:else}
					<p class="text-gray-500 italic">Sin horarios configurados</p>
				{/if}
			</div>
		{:else}
			<p class="text-gray-500 italic text-center py-8">No hay secciones configuradas para esta materia</p>
		{/each}
	</div>
	<div class="flex justify-end mt-4">
		<Button onclick={() => (showAsignaciones = false)}>Cerrar</Button>
	</div>
</Modal>

<Modal bind:open={showAyuda} size="md">
	{#snippet header()}
		<h3 class="text-lg font-bold">Guía de expresiones</h3>
	{/snippet}
	<ul class="list-disc list-inside space-y-1">
		<li><code>ht &gt; 4</code> → Horas teóricas mayores a 4</li>
		<li><code>hp &lt;= 5</code> → Horas prácticas menores o iguales a 5</li>
		<li><code>maximo != 30</code> → Máximo de estudiantes distinto de 30</li>
		<li><code>semestre == 9</code> → Solo materias del 9° semestre</li>
		<li><code>nombre == 'mate'</code> → Nombre contiene "mate"</li>
	</ul>
	{#snippet footer()}
		<div>
			<Button onclick={() => (showAyuda = false)}>Cerrar</Button>
		</div>
	{/snippet}
</Modal>

<!-- Modal de confirmación de eliminación -->
<ConfirmDeleteModal
	bind:open={deleteModalOpen}
	title="Eliminar Materia"
	message="¿Estás seguro de que deseas eliminar la materia {selectedMateriaForDelete?.nombre}? Esta acción no se puede deshacer."
	action="?/delete"
	formData={{ id: selectedMateriaForDelete?.id || '' }}
	onSuccess={() => {
		selectedMateriaForDelete = null;
	}}
/>