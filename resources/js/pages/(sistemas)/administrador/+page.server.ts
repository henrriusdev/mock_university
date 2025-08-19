import {
	obtenerPagosPorDia,
	obtenerPagosPorTipo,
	obtenerTotalRecaudado
} from '$lib/servicios/pagos';
import { obtenerEstudiantes } from '$lib/servicios/estudiantes';
import { obtenerPeticiones } from '$lib/servicios/peticiones';

export const load = async ({ fetch }) => {
	const today = new Date().toISOString().split('T')[0];

	const [estudiantes, peticiones, totalRecaudado, pagosPorTipo, pagosPorDia] = await Promise.all([
		obtenerEstudiantes(fetch),
		obtenerPeticiones(fetch),
		obtenerTotalRecaudado(fetch, today, today),
		obtenerPagosPorTipo(fetch, today, today),
		obtenerPagosPorDia(fetch, 7)
	]);

	return {
		totalEstudiantes: estudiantes.length,
		totalPeticiones: peticiones.length,
		peticionesPorEstado: peticiones.reduce(
			(acc, p) => {
				acc[p.peticion.estado] = (acc[p.peticion.estado] || 0) + 1;
				return acc;
			},
			{} as Record<string, number>
		),
		totalRecaudado,
		pagosPorTipo,
		pagosPorDia
	};
};
