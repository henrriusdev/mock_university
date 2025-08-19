import type { PageServerLoad } from './$types';
import { obtenerHistoricoMaterias } from '$lib';

export const load = (async ({fetch}) => {
    const historico = await obtenerHistoricoMaterias(fetch);
    return {historico};
}) satisfies PageServerLoad;