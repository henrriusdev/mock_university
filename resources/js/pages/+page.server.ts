import { login } from '$lib';
import { redirect } from '@sveltejs/kit';

// Define enhanced error type
interface EnhancedError {
	message: string;
	code: string;
	status: number;
}

export const load = async ({ locals }) => {
	if (locals.usuario) {
		throw redirect(302, `/${locals.usuario.rol.nombre.toLowerCase()}`);
	}
};

export const actions = {
	default: async ({ request, fetch, cookies }) => {
		const data = await request.formData();
		const correo = data.get('correo') as string;
		const password = data.get('password') as string;

		let destino = '/';
		try {
			const { usuario, token } = await login(fetch, correo, password);

                        cookies.set('sesion', token, {
                                path: '/',
                                httpOnly: true,
                                sameSite: 'lax',
                                secure: false, // true en producción
                                maxAge: 60 * 60 * 24 * 7 // 7 days to match JWT expiration
                        });

			destino = `/${usuario.rol.nombre.toLowerCase()}`;
			return {
				type: 'success',
				message: 'Inicio de sesión exitoso',
				invalidate: destino
			};
		} catch (e: unknown) {
			console.error('Error en inicio de sesión:', e);
			
			// Handle enhanced error format
			if (e && typeof e === 'object' && 'code' in e && 'message' in e) {
				const enhancedError = e as EnhancedError;
				return {
					type: 'failure',
					message: enhancedError.message,
					error_code: enhancedError.code
				};
			}
			
			// Handle standard Error objects
			if (e instanceof Error) {
				return {
					type: 'failure',
					message: e.message
				};
			}
			
			// Fallback for unknown error types
			return {
				type: 'failure',
				message: 'Error desconocido al iniciar sesión'
			};
		}
	}
};
