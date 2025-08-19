<script lang="ts">
	import {enhance} from '$app/forms';
	import {Button, Modal, Popover, Spinner} from 'flowbite-svelte';
	import {CalendarWeekOutline, PlusOutline} from 'flowbite-svelte-icons';
	import {slide} from 'svelte/transition';
	import type {MateriaDisponible} from '../../../../app';
	import type {PageData} from './$types';
	import {addToast, GrillaHorario} from '$lib';
	import {resolver} from '$lib/utilidades/resolver';
	import type {SubmitFunction} from '@sveltejs/kit';
	import {getRandomColor} from "$lib/utilidades/colors";

	// Tipos
	type HorarioMateria = {
		id: string;
		nombre: string;
		dia: 'Lunes' | 'Martes' | 'Miércoles' | 'Jueves' | 'Viernes' | 'Sábado';
		hora_inicio: string;
		hora_fin: string;
		color?: string;
		conflicto?: boolean;
		unidad_credito?: number;
		prelacion?: string | null;
		carrera?: string;
	};

	let { data }: { data: PageData } = $props();

	let materiasDisponibles = $state<MateriaDisponible[]>(data.materiasDisponibles);
	let materias: (HorarioMateria & { editable: boolean })[] = $state([]);

	let openModal = $state(false);
	let isLoading = $state(false);

	function horaAMinutos(hora: string): number {
		const [horas, minutos] = hora.split(':').map(Number);
		return horas * 60 + minutos;
	}

	function verificarConflictos() {
		let conflictoDetectado = false;

		for (let i = 0; i < materias.length; i++) {
			materias[i].conflicto = false;
		}

		for (let i = 0; i < materias.length; i++) {
			for (let j = i + 1; j < materias.length; j++) {
				const materiaA = materias[i];
				const materiaB = materias[j];

				if (materiaA.dia === materiaB.dia) {
					const inicioA = horaAMinutos(materiaA.hora_inicio);
					const finA = horaAMinutos(materiaA.hora_fin);
					const inicioB = horaAMinutos(materiaB.hora_inicio);
					const finB = horaAMinutos(materiaB.hora_fin);

					if ((inicioA < finB && finA > inicioB) || (inicioB < finA && finB > inicioA)) {
						materiaA.conflicto = true;
						materiaB.conflicto = true;
						conflictoDetectado = true;
					}
				}
			}
		}

		if (conflictoDetectado) {
			addToast({
				type: 'error',
				message: '¡Existen conflictos en el horario!'
			});
		}
	}

	function agregarSeccion(seccion: any) {
		// Check if a section of the same materia is already selected
		const materiaName = seccion.materia_nombre;
		const existingMateriaSections = materiasDisponibles.filter(s => 
			s.materia_nombre === materiaName && 
			materias.some(m => m.id === s.id)
		);
		
		// If a section of this materia is already selected, remove it first
		if (existingMateriaSections.length > 0) {
			existingMateriaSections.forEach(existingSection => {
				quitarSeccionPorID(existingSection.id);
				addToast({
					type: 'info',
					message: `Se ha reemplazado la sección ${existingSection.seccion_nombre} de ${materiaName}`
				});
			});
		}
		
		// Now add the new section
		const horarios = seccion.horarios || [];
		horarios.forEach((horario: any) => {
			materias.push({
				id: seccion.id, // Using asignacion ID
				nombre: seccion.nombre, // Full name with section
				dia: horario.dia,
				hora_inicio: horario.hora_inicio,
				hora_fin: horario.hora_fin,
				color: getRandomColor(),
				conflicto: false,
				unidad_credito: seccion.creditos,
				prelacion: seccion.prelacion,
				carrera: seccion.carrera.nombre,
				editable: true
			});
		});
		replaceMateriaColor();
		verificarConflictos();
	}

	function replaceMateriaColor() {
		for (let i = 0; i < materias.length; i++) {
			let m = materias.find((mat) => mat.id === materias[i].id);
			materias[i].color = m?.color;
		}
	}

	$effect(() => {
		if (data.materiasInscritas.length > 0 && materias.length === 0) {
			data.materiasInscritas.forEach((materia) => {
				const horarios = materia.horarios || [];
				horarios.forEach((horario: any) => {
					materias.push({
						id: materia.id,
						nombre: materia.nombre,
						dia: horario.dia,
						hora_inicio: horario.hora_inicio,
						hora_fin: horario.hora_fin,
						color: getRandomColor(),
						conflicto: false,
						editable: false
					});
				});
			});
			replaceMateriaColor();
			verificarConflictos();
		}
	});

	function quitarSeccionPorID(seccionId: number) {
		materias = materias.filter((m) => m.id !== seccionId);
		verificarConflictos();
	}

	// Handle double-click on materia in the grid to open details or perform an action
	function handleMateriaDoubleClick(materiaData: any) {
		// Find the materia in the available materias
		const materia = materiasDisponibles.find(m => m.id === materiaData.id);
		if (materia) {
			// Open modal to show details or perform an action
			openModal = true;
		}
	}

	const handleSubmit: SubmitFunction = () => {
		return resolver(() => (isLoading = false));
	};
</script>

<!-- Vista -->
<div class="container mx-auto p-4 bg-white">
	<h1 class="text-2xl font-bold text-center mb-6">Horario de clases</h1>

	<GrillaHorario {materias} docente={false} onMateriaDoubleClick={handleMateriaDoubleClick} />

	<!-- Botones -->
	<div class="flex flex-wrap justify-center gap-4 mt-6">
		<Button color="blue" onclick={() => (openModal = true)} disabled={!data.inscripcionAbierta}>
			<CalendarWeekOutline class="mr-2 h-5 w-5" />
			Seleccionar Secciones
		</Button>
		<form use:enhance={handleSubmit} method="post">
			{#each materias as materia}
				<input type="hidden" name="asignaciones" value={materia.id} />
			{/each}
			<Button
				type="submit"
				color="green"
				disabled={materias.some((m) => m.conflicto) ||
					!data.inscripcionAbierta ||
					materias.length === 0 ||
					isLoading}
			>
				{#if isLoading}
					<Spinner class="me-3" size="4" color="gray" />
					Cargando ...
				{:else}
					<CalendarWeekOutline class="mr-2 h-5 w-5" />
					Registrar Horario
				{/if}
			</Button>
		</form>
	</div>

	<!-- Modal Seleccionar Secciones -->
	<Modal bind:open={openModal} size="xl">
		{#snippet header()}
			<div class="text-lg font-semibold">Seleccionar Secciones</div>
		{/snippet}

		<div class="p-4 flex flex-wrap justify-start gap-4">
			{#each materiasDisponibles as seccion}
				<div class="relative">
					<Button color="light" type="button" id="seccion-{seccion.id}">
						<PlusOutline class="mr-2 h-5 w-5" />
						{seccion.nombre}
					</Button>

					<Popover
						triggeredBy="#seccion-{seccion.id}"
						placement="bottom"
						class="w-80 z-50 text-sm"
						transition={slide}
					>
						<div class="p-2">
							<p><span class="font-semibold">Materia:</span> {seccion.materia_nombre}</p>
							<p><span class="font-semibold">Sección:</span> {seccion.seccion_nombre}</p>
							<p><span class="font-semibold">Créditos:</span> {seccion.creditos}</p>
							<p><span class="font-semibold">Prelación:</span> {seccion.prelacion || 'Ninguna'}</p>
							<p><span class="font-semibold">Profesor:</span> {seccion.profesor?.nombre || 'No asignado'}</p>
							<p><span class="font-semibold">Estudiantes:</span> {seccion.cantidad_estudiantes}/{seccion.maximo}</p>
							<p><span class="font-semibold">Carrera:</span> {seccion.carrera.nombre}</p>
							<div class="mt-2">
								<p class="font-semibold">Horarios:</p>
								{#each seccion.horarios as horario}
									<p class="text-xs">{horario.dia}: {horario.hora_inicio} - {horario.hora_fin}</p>
								{/each}
							</div>
							{#if materias.some((m) => m.id === seccion.id)}
								<Button
									size="xs"
									class="mt-2"
									type="button"
									onclick={() => quitarSeccionPorID(seccion.id)}>Quitar</Button
								>
							{:else}
								<Button size="xs" class="mt-2" type="button" onclick={() => agregarSeccion(seccion)}
									>Agregar</Button
								>
							{/if}
						</div>
					</Popover>
				</div>
			{/each}
		</div>

		{#snippet footer()}
			<div class="flex justify-end p-4">
				<Button color="alternative" onclick={() => (openModal = false)}>Cerrar</Button>
			</div>
		{/snippet}
	</Modal>
</div>