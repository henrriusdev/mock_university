import { actualizarConfiguracion, crearConfiguracion, obtenerConfiguracion } from '$lib';
import { redirect } from '@sveltejs/kit';
import { format } from 'date-fns';
import type { Configuracion } from '../../../app';
import type { Actions, PageServerLoad } from './$types';

export const load = (async ({ fetch, locals: { usuario } }) => {
	if (!['administrador', 'coordinador'].includes(usuario!.rol.nombre)) {
		throw redirect(302, '/' + usuario!.rol.nombre);
	}
	const configuracion = await obtenerConfiguracion(fetch);
	return { configuracion };
}) satisfies PageServerLoad;

export const actions: Actions = {
	default: async ({ request, fetch }) => {
		const data = await request.formData();

		const ciclo = `${data.get('cicloYear')}-${data.get('cicloPeriodo')}`;
		const num_porcentaje = Number(data.get('num_porcentaje'));
		const num_cuotas = Number(data.get('num_cuotas'));

		const porcentajes = Array.from({ length: num_porcentaje }, (_, i) =>
			Number(data.get(`porcentajes[${i}]`))
		);

		const cuotas = data.getAll('cuotas').map((cuota) => cuota.toString());

		const horario_inicio = new Date(data.get('horario_inicio')!.toString());
		const horario_fin = new Date(data.get('horario_fin')!.toString());
		
		const payload: Configuracion = {
			ciclo,
			num_porcentaje,
			num_cuotas,
			porcentajes,
			cuotas: cuotas,
			horario_inicio: format(horario_inicio, 'dd/MM/yyyy'),
			horario_fin: format(horario_fin, 'dd/MM/yyyy')
		};
		
		// Verificamos si ya existe config para decidir acción
		const configActual = await obtenerConfiguracion(fetch);
		try {
			if (configActual) {
				await actualizarConfiguracion(fetch, payload);
			} else {
				await crearConfiguracion(fetch, payload);
			}
			return {
				type: 'success',
				message: 'Configuración actualizada exitosamente',
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al actualizar configuración:', e);
			return {
				type: 'failure',
				message: e.message
			};
		}
	}
};