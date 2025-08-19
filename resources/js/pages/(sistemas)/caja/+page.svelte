<script lang="ts">
  import {Chart} from '@flowbite-svelte-plugins/chart';
  import type {PageData} from './$types';

  let {data}: { data: PageData } = $props();

  const pagosPorDia = $derived(Object.entries(data.pagosPorDia));
  const pagosLabels = $derived(pagosPorDia.map(([fecha]) => fecha));
  const pagosData = $derived(pagosPorDia.map(([, monto]) => monto));

  const pagosPorTipoData = $derived(Object.entries(data.pagosPorTipo));

  const donutChartOptions: ApexCharts.ApexOptions = $derived({
    series: pagosPorTipoData.map(([, monto]) => monto),
    colors: ['#1C64F2', '#16BDCA', '#9061F9'], // Customize colors
    chart: {
      height: 420,
      width: '100%',
      type: 'donut'
    },
    labels: pagosPorTipoData.map(([tipo]) => tipo),
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
</script>

<div class="p-4 space-y-8">
    <div class="flex justify-between">
        <h1 class="text-2xl font-bold">Bienvenido, {data.nombre}</h1>
    </div>

    <!-- Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div class="bg-blue-50 p-4 rounded shadow">
            <p class="text-sm text-gray-500">Estudiantes registrados</p>
            <p class="text-2xl font-bold">{data.totalEstudiantes}</p>
        </div>
        <div class="bg-blue-50 p-4 rounded shadow">
            <p class="text-sm text-gray-500">Recaudado hoy</p>
            <p class="text-2xl font-bold">{data.totalRecaudado.toFixed(2)} Bs.</p>
        </div>
    </div>

    <!-- Pagos por tipo (Torta) -->
    <div class="bg-blue-50 p-4 rounded shadow">
        <h2 class="text-lg font-semibold mb-4">Pagos por tipo</h2>
        <Chart
                options={donutChartOptions}
        />
    </div>

    <!-- Pagos por día (Barras) -->
    <div class="bg-blue-50 p-4 rounded shadow">
        <h2 class="text-lg font-semibold mb-4">Pagos últimos 7 días</h2>
        <Chart
                options={{
        chart: { type: 'bar', height: 350, toolbar: { show: false } },
        xaxis: { categories: pagosLabels, title: { text: 'Fecha' } },
        yaxis: { title: { text: 'Monto ($)' } },
        colors: ['#1a56db'],
        tooltip: { y: { formatter: (v: number) => `$${v.toFixed(2)}` } },
        series: [{ name: 'Pagos', data: pagosData }]
      }}
        />
    </div>
</div>
