import type { Actions, PageServerLoad } from './$types';
import { obtenerConfiguracion, obtenerHistoricoMaterias } from '$lib';
import { obtenerPlanificacion } from '$lib/servicios/archivos';

export const load: PageServerLoad = async ({fetch}) => {
	const historico = await obtenerHistoricoMaterias(fetch);
	const config = await obtenerConfiguracion(fetch);
	const materiasActuales = historico?.filter(m => config?.ciclo ? m.ciclo === config?.ciclo : true);
	return {historico: materiasActuales};
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
	}
};