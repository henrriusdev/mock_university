import type { Actions } from './$types';
import { obtenerPreguntasSeguridad, verificarPreguntaSeguridad } from '$lib/servicios/pregunta-seguridad';
import { recuperarContrasena } from '$lib/servicios/autenticacion';
import type { PreguntaSeguridad } from '../../app';

// Almacenamiento en memoria de intentos por correo
const intentosPorCorreo = new Map<string, { count: number; lastAttempt: Date }>();

function getIntentos(correo: string): number {
	const intentos = intentosPorCorreo.get(correo);
	return intentos?.count || 0;
}

function checkIntentos(correo: string): boolean {
	const now = new Date();
	const intentos = intentosPorCorreo.get(correo);

	if (!intentos) {
		intentosPorCorreo.set(correo, { count: 0, lastAttempt: now });
		return true;
	}

	// Si ha pasado más de 24 horas, reiniciar contador
	if (now.getTime() - intentos.lastAttempt.getTime() > 24 * 60 * 60 * 1000) {
		intentosPorCorreo.set(correo, { count: 0, lastAttempt: now });
		return true;
	}

	// Si ya usó los 3 intentos en las últimas 24 horas, bloquear
	if (intentos.count >= 3) {
		return false;
	}

	// Si aún tiene intentos disponibles
	return intentos.count < 3;
}

// Función para incrementar intentos
function incrementarIntentos(correo: string): void {
	const now = new Date();
	const intentos = intentosPorCorreo.get(correo);
	if (intentos) {
		intentosPorCorreo.set(correo, {
			count: intentos.count + 1,
			lastAttempt: now
		});
	} else {
		intentosPorCorreo.set(correo, { count: 1, lastAttempt: now });
	}
}

export const actions: Actions = {
	default: async ({ fetch, request }) => {
		const formData = await request.formData();
		const correo = formData.get('correo')?.toString();
		const respuesta = formData.get('respuesta')?.toString();
		const newPassword = formData.get('newPassword')?.toString();
		const preguntaActual = formData.get('preguntaActual');

		// Si viene correo, obtener preguntas
		if (correo && !respuesta && !newPassword) {
			try {
				const result = await obtenerPreguntasSeguridad(fetch, correo);
				// ordenar las preguntas
				const preguntasOrdenadas = result.data.preguntas.sort((a: PreguntaSeguridad, b: PreguntaSeguridad) => a.orden - b.orden).map((pregunta: PreguntaSeguridad) => pregunta.pregunta);
				return {
						type: 'success',
						message: 'Preguntas de seguridad obtenidas',
						preguntas: preguntasOrdenadas
				};
			} catch (e) {
				return {
						type: 'failure',
						message: e instanceof Error ? e.message : 'Error al obtener las preguntas de seguridad'
				};
			}
		}

		// Si viene respuesta, verificarla
		if (correo && respuesta && !newPassword && preguntaActual) {
			// Verificar intentos primero
			if (!checkIntentos(correo)) {
				return {
						type: 'failure',
						message: 'Has excedido el número máximo de intentos. Intenta de nuevo mañana.'
				};
			}

			try {
				await verificarPreguntaSeguridad(fetch, correo, respuesta, Number(preguntaActual));
				return {
						type: 'success',
						message: 'Respuesta verificada correctamente'
				};
			} catch (e) {
				// Incrementar intentos solo si la respuesta es incorrecta (error 400)
				if (e instanceof Error && e.message === 'Respuesta incorrecta') {
					incrementarIntentos(correo);
					const intentosRestantes = Math.max(0, 3 - getIntentos(correo));
					const mensaje = intentosRestantes > 0
						? `Respuesta incorrecta. Te quedan ${intentosRestantes} ${intentosRestantes === 1 ? 'intento' : 'intentos'}.`
						: 'Has excedido el número máximo de intentos. Intenta de nuevo mañana.';
					return {
							type: 'failure',
							message: mensaje
					};
				}
				return {
						type: 'failure',
						message: e instanceof Error ? e.message : 'Error al verificar la respuesta'
				};
			}
		}

		// Si viene nueva contraseña, actualizarla
		if (correo && newPassword) {
			try {
				await recuperarContrasena(fetch, correo, newPassword);
				return {
						type: 'success',
						message: 'Contraseña actualizada correctamente',
						location: '/'
				};
			} catch (e) {
				return {
						type: 'failure',
						message: e instanceof Error ? e.message : 'Error al cambiar la contraseña'
				};
			}
		}

		// Si llegamos aquí, validar qué datos faltan
		const errors = [];
		if (!correo) errors.push('Falta el correo electrónico');
		
		// Si hay respuesta pero falta preguntaActual
		if (respuesta && !preguntaActual) errors.push('Falta el índice de la pregunta actual');
		
		// Si hay preguntaActual pero falta respuesta
		if (preguntaActual && !respuesta) errors.push('Falta la respuesta a la pregunta de seguridad');
		
		// Si hay newPassword pero faltan validaciones previas
		if (newPassword && (!preguntaActual || !respuesta)) errors.push('Debes completar la verificación de preguntas primero');
		
		return {
				type: 'failure',
				message: errors.length > 0 ? errors.join('. ') : 'Datos del formulario inválidos'
		};
	}
} satisfies Actions;
