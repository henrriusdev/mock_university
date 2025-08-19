import { obtenerMateriasAsignadas } from '$lib';
import type { MateriaDocente } from '../../../app';
import type { PageServerLoad } from './$types';

export const load = (async ({fetch}) => {
    try {
        const res = await obtenerMateriasAsignadas(fetch)
        return {materiasAsignadas: res}
    } catch (e) {
        console.log(e)
        return {materiasAsignadas: [] as MateriaDocente[]}
    }
}) satisfies PageServerLoad;