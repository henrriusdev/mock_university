import { actualizarNota, obtenerMateria } from '$lib';
import { redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import { crearPeticion } from '$lib/servicios/peticiones';
import type { Peticion } from '../../../../app';
import { obtenerReporteNotas, subirPlanificacion } from '$lib/servicios/archivos';

export const load = (async ({ fetch, parent, params }) => {
	const { rol } = await parent();
	if (['estudiante', 'caja'].includes(rol)) {
		throw redirect(302, '/' + rol);
	}

	const materia = await obtenerMateria(fetch, params.materia);
	return { materia };
}) satisfies PageServerLoad;

export const actions: Actions = {
	editar: async ({ fetch, request, locals: { usuario } }) => {
		const form = await request.formData();

		let payload = Object.fromEntries(form) as unknown as {
			cedula_estudiante: string;
			nombre_campo: string;
			valor: string | number;
			materia: string;
			peticion?: string;
			observacion?: string;
		};
		const { peticion, observacion } = payload;

		payload = {
			...payload,
			peticion: undefined,
			observacion: undefined,
			valor: Number(payload.valor)
		};

		if (peticion === 'true') {
			const body = {
				campo: payload.nombre_campo,
				descripcion: observacion!,
				estado: "Pendiente",
				id_docente: usuario?.id ?? 1,
				id_estudiante: payload.cedula_estudiante,
				id_materia: payload.materia,
				valor: payload.valor
			};

			try {
				await crearPeticion(fetch, body as Omit<Peticion, "id">);
				return {
					type: 'success',
					message: 'Petición creada exitosamente',
					invalidate: true
				};
			} catch (error: any) {
				console.error('Error al crear petición:', error);
				return {
					type: 'failure',
					message: error.message
				};
			}
		}

		try {
			await actualizarNota(fetch, payload);
			return {
				type: 'success',
				message: 'Nota actualizada exitosamente',
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al actualizar nota:', error);
			return {
				type: 'failure',
				message: error.message
			};
		}
	},

	subir: async ({ fetch, request }) => {
		const formData = await request.formData();

		try {
			await subirPlanificacion(fetch, formData);
			return {
				type: 'success',
				message: 'Planificación subida exitosamente',
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al subir planificación:', error);
			return {
				type: 'failure',
				message: error.message
			};
		}
	},

	notas: async ({ fetch, params }) => {
		try {
			const { base64, type } = await obtenerReporteNotas(fetch, params.materia);
			return {
				type,
				base64,
				message: 'Reporte de notas obtenido exitosamente',
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al obtener reporte de notas:', error);
			return {
				type: 'failure',
				message: error.message
			};
		}
	}
};
