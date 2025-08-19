import { obtenerPagosPorCedula } from '$lib';
import type { PageServerLoad } from './$types';

export const load = (async ({fetch, params}) => {
    const { estudiante } = params;
    const res = await obtenerPagosPorCedula(fetch, estudiante);
    return {pagos: res.pagos, nombre: res.nombre};
}) satisfies PageServerLoad;