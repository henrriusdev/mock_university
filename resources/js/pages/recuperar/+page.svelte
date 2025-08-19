<script lang="ts">
	import { enhance } from '$app/forms';
	import { resolver } from '$lib/utilidades/resolver';
	import type { SubmitFunction } from '@sveltejs/kit';
	import { Alert, Button, Input, Label, StepIndicator } from 'flowbite-svelte';
	import { EyeSlashSolid, EyeSolid } from 'flowbite-svelte-icons';

	// Estado
	let loading = $state(false);
	let error = $state('');
	let currentStep = $state(0);
	let intentos = $state(0);
	let passwordErrors = $state<string[]>([]);
	let showPassword = $state(false);
	let showConfirmPassword = $state(false);

	function validatePassword(password: string): string[] {
		const errors: string[] = [];

		if (password.length < 8) {
			errors.push('La contraseña debe tener al menos 8 caracteres');
		}
		if (!/[A-Z]/.test(password)) {
			errors.push('Debe contener al menos una mayúscula');
		}
		if (!/[a-z]/.test(password)) {
			errors.push('Debe contener al menos una minúscula');
		}
		if (!/\d/.test(password)) {
			errors.push('Debe contener al menos un número');
		}

		return errors;
	}

	// Datos del formulario
	let correo = $state('');
	let preguntas = $state<string[]>([]);
	let respuestas = $state<string[]>(['', '', '']);
	let preguntaActual = $state(0);
	let newPassword = $state('');
	let confirmPassword = $state('');

	// Validaciones
	let isPasswordValid = $derived(
		newPassword === confirmPassword &&
			confirmPassword.length > 0 &&
			validatePassword(newPassword).length === 0
	);
	let canSubmitEmail = $derived(correo.length > 0);
	let canSubmitAnswer = $derived(respuestas[preguntaActual]?.length > 0);
	let canSubmitPassword = $derived(isPasswordValid);

	// Handlers
	const handleEmailSubmit: SubmitFunction = () => {
		if (!canSubmitEmail) return;
		loading = true;

		return async ({ result }) => {
			loading = false;
			if (result?.data?.type === 'success') {
				preguntas = result.data.preguntas || [];
				preguntaActual = 0;
				currentStep = 1;
				error = '';
			} else if (result?.data?.type === 'failure') {
				error = result.data.message;
			}
		};
	};

	const handleAnswerSubmit: SubmitFunction = () => {
		if (!canSubmitAnswer) return;
		loading = true;

		return async ({ result }) => {
			loading = false;
			
			// Si es un error, mostrar mensaje y no avanzar
			if (result?.data?.type === 'failure') {
				error = result.data.message;
				// Si el error indica que se excedieron los intentos, deshabilitar el formulario
				if (error.includes('excedido el número máximo de intentos')) {
					intentos = 3;
				}
				return; // No avanzar si hay error
			}

			// Solo avanzar si la respuesta fue exitosa
			if (result?.data?.type === 'success') {
				if (preguntaActual < 2) {
					preguntaActual++;
					error = '';
					respuestas[preguntaActual] = ''; // Limpiar la respuesta para la siguiente pregunta
				} else {
					currentStep = 2;
					error = '';
				}
			}
		};
	};

	const handlePasswordSubmit: SubmitFunction = async ({ cancel }) => {
		if (!canSubmitPassword) return;

		// Validar que las contraseñas coincidan
		if (newPassword !== confirmPassword) {
			error = 'Las contraseñas no coinciden';
			return cancel();
		}

		// Validar requisitos de la contraseña
		const errors = validatePassword(newPassword);
		if (errors.length > 0) {
			passwordErrors = errors;
			return cancel();
		}

		loading = true;
		return resolver(() => {
			loading = false;
		});
	};

	$effect(() => {
		if (newPassword.length > 0) {
			passwordErrors = validatePassword(newPassword);
		}
	});
</script>

<div class="min-h-screen flex items-center justify-center bg-primary-50 dark:bg-gray-900">
	<div class="w-full max-w-md bg-white dark:bg-gray-800 p-6 rounded-2xl shadow-xl space-y-6">
		<h1 class="text-2xl font-bold text-center text-primary-700 dark:text-white">
			Recuperar Contraseña
		</h1>

		<StepIndicator
			{currentStep}
			steps={['Correo', 'Pregunta', 'Contraseña']}
			color="blue"
			class="mb-4"
		/>

		{#if error}
			<Alert color="red" class="text-sm">{error}</Alert>
		{/if}

		{#if currentStep === 0}
			<div>
				<h3 class="mb-4">Paso 1: Ingresa tu correo</h3>
				<form method="POST" use:enhance={handleEmailSubmit} class="space-y-6">
					<div>
						<Label for="correo">Correo Electrónico</Label>
						<Input
							type="email"
							id="correo"
							name="correo"
							bind:value={correo}
							required
							placeholder="usuario@ejemplo.com"
						/>
					</div>

					<div class="flex justify-between gap-4">
						<Button href="/" color="alternative">Cancelar</Button>
						<Button type="submit" disabled={!canSubmitEmail || loading}>
							{loading ? 'Verificando...' : 'Siguiente'}
						</Button>
					</div>
				</form>
			</div>
		{:else if currentStep === 1}
			<div>
				<h3 class="mb-4">Paso 2: Responde tu pregunta de seguridad</h3>
				<form method="POST" use:enhance={handleAnswerSubmit} class="space-y-6">
					<input type="hidden" name="correo" value={correo} />
					<input type="hidden" name="preguntaActual" value={preguntaActual} />
					<div>
						<Label>Pregunta de Seguridad {preguntaActual + 1} de 3</Label>
						<p class="text-gray-700 dark:text-gray-300 mb-4">{preguntas[preguntaActual]}</p>

						<Label for="respuesta">Tu Respuesta</Label>
						<Input
							type="text"
							id="respuesta"
							name="respuesta"
							bind:value={respuestas[preguntaActual]}
							required
							placeholder="Tu respuesta"
						/>
					</div>

					<div class="flex justify-between gap-4">
						<Button color="alternative" onclick={() => (currentStep = Math.max(0, currentStep - 1))}
							>Atrás</Button
						>
						<Button type="submit" disabled={!canSubmitAnswer || loading || intentos >= 3}>
							{loading ? 'Verificando...' : 'Siguiente'}
						</Button>
					</div>
				</form>
			</div>
		{:else if currentStep === 2}
			<div>
				<h3 class="mb-4">Paso 3: Establece tu nueva contraseña</h3>
				<form method="POST" use:enhance={handlePasswordSubmit} class="space-y-6">
					<input type="hidden" name="correo" value={correo} />
					<div>
						<Label for="newPassword">Nueva Contraseña</Label>
						<Input
							type={showPassword ? 'text' : 'password'}
							id="newPassword"
							name="newPassword"
							bind:value={newPassword}
							required
							color={passwordErrors.length === 0 && newPassword.length > 0 ? 'green' : undefined}
						>
							{#snippet right()}
								<Button
									type="button"
									outline
									size="xs"
									class="!p-2"
									onclick={() => (showPassword = !showPassword)}
								>
									{#if showPassword}
										<EyeSlashSolid />
									{:else}
										<EyeSolid />
									{/if}
								</Button>
							{/snippet}
						</Input>
						{#if passwordErrors.length > 0}
							<ul class="mt-2 text-sm text-red-500">
								{#each passwordErrors as error}
									<li>• {error}</li>
								{/each}
							</ul>
						{/if}
					</div>

					<div>
						<Label for="confirmPassword">Confirmar Nueva Contraseña</Label>
						<Input
							type={showConfirmPassword ? 'text' : 'password'}
							id="confirmPassword"
							name="confirmPassword"
							bind:value={confirmPassword}
							required
							color={newPassword === confirmPassword && confirmPassword.length > 0
								? 'green'
								: 'red'}
						>
							{#snippet right()}
								<Button
									type="button"
									outline
									size="xs"
									class="!p-2"
									onclick={() => (showConfirmPassword = !showConfirmPassword)}
								>
									{#if showConfirmPassword}
										<EyeSlashSolid />
									{:else}
										<EyeSolid />
									{/if}
								</Button>
							{/snippet}
						</Input>
						{#if confirmPassword.length > 0 && newPassword !== confirmPassword}
							<p class="mt-2 text-sm text-red-500">Las contraseñas no coinciden</p>
						{/if}
					</div>

					<div class="flex justify-between gap-4">
						<Button color="alternative" onclick={() => (currentStep = Math.max(0, currentStep - 1))}
							>Atrás</Button
						>
						<Button type="submit" disabled={!canSubmitPassword || loading}>
							{loading ? 'Guardando...' : 'Cambiar Contraseña'}
						</Button>
					</div>
				</form>
			</div>
		{/if}
	</div>
</div>
