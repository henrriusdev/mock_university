import {
	actualizarEstudiante,
	actualizarUsuario,
	crearEstudiante,
	crearUsuario,
	eliminarEstudiante,
	toggleEstudianteStatus,
	obtenerCarreras,
	obtenerEstudiantes,
	type EstudianteReq
} from '$lib';
import { redirect } from '@sveltejs/kit';
import { format } from 'date-fns';
import type { Usuario } from '../../../app';
import type { Actions, PageServerLoad } from './$types';

type ErroresEstudiante = {
	cedula?: string;
	nombre?: string;
	correo?: string;
	fecha_nac?: string;
	carrera?: string;
	sexo?: string;
	semestre?: string;
	promedio?: string;
	direccion?: string;
	password?: string;
};

export const load: PageServerLoad = async ({ fetch, parent }) => {
	const { rol } = await parent();
	if (!['caja', 'administrador', 'coordinador', 'control'].includes(rol)) {
		redirect(302, '/'+rol);
	}
	try {
		const res = await obtenerEstudiantes(fetch);
		const estudiantes = res.map((est) => ({
			...est,
			fecha_nacimiento: format(new Date(est.fecha_nacimiento), 'dd/MM/yyyy')
		}));
		const carreras = await obtenerCarreras(fetch);
		return { carreras, estudiantes };
	} catch (error) {
		console.error('Error al obtener carreras:', error);
		return { carreras: [], estudiantes: [] };
	}
};

export const actions: Actions = {
	create: async ({ request, fetch }) => {
		const payload = Object.fromEntries(await request.formData()) as unknown as EstudianteReq &
			(Usuario);
		
		const errores = validarPayload(payload as unknown as Record<string, string | number | boolean>, false);
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
			rol_id: 5
		};

		try {
			const { data }: { data: Usuario } = await crearUsuario(fetch, usuario);
			usuario.id = data.id;
		} catch (error) {
			console.error('Error al crear usuario:', error);
			return { 
				type: 'failure', 
				message: error instanceof Error ? error.message : 'Error desconocido' 
			};
		}

		const estudiante: EstudianteReq = {
			carrera: payload.carrera,
			edad: payload.edad,
			fecha_nac: format(new Date(payload.fecha_nac), "dd/MM/yyyy"),
			direccion: payload.direccion,
			semestre: payload.semestre,
			sexo: payload.sexo,
			promedio: payload.promedio,
			usuario: usuario.id
		};

		try {
			await crearEstudiante(fetch, estudiante);
			return {
				type: 'success',
				message: 'Estudiante creado exitosamente',
				invalidate: true
			};
		} catch (e) {
			console.error('Error al crear estudiante:', e);
			return { 
				type: 'failure', 
				message: e instanceof Error ? e.message : 'Error desconocido' 
			};
		}
	},

	edit: async ({ request, fetch }) => {
		const payload = Object.fromEntries(await request.formData()) as unknown as EstudianteReq &
			Usuario & { id_estudiante: number };
		
		const errores = validarPayload(payload as unknown as Record<string, string | number | boolean>, true);
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
			cedula: payload.cedula,
			correo: payload.correo,
			nombre: payload.nombre,
			rol_id: 5
		};

		try {
			await actualizarUsuario(fetch, payload.id, usuario);
		} catch (error) {
			console.error('Error al editar usuario:', error);
			return { 
				type: 'failure', 
				message: error instanceof Error ? error.message : 'Error desconocido' 
			};
		}

		const estudiante: EstudianteReq = {
			carrera: payload.carrera,
			edad: payload.edad,
			fecha_nac: format(new Date(payload.fecha_nac), 'dd/MM/yyyy'),
			direccion: payload.direccion,
			semestre: payload.semestre,
			sexo: payload.sexo,
			promedio: payload.promedio,
			usuario: payload.id
		};

		try {
			await actualizarEstudiante(fetch, payload.id_estudiante, estudiante);
			return {
				type: 'success',
				message: 'Estudiante actualizado exitosamente',
				invalidate: true
			};
		} catch (e: any) {
			console.error('Error al actualizar estudiante:', e);
			return { 
				type: 'failure', 
				message: e instanceof Error ? e.message : 'Error desconocido' 
			};
		}
	},

	delete: async ({ request, fetch }) => {
		const formData = await request.formData();
		const cedula = formData.get('cedula')?.toString() || '';

		try {
			await eliminarEstudiante(fetch, cedula);
			return {
				type: 'success',
				message: 'Estudiante eliminado exitosamente',
				invalidate: true
			};
		} catch (error) {
			console.error('Error al eliminar estudiante:', error);
			return { 
				type: 'failure', 
				message: error instanceof Error ? error.message : 'Error desconocido' 
			};
		}
	},

	toggleStatus: async ({ request, fetch }) => {
		const formData = await request.formData();
		const cedula = formData.get('cedula')?.toString() || '';

		try {
			const response = await toggleEstudianteStatus(fetch, cedula);
			if (!response.ok) {
				return {
					type: 'failure',
					message: response.data?.message || 'Error al cambiar el estado del estudiante'
				};
			}
			
			const isActivating = response.data?.activo;
			const statusMessage = isActivating ? 'Estudiante activado exitosamente' : 'Estudiante inactivado exitosamente';
			
			return {
				type: 'success',
				message: statusMessage,
				invalidate: true
			};
		} catch (error) {
			console.error('Error al cambiar el estado del estudiante:', error);
			return { 
				type: 'failure', 
				message: error instanceof Error ? error.message : 'Error desconocido' 
			};
		}
	}
};

function validarPayload(payload: Record<string, string | number | boolean>, isEditing: boolean): ErroresEstudiante {
	const errores: ErroresEstudiante = {};

	const camposBase: (keyof ErroresEstudiante)[] = [
		'cedula',
		'nombre',
		'correo',
		'fecha_nac',
		'carrera',
		'sexo',
		'semestre',
		'promedio',
		'direccion'
	];

	for (const campo of camposBase) {
		const valor = payload[campo];
		if (!valor || valor.toString().trim() === '') {
			errores[campo] = `El campo ${campo} es requerido`;
		}
	}

	if (payload.correo && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(payload.correo as string)) {
		errores.correo = 'Correo inválido';
	}

	if (
		payload.promedio &&
		(isNaN(payload.promedio as number) || Number(payload.promedio) < 0 || Number(payload.promedio) > 20)
	) {
		errores.promedio = 'Promedio debe estar entre 0 y 20';
	}

	return errores;
}
