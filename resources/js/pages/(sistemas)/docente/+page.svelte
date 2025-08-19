<script lang="ts">
  import {GrillaHorario} from '$lib';
  import { getRandomColor } from "$lib/utilidades/colors";
  import type { MateriaDocente } from '../../../app';

  let {data} = $props();

  const materias = $derived(
    (data.materiasAsignadas || []).map((materia: MateriaDocente) => ({
      ...materia,
      id: `${materia.id}`,
      originalId: materia.id,
      nombre: materia.nombre + " - " + materia.seccion,
      color: materia.color || getRandomColor()
    }))
  );

  function handleMateriaDoubleClick(materiaData: any) {
    console.log('Materia clicked:', materiaData);
  }
</script>

<div class="container mx-auto p-4">
  <h1 class="text-2xl font-bold mb-6">Hola de nuevo, {data.nombre}!</h1>
  
  <!-- Grilla de horarios con información de conteo -->
  <div class="mb-4">
    <p class="text-sm text-gray-600">Materias asignadas: {data.materiasAsignadas?.length || 0} | Horarios en grilla: {materias.length}</p>
  </div>
  
  <GrillaHorario {materias} docente={true} onMateriaDoubleClick={handleMateriaDoubleClick} />
</div>
