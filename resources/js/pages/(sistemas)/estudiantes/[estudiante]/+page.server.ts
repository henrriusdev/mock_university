import { error, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { 
	obtenerEstudiantes,
	obtenerEstudiantePorCedula,
	obtenerNotasEstudiantePorCedula
} from '$lib';
import type { Estudiante } from '../../../../app';

export const load: PageServerLoad = async ({ params, fetch, parent }) => {
	const { rol } = await parent();
	
	// Only allow these roles to view student details
	if (!['administrador', 'coordinador', 'control'].includes(rol)) {
		redirect(302, '/' + rol);
	}

	const cedula = params.estudiante;

	try {
		let estudiante: Estudiante | undefined;
		let notasEstudiante: { 
			notas: Array<{
				materia: string;
				id: number;
				notas: number[];
				promedio: number;
			}>;
			ciclo: string;
		} | null = null;
		
		// Try to get student by cedula, fallback to searching all students
		try {
			estudiante = await obtenerEstudiantePorCedula(fetch, cedula);
			notasEstudiante = await obtenerNotasEstudiantePorCedula(fetch, cedula);
			console.log(estudiante, notasEstudiante);
		} catch (e) {
			console.warn('Direct student lookup failed, searching all students:', e);
			const estudiantes = await obtenerEstudiantes(fetch);
			estudiante = estudiantes.find(est => est.cedula === cedula);
			
			if (!estudiante) {
				throw error(404, 'Estudiante no encontrado');
			}
		}

		return {
			estudiante,
			notasEstudiante,
		};
	} catch (err) {
		console.error('Error loading student data:', err);
		
		// If it's already an error with status, re-throw
		if (err.status) {
			throw err;
		}
		
		// Otherwise create a 404 error
		throw error(404, 'Estudiante no encontrado');
	}
};