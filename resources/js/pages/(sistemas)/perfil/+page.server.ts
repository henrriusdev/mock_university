import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import { cambiarPassword, subirFotoPerfil, eliminarFotoPerfil } from '$lib/servicios/autenticacion';
import { configurarPreguntaSeguridad } from '$lib/servicios/pregunta-seguridad';

export const load = (async ({locals: { usuario } }) => {
	const photoUrl = usuario?.foto ? `http://127.0.0.1:8000/api/usuario/photo/${usuario.foto}` : null;
	
	return {
		usuario,
		photoUrl
	}
}) satisfies PageServerLoad;

export const actions: Actions = {
	cambiarPassword: async ({ request, fetch }) => {
		const data = await request.formData();
		const currentPassword = data.get('currentPassword') as string;
		const newPassword = data.get('newPassword') as string;
		const confirmPassword = data.get('confirmPassword') as string;

		if (!currentPassword || !newPassword || !confirmPassword) {
			return fail(400, {
				type: 'error',
				message: 'Todos los campos son requeridos'
			});
		}

		if (newPassword !== confirmPassword) {
			return fail(400, {
				type: 'error',
				message: 'Las contraseñas no coinciden'
			});
		}

		try {
			await cambiarPassword(fetch, currentPassword, newPassword);
			return {
				type: 'success',
				message: 'Contraseña actualizada correctamente'
			};
		} catch (e) {
			console.error(e);
			return fail(400, {
				type: 'error',
				message: e instanceof Error ? e.message : 'Error al actualizar la contraseña'
			});
		}
	},

	configurarPregunta: async ({ request, fetch }) => {
		const data = await request.formData();
		const preguntas = [];

		// Recolectar las 3 preguntas y respuestas
		for (let i = 0; i < 3; i++) {
			const pregunta = data.get(`pregunta${i}`) as string;
			const respuesta = data.get(`respuesta${i}`) as string;

			if (!pregunta || !respuesta) {
				return {
					type: 'failure',
					message: 'Todas las preguntas y respuestas son requeridas'
				}
			}

			preguntas.push({ pregunta, respuesta });
		}

		try {
			// Configurar todas las preguntas en un solo request
			await configurarPreguntaSeguridad(fetch, preguntas);

			return {
				type: 'success',
				message: 'Preguntas de seguridad configuradas correctamente'
			};
		} catch (e) {
			console.error(e);
			return {
				type: 'failure',
				message: e instanceof Error ? e.message : 'Error al configurar las preguntas de seguridad'
			};
		}
	},

	subirFoto: async ({ request, fetch }) => {
		const data = await request.formData();
		const file = data.get('file') as File;

		if (!file || file.size === 0) {
			return fail(400, {
				type: 'error',
				message: 'No se seleccionó ningún archivo'
			});
		}

		try {
			const filename = await subirFotoPerfil(fetch, file);
			return {
				type: 'success',
				message: 'Foto subida exitosamente',
				filename
			};
		} catch (e) {
			console.error(e);
			return fail(400, {
				type: 'error',
				message: e instanceof Error ? e.message : 'Error al subir la foto'
			});
		}
	},

	eliminarFoto: async ({ fetch }) => {
		try {
			await eliminarFotoPerfil(fetch);
			return {
				type: 'success',
				message: 'Foto eliminada exitosamente'
			};
		} catch (e) {
			console.error(e);
			return fail(400, {
				type: 'error',
				message: e instanceof Error ? e.message : 'Error al eliminar la foto'
			});
		}
	}
};
