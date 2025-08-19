<script lang="ts">
  import {enhance} from '$app/forms';
	import { addToast } from '$lib';
  import {resolver} from '$lib/utilidades/resolver';
  import type {SubmitFunction} from '@sveltejs/kit';
  import {Alert, Button, Input, Label} from 'flowbite-svelte';
  import {EyeSlashSolid, EyeSolid, ExclamationCircleOutline, LockSolid, EnvelopeSolid} from 'flowbite-svelte-icons';

  let password = $state('');
  let visible = $state(false);
  let loading = $state(false);

  const toggleVisibility = () => (visible = !visible);

  const handleSubmit: SubmitFunction = () => {
    loading = true;
    
    return async ({ result, update }) => {
        await update();
      loading = false;
      console.log(result);
      if (result.data.type === 'failure') {
        addToast({
          type: 'error',
          message: result.data.message
        });
      } else{
        addToast({
          type: 'success',
          message: 'Inicio de sesión exitoso'
        });
      }
    };
  };
</script>

<div class="min-h-screen flex flex-col items-center justify-center bg-primary-50 dark:bg-gray-900 relative">

    <img src="/vertice-logo-full.svg" alt="VÉRTICE" class="h-80 aspect-video -mb-20 -mt-20">
    <form
            method="post"
            use:enhance={handleSubmit}
            class="w-full max-w-md bg-white dark:bg-gray-800 p-6 rounded-2xl shadow-xl space-y-6"
    >
        <h1 class="text-2xl font-bold text-center text-primary-700 dark:text-white">Iniciar sesión</h1>

        <div>
            <Label for="correo">Correo</Label>
            <div class="relative">
                <Input id="correo" name="correo" type="email" placeholder="usuario@ejemplo.com" required>
                </Input>
            </div>
        </div>

        <div>
            <Label for="password">Contraseña</Label>
            <div class="relative">
                <Input
                        id="password"
                        bind:value={password}
                        type={visible ? 'text' : 'password'}
                        name="password"
                        placeholder="••••••••"
                        required
                >
                    {#snippet right()}
                        <Button
                                type="button"
                                outline
                                size="xs"
                                class="!p-2 mr-[-10px]"
                                onclick={toggleVisibility}
                        >
                            {#if visible}
                                <EyeSlashSolid/>
                            {:else}
                                <EyeSolid/>
                            {/if}
                        </Button>
                    {/snippet}
                </Input>
            </div>
        </div>

        <Button type="submit" color="blue" class="w-full" disabled={loading}>Entrar</Button>

        <div class="text-center">
            <a href="/recuperar" class="text-sm text-primary-600 hover:underline dark:text-primary-500">
                ¿Olvidaste tu contraseña?
            </a>
        </div>
    </form>
</div>