import {
	actualizarDocente,
	actualizarUsuario,
	crearDocente,
	crearUsuario,
	eliminarDocente,
	obtenerDocentes,
	toggleDocenteStatus,
	type DocenteReq
} from '$lib';
import { redirect } from '@sveltejs/kit';
import type { Usuario } from '../../../app';
import type { Actions, PageServerLoad } from './$types';
import { format, parse } from 'date-fns';

export type ErroresDocente = {
	cedula?: string;
	nombre?: string;
	correo?: string;
	password?: string;
	titulo?: string;
	fecha_ingreso?: string;
	activo?: string;
};

export const load: PageServerLoad = async ({ fetch, parent }) => {
	const { rol } = await parent();
	// solo administradores y coordinadores pueden acceder a la vista de docentes
	if (!['administrador', 'coordinador'].includes(rol)) {
		redirect(302, '/' + rol);
	}
	try {
		const res = await obtenerDocentes(fetch);
		return { docentes: res };
	} catch (error) {
		console.error('Error al obtener carreras:', error);
		return { carreras: [], docentes: [] };
	}
};

export const actions: Actions = {
	// Acción para cambiar el estado del docente
	toggleStatus: async ({ request, fetch }) => {
		const formData = await request.formData();
		const cedula = formData.get('cedula')?.toString() || '';

		try {
			const docente = await toggleDocenteStatus(fetch, cedula);
			if (!docente) {
				return {
					type: 'failure',
					message: 'Error al cambiar el estado del docente'
				};
			}

			const statusMessage = docente.activo ? 'Docente activado exitosamente' : 'Docente inactivado exitosamente';
			
			return {
				type: 'success',
				message: statusMessage,
				invalidate: true
			};
		} catch (error) {
			console.error('Error al cambiar el estado del docente:', error);
			return { 
				type: 'failure', 
				message: error instanceof Error ? error.message : 'Error desconocido' 
			};
		}
	},
	// Acción para crear un usuario y un docente
	create: async ({ request, fetch }) => {
		const payload = Object.fromEntries(await request.formData()) as unknown as DocenteReq & Usuario;

		const errores = validarPayload(
			payload as unknown as Record<string, string | number | boolean>,
			false
		);
		if (Object.keys(errores).length > 0) {
			const errorString = Object.entries(errores)
				.map(([campo, mensaje]) => `${campo}: ${mensaje}`)
				.join(', ');
			return {
				type: 'failure',
				message: errorString
			};
		}
		const usuario: Partial<Usuario & { rol_id: number }> = {
			id: 0,
			cedula: payload.cedula,
			correo: payload.correo,
			activo: true,
			nombre: payload.nombre,
			rol_id: 6
		};

		try {
			const { data }: { data: Usuario } = await crearUsuario(fetch, usuario);
			usuario.id = data.id;
		} catch (error) {
			console.error('Error al crear usuario:', error);
			return {
				success: false,
				message: error instanceof Error ? error.message : 'Error desconocido'
			};
		}

		const docente: DocenteReq = {
			titulo: payload.titulo,
			fecha_ingreso: format(new Date(payload.fecha_ingreso), 'dd/MM/yyyy'),
			usuario_id: usuario.id
		};

		try {
			await crearDocente(fetch, docente);
			return {
				type: 'success',
				message: `Docente ${usuario.nombre} creado exitosamente`,
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al crear docente:', e);
			return {
				type: 'failure',
				message: e.message
			};
		}
	},

	edit: async ({ request, fetch }) => {
		const payload = Object.fromEntries(await request.formData()) as unknown as DocenteReq &
			Usuario & { id_docente: number };

		const errores = validarPayload(
			payload as unknown as Record<string, string | number | boolean>,
			true
		);
		if (Object.keys(errores).length > 0) {
			return {
				type: 'failure',
				message: Object.entries(errores)
					.map(([campo, mensaje]) => `${campo}: ${mensaje}`)
					.join(', ')
			};
		}

		const usuario: Partial<Usuario & { rol_id: number }> = {
			cedula: payload.cedula,
			correo: payload.correo,
			activo: true,
			nombre: payload.nombre,
			rol_id: 6
		};

		try {
			await actualizarUsuario(fetch, payload.id, usuario);
		} catch (error) {
			console.error('Error al editar usuario:', error);
			return {
				success: false,
				message: error instanceof Error ? error.message : 'Error desconocido'
			};
		}

		const docente: DocenteReq = {
			titulo: payload.titulo,
			fecha_ingreso: format(new Date(payload.fecha_ingreso), 'dd/MM/yyyy'),
			usuario_id: payload.id
		};

		try {
			await actualizarDocente(fetch, payload.id_docente, docente);
			return {
				type: 'success',
				message: `Docente ${usuario.nombre} actualizado exitosamente`,
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al actualizar docente:', e);
			return {
				type: 'failure',
				message: e.message
			};
		}
	},

	delete: async ({ request, fetch }) => {
		const formData = await request.formData();
		const cedula = formData.get('cedula')?.toString() || '';

		try {
			await eliminarDocente(fetch, cedula);
			return {
				type: 'success',
				message: 'Docente eliminado exitosamente',
				invalidate: true
			};
		} catch (error) {
			console.error('Error al eliminar docente:', error);
			return {
				success: false,
				message: error instanceof Error ? error.message : 'Error desconocido'
			};
		}
	}
};

function validarPayload(
	payload: Record<string, string | number | boolean>,
	isEditing: boolean
): ErroresDocente {
	const errores: ErroresDocente = {};

	// Campos requeridos comunes
	const camposBase: (keyof ErroresDocente)[] = [
		'cedula',
		'nombre',
		'correo',
		'titulo',
		'fecha_ingreso'
	];

	for (const campo of camposBase) {
		const valor = payload[campo];
		if (!valor || valor.toString().trim() === '') {
			errores[campo] = `El campo ${campo} es requerido`;
		}
	}

	// Validaciones específicas
	if (payload.correo && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(payload.correo as string)) {
		errores.correo = 'Correo inválido';
	}

	return errores;
}
