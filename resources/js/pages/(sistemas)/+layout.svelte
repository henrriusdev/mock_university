<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { addToast } from '$lib';
	import {
		Avatar,
		Breadcrumb,
		BreadcrumbItem,
		Button,
		Drawer,
		Dropdown,
		DropdownDivider,
		DropdownItem,
		Input,
		Label,
		Navbar,
		Sidebar,
		SidebarGroup,
		SidebarItem,
		SidebarWrapper
	} from 'flowbite-svelte';
	import {
		BookOpenOutline,
		BookOutline,
		BuildingOutline,
		CalendarMonthOutline,
		ChalkboardUserOutline,
		ChevronLeftOutline,
		ChevronRightOutline,
		CogOutline,
		HomeOutline,
		ListOutline,
		ReceiptOutline,
		ShieldCheckOutline,
		UserCircleOutline,
		UserHeadsetOutline,
		UsersGroupOutline,
		UsersOutline
	} from 'flowbite-svelte-icons';
	import type { LayoutData } from './$types';
	import { enhance } from '$app/forms';
	import type { SubmitFunction } from '@sveltejs/kit';
	import { resolver } from '$lib/utilidades/resolver';

	// Obtener los datos del usuario desde los datos proporcionados por +layout.server.ts
	let { data, children } = $props<{ data: LayoutData }>();
	let cambiarClave = $state(!data.cambiarClave);

	// Estado para controlar si el sidebar está abierto o cerrado
	let sidebarOpen = $state(true);

	// Estado para mostrar alertas
	let userDropdownOpen = $state(false);

	// Form state
	let newPassword = $state('');
	let confirmPassword = $state('');
	let loading = $state(false);
	let error = $state('');
	let passwordErrors = $state<string[]>([]);

	function validatePassword(password: string): string[] {
		const errors: string[] = [];

		if (password.length < 6) {
			errors.push('La contraseña debe tener al menos 6 caracteres');
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

	const handleSubmit: SubmitFunction = async ({ cancel }) => {
		error = '';
		passwordErrors = [];

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
			cambiarClave = false;
			window.location.reload();
		});
	};

	// Función para alternar el estado del sidebar
	function toggleSidebar() {
		sidebarOpen = !sidebarOpen;
	}

	function toggleUserDropdown() {
		userDropdownOpen = !userDropdownOpen;
	}

	const elementosNav = [
		{
			titulo: 'Inicio',
			icono: HomeOutline,
			href: '/' + data.rol,
			roles: [] // Todos pueden ver el inicio
		},
		{
			titulo: 'Docentes',
			icono: UsersOutline,
			href: '/docentes',
			roles: ['administrador', 'coordinador']
		},
		{
			titulo: 'Estudiantes',
			icono: UsersGroupOutline,
			href: '/estudiantes',
			roles: ['administrador', 'caja', 'coordinador', 'control']
		},
		{
			titulo: 'Materias',
			icono: ListOutline,
			href: '/materias',
			roles: ['administrador', 'control']
		},
		{
			titulo: 'Gestión de Materias',
			icono: ListOutline,
			href: '/coordinador/materias',
			roles: ['coordinador']
		},
		{
			titulo: 'Usuarios',
			icono: UsersOutline,
			href: '/administrador/usuarios',
			roles: ['administrador']
		},
		{
			titulo: 'Coordinadores',
			icono: UserHeadsetOutline,
			href: '/administrador/coordinadores',
			roles: ['administrador']
		},
		{
			titulo: 'Carreras',
			icono: BuildingOutline,
			href: '/administrador/carreras',
			roles: ['administrador']
		},
		{
			titulo: 'Trazabilidad',
			icono: BookOpenOutline,
			href: '/administrador/trazabilidad',
			roles: ['administrador']
		},
		{
			titulo: 'Pagos',
			icono: ReceiptOutline,
			href: '/caja/pagos/registrar',
			roles: ['caja']
		},
		{
			titulo: 'Reportes de Pagos',
			icono: ReceiptOutline,
			href: '/caja/pagos/reportes',
			roles: ['caja']
		},
		{
			titulo: 'Configuración',
			icono: CogOutline,
			href: '/configuracion',
			roles: ['administrador', 'coordinador']
		},
		{
			titulo: 'Horario',
			icono: CalendarMonthOutline,
			href: '/estudiante/horario',
			roles: ['estudiante']
		},
		{
			titulo: 'Histórico de notas',
			icono: ShieldCheckOutline,
			href: '/estudiante/notas',
			roles: ['estudiante']
		},
		{
			titulo: 'Cursando',
			icono: BookOpenOutline,
			href: '/estudiante/materias',
			roles: ['estudiante']
		},
		{
			titulo: 'Peticiones',
			icono: ChalkboardUserOutline,
			href: '/peticiones',
			roles: ['control']
		}
	];

	// Función para generar las migas de pan basadas en la URL
	let rutaActual = $derived(page.url.pathname);
	let migasDePan = $derived(generarMigasDePan(rutaActual));

	function generarMigasDePan(ruta: string) {
		// Ignorar la primera barra y dividir la ruta
		const segmentos = ruta.split('/').filter(Boolean);

		if (segmentos.length === 0) {
			return [{ titulo: 'Inicio', href: '/', actual: true }];
		}

		// Crear las migas de pan
		const migas: { titulo: string; href: string; actual: boolean }[] = [];

		let rutaAcumulada = '';
		segmentos.forEach((segmento, index) => {
			rutaAcumulada += `/${segmento}`;

			// Convertir el segmento a un formato más legible
			const titulo = segmento.charAt(0).toUpperCase() + segmento.slice(1);

			migas.push({
				titulo,
				href: rutaAcumulada,
				actual: index === segmentos.length - 1
			});
		});

		return migas;
	}

	// Mapeo de rutas a nombres más amigables en español
	const mapeoRutas: { [key: string]: string } = {
		'/': 'Inicio',
		'/usuarios': 'Gestión de Usuarios',
		'/configuracion': 'Configuración del Sistema'
	};

	// Verificar si el usuario tiene acceso a una ruta específica
	function hasAccess(requiredRoles: string[], userRole: string) {
		if (!requiredRoles || requiredRoles.length === 0) {
			return true; // Si no hay roles requeridos, todos tienen acceso
		}

		return requiredRoles.includes(userRole);
	}

	// Verificar acceso a la ruta actual
	$effect(() => {
		if (data && data.nombre) {
			const rutaActual = page.url.pathname;
			const rutaEncontrada = elementosNav.find((item) => item.href === rutaActual);

			if (rutaEncontrada && !hasAccess(rutaEncontrada.roles, data.rol)) {
				// mostrar toast de acceso denegado
				addToast({
					type: 'error',
					message: 'Acceso denegado a esta sección.'
				});

				// Redirigir al inicio después de mostrar la alerta
				setTimeout(() => {
					goto('/');
				}, 2000);
			}
		}
	});
</script>

<!-- Main layout container with fixed height and no overflow -->
<div class="flex w-full h-screen overflow-hidden bg-gray-50 dark:bg-gray-900">
	<!-- Sidebar - fixed height with its own scrollbar -->
	<Sidebar
		activeUrl={rutaActual}
		class="relative z-40 h-screen transition-all duration-300 {sidebarOpen
			? 'w-72 min-w-72'
			: 'w-16 min-w-16'} border-r border-gray-200 dark:border-gray-700 bg-primary-50"
			activeClass="bg-primary-500"
			nonActiveClass="hover:bg-primary-200 transition-colors duration-300"
	>
		<SidebarWrapper class="h-full py-4 overflow-y-auto bg-primary-50">
			<div class="flex items-center mb-5 justify-start">
				{#if sidebarOpen}
					<img src="/vertice-logo-full.svg" alt="VÉRTICE" class="w-auto aspect-video" />
				{:else}
					<img src="/vertice-logo-v.svg" alt="V" class="w-auto" />
				{/if}
			</div>

			{#if data && data.nombre}
				<SidebarGroup>
					{#each elementosNav as item}
						{#if hasAccess(item.roles, data.rol)}
							{#if sidebarOpen}
								<SidebarItem href={item.href} label={item.titulo} >
									{#snippet icon()}
										<item.icono class="w-5 h-5" />
									{/snippet}
								</SidebarItem>
							{:else}
								<div class="mb-2">
									<SidebarItem href={item.href} title={item.titulo} active={item.href === rutaActual}>
										{#snippet icon()}
											<item.icono class="w-5 h-5" />
										{/snippet}
									</SidebarItem>
								</div>
							{/if}
						{/if}
					{/each}
				</SidebarGroup>

				<!-- Perfil de usuario en la parte inferior -->
				<!-- User profile at the bottom - now with popup -->
				<div
					class="absolute bottom-0 left-0 w-full p-4 border-t border-gray-200 dark:border-gray-700"
				>
					<!-- User profile button that triggers the dropdown -->
					<div class="relative">
						<Button
							color="light"
							class="flex items-center w-full gap-2 text-sm "
							onclick={toggleUserDropdown}
						>
							<Avatar src={data.photoUrl || `https://unavatar.io/${data.correo}`} class="mr-3" />
							{#if sidebarOpen}
								<div class="flex-1 text-left">
									<span class="font-medium truncate">{data.nombre}</span>
								</div>
							{/if}
						</Button>

						<!-- Dropdown menu that appears as a popup -->
						<Dropdown
							bind:isOpen={userDropdownOpen}
							class="z-50 {sidebarOpen ? 'left-0' : 'left-16'} bottom-14 w-56"
							onclick={() => (userDropdownOpen = false)}
						>
							<div class="px-4 py-3 text-sm text-gray-900 dark:text-white">
								<div class="font-medium">{data.nombre}</div>
								<div class="truncate text-gray-500">Rol: {data.rol}</div>
							</div>
							<DropdownDivider />
							<DropdownItem href="/perfil">
								<div class="flex items-center">
									<UserCircleOutline class="w-5 h-5 mr-2" />
									Mi Perfil
								</div>
							</DropdownItem>
							<DropdownItem href="/manual">
								<div class="flex items-center">
									<BookOutline class="w-5 h-5 mr-2" />
									Manual de Usuario
								</div>
							</DropdownItem>
							<DropdownDivider />
							<DropdownItem href="/logout">
								<div class="flex items-center">
									<ShieldCheckOutline class="w-5 h-5 mr-2" />
									Cerrar Sesión
								</div>
							</DropdownItem>
						</Dropdown>
					</div>
				</div>
			{:else}
				<!-- Si no está autenticado, no mostrar nada en el sidebar -->
				<div class="p-4 text-center text-gray-500">
					{#if sidebarOpen}
						No has iniciado sesión
					{:else}
						<span>!</span>
					{/if}
				</div>
			{/if}
		</SidebarWrapper>
	</Sidebar>

	<!-- Content area - flex column with fixed height -->
	<div class="flex flex-col flex-1 h-screen overflow-hidden">
		<!-- Navbar - fixed at top -->
		<Navbar fluid class="w-full border-b border-gray-200 dark:border-gray-700 shrink-0 bg-blue-50">
			<div class="flex items-center justify-start gap-6 w-full">
				<Button color="light" class="!p-2" size="xs" onclick={toggleSidebar}>
					{#if sidebarOpen}
						<ChevronLeftOutline class="w-5 h-5" />
					{:else}
						<ChevronRightOutline class="w-5 h-5" />
					{/if}
				</Button>
				<Breadcrumb aria-label="Migas de pan" class="hidden md:flex">
					{#each migasDePan as miga, index}
						<BreadcrumbItem href={miga.href} home={index === 0}>
							{mapeoRutas[miga.href] || miga.titulo}
						</BreadcrumbItem>
					{/each}
				</Breadcrumb>
			</div>
		</Navbar>

		<!-- Main content - only this should scroll vertically -->
		<main class="flex-1 overflow-y-auto p-6 w-full">
			<Drawer placement="right" id="drawer-cambiar-clave" bind:hidden={cambiarClave} class="w-1/4">
				<div class="flex h-full flex-col justify-between p-4">
					<div>
						<h5 class="mb-4 text-xl font-medium text-gray-500 dark:text-gray-400">
							Cambiar Contraseña
						</h5>
						<form
							class="flex flex-col gap-4"
							use:enhance={handleSubmit}
							action="/password"
							method="POST"
						>
							<div>
								<Label for="new-password" class="mb-2">Nueva Contraseña</Label>
								<Input
									id="new-password"
									name="new_password"
									type="password"
									required
									bind:value={newPassword}
									placeholder="••••••••"
									class={passwordErrors.length > 0 ? 'border-red-500' : ''}
								/>
								{#if passwordErrors.length > 0}
									<ul class="mt-2 text-sm text-red-500">
										{#each passwordErrors as error}
											<li>• {error}</li>
										{/each}
									</ul>
								{/if}
							</div>
							<div>
								<Label for="confirm-password" class="mb-2">Confirmar Nueva Contraseña</Label>
								<Input
									id="confirm-password"
									type="password"
									required
									bind:value={confirmPassword}
									placeholder="••••••••"
									class={error ? 'border-red-500' : ''}
								/>
								{#if error}
									<p class="mt-2 text-sm text-red-500">{error}</p>
								{/if}
							</div>
							<div class="flex justify-end gap-4">
								<Button type="submit" disabled={loading}>
									{loading ? 'Actualizando...' : 'Actualizar Contraseña'}
								</Button>
							</div>
						</form>
					</div>
				</div>
			</Drawer>
			{@render children()}
		</main>
	</div>
</div>

<style>
	:global(.absolute.bottom-0 li::marker) {
		content: none;
	}

	:global(input + div.flex.absolute) {
		padding: 0 !important;
	}
	
	:global(button[type='submit']) {
		margin-bottom: 0;
	}
	:global(aside > div){
		background-color: #eff6ff;
	}
</style>
