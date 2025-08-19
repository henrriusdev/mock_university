import {
	actualizarCoordinador,
	actualizarUsuario,
	type CoordinadorReq,
	crearCoordinador,
	crearUsuario,
	eliminarCoordinador,
	toggleStatusCoordinador,
	obtenerCarreras,
	obtenerCoordinadores
} from '$lib';
import type { Usuario } from '../../../../app';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
	try {
		const carreras = await obtenerCarreras(fetch);
		const res = await obtenerCoordinadores(fetch);
		return { coordinadores: res, carreras };
	} catch (error) {
		console.error('Error al obtener carreras:', error);
		return { carreras: [], coordinadores: [] };
	}
};

export const actions: Actions = {
	// Acción para crear un usuario y un coordinador
	create: async ({ request, fetch }) => {
		const payload = Object.fromEntries(await request.formData()) as unknown as CoordinadorReq &
			Usuario;

		const usuario: Partial<Usuario & { rol_id: number }> = {
			id: 0,
			cedula: payload.cedula,
			correo: payload.correo,
			activo: true,
			nombre: payload.nombre,
			rol_id: 4
		};

		try {
			const { data }: { data: Usuario } = await crearUsuario(fetch, usuario);
			usuario.id = data.id;
		} catch (error: any) {
			console.error('Error al crear usuario:', error);
			return {
				type: 'failure',
				message: 'Error al crear el usuario'
			};
		}

		const coordinador: CoordinadorReq = {
			carrera_id: payload.carrera_id,
			usuario_id: usuario.id,
			telefono: payload.telefono
		};

		try {
			await crearCoordinador(fetch, coordinador);
			return {
				type: 'success',
				message: 'Coordinador creado exitosamente',
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al crear coordinador:', e);
			return {
				type: 'failure',
				message: 'Error al crear el coordinador'
			};
		}
	},

	edit: async ({ request, fetch }) => {
		const payload = Object.fromEntries(await request.formData()) as unknown as CoordinadorReq &
			Usuario & { id_coordinador: number };

		const usuario: Partial<Usuario & { rol_id: number }> = {
			cedula: payload.cedula,
			correo: payload.correo,
			activo: true,
			nombre: payload.nombre,
			rol_id: 4
		};

		try {
			await actualizarUsuario(fetch, payload.id, usuario);
		} catch (error: any) {
			console.error('Error al editar usuario:', error);
			return {
				type: 'failure',
				message: 'Error al editar el usuario'
			};
		}

		const coordinador: CoordinadorReq = {
			carrera_id: payload.carrera_id,
			telefono: payload.telefono
		};

		try {
			await actualizarCoordinador(fetch, payload.id_coordinador, coordinador);
			return {
				type: 'success',
				message: 'Coordinador actualizado exitosamente',
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al actualizar coordinador:', e);
			return {
				type: 'failure',
				message: 'Error al actualizar el coordinador'
			};
		}
	},

	delete: async ({ request, fetch }) => {
		const formData = await request.formData();
		const cedula = formData.get('cedula')?.toString() || '';

		try {
			await eliminarCoordinador(fetch, cedula);
			return {
				type: 'success',
				message: 'Coordinador eliminado exitosamente',
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al eliminar coordinador:', error);
			return {
				type: 'failure',
				message: 'Error al eliminar el coordinador'
			};
		}
	},

	toggleStatus: async ({ request, fetch }) => {
		const formData = await request.formData();
		const cedula = formData.get('cedula')?.toString() || '';

		try {
			await toggleStatusCoordinador(fetch, cedula);
			return {
				type: 'success',
				message: 'Estado del coordinador actualizado exitosamente',
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al cambiar estado del coordinador:', error);
			return {
				type: 'failure',
				message: 'Error al cambiar estado del coordinador'
			};
		}
	}
};
