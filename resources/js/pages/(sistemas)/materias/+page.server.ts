import {
	crearMateria,
	actualizarMateria,
	eliminarMateria,
	obtenerMaterias,
	obtenerCarreras,
	obtenerDocentes,
} from '$lib';
import { redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import type { MateriaReq, AsignacionReq } from '$lib/types';

export const load: PageServerLoad = async ({ fetch, parent }) => {
        const { rol } = await parent();
        // la gestión de materias solo está disponible para administradores y coordinadores
        if (!['administrador', 'coordinador', 'control'].includes(rol)) {
                redirect(302, '/' + rol);
        }

	try {
		const data = await obtenerMaterias(fetch);
		const carreras = await obtenerCarreras(fetch);
		const docentes = await obtenerDocentes(fetch);

		return { materias: data.materias, carreras, docentes };
	} catch (error) {
		console.error('Error al obtener las materias:', error);
		return { materias: [], carreras: [], docentes: [] };
	}
};

export const actions: Actions = {
	create: async ({ request, fetch }) => {
		const form = await request.formData();
		// Helper function for safe number conversion
		const safeNumber = (value: FormDataEntryValue | null, defaultValue = 0): number => {
			if (!value) return defaultValue;
			const num = Number(value);
			return isNaN(num) ? defaultValue : num;
		};

		const asignaciones: AsignacionReq[] = JSON.parse(form.get('asignaciones')?.toString() || '[]');

		const payload: MateriaReq = {
			id: form.get('id')?.toString() || '',
			nombre: form.get('nombre')?.toString() || '',
			prelacion: form.get('prelacion')?.toString() || '',
			unidad_credito: safeNumber(form.get('unidad_credito')),
			hp: safeNumber(form.get('hp')),
			ht: safeNumber(form.get('ht')),
			semestre: safeNumber(form.get('semestre'), 1),
			id_carrera: safeNumber(form.get('id_carrera')).toString(),
			ciclo: form.get('ciclo')?.toString() || '',
			maximo: safeNumber(form.get('maximo'), 30),
			asignaciones: asignaciones
                };

		const errores = validarPayload(payload);
		if (Object.keys(errores).length > 0) {
			const errorMessages = Object.values(errores).join(', ');
			return {
				type: 'failure',
				message: errorMessages
			};
		}

		try {
			await crearMateria(fetch, payload);
			return {
				type: 'success',
				message: 'Materia creada exitosamente',
				invalidate: true
			};
		} catch (error) {
			console.error('Error al crear materia:', error);
			return { type: 'failure', message: error instanceof Error ? error.message : 'Error desconocido' };
		}
	},

	// Acción para editar una materia
	edit: async ({ request, fetch }) => {
		const form = await request.formData();
		// Helper function for safe number conversion
		const safeNumber = (value: FormDataEntryValue | null, defaultValue = 0): number => {
			if (!value) return defaultValue;
			const num = Number(value);
			return isNaN(num) ? defaultValue : num;
		};

		const asignaciones: AsignacionReq[] = JSON.parse(form.get('asignaciones')?.toString() || '[]');

		const payload: MateriaReq & { id: string } = {
			id: form.get('id')?.toString() as string,
			nombre: form.get('nombre')?.toString() as string,
			prelacion: form.get('prelacion')?.toString() as string,
			unidad_credito: safeNumber(form.get('unidad_credito')),
			hp: safeNumber(form.get('hp')),
			ht: safeNumber(form.get('ht')),
			semestre: safeNumber(form.get('semestre'), 1),
			id_carrera: safeNumber(form.get('id_carrera')).toString(),
			ciclo: form.get('ciclo')?.toString() as string,
			maximo: safeNumber(form.get('maximo'), 30),
			asignaciones: asignaciones
                };

		const errores = validarPayload(payload);
		if (Object.keys(errores).length > 0) {
			const errorMessages = Object.values(errores).join(', ');
			return {
				type: 'failure',
				message: errorMessages
			};
		}

		try {
			await actualizarMateria(fetch, payload.id, payload);
			return {
				type: 'success',
				message: 'Materia actualizada exitosamente',
				invalidate: true
			};
		} catch (error) {
			console.error('Error al editar materia:', error);
			return { type: 'failure', message: error instanceof Error ? error.message : 'Error desconocido' };
		}
	},

	// Acción para eliminar una materia
	delete: async ({ request, fetch }) => {
		const formData = await request.formData();
		const id = formData.get('id')?.toString() || '';

		try {
			await eliminarMateria(fetch, id);
			return {
				type: 'success',
				message: 'Materia eliminada exitosamente',
				invalidate: true
			};
		} catch (error) {
			console.error('Error al eliminar materia:', error);
			return { type: 'failure', message: error instanceof Error ? error.message : 'Error desconocido' };
		}
	}
};

function validarPayload(
	payload: Record<string, string | number | boolean | AsignacionReq[]>
): Record<string, string> {
	const errores: Record<string, string> = {};
	const camposBase: (keyof typeof payload)[] = [
		'nombre',
                'semestre',
                'id_carrera',
                'maximo'
        ];

	for (const campo of camposBase) {
		const valor = payload[campo];
		if (!valor || valor.toString().trim() === '') {
			errores[campo] = `El campo ${campo} es requerido`;
		}
	}

	return errores;
}
