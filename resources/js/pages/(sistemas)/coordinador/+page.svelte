<script lang="ts">
	import { Chart } from '@flowbite-svelte-plugins/chart';
	import { Card } from 'flowbite-svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	// Estadísticas de la(s) carrera(s) del coordinador
	let totalEstudiantes = $derived(data.estadisticas?.estudiantes || 0);
	let totalDocentes = $derived(data.estadisticas?.docentes || 0);
	let totalMaterias = $derived(data.estadisticas?.materias || 0);

	// Gráfico de distribución por carrera (solo las del coordinador)
	let distribucionOptions: ApexCharts.ApexOptions = $derived({
		series: [data.distribucionCarreras?.estudiantes ?? 0],
		colors: ['#1C64F2', '#16BDCA', '#9061F9'], // Customize colors
		chart: {
			height: 420,
			width: '100%',
			type: 'donut' as const
		},
		labels: [data.distribucionCarreras?.nombre ?? 'Carrera'],
		legend: {
			position: 'bottom',
			fontFamily: 'Inter, sans-serif'
		},
		dataLabels: {
			enabled: true,
			style: {
				fontFamily: 'Inter, sans-serif'
			}
		}
	});

	// Gráfico de promedios por carrera (solo las del coordinador)
	let promediosOptions = $derived({
		chart: {
			type: 'line' as const,
			height: 350,
			toolbar: { show: false }
		},
		series: [
			{
				name: 'Promedio',
				data: [data.promediosCarreras?.promedio ?? 0]
			}
		],
		xaxis: {
			categories: [data.promediosCarreras?.nombre ?? 'Carrera'],
			title: { text: 'Carrera' }
		},
		yaxis: { min: 0, max: 20, title: { text: 'Promedio' } },
		colors: ['#16a34a'],
		stroke: { curve: 'smooth' as const, width: 3 },
		markers: { size: 5 }
	});
</script>

<div class="w-full p-4">
	<h2 class="text-2xl font-bold mb-4">Panel de Coordinador</h2>

	<div class="grid md:grid-cols-6 gap-4 mb-6">
		<!-- Estadísticas -->
		<Card class="col-span-6 p-4 max-w-full bg-blue-50">
			<div class="grid grid-cols-1 md:grid-cols-3 gap-4">
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
			</div>
		</Card>

		<!-- Gráfico de distribución por carrera -->
		<Card class="col-span-3 p-4 max-w-full bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Distribución de Estudiantes por Carrera</h3>
			<div class="w-1/2 mx-auto">
				<Chart options={distribucionOptions} />
			</div>
		</Card>

		<!-- Gráfico de promedios por carrera -->
		<Card class="col-span-3 p-4 max-w-full bg-blue-50">
			<h3 class="text-lg font-semibold mb-4">Promedio por Carrera</h3>
			<Chart options={promediosOptions} />
		</Card>
	</div>
</div>
