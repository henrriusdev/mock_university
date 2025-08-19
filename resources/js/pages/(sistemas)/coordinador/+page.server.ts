import type { PageServerLoad } from './$types';
import { obtenerDocentes, obtenerEstudiantes, obtenerMaterias } from '$lib/servicios';
import type { DistribucionCarrera, PromedioCarrera } from '$lib/types';

export const load: PageServerLoad = async ({ fetch }) => {

    // 1. Usar los servicios centralizados
    const [docentes, estudiantes, materiasResp] = await Promise.all([
        obtenerDocentes(fetch),
        obtenerEstudiantes(fetch),
        obtenerMaterias(fetch)
    ]);

    // Materias puede venir como { materias: Materia[] }
    const materias = Array.isArray(materiasResp) ? materiasResp : materiasResp.materias ?? [];

    // 2. Estadísticas simples
    const estadisticas = {
        estudiantes: estudiantes.length,
        docentes: docentes.length,
        materias: materias.length
    };

    // 3. Distribución de estudiantes por carrera (ajusta si backend retorna varias carreras)
    let distribucionCarreras: DistribucionCarrera = {} as DistribucionCarrera;
    let promediosCarreras: PromedioCarrera = {} as PromedioCarrera;
    if (estudiantes.length > 0) {
        const carreraNombre = (estudiantes[0]?.carrera as unknown as string) || 'Carrera';
        distribucionCarreras = { nombre: carreraNombre , estudiantes: estudiantes.length };
        const promedio = estudiantes.reduce((acc, e) => acc + (e.promedio || 0), 0) / estudiantes.length;
        promediosCarreras = { nombre: carreraNombre, promedio: Math.round(promedio * 100) / 100 };
    }

    return {
        estadisticas,
        distribucionCarreras,
        promediosCarreras
    };
};