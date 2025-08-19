import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';
import { forzarCambioPassword } from '$lib/servicios/autenticacion';

export const actions: Actions = {
	default: async ({ request, locals, fetch }) => {
		if (!locals.usuario) {
			throw redirect(302, '/');
		}

		try {
			const data = await request.formData();
			const newPassword = data.get('new_password') as string;

			await forzarCambioPassword(fetch, newPassword);

			const location = '/' + locals.usuario.rol.nombre.toLowerCase();
			return {
				invalidate: true,
				type: 'success',
				message: 'Contraseña actualizada correctamente',
				location
			};
		} catch (e) {
			console.error(e);
			return {
				type: 'failure',
				message: e instanceof Error ? e.message : 'Error al actualizar la contraseña'
			};
		}
	}
};
