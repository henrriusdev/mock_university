import { generarReporte } from '$lib';
import type { Actions } from './$types';

export const actions: Actions = {
	generarReporte: async ({ fetch, request }) => {
		const formData = await request.formData();

		const tipo = formData.get('tipo');
		const filtro = formData.get('filtro');
		const fecha = formData.get('fecha');
		const fi = formData.get('fi');
                const ff = formData.get('ff');

		const params = new URLSearchParams();
		if (tipo) params.set('tipo', tipo.toString());
		if (filtro && filtro !== 'todos') params.set('f', filtro.toString());
		if (tipo === 'dia' && fecha) params.set('d', fecha.toString());
		if ((tipo === 'fechas' || tipo === 'monto') && fi && ff) {
			params.set('fi', fi.toString());
			params.set('ff', ff.toString());
		}

		try {
			const { base64, filename } = await generarReporte(fetch, params.toString());
			return {
				base64,
				type: 'application/pdf',
				filename,
				message: 'Reporte generado exitosamente',
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al generar reporte:', e);
			return {
				type: 'failure',
				message: e.message
			};
		}
	}
};
