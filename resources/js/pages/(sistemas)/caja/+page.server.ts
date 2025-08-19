import type { PageServerLoad } from './$types';
import {
	obtenerEstudiantes,
	obtenerPagosPorDia,
	obtenerPagosPorTipo,
	obtenerTotalRecaudado
} from '$lib';

export const load: PageServerLoad = async ({ fetch }) => {
	const today = new Date().toLocaleDateString("en-CA").split('T')[0];

	const [estudiantes, totalRecaudado, pagosPorTipo, pagosPorDia] = await Promise.all([
		obtenerEstudiantes(fetch),
		obtenerTotalRecaudado(fetch, today, today),
		obtenerPagosPorTipo(fetch, today, today),
		obtenerPagosPorDia(fetch, 7)
	]);

	return {
		totalEstudiantes: estudiantes.length,
		totalRecaudado,
		pagosPorTipo,
		pagosPorDia
	};
};
