import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals, cookies }) => {
  locals.usuario = null;
  cookies.delete('sesion', { path: '/' });
  redirect(302, '/');
};