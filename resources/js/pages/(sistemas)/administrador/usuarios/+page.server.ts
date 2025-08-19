import { actualizarUsuario, crearUsuario, eliminarUsuario, obtenerUsuarios, toggleUsuarioStatus } from '$lib';
import { descargarExcel, subirExcelUsuarios } from '$lib/servicios/archivos';
import type { Usuario } from '../../../../app';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
	try {
		const res = await obtenerUsuarios(fetch);
		return { usuarios: res };
	} catch (error) {
		console.error('Error al obtener carreras:', error);
		return { usuarios: [] };
	}
};

export const actions: Actions = {
	// Acción para cambiar el estado del usuario
	toggleStatus: async ({ request, fetch }) => {
		const formData = await request.formData();
		const cedula = formData.get('cedula')?.toString() || '';

		try {
			const usuario = await toggleUsuarioStatus(fetch, cedula);
			if (!usuario) {
				return {
					type: 'failure',
					message: 'Error al cambiar el estado del usuario'
				};
			}

			const statusMessage = usuario.activo ? 'Usuario activado exitosamente' : 'Usuario inactivado exitosamente';
			
			return {
				type: 'success',
				message: statusMessage,
				invalidate: true
			};
		} catch (error) {
			console.error('Error al cambiar el estado del usuario:', error);
			return { 
				type: 'failure', 
				message: error instanceof Error ? error.message : 'Error desconocido' 
			};
		}
	},
	// Acción para crear un usuario y un usuario
	create: async ({ request, fetch }) => {
		const payload = Object.fromEntries(await request.formData()) as unknown as Omit<Usuario, "rol"> & { rol: number };

		const usuario: Partial<Usuario & { rol_id: number }> = {
			id: 0,
			cedula: payload.cedula,
			correo: payload.correo,
			activo: true,
			nombre: payload.nombre,
			rol_id: payload.rol
		};

		try {
			const { data }: { data: Usuario } = await crearUsuario(fetch, usuario);
			return {
				type: 'success',
				message: `Usuario ${data.nombre} creado exitosamente`,
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al crear usuario:', error);
			return {
				type: 'failure',
				message: error.message
			};
		}
	},

	edit: async ({ request, fetch }) => {
		const payload = Object.fromEntries(await request.formData()) as unknown as Omit<
			Usuario,
			'rol'
		> & {
			rol: number;
		};

		const usuario: Partial<Usuario & { rol_id: number }> = {
			cedula: payload.cedula,
			correo: payload.correo,
			activo: true,
			nombre: payload.nombre,
			rol_id: payload.rol
		};

		try {
			await actualizarUsuario(fetch, payload.id, usuario);
			return {
				type: 'success',
				message: `Usuario ${usuario.nombre} actualizado exitosamente`,
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al editar usuario:', error);
			return {
				type: 'failure',
				message: error.message
			};
		}
	},

	delete: async ({ request, fetch }) => {
		const formData = await request.formData();
		const cedula = formData.get('cedula')?.toString() || '';

		try {
			await eliminarUsuario(fetch, cedula);
			return {
				type: 'success',
				message: 'Usuario eliminado exitosamente',
				invalidate: true
			};
		} catch (error: any) {
			console.error('Error al eliminar usuario:', error);
			return {
				type: 'failure',
				message: error.message
			};
		}
	},

	descargar: async ({fetch }) => {
		try {
			const { base64, type } = await descargarExcel(fetch);
			return {
				type,
				base64,
				message: 'Archivo descargado exitosamente'
			}
		} catch (error) {
			console.error('Error al descargar el archivo:', error);
			return {
				type: 'failure',
				message: 'Error al descargar el archivo'
			};
		}
	},
	cargar: async ({ request, fetch }) => {
		const formData = await request.formData();
		const archivo = formData.get('archivo');

		if (!archivo || !(archivo instanceof File)) {
			return {type: 'failure', message: 'No se ha seleccionado un archivo válido'};
		}

		const nuevoForm = new FormData();
		nuevoForm.set('file', archivo);

		try {
			await subirExcelUsuarios(fetch, nuevoForm);
			return {
				type: 'success',
				message: 'Archivo subido exitosamente',
				invalidate: true
			}
		} catch (err) {
			console.error(err);
			return {
				type: 'failure',
				message: 'Error al subir el archivo'
			}
		}
	}
};
