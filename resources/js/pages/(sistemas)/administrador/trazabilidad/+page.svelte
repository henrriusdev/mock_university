<script lang="ts">
	import { enhance } from '$app/forms';
	import { resolver } from '$lib/utilidades/resolver';
	import { format } from 'date-fns';
	import { es } from 'date-fns/locale';
	import {
		Button,
		Datepicker,
		Input,
		Label,
		Select,
		Spinner,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Tooltip
	} from 'flowbite-svelte';
	import {
		CalendarEditOutline,
		FileExportOutline,
		FileLinesOutline,
		FilePdfOutline,
		InfoCircleOutline,
		SearchOutline
	} from 'flowbite-svelte-icons';
	import { tick } from 'svelte';

	// Definir props
	let { data } = $props();

	// Estado local con runas
	let busqueda = $state('');
	let fechaDesde = $state(new Date());
	let fechaHasta = $state(new Date());
	let rolSeleccionado = $state('');
	let paginaActual = $state(1);
	let elementosPorPagina = $state(10);
	let cargando = $state(false);
	let formatoExportacion = $state('');
	let exportarForm: HTMLFormElement | null = $state(null);
	let filtroForm: HTMLFormElement | null = $state(null);

	// Roles disponibles
	const roles = [
		{ value: '', name: 'Todos los roles' },
		{ value: 'administrador', name: 'Administrador' },
		{ value: 'caja', name: 'Personal de caja' },
		{ value: 'coordinador', name: 'Coordinador' },
		{ value: 'docente', name: 'Docente' },
		{ value: 'control', name: 'Control de estudios' },
		{ value: 'estudiante', name: 'Estudiante' }
	];

	// Calcular datos paginados con runas
	const registrosPaginados = $derived.by(() => {
		const inicio = (paginaActual - 1) * elementosPorPagina;
		const fin = inicio + elementosPorPagina;
		return data.registros?.slice(inicio, fin);
	});

	const totalPaginas = $derived.by(() => Math.ceil(data.registros.length / elementosPorPagina));

	// Formatear fecha para mostrar
	function formatearFecha(fecha: string) {
		return format(new Date(fecha), 'dd/MM/yyyy HH:mm:ss', { locale: es });
	}

	// Manejar cambio de página
	function cambiarPagina(pagina: number) {
		paginaActual = pagina;
	}

	// Función para exportar usando el formulario dedicado
	async function exportar(formato: string) {
		formatoExportacion = formato;
		await tick();
		exportarForm!.requestSubmit();
	}

	// Limpiar filtros
	function limpiarFiltros() {
		busqueda = '';
		fechaDesde = new Date();
		fechaHasta = new Date();
		rolSeleccionado = '';
	}
</script>

<div class="container mx-auto px-4 py-8">
	<h1 class="text-3xl font-bold mb-6">Trazabilidad del Sistema</h1>

	<!-- Filtros -->
	<div class="bg-blue-50 rounded-lg shadow p-6 mb-6">
		<h2 class="text-xl font-semibold mb-4">Filtros</h2>

		<form
			bind:this={filtroForm}
			method="POST"
			action="?/filtrar"
			use:enhance={() => {
				cargando = true;
				return resolver(() => {
					cargando = false;
				});
			}}
		>
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
				<!-- Búsqueda por texto -->
				<div>
					<Label for="busqueda" class="mb-2">Búsqueda</Label>
					<div class="relative">
						<div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
							<SearchOutline class="w-4 h-4 text-gray-500" />
						</div>
						<Input
							id="busqueda"
							name="busqueda"
							bind:value={busqueda}
							placeholder="Usuario, módulo, acción..."
							class="pl-10"
						/>
					</div>
				</div>

				<!-- Fecha desde -->
				<div>
					<Label for="fechaDesde" class="mb-2">Fecha desde</Label>
					<div class="relative">
						<div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
							<CalendarEditOutline class="w-4 h-4 text-gray-500" />
						</div>
						<Datepicker
							bind:value={fechaDesde}
							availableTo={new Date()}
							placeholder="Seleccione una fecha"
							translationLocale="es-VE"
							locale="fr-FR"
							dateFormat={{ year: 'numeric', month: '2-digit', day: '2-digit' }}
						/>
						<input type="hidden" name="fechaDesde" bind:value={fechaDesde} class="hidden" />
					</div>
				</div>

				<!-- Fecha hasta -->
				<div>
					<Label for="fechaHasta" class="mb-2">Fecha hasta</Label>
					<div class="relative">
						<Datepicker
							bind:value={fechaHasta}
							availableFrom={fechaDesde}
							availableTo={new Date()}
							placeholder="Seleccione una fecha"
							translationLocale="es-VE"
							locale="fr-FR"
							dateFormat={{ year: 'numeric', month: '2-digit', day: '2-digit' }}
						/>
						<input type="hidden" name="fechaHasta" bind:value={fechaHasta} class="hidden" />
					</div>
				</div>

				<!-- Rol -->
				<div>
					<Label for="rol" class="mb-2">Rol</Label>
					<Select
						id="rol"
						name="rol"
						bind:value={rolSeleccionado}
						items={roles}
						placeholder="Seleccionar"
					/>
				</div>
			</div>

			<div class="flex flex-wrap gap-2 justify-between">
				<div>
					<Button type="submit" color="blue" disabled={cargando}>
						<SearchOutline class="w-4 h-4 mr-2" />
						{cargando ? 'Filtrando...' : 'Filtrar'}
					</Button>
					<Button type="button" color="light" onclick={() => limpiarFiltros()} disabled={cargando}>
						Limpiar filtros
					</Button>
				</div>

				<!-- Botones de exportación -->
				<div class="flex gap-2">
					<div class="relative">
						<Button color="green" onclick={() => exportar('csv')} disabled={cargando}>
							<FileLinesOutline class="w-4 h-4 mr-2" />
							CSV
						</Button>
						<Tooltip placement="top">Exportar como CSV</Tooltip>
					</div>
					<div class="relative">
						<Button color="green" onclick={() => exportar('xlsx')} disabled={cargando}>
							<FileExportOutline class="w-4 h-4 mr-2" />
							XLSX
						</Button>
						<Tooltip placement="top">Exportar como Excel</Tooltip>
					</div>
					<div class="relative">
						<Button color="green" onclick={() => exportar('pdf')} disabled={cargando}>
							<FilePdfOutline class="w-4 h-4 mr-2" />
							PDF
						</Button>
						<Tooltip placement="top">Exportar como PDF</Tooltip>
					</div>
				</div>
			</div>
		</form>
	</div>

	<!-- Formulario de exportación -->
	<form
		bind:this={exportarForm}
		method="POST"
		action="?/exportar"
		use:enhance={() => {
			return resolver(() => {});
		}}
		style="display: none;"
	>
		<input type="hidden" name="formato" bind:value={formatoExportacion} />
		<input type="hidden" name="busqueda" bind:value={busqueda} />
		<input type="hidden" name="fechaDesde" bind:value={fechaDesde} />
		<input type="hidden" name="fechaHasta" bind:value={fechaHasta} />
		<input type="hidden" name="rol" bind:value={rolSeleccionado} />
	</form>

	<!-- Tabla de resultados -->
	<div class="bg-white rounded-lg shadow overflow-hidden">
		{#if cargando}
			<div class="flex justify-center items-center p-12">
				<Spinner size="12" />
				<p class="ml-4 text-lg">Cargando resultados...</p>
			</div>
		{:else if data.registros?.length === 0}
			<div class="flex flex-col items-center justify-center p-12 text-gray-500">
				<InfoCircleOutline class="w-12 h-12 mb-4" />
				<p class="text-lg font-medium">Sin resultados</p>
				<p>No se encontraron registros que coincidan con los filtros aplicados.</p>
			</div>
		{:else}
			<div class="overflow-x-auto">
				<Table striped={true}>
					<TableHead>
						<TableHeadCell>Usuario</TableHeadCell>
						<TableHeadCell>Rol</TableHeadCell>
						<TableHeadCell>Módulo</TableHeadCell>
						<TableHeadCell>Acción</TableHeadCell>
						<TableHeadCell>Fecha y hora</TableHeadCell>
					</TableHead>
					<TableBody>
						{#each registrosPaginados as registro, index (index)}
							<TableBodyRow>
								<TableBodyCell>{registro.usuario}</TableBodyCell>
								<TableBodyCell>{registro.rol}</TableBodyCell>
								<TableBodyCell>{registro.modulo}</TableBodyCell>
								<TableBodyCell>{registro.accion}</TableBodyCell>
								<TableBodyCell>{formatearFecha(registro.fecha)}</TableBodyCell>
							</TableBodyRow>
						{/each}
					</TableBody>
				</Table>
			</div>

			<!-- Paginación -->
			<div class="p-4 flex justify-between items-center flex-wrap gap-4">
				<div>
					<p class="text-sm text-gray-700">
						Mostrando <span class="font-medium">{(paginaActual - 1) * elementosPorPagina + 1}</span>
						a
						<span class="font-medium"
							>{Math.min(paginaActual * elementosPorPagina, data.registros.length)}</span
						>
						de <span class="font-medium">{data.registros.length}</span> registros
					</p>
				</div>

				<div class="flex items-center gap-2">
					<Button
						size="sm"
						color="alternative"
						disabled={paginaActual === 1}
						onclick={() => cambiarPagina(paginaActual - 1)}
					>
						Anterior
					</Button>

					{#each Array.from({ length: Math.min(10, totalPaginas) }, (_, i) => {
						// Mostrar 5 páginas centradas alrededor de la página actual
						let start = Math.max(1, paginaActual - 2);
						let end = Math.min(totalPaginas, start + 4);
						start = Math.max(1, end - 4);
						return start + i;
					}).filter((p) => p <= totalPaginas) as pagina (pagina)}
						<Button
							size="sm"
							color={pagina === paginaActual ? 'blue' : 'alternative'}
							onclick={() => cambiarPagina(pagina)}
						>
							{pagina}
						</Button>
					{/each}

					<Button
						size="sm"
						color="alternative"
						disabled={paginaActual === totalPaginas}
						onclick={() => cambiarPagina(paginaActual + 1)}
					>
						Siguiente
					</Button>
				</div>
			</div>
		{/if}
	</div>
</div>
