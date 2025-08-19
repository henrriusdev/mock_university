import { actualizarCarrera, crearCarrera, eliminarCarrera, obtenerCarreras } from "$lib";
import type { Carrera } from "../../../../app";
import type { Actions, PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  try {
    const res = await obtenerCarreras(fetch);
    return { carreras: res };
  } catch (error) {
    console.error('Error al obtener carreras:', error);
    return { carreras: [] };
  }
};

export const actions: Actions = {
  // Acción para crear una carrera
  create: async ({ request, fetch }) => {
    const formData = await request.formData();
    const carrera: Carrera = {
      id: 0,
      nombre: formData.get('nombre')?.toString() || '',
    };

    try {
      await crearCarrera(fetch, carrera);
      return {
        type: 'success',
        message: 'Carrera creada exitosamente',
        invalidate: true
      };
    } catch (error: any) {
      console.error('Error al crear carrera:', error);
      return {
        type: 'failure',
        message: error.message
      };
    }
  },

  edit: async ({ request, fetch }) => {
    const formData = await request.formData();
    const carrera: Carrera = {
      id: parseInt(formData.get('id')?.toString() || '0'),
      nombre: formData.get('nombre')?.toString() || '',
    };

    try {
      await actualizarCarrera(fetch, carrera.id, carrera);
      return {
        type: 'success',
        message: 'Carrera actualizada exitosamente',
        invalidate: true
      };
    } catch (error: any) {
      console.error('Error al actualizar carrera:', error);
      return {
        type: 'failure',
        message: error.message
      };
    }
  },

  delete: async ({ request, fetch }) => {
    const formData = await request.formData();
    const id = parseInt(formData.get('id')?.toString() || '0');

    try {
      await eliminarCarrera(fetch, id);
      return {
        type: 'success',
        message: 'Carrera eliminada exitosamente',
        invalidate: true
      };
    } catch (error: any) {
      console.error('Error al eliminar carrera:', error);
      return {
        type: 'failure',
        message: error.message
      };
    }
  }
};
