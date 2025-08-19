<script lang="ts">
  import {enhance} from '$app/forms';
  import {goto} from '$app/navigation';
  import {
    Button,
    Card,
    Datepicker,
    Heading,
    Input,
    Label,
    Modal,
    Select,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
    Toggle
  } from 'flowbite-svelte';
  import {addToast, resolver} from '$lib/utilidades';
  import {
    CreditCardOutline,
    DownloadOutline,
    PlusOutline,
    SearchOutline,
    TrashBinOutline
  } from 'flowbite-svelte-icons';
  import {cedulaMask} from '$lib';
  import {imask} from '@imask/svelte';
  import ToastContainer from '$lib/componentes/ToastContainer.svelte';

  import type {PageData} from './$types';

  let {data}: { data: PageData } = $props();

  // Main view state
  let searchQuery = $state('');
  let currentCycleOnly = $state(true);

  // Payment concepts
  const paymentConcepts = [
    {value: 'inscripcion', name: 'Inscripción'},
    {value: 'servicios', name: 'Servicios Estudiantiles'},
    {value: 'cuota1', name: 'Cuota 1'},
    {value: 'cuota2', name: 'Cuota 2'},
    {value: 'cuota3', name: 'Cuota 3'},
    {value: 'cuota4', name: 'Cuota 4'},
    {value: 'cuota5', name: 'Cuota 5'}
  ];

  // Payment methods
  const paymentMethods = [
    {value: 'transfer', name: 'Transferencia'},
    {value: 'cash', name: 'Efectivo'},
    {value: 'point', name: 'Punto de Venta'}
  ];

  let serial: string = $state('');
  let monto: number = $state(1);
  let billetes: { serial: string; monto: number }[] = $state([]);

  // Form state variables
  let student: string = $state('');
  let paymentConcept: string = $state('inscripcion');
  let paymentDate = $state<Date | undefined>(undefined);
  let paymentMethod: string = $state('transfer');
  let amount: string = $state('');
  let exchangeRate: string = $state('');
  let paymentReference: string = $state('');
  let showPaymentModal: boolean = $state(false);
  let form: HTMLFormElement;

  // PDF data for modal
  let pdfData = $state<{ base64: string; filename: string } | null>(null);
  let pdfModalOpen = $state(false);

  // Download PDF function
  function downloadPdf() {
    if (!pdfData) return;

    const linkSource = `data:application/pdf;base64,${pdfData.base64}`;
    const downloadLink = document.createElement('a');
    const fileName = pdfData.filename;

    downloadLink.href = linkSource;
    downloadLink.download = fileName;
    downloadLink.click();
  }

  // Derived state for form validation
  const isFormValid = $derived(
    student.trim() !== '' &&
    paymentConcept !== '' &&
    paymentDate !== null &&
    paymentMethod !== '' &&
    amount !== ''
  );

  $effect(() => {
    if (!showPaymentModal) {
      // Reset form fields when modal is closed
      student = '';
      paymentConcept = 'inscripcion';
      paymentDate = undefined;
      paymentMethod = 'transfer';
      amount = '';
      exchangeRate = '';
      billetes = [];
      serial = '';
    }
  });

  // Functions for payment form
  function openPaymentModal() {
    // Reset form fields when opening modal
    student = '';
    paymentConcept = 'inscripcion';
    paymentDate = undefined;
    paymentMethod = 'transfer';
    amount = '';
    exchangeRate = '';
    showPaymentModal = true;
  }

  function closePaymentModal() {
    showPaymentModal = false;
  }

  // Close PDF modal
  function closePdfModal() {
    pdfModalOpen = false;
  }

  const handleSubmit = ({formData, cancel}: { formData: FormData; cancel: () => void }) => {
    // Validación de campos
    if (!student) {
      addToast({type: 'error', message: 'Por favor ingresa la cédula del estudiante'});
      return cancel();
    }
    if (!data.estudiantes.includes(student)) {
      alert('No hay algún estudiante con esa cédula en nuestro sistema');
      return cancel();
    }

    if (paymentMethod === 'cash' && billetes.length > 0) {
      billetes.forEach((billete) => {
        formData.append('billetes', JSON.stringify(billete));
      });
    }

    if (paymentMethod === 'transfer' && paymentReference.trim()) {
      formData.append('referencia', paymentReference.trim());
    }

    // Append payment method
    formData.append('payment-method', paymentMethod);
    if (paymentMethod === 'cash') {
      formData.append('tasa_divisa', exchangeRate);
    } else {
      formData.append('tasa_divisa', '');
    }

    // Format payment date
    if (paymentDate instanceof Date) {
      formData.append('fecha_pago', paymentDate.toISOString().split('T')[0]);
    } else {
      // If no date is selected, use today's date as fallback
      formData.append('fecha_pago', new Date().toISOString().split('T')[0]);
    }

    return resolver(()=>{
    })
  };

  function agregarBillete() {
    if (!serial.trim()) {
      alert('Por favor, introduzca el serial del billete');
      return;
    }
    billetes.push({serial: serial, monto: monto});
    serial = '';
  }

  function borrarBillete(i: number) {
    billetes = billetes.filter((bil, index) => i !== index);
  }

  function viewPayment() {
    if (!searchQuery) return;
    const params = new URLSearchParams();
    if (currentCycleOnly) params.set('ciclo_actual', 'true');
    searchQuery = searchQuery.split('_')[0];
    goto(`/caja/pagos/${searchQuery}?${params.toString()}`);
  }
</script>

<div class="w-full h-full p-6">
    <!-- ... (rest of the code remains the same) -->
    <div class="flex justify-between items-center mb-8">
        <Heading tag="h1" class="text-2xl font-bold text-blue-700">Gestión de Pagos</Heading>
        <Button color="blue" size="lg" onclick={openPaymentModal}>
            <PlusOutline class="mr-2 h-5 w-5"/>
            Registrar Pago
        </Button>
    </div>

    <div class="flex justify-center items-center h-11/12 gap-8">
        <!-- Search Payments Section -->
        <Card class="p-6 shadow-lg">
            <Heading tag="h2" class="text-xl font-bold text-blue-600 mb-6">Buscar Pagos</Heading>

            <div class="space-y-6">
                <div>
                    <Label for="search-payments" class="mb-2 font-medium">Seleccione el estudiante</Label>
                    <div class="relative">
						<span class="absolute inset-y-0 left-0 flex items-center pl-2">
							<SearchOutline class="h-5 w-5 text-gray-400"/>
						</span>
                        <input
                                id="search-payments"
                                placeholder="Ingrese cédula del estudiante..."
                                bind:value={searchQuery}
                                use:imask={cedulaMask as any}
                                oninput={(e) => {
								searchQuery = e.currentTarget.value;
							}}
                                class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 pl-10"
                        />
                    </div>
                </div>

                <div class="flex items-center">
                    <Toggle bind:checked={currentCycleOnly} class="mr-3"/>
                    <Label>Mostrar solo pagos del ciclo actual</Label>
                </div>

                <Button
                        color="primary"
                        size="lg"
                        class="w-full"
                        onclick={viewPayment}
                        disabled={!searchQuery}
                >
                    Ver pago
                </Button>
            </div>
        </Card>
    </div>

    <!-- Payment Form Modal -->
    <Modal bind:open={showPaymentModal} size="lg" autoclose={false}>
        {#snippet header()}
            <div class="flex items-center">
                <CreditCardOutline class="w-6 h-6 mr-2 text-blue-600"/>
                <h3 class="text-xl font-medium text-blue-700">Registrar Nuevo Pago</h3>
            </div>
        {/snippet}
        <form
                bind:this={form}
                method="post"
                use:enhance={handleSubmit}
                action="?/crear"
                class="space-y-6"
        >
            <!-- Form sections with visual separation -->
            <div class="bg-gray-50 p-4 rounded-lg border border-gray-200 mb-4">
                <h4 class="text-gray-700 font-medium mb-3">Información del Estudiante</h4>
                <div class="grid md:grid-cols-2 gap-4">
                    <!-- Student Field -->
                    <div>
                        <Label for="student" class="mb-2 font-medium">Cédula del Estudiante</Label>
                        <div class="relative">
							<span class="absolute inset-y-0 left-0 flex items-center pl-2">
								<SearchOutline class="h-5 w-5 text-gray-400"/>
							</span>
                            <input
                                    id="student"
                                    name="student"
                                    placeholder="Ej: V-12345678"
                                    bind:value={student}
                                    oninput={(e) => {
									// Ensure the value is properly updated
									student = e.currentTarget.value;
								}}
                                    use:imask={cedulaMask}
                                    class="block w-full rounded-lg border border-gray-300 bg-white p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 pl-10"
                            />
                        </div>
                    </div>

                    <!-- Payment Concept -->
                    <div>
                        <Label for="payment-concept" class="mb-2 font-medium">Concepto del pago</Label>
                        <Select
                                id="payment-concept"
                                name="payment-concept"
                                placeholder="Seleccionar"
                                items={paymentConcepts}
                                bind:value={paymentConcept}
                                class="bg-white border-gray-300"
                        />
                    </div>
                </div>
            </div>

            <!-- Payment Details Section -->
            <div class="bg-white p-4 rounded-lg border border-gray-200 mb-4">
                <h4 class="text-gray-700 font-medium mb-3">Detalles del Pago</h4>
                <div class="grid md:grid-cols-3 gap-4">
                    <!-- Payment Date -->
                    <div>
                        <Label for="payment-date" class="mb-2 font-medium">Fecha del Pago</Label>
                        <Datepicker
                                bind:value={paymentDate}
                                availableTo={new Date()}
                                placeholder="Seleccione una fecha"
                                translationLocale="es-VE" locale="fr-FR"
                                dateFormat={{year: 'numeric', month: '2-digit', day: '2-digit'}}
                        />
                        <input
                                type="hidden"
                                name="payment-date"
                                value={paymentDate?.toISOString().split('T')[0] ?? ''}
                        />
                    </div>

                    <!-- Amount -->
                    <div>
                        <Label for="amount" class="mb-2 font-medium">Monto (Bs.)</Label>
                        <Input
                                id="amount"
                                name="amount"
                                type="number"
                                placeholder="Monto en bolívares"
                                bind:value={amount}
                                min="0"
                                class="bg-white border-gray-300"
                        />
                    </div>

                    <!-- Exchange Rate - only shown for cash payment -->
                    {#if paymentMethod === 'cash'}
                        <div>
                            <Label for="exchange" class="mb-2 font-medium">Tasa de cambio ($)</Label>
                            <Input
                                    id="exchange"
                                    name="tasa_divisa"
                                    type="number"
                                    step="0.01"
                                    min="0"
                                    placeholder="Tasa actual"
                                    bind:value={exchangeRate}
                                    class="bg-white border-gray-300"
                            />
                        </div>
                    {/if}
                </div>
            </div>

            <!-- Payment Method Section -->
            <div class="bg-gray-50 p-4 rounded-lg border border-gray-200 mb-4">
                <h4 class="text-gray-700 font-medium mb-3">Método de Pago</h4>

                <!-- Payment Method Selector -->
                <div class="mb-4">
                    <Label for="payment-method" class="mb-2 font-medium">Seleccione el método de pago</Label>
                    <Select
                            id="payment-method"
                            name="payment-method"
                            items={paymentMethods}
                            placeholder="Seleccionar"
                            bind:value={paymentMethod}
                            class="bg-white border-gray-300"
                    />
                </div>

                <!-- Billetes (Cash) -->
                {#if paymentMethod === 'cash'}
                    <div class="bg-white p-3 rounded-lg border border-gray-200 mb-4">
                        <div class="flex gap-4 mb-4">
                            <div class="flex-1">
                                <Label for="serial" class="mb-2 font-medium">Serial del billete</Label>
                                <Input
                                        id="serial"
                                        type="text"
                                        bind:value={serial}
                                        placeholder="Ingrese el serial"
                                />
                            </div>
                            <div class="flex-1">
                                <Label for="monto" class="mb-2 font-medium">Denominación</Label>
                                <Select
                                placeholder="Seleccionar"
                                        id="monto"
                                        bind:value={monto}
                                        items={[
										{ value: 1, name: '1$' },
										{ value: 2, name: '2$' },
										{ value: 5, name: '5$' },
										{ value: 10, name: '10$' },
										{ value: 20, name: '20$' },
										{ value: 50, name: '50$' },
										{ value: 100, name: '100$' }
									]}
                                />
                            </div>
                        </div>
                        <Button color="purple" class="w-full" outline onclick={agregarBillete}>
                            Añadir billete
                        </Button>
                    </div>

                    {#if billetes.length > 0}
                        <Table>
                            <TableHead>
                                <TableHeadCell>Serial</TableHeadCell>
                                <TableHeadCell>Monto</TableHeadCell>
                                <TableHeadCell>Acciones</TableHeadCell>
                            </TableHead>
                            <TableBody class="divide-y">
                                {#each billetes as billete, i}
                                    <TableBodyRow>
                                        <TableBodyCell>{billete.serial}</TableBodyCell>
                                        <TableBodyCell>{billete.monto}$</TableBodyCell>
                                        <TableBodyCell>
                                            <Button color="red" onclick={() => borrarBillete(i)} size="xs" outline>
                                                <TrashBinOutline class="w-5 h-5"/>
                                            </Button>
                                        </TableBodyCell>
                                    </TableBodyRow>
                                {/each}
                            </TableBody>
                        </Table>
                    {/if}
                {:else if paymentMethod === 'transfer'}
                    <div class="bg-white p-3 rounded-lg border border-gray-200">
                        <Label for="payment-reference" class="mb-2 font-medium"
                        >Referencia de transferencia</Label
                        >
                        <Input
                                id="payment-reference"
                                type="text"
                                placeholder="Ingrese el número de referencia"
                                bind:value={paymentReference}
                        />
                    </div>
                {/if}
            </div>
            <!-- Form Actions -->
            <div class="flex justify-start space-x-3 mt-6">
                <Button color="alternative" onclick={closePaymentModal}>Cancelar</Button>
                <Button type="submit" color="blue" disabled={!isFormValid}>
                    <CreditCardOutline class="w-5 h-5 mr-2"/>
                    Registrar Pago
                </Button>
            </div>
        </form>

        <ToastContainer/>
    </Modal>

    <!-- PDF Modal -->
    <Modal bind:open={pdfModalOpen} size="lg" autoclose={false} class="w-full">
        <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t">
            <Heading tag="h3" class="text-xl font-semibold text-gray-900">Constancia de Pago</Heading>
            <Button
                    class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center"
                    onclick={closePdfModal}
            >
                <svg
                        class="w-3 h-3"
                        aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 14 14"
                >
                    <path
                            stroke="currentColor"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
                    />
                </svg>
            </Button>
        </div>

        <div class="p-4 md:p-5 space-y-4">
            {#if pdfData}
                <div class="flex flex-col items-center">
                    <p class="mb-4 text-center text-gray-700">
                        La constancia de pago ha sido generada exitosamente.
                    </p>

                    <div class="w-full h-[500px] mb-4 border border-gray-300 rounded overflow-hidden">
                        <iframe
                                title="Constancia de Pago"
                                class="w-full h-full"
                                src={`data:application/pdf;base64,${pdfData?.base64 || ''}`}
                        >
                        </iframe>
                    </div>

                    <Button color="blue" onclick={downloadPdf}>
                        <DownloadOutline class="mr-2 h-5 w-5"/>
                        Descargar Constancia
                    </Button>
                </div>
            {:else}
                <p class="text-center text-gray-700">No se ha generado ninguna constancia de pago.</p>
            {/if}
        </div>
    </Modal>
</div>
