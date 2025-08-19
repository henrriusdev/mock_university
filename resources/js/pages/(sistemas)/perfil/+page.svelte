<script lang="ts">
	import { enhance } from '$app/forms';
	import { PREGUNTAS_SEGURIDAD } from '$lib/servicios/pregunta-seguridad';

	import { resolver } from '$lib/utilidades/resolver.js';
	import type { SubmitFunction } from '@sveltejs/kit';
	import { Avatar, Button, Card, Input, Label, Modal, Select } from 'flowbite-svelte';
	import { EnvelopeOutline, ShieldCheckOutline, CameraPhotoOutline, TrashBinOutline } from 'flowbite-svelte-icons';

	// Props
	let { data } = $props();

	let usuario = $derived(data.usuario);

	// Estado
	let passwordModal = $state(false);
	let securityQuestionsModal = $state(false);
	let photoModal = $state(false);
	let currentPassword = $state('');
	let newPassword = $state('');
	let confirmPassword = $state('');
	let loading = $state(false);
	let uploadingPhoto = $state(false);
	let selectedFile: File | null = $state(null);

	// Preguntas de seguridad
	let securityQuestions = $state([
		{ pregunta: '', respuesta: '' },
		{ pregunta: '', respuesta: '' },
		{ pregunta: '', respuesta: '' }
	]);

	// Estado derivado
	let isPasswordValid = $derived(newPassword === confirmPassword && newPassword.length > 0);
	let canSubmitPassword = $derived(isPasswordValid);
	let canSubmitQuestions = $derived(securityQuestions.every((q) => q.pregunta && q.respuesta));
	let avatarSrc = $derived(
		data.photoUrl || `https://unavatar.io/${usuario?.correo}`
	);

	const handleSubmit: SubmitFunction = () => {
		loading = true;
		return resolver(() => {
			loading = false;
			passwordModal = false;
			securityQuestionsModal = false;
		});
	};

	// Form action handlers
	const handlePhotoUpload: SubmitFunction = () => {
		uploadingPhoto = true;
		return async ({ result, update }) => {
			if (result.type === 'success') {
				photoModal = false;
				selectedFile = null;
				await update();
			} else if (result.type === 'failure') {
				alert(result.data?.message || 'Error al subir la foto');
			}
			uploadingPhoto = false;
		};
	};

	const handlePhotoDelete: SubmitFunction = () => {
		uploadingPhoto = true;
		return async ({ result, update }) => {
			if (result.type === 'success') {
				await update();
			} else if (result.type === 'failure') {
				alert(result.data?.message || 'Error al eliminar la foto');
			}
			uploadingPhoto = false;
		};
	};

	function handleCameraClick() {
		photoModal = true;
	}

	function handleDrop(event: DragEvent) {
		event.preventDefault();
		const target = event.currentTarget as HTMLElement;
		target.classList.remove('border-primary-400', 'bg-primary-50');
		
		const files = event.dataTransfer?.files;
		if (files && files[0]) {
			selectedFile = files[0];
		}
	}

	function onFileSelected(event: Event) {
		const input = event.target as HTMLInputElement;
		if (input.files && input.files[0]) {
			selectedFile = input.files[0];
		}
	}
</script>

<div class="h-full bg-gray-50">
	<div class="container grid place-items-center h-full">
		<div class="max-w-4xl mx-auto space-y-8">
			<h1 class="text-2xl font-bold text-center">Perfil</h1>
			<!-- Información del Usuario -->
			<Card class="p-8 shadow-lg max-w-full! bg-primary-50">
				<div class="flex flex-col items-center md:flex-row md:items-start gap-12">
					<!-- Avatar grande -->
					<div class="flex flex-col items-center gap-6">
						<div class="relative">
							<Avatar
								src={avatarSrc}
								size="xl"
								class="w-40 h-40 ring-4 ring-primary-500 ring-offset-4 ring-offset-white"
							/>
							<div class="absolute bottom-2 right-2 flex gap-2">
								<Button 
									size="xs" 
									pill 
									color="primary"
									onclick={handleCameraClick}
									disabled={uploadingPhoto}
								>
									<CameraPhotoOutline class="w-4 h-4" />
								</Button>
								{#if usuario?.foto}
							<form method="POST" action="?/eliminarFoto" use:enhance={handlePhotoDelete}>
								<Button 
									type="submit"
									size="xs" 
									pill 
									color="red"
									disabled={uploadingPhoto}
								>
									<TrashBinOutline class="w-4 h-4" />
								</Button>
							</form>
						{/if}
							</div>
							<div
								class="absolute -top-4 left-1/2 -translate-x-1/2 px-3 py-1 bg-primary-500 text-white text-sm font-medium rounded-full"
							>
								{usuario?.activo ? 'Activo' : 'Inactivo'}
							</div>
						</div>
						<div class="text-center">
							<h2 class="text-2xl font-bold text-gray-900">{usuario?.nombre}</h2>
							<p class="text-primary-600 font-medium capitalize mt-1">{usuario?.rol.nombre}</p>
						</div>
					</div>

					<!-- Detalles del usuario -->
					<div class="flex-1 space-y-6">
						<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
							<div class="p-4 bg-white rounded-lg border border-gray-200">
								<div class="flex items-center gap-3 mb-1">
									<div class="p-2 bg-primary-100 rounded-lg">
										<EnvelopeOutline class="w-5 h-5 text-primary-600" />
									</div>
									<p class="text-sm font-medium text-gray-600">Información de Contacto</p>
								</div>
								<div class="space-y-2 ml-11">
									<p class="text-gray-900 font-medium">Correo: {usuario?.correo}</p>
									<p class="text-gray-900 font-medium">Cédula: {usuario?.cedula}</p>
								</div>
							</div>
							<div class="p-4 bg-white rounded-lg border border-gray-200">
								<div class="flex items-center gap-3 mb-1">
									<div class="p-2 bg-primary-100 rounded-lg">
										<ShieldCheckOutline class="w-5 h-5 text-primary-600" />
									</div>
									<p class="text-sm font-medium text-gray-600">Estado de la Cuenta</p>
								</div>
								<div class="space-y-2 ml-11">
									<p class="text-gray-900 font-medium">Miembro desde: {usuario?.fecha_creacion}</p>
									<p class="text-gray-900 font-medium">
										Pregunta de seguridad: {usuario?.pregunta_configurada
											? 'Configurada'
											: 'Pendiente'}
									</p>
									<p class="text-gray-900 font-medium">
										Contraseña: {usuario?.cambiar_clave
											? 'Debe cambiar su contraseña'
											: 'Actualizada'}
									</p>
								</div>
							</div>
						</div>
						<!-- Botones de acción -->
						<div class="flex flex-col md:flex-row gap-4">
							<Button onclick={() => (passwordModal = true)} class="w-full">
								Cambio de contraseña
							</Button>
							<Button onclick={() => (securityQuestionsModal = true)} class="w-full">
								Preguntas de seguridad
							</Button>
						</div>
					</div>
				</div>
			</Card>

			<!-- Modal de cambio de contraseña -->
			<Modal bind:open={passwordModal} size="md" autoclose={false}>
				<form method="POST" action="?/cambiarPassword" use:enhance={handleSubmit} class="space-y-6">
					<h3 class="text-xl font-medium">Cambiar Contraseña</h3>
					<div class="space-y-4">
						<div>
							<Label for="currentPassword" class="space-y-2">Contraseña Actual</Label>
							<Input
								type="password"
								name="currentPassword"
								id="currentPassword"
								required
								bind:value={currentPassword}
							/>
						</div>
						<div>
							<Label for="newPassword" class="space-y-2">Nueva Contraseña</Label>
							<Input
								type="password"
								name="newPassword"
								id="newPassword"
								required
								bind:value={newPassword}
							/>
						</div>
						<div>
							<Label for="confirmPassword" class="space-y-2">Confirmar Contraseña</Label>
							<Input
								type="password"
								name="confirmPassword"
								id="confirmPassword"
								required
								bind:value={confirmPassword}
							/>
						</div>
						<div class="flex justify-end gap-4">
							<Button color="alternative" onclick={() => (passwordModal = false)}>Cancelar</Button>
							<Button type="submit" disabled={!canSubmitPassword || loading}>
								{loading ? 'Guardando...' : 'Cambiar Contraseña'}
							</Button>
						</div>
					</div>
				</form>
			</Modal>

			<!-- Modal de preguntas de seguridad -->
			<Modal bind:open={securityQuestionsModal} size="md" autoclose={false}>
				<form
					method="POST"
					action="?/configurarPregunta"
					use:enhance={handleSubmit}
					class="space-y-6"
				>
					<h3 class="text-xl font-medium">Configurar Preguntas de Seguridad</h3>
					<div class="space-y-4">
						{#each securityQuestions as question, i}
							<div class="space-y-4">
								<div>
									<Label for="pregunta{i}" class="space-y-2">Pregunta {i + 1}</Label>
									<Select
										name="pregunta{i}"
										id="pregunta{i}"
										bind:value={question.pregunta}
										placeholder="Seleccionar"
										required
									>
										{#each PREGUNTAS_SEGURIDAD.filter(p => 
										!securityQuestions.some((q, index) => index !== i && q.pregunta === p)
									) as pregunta}
										<option value={pregunta}>{pregunta}</option>
									{/each}
									</Select>
								</div>
								<div>
									<Label for="respuesta{i}" class="space-y-2">Respuesta {i + 1}</Label>
									<Input
										type="text"
										name="respuesta{i}"
										id="respuesta{i}"
										bind:value={question.respuesta}
										required
									/>
								</div>
							</div>
						{/each}
						<div class="flex justify-end gap-4">
							<Button color="alternative" onclick={() => (securityQuestionsModal = false)}
								>Cancelar</Button
							>
							<Button type="submit" disabled={!canSubmitQuestions || loading}>
								{loading ? 'Guardando...' : 'Guardar Preguntas'}
							</Button>
						</div>
					</div>
				</form>
			</Modal>

			<!-- Modal de foto de perfil -->
			<Modal bind:open={photoModal} size="sm" autoclose={false}>
				<form method="POST" action="?/subirFoto" enctype="multipart/form-data" use:enhance={handlePhotoUpload}>
					<div class="space-y-6">
						<h3 class="text-lg font-semibold text-gray-900">Cambiar foto de perfil</h3>
						
						<div class="space-y-4">
							<div>
								<Label class="mb-2">Seleccionar imagen</Label>
								<div 
									class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center hover:border-gray-400 transition-colors cursor-pointer"
									role="button"
									tabindex="0"
									ondragover={(e) => { e.preventDefault(); e.currentTarget.classList.add('border-primary-400', 'bg-primary-50'); }}
									ondragleave={(e) => { e.currentTarget.classList.remove('border-primary-400', 'bg-primary-50'); }}
									ondrop={handleDrop}
									onclick={() => document.getElementById('photo-upload')?.click()}
									onkeydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); document.getElementById('photo-upload')?.click(); } }}
								>
									<div class="space-y-2">
										<svg class="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48">
											<path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
										</svg>
										<p class="text-sm text-gray-600">
											<span class="font-medium text-primary-600 hover:text-primary-500">Haz click para seleccionar</span>
											 una imagen
										</p>
										<p class="text-xs text-gray-500">
											PNG, JPG, JPEG, GIF hasta 5MB
										</p>
									</div>
								</div>
								<input 
									id="photo-upload"
									name="file"
									type="file"
									accept="image/*"
									onchange={onFileSelected}
									class="hidden"
									required
								/>
							</div>
							
							{#if selectedFile}
								<div class="text-sm text-gray-700">
									Archivo seleccionado: {selectedFile.name}
								</div>
							{/if}
						</div>

						<div class="flex justify-end gap-4">
							<Button type="button" color="alternative" onclick={() => {
								photoModal = false;
								selectedFile = null;
							}}>
								Cancelar
							</Button>
							<Button 
								type="submit"
								color="primary" 
								disabled={!selectedFile || uploadingPhoto}
							>
								{uploadingPhoto ? 'Subiendo...' : 'Subir Foto'}
							</Button>
						</div>
					</div>
				</form>
			</Modal>
		</div>
	</div>
</div>
