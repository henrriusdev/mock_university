<script lang="ts">
  import type { PageData } from './$types';
  import { Chart } from '@flowbite-svelte-plugins/chart';

  let { data }: { data: PageData } = $props();

  const pagosPorDia = $derived(Object.entries(data.pagosPorDia));
  const pagosLabels = $derived(pagosPorDia.map(([fecha]) => fecha));
  const pagosData = $derived(pagosPorDia.map(([, monto]) => monto));
</script>

<div class="p-4 space-y-8">
  <h1 class="text-2xl font-bold">Bienvenido, {data.nombre}</h1>

  <!-- Cards -->
  <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
    <div class="bg-blue-50 p-4 rounded shadow">
      <p class="text-sm text-gray-500">Estudiantes registrados</p>
      <p class="text-2xl font-bold">{data.totalEstudiantes}</p>
    </div>
    <div class="bg-blue-50 p-4 rounded shadow">
      <p class="text-sm text-gray-500">Peticiones totales</p>
      <p class="text-2xl font-bold">{data.totalPeticiones}</p>
    </div>
    <div class="bg-blue-50 p-4 rounded shadow">
      <p class="text-sm text-gray-500">Recaudado hoy</p>
      <p class="text-2xl font-bold">{data.totalRecaudado.toFixed(2)} Bs.</p>
    </div>
    <div class="bg-blue-50 p-4 rounded shadow">
      <p class="text-sm text-gray-500">Peticiones por estado</p>
      <ul class="text-sm mt-1">
        {#each Object.entries(data.peticionesPorEstado) as [estado, cantidad] (estado)}
          <li>{estado}: {cantidad}</li>
        {/each}
      </ul>
    </div>
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
