import { obtenerCarreras } from '$lib/servicios/carreras';
import { obtenerDocentes } from '$lib/servicios/docentes';
import { obtenerEstudiantes } from '$lib/servicios/estudiantes';
import { obtenerMaterias } from '$lib/servicios/materias';
import { obtenerPagosPorDia } from '$lib/servicios/pagos';
import { obtenerPeticiones } from '$lib/servicios/peticiones';
import type { PageServerLoad } from './$types';

export const load = (async ({ fetch }) => {
    // Obtener datos para estadísticas
    const [estudiantes, docentes, { materias }, carreras] = await Promise.all([
        obtenerEstudiantes(fetch),
        obtenerDocentes(fetch),
        obtenerMaterias(fetch),
        obtenerCarreras(fetch)
    ]);

    // Calcular estadísticas
    const estadisticas = {
        estudiantes: estudiantes.length,
        docentes: docentes.length,
        materias: materias.length,
        carreras: carreras.length
    };

    // Calcular distribución por carrera
    const distribucionCarreras = carreras.map(carrera => ({
        nombre: carrera.nombre,
        estudiantes: estudiantes.filter(e => e.carrera === carrera.nombre).length
    }));

    // Calcular promedios por carrera
    const promediosCarreras = carreras.map(carrera => {
        const estudiantesCarrera = estudiantes.filter(e => e.carrera === carrera.nombre);
        const promedio = estudiantesCarrera.reduce((acc, e) => acc + (e.promedio || 0), 0) / estudiantesCarrera.length;
        return {
            nombre: carrera.nombre,
            promedio: promedio || 0
        };
    });

    // Obtener últimos pagos
    const pagos = await obtenerPagosPorDia(fetch, 5);

    // Obtener peticiones pendientes
    const peticiones = await obtenerPeticiones(fetch);

    return {
        estadisticas,
        distribucionCarreras,
        promediosCarreras,
        pagos: Object.entries(pagos).map(([fecha, monto]) => ({
            estudiante: fecha,
            monto: Number(monto),
            estado: 'aprobado' as const
        })).slice(0, 5),
        peticiones: peticiones.slice(0, 5).filter(p => p.peticion.estado === 'Pendiente').map(p => ({
            estudiante: p.estudiante.nombre,
            campo: p.peticion.campo,
            estado: p.peticion.estado
        }))
    };
}) satisfies PageServerLoad;