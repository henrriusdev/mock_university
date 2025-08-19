<script lang="ts">
	import { enhance } from '$app/forms';
	import { resolver } from '$lib/utilidades/resolver';
	import type { SubmitFunction } from '@sveltejs/kit';
	import {
		Alert,
		Badge,
		Button,
		Card,
		Spinner,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Tooltip
	} from 'flowbite-svelte';
	import { ArrowDownToBracketOutline } from 'flowbite-svelte-icons';
	import type { PageData } from './$types';
	import { Chart } from '@flowbite-svelte-plugins/chart';

	let { data }: { data: PageData } = $props();

	// Calcular estadísticas
	let creditosObtenidos = $derived(
		data.historicoMaterias
			?.filter((materia) => materia.estatus === 'Aprobada')
			?.reduce((total, materia) => total + materia.creditos, 0) || 0
	);

	let materiasAplazadas = $derived(
		data?.historicoMaterias?.filter((materia) => materia.estatus === 'Reprobada')?.length || 0
	);

	let promedioGeneral = $derived(
		data.historicoMaterias
			?.filter((materia) => materia.estatus === 'Aprobada')
			?.reduce((sum, materia) => sum + materia.nota_final, 0) /
			data.historicoMaterias?.filter((materia) => materia.estatus === 'Aprobada')?.length || 0
	);

	// Preparar datos para el gráfico
	let cicloLabels = $derived(
		Array.from(new Set(data.historicoMaterias?.map((materia) => materia.ciclo)))
	);

	let promedioData = $derived(
		cicloLabels?.map((cicloLabel) => {
			const materiasSemestre = data.historicoMaterias.filter(
				(materia) => materia.ciclo === cicloLabel && materia.promedio >= 9.5
			);

			const promedio =
				materiasSemestre.length > 0
					? materiasSemestre.reduce((sum, materia) => sum + materia.promedio, 0) /
						materiasSemestre.length
					: 0;

			return parseFloat(promedio.toFixed(2));
		}) || []
	);

	let loadingConstancia = $state(false);
	let loadingPlanificacion = $state(false);
	// Configuración del gráfico
	let options: ApexCharts.ApexOptions = $derived({
		chart: {
			type: 'line',
			height: 350,
			toolbar: {
				show: false
			}
		},
		series: [
			{
				name: 'Promedio',
				data: promedioData
			}
		],
		stroke: {
			curve: 'smooth',
			width: 3
		},
		colors: ['#1a56db'],
		markers: {
			size: 5,
			colors: ['#1a56db'],
			strokeWidth: 2,
			strokeColors: '#fff'
		},
		xaxis: {
			categories: cicloLabels,
			title: {
				text: 'Semestre'
			}
		},
		yaxis: {
			min: 0,
			max: 20,
			title: {
				text: 'Promedio'
			}
		},
		tooltip: {
			y: {
				formatter: (value: number) => value.toFixed(2)
			}
		}
	});

	const handleSubmit: SubmitFunction = () => {
		return resolver(() => loadingConstancia = false);
	};
</script>

<div class="w-full p-4">
	<div class="flex justify-between items-center">
		<h2 class="text-2xl font-bold mb-4">Hola de nuevo, {data.estudiante.nombre}!</h2>
		<form
			method="POST"
			action="?/constancia"
			use:enhance={handleSubmit}
			class="space-y-6"
		>
			<Button
				color="primary"
				class="flex justify-center gap-x-3"
				type="submit"
				disabled={loadingConstancia}
			>
				{#if loadingConstancia}
					<Spinner class="me-3" size="4" color="gray" />
					Cargando ...
				{:else}
					<ArrowDownToBracketOutline class="h-5 w-5" />
					Generar constancia
				{/if}
			</Button>
		</form>
	</div>

	<div class="grid md:grid-cols-6 gap-4 mb-6">
		<!-- Resumen estadístico -->
		<Card class="col-span-2 p-4 max-w-full bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Resumen Académico</h3>

			<div class="grid grid-cols-1 md:grid-cols-3 gap-4">
				<div class="bg-gray-50 p-4 rounded-lg text-center">
					<p class="text-3xl font-bold text-primary-700">{creditosObtenidos}</p>
					<p class="text-sm">Créditos Obtenidos</p>
				</div>

				<div class="bg-gray-50 p-4 rounded-lg text-center">
					<p class="text-3xl font-bold text-red-600">{materiasAplazadas}</p>
					<p class="text-sm">Materias Aplazadas</p>
				</div>

				<div class="bg-gray-50 p-4 rounded-lg text-center">
					<p class="text-3xl font-bold text-primary-700">{promedioGeneral.toFixed(2)}</p>
					<p class="text-sm">Promedio General</p>
				</div>
			</div>
			<h3 class="text-lg font-semibold mr-3">Estado de inscripción:</h3>
			{#if data.inscripcionAbierta}
				<Badge color="green" class="w-fit">Abierta</Badge>
			{:else}
				<Badge color="red" class="w-fit">Cerrada</Badge>
			{/if}
		</Card>

		<!-- Gráfico de promedio por semestre -->
		<Card class="col-span-4 p-4 max-w-full bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Evolución del Promedio</h3>
			<Chart {options} />
		</Card>
		<!-- Materias inscritas -->
		<Card class="mb-6 p-4 max-w-full {data.materiasDisponibles.length > 0 ? 'col-span-2' : 'col-span-3'} bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Materias Inscritas</h3>

			{#if data.materiasInscritas.length > 0}
				<Table color="secondary" striped={true}>
					<TableHead class="text-black font-extrabold">
						<TableHeadCell>Materia</TableHeadCell>
						<TableHeadCell>Código</TableHeadCell>
						<TableHeadCell>Docente</TableHeadCell>
						<TableHeadCell>Créditos</TableHeadCell>
						<TableHeadCell>Acciones</TableHeadCell>
					</TableHead>
					<TableBody>
						{#each data.materiasInscritas as materia}
							<TableBodyRow>
								<TableBodyCell>{materia.nombre}</TableBodyCell>
								<TableBodyCell>{materia.id}</TableBodyCell>
								<TableBodyCell>{materia.docente}</TableBodyCell>
								<TableBodyCell>{materia.unidad_credito}</TableBodyCell>
								<TableBodyCell>
									<form
										method="POST"
										action="?/planificacion"
										use:enhance={({formData}) =>{
											formData.set('materia', materia.id);
											return resolver(loadingPlanificacion);
										}}
										class="space-y-6"
									>
										<Button
											color="light"
											class="p-2! grid place-content-center"
											pill
											type="submit"
											disabled={loadingPlanificacion}
										>
											{#if loadingPlanificacion}
												<Spinner class="me-3" size="4" color="gray" />
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
								</TableBodyCell>
							</TableBodyRow>
						{/each}
					</TableBody>
				</Table>
			{:else}
				<Alert color="primary">No tienes materias inscritas actualmente</Alert>
			{/if}

			{#if data.inscripcionAbierta}
				<div class="mt-6">
					<Button color="primary" href="/estudiante/horario">Ir a inscripción de materias</Button>
				</div>
			{/if}
		</Card>

		<!-- Histórico de materias -->
		{#if data.historicoMaterias?.length > 0}
			<Card class="{data.materiasDisponibles.length > 0 ? 'col-span-2' : 'col-span-3'} max-w-full p-4 bg-blue-50">
				<h3 class="text-lg font-semibold mb-4">Histórico de Materias</h3>

				<Table color="secondary" striped={true}>
					<TableHead class="text-black font-extrabold">
						<TableHeadCell>Materia</TableHeadCell>
						<TableHeadCell>Ciclo</TableHeadCell>
						<TableHeadCell>Docente</TableHeadCell>
						<TableHeadCell>Nota</TableHeadCell>
						<TableHeadCell>Estatus</TableHeadCell>
					</TableHead>
					<TableBody>
						{#each data.historicoMaterias as materia}
							<TableBodyRow>
								<TableBodyCell>{materia.id}</TableBodyCell>
								<TableBodyCell>{materia.ciclo}</TableBodyCell>
								<TableBodyCell>{materia.docente}</TableBodyCell>
								<TableBodyCell>{materia.promedio}</TableBodyCell>
								<TableBodyCell>
									{#if materia.estatus === 'Aprobada'}
										<Badge color="green">{materia.estatus}</Badge>
									{:else if materia.estatus === 'Reprobada'}
										<Badge color="red">{materia.estatus}</Badge>
									{:else}
										<Badge color="blue">{materia.estatus}</Badge>
									{/if}
								</TableBodyCell>
							</TableBodyRow>
						{/each}
					</TableBody>
				</Table>
			</Card>
		{/if}

		<!-- Materias disponibles -->
		{#if data.materiasDisponibles.length > 0}
		<Card class="col-span-2 max-w-full p-4 bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Materias Disponibles</h3>
			
				<Table color="secondary" striped={true}>
					<TableHead class="text-black font-extrabold">
						<TableHeadCell>Materia</TableHeadCell>
						<TableHeadCell>Código</TableHeadCell>
						<TableHeadCell>U. C.</TableHeadCell>
						<TableHeadCell>Prelación</TableHeadCell>
						<TableHeadCell>Carrera</TableHeadCell>
					</TableHead>
					<TableBody>
						{#each data.materiasDisponibles as materia}
							<TableBodyRow>
								<TableBodyCell>{materia.nombre}</TableBodyCell>
								<TableBodyCell>{materia.id}</TableBodyCell>
								<TableBodyCell>{materia.unidad_credito}</TableBodyCell>
								<TableBodyCell>{materia.prelacion || 'Ninguna'}</TableBodyCell>
								<TableBodyCell>{materia.carrera.nombre}</TableBodyCell>
							</TableBodyRow>
						{/each}
					</TableBody>
				</Table>
			</Card>
			{/if}
	</div>
</div>
