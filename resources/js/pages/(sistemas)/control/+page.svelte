<script lang="ts">
	import {
		Alert,
		Badge,
		Card,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import type { PageData } from './$types';
	import type { DistribucionCarrera, PromedioCarrera } from './types';
	import { Chart } from '@flowbite-svelte-plugins/chart';

	let { data }: { data: PageData } = $props();

	// Estadísticas
	let totalEstudiantes = $derived(data.estadisticas?.estudiantes || 0);
	let totalDocentes = $derived(data.estadisticas?.docentes || 0);
	let totalMaterias = $derived(data.estadisticas?.materias || 0);
	let totalCarreras = $derived(data.estadisticas?.carreras || 0);

	// Gráfico de distribución por carrera
	let distribucionOptions = $derived({
		chart: {
			type: 'bar' as const,
			height: 350,
			toolbar: {
				show: false
			}
		},
		series: [
			{
				name: 'Estudiantes',
				data: data.distribucionCarreras?.map((c: DistribucionCarrera) => c.estudiantes) || []
			}
		],
		xaxis: {
			categories: data.distribucionCarreras?.map((c: DistribucionCarrera) => c.nombre) || [],
			title: {
				text: 'Carrera'
			}
		},
		yaxis: {
			title: {
				text: 'Estudiantes'
			}
		},
		colors: ['#1a56db']
	});

	// Gráfico de promedios por carrera
	let promediosOptions = $derived({
		chart: {
			type: 'line' as const,
			height: 350,
			toolbar: {
				show: false
			}
		},
		series: [
			{
				name: 'Promedio',
				data: data.promediosCarreras?.map((c: PromedioCarrera) => parseFloat(c.promedio.toFixed(2))) || []
			}
		],
		xaxis: {
			categories: data.promediosCarreras?.map((c: PromedioCarrera) => c.nombre) || [],
			title: {
				text: 'Carrera'
			}
		},
		yaxis: {
			min: 0,
			max: 20,
			title: {
				text: 'Promedio'
			}
		},
		colors: ['#16a34a'],
		stroke: {
			curve: 'smooth' as const,
			width: 3
		},
		markers: {
			size: 5
		}
	});
</script>

<div class="w-full p-4">
	<h2 class="text-2xl font-bold mb-4">Panel de Control</h2>

	<div class="grid md:grid-cols-6 gap-4 mb-6">
		<!-- Estadísticas -->
		<Card class="col-span-6 p-4 max-w-full bg-white">
			<div class="grid grid-cols-1 md:grid-cols-4 gap-4">
				<div class="bg-blue-50 p-4 rounded-lg text-center">
					<p class="text-3xl font-bold text-primary-700">{totalEstudiantes}</p>
					<p class="text-sm">Estudiantes Activos</p>
				</div>

				<div class="bg-blue-50 p-4 rounded-lg text-center">
					<p class="text-3xl font-bold text-green-600">{totalDocentes}</p>
					<p class="text-sm">Docentes Activos</p>
				</div>

				<div class="bg-blue-50 p-4 rounded-lg text-center">
					<p class="text-3xl font-bold text-blue-600">{totalMaterias}</p>
					<p class="text-sm">Materias Activas</p>
				</div>

				<div class="bg-blue-50 p-4 rounded-lg text-center">
					<p class="text-3xl font-bold text-purple-600">{totalCarreras}</p>
					<p class="text-sm">Carreras</p>
				</div>
			</div>
		</Card>

		<!-- Gráfico de distribución por carrera -->
		<Card class="col-span-3 p-4 max-w-full bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Distribución de Estudiantes por Carrera</h3>
			<Chart options={distribucionOptions} />
		</Card>

		<!-- Gráfico de promedios por carrera -->
		<Card class="col-span-3 p-4 max-w-full bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Promedio por Carrera</h3>
			<Chart options={promediosOptions} />
		</Card>

		<!-- Peticiones pendientes -->
		<Card class="col-span-6 p-4 max-w-full bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Peticiones Pendientes</h3>
			{#if data.peticiones?.length > 0}
				<Table striped={true}>
					<TableHead>
						<TableHeadCell>Estudiante</TableHeadCell>
						<TableHeadCell>Corte</TableHeadCell>
						<TableHeadCell>Estado</TableHeadCell>
					</TableHead>
					<TableBody>
						{#each data.peticiones as peticion}
							<TableBodyRow>
								<TableBodyCell>{peticion.estudiante}</TableBodyCell>
								<TableBodyCell>{peticion.campo}</TableBodyCell>
								<TableBodyCell>
									<Badge color="yellow">{peticion.estado}</Badge>
								</TableBodyCell>
							</TableBodyRow>
						{/each}
					</TableBody>
				</Table>
			{:else}
				<Alert color="blue">No hay peticiones pendientes</Alert>
			{/if}
		</Card>
	</div>
</div>