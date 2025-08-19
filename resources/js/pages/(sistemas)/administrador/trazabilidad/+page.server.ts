import { exportarTrazabilidad } from '$lib/servicios/archivos';
import { filtrarTrazabilidad } from '$lib/servicios/trazabilidad';
import { format } from 'date-fns';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, url }) => {
	const busqueda = url.searchParams.get('busqueda') || '';
	const fechaDesde = url.searchParams.get('fechaDesde') || '';
	const fechaHasta = url.searchParams.get('fechaHasta') || '';
	const rol = url.searchParams.get('rol') || '';

	try {
		const registros = await filtrarTrazabilidad({ busqueda, fechaDesde, fechaHasta, rol }, fetch);
		return { registros };
	} catch (error) {
		console.error('Error al cargar los datos de trazabilidad:', error);
		return { registros: [] };
	}
};

export const actions: Actions = {
	filtrar: async ({ request }) => {
		const formData = await request.formData();
		const busqueda = formData.get('busqueda')?.toString() || '';
		const fechaDesde = formData.get('fechaDesde')?.toString() || '';
		const fechaHasta = formData.get('fechaHasta')?.toString() || '';
		const rol = formData.get('rol')?.toString() || '';

		const fechaDesdeDate = format(new Date(fechaDesde), 'yyyy-MM-dd');
		const fechaHastaDate = format(new Date(fechaHasta), 'yyyy-MM-dd');

		const query = new URLSearchParams({
			busqueda,
			fechaDesde: fechaDesdeDate,
			fechaHasta: fechaHastaDate,
			rol
		}).toString();

		// Redirigir a la misma página con los parámetros de búsqueda
		const location = `/administrador/trazabilidad?${query}`;
		return { success: true, invalidate: true, bypasstoast: true, location };
	},

	exportar: async ({ request, fetch }) => {
		const formData = await request.formData();

		const formato = formData.get('formato')?.toString() || 'csv';
		const busqueda = formData.get('busqueda')?.toString() || '';
		const fechaDesde = formData.get('fechaDesde')?.toString() || '';
		const fechaHasta = formData.get('fechaHasta')?.toString() || '';
		const rol = formData.get('rol')?.toString() || '';

		try {
			// Verificar que hay datos para exportar
			const registros = await filtrarTrazabilidad(
				{
					busqueda,
					fechaDesde: format(new Date(fechaDesde), 'yyyy-MM-dd'),
					fechaHasta: format(new Date(fechaHasta), 'yyyy-MM-dd'),
					rol
				},
				fetch
			);

			if (registros.length === 0) {
				return {
					type: 'failure',
					message: 'No hay datos para exportar'
				}
			}

			// Generar el archivo de exportación
			const { base64, type } = await exportarTrazabilidad(
				{
					formato,
					busqueda,
					fechaDesde: format(new Date(fechaDesde), 'yyyy-MM-dd'),
					fechaHasta: format(new Date(fechaHasta), 'yyyy-MM-dd'),
					rol
				},
				fetch
			);

			return {
				base64,
				type,
				message: 'Archivo exportado exitosamente',
				invalidate: true
			};
		} catch (error) {
			console.error('Error al exportar los datos:', error);
			return {
				type: 'failure',
				message: 'Error al exportar los datos. Por favor, inténtelo de nuevo.'
			};
		}
	}
};
