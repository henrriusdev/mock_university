import { actualizarPeticion, obtenerPeticiones } from '$lib/servicios/peticiones';
import type { Actions, PageServerLoad } from './$types';
import { obtenerDocentes, obtenerEstudiantes } from '$lib';

export const load = (async ({ fetch, }) => {
	const peticiones = await obtenerPeticiones(fetch);
	const docentes = await obtenerDocentes(fetch);
	const estudiantes = await obtenerEstudiantes(fetch);
	return { peticiones, docentes, estudiantes };
}) satisfies PageServerLoad;

export const actions: Actions = {
	aprobar: async ({ fetch, request }) => {
		try {
			const data = await request.formData();
			const id = Number(data.get('id'));
			await actualizarPeticion(fetch, id, { estado: 'Aprobado' });
			return {
				type: 'success',
				message: 'Petición aprobada exitosamente',
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al aprobar petición:', e);
			return {
				type: 'failure',
				message: e.message
			};
		}
	},

	rechazar: async ({ fetch, request }) => {
		try {
			const data = await request.formData();
			const id = Number(data.get('id'));
			await actualizarPeticion(fetch, id, { estado: 'Denegado' });
			return {
				type: 'success',
				message: 'Petición denegada exitosamente',
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al denegar petición:', e);
			return {
				type: 'failure',
				message: e.message
			};
		}
	}
};