import { inscribirMaterias, obtenerConfiguracion, obtenerMateriasInscritas } from '$lib';
import { obtenerMateriasDisponibles } from '$lib/servicios/materias';
import { isAfter } from 'date-fns';
import type { Actions, PageServerLoad } from './$types';
import type { MateriaDisponible } from '../../../../app';
import { addToast } from '$lib';

export const load: PageServerLoad = async ({ fetch, locals: { usuario } }) => {
	const config = await obtenerConfiguracion(fetch);
	const { horario_inicio: inicio, horario_fin: fin }: { horario_inicio: Date; horario_fin: Date } =
		config;
	const materiasInscritas = await obtenerMateriasInscritas(fetch);
	if (isAfter(inicio, new Date()) || isAfter(new Date(), fin)) {
		return {
			materiasDisponibles: [] as MateriaDisponible[],
			materiasInscritas,
			inscripcionAbierta: false
		};
	}

	if (materiasInscritas.length > 0) {
		addToast({
			type: 'info',
			message: 'Ya tienes tu horario inscrito'
		});
		return {
			materiasDisponibles: [] as MateriaDisponible[],
			inscripcionAbierta: false,
			materiasInscritas
		};
	}
	const materiasDisponibles = await obtenerMateriasDisponibles(fetch, usuario?.cedula ?? '');

	return {
		materiasDisponibles,
		materiasInscritas,
		inscripcionAbierta: true
	};
};

export const actions: Actions = {
	default: async ({ request, fetch }) => {
		const form = await request.formData();
		const asignaciones: number[] = [];
		
		// Handle both 'asignaciones' and 'materias' for backward compatibility
		const formAsignaciones = form.getAll('asignaciones').length > 0 
			? form.getAll('asignaciones') 
			: form.getAll('materias');
		
		for (const asignacion of formAsignaciones) {
			const id = parseInt(asignacion.toString());
			if (!isNaN(id) && !asignaciones.includes(id)) {
				asignaciones.push(id);
			}
		}

		try {
			const payload = {
				asignaciones: asignaciones
			};

			await inscribirMaterias(fetch, payload);
			return {
				type: 'success',
				message: 'Secciones inscritas exitosamente',
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al inscribir secciones:', e);
			return {
				type: 'failure',
				message: e.message
			};
		}
	}
};
