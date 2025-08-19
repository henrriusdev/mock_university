import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
	if (!locals.usuario) {
		throw redirect(302, '/');
	}

	const photoUrl = locals.usuario.foto ? `http://127.0.0.1:8000/api/usuario/photo/${locals.usuario.foto}` : null;

	return {
		nombre: locals.usuario.nombre,
		rol: locals.usuario.rol.nombre,
		correo: locals.usuario.correo,
		cambiarClave: locals.usuario.cambiar_clave,
		photoUrl
	}
};
