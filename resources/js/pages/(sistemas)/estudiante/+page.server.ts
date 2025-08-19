import {
	obtenerHistoricoMaterias,
	obtenerMateriasDisponibles,
	obtenerMateriasInscritas
} from '$lib';
import { error } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import type { MateriaDisponible } from '../../../app';
import { obtenerConstancia, obtenerPlanificacion } from '$lib/servicios/archivos';

export const load: PageServerLoad = async ({ fetch, locals: { usuario } }) => {
	try {
		// Obtener datos
		const materiasInscritas = await obtenerMateriasInscritas(fetch);
		const historicoMaterias = await obtenerHistoricoMaterias(fetch);
		if (materiasInscritas.length > 2) {
			return {
				estudiante: usuario,
				materiasInscritas,
				historicoMaterias,
				inscripcionAbierta: false,
				materiasDisponibles: [] as MateriaDisponible[]
			};
		}

		const materiasDisponibles = await obtenerMateriasDisponibles(fetch, usuario?.cedula ?? '');
		return {
			estudiante: usuario,
			materiasInscritas,
			historicoMaterias,
			materiasDisponibles,
			inscripcionAbierta: true
		};
	} catch (err) {
		console.error(err);
		throw error(500, 'Error cargando datos del estudiante');
	}
};

export const actions: Actions = {
	planificacion: async ({ request, fetch }) => {
		const formData = await request.formData();
		const materia = formData.get('materia') as string;

		try {
			const { base64, type } = await obtenerPlanificacion(fetch, materia);
			return {
				base64,
				type,
				message: 'Planificación obtenida exitosamente',
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al obtener planificación:', error);
			return {
				type: 'failure',
				message: error.message
			};
		}
	},

	constancia: async ({ locals: { usuario }, fetch }) => {
		const cedula = usuario?.cedula ?? '';
		try {
			const { base64, type } = await obtenerConstancia(fetch, cedula);
			return {
				type: type,
				base64,
				message: 'Constancia obtenida exitosamente',
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al obtener constancia:', error);
			return {
				type: 'failure',
				message: error.message
			};
		}
	}
};