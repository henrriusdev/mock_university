<script>
    import StudentLayout from "$lib/layouts/StudentLayout.svelte";
    import * as Table from "$lib/components/ui/table/index.js";

    const days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
    const hours = ["7:00", "7:30", "8:00", "8:30", "9:00", "9:30", "10:00", "10:30", "11:00", "11:30", "12:00", "12:30", "13:00", "13:30", "14:00", "14:30", "15:00", "15:30", "16:00", "16:30", "17:00", "17:30", "18:00", "18:30", "19:00"];

    /** @typedef {Object} ScheduleSubjectDto
     * @property {number} id
     * @property {string} name
     * @property {string} description
     * @property {string} code
     * @property {number} semester
     * @property {number} credits
     * @property {number} pHours
     * @property {number} tHours
     * @property {number} lHours
     * @property {Object<string, string[]>} classSchedule
     */

    /** @type {ScheduleSubjectDto[]} */
    export let subjects = [];

    // Convertir hora (HH:MM) a minutos desde la medianoche
    /**
     * @param {string} time
     * @returns {number}
     */
    function timeToMinutes(time) {
        const [hours, minutes] = time.split(":").map(Number);
        return hours * 60 + minutes;
    }

    // Calcular cuántas filas debe ocupar una materia
    /**
     * @param {string} startTime
     * @param {string} endTime
     * @returns {number}
     */
    function calculateRowSpan(startTime, endTime) {
        const start = timeToMinutes(startTime);
        const end = timeToMinutes(endTime);
        return Math.ceil((end - start) / 30); // Cada fila representa 30 minutos
    }

    // Generar clases de color usando Tailwind CSS para modo claro y oscuro
    /**
     * @param {number} id
     * @returns {string}
     */
    function generateTailwindColorClass(id) {
        const colors = [
            'bg-blue-500 text-white dark:bg-blue-700 dark:text-white',
            'bg-green-500 text-white dark:bg-green-700 dark:text-white',
            'bg-pink-500 text-white dark:bg-pink-700 dark:text-white',
            'bg-yellow-500 text-black dark:bg-yellow-700 dark:text-black',
            'bg-purple-500 text-white dark:bg-purple-700 dark:text-white',
            'bg-red-500 text-white dark:bg-red-700 dark:text-white',
            'bg-indigo-500 text-white dark:bg-indigo-700 dark:text-white'
        ];
        return colors[id % colors.length]; // Usa el ID para generar una clase de color de Tailwind
    }

    // Preprocesar los datos para crear una lista plana de celdas (días y horas)
    /**
     * @param {ScheduleSubjectDto[]} subjects
     * @returns {Array<{day: string, hour: string, subject: ScheduleSubjectDto | null}>}
     */
    function processSchedule(subjects) {
        /** @type {Array<{day: string, hour: string, subject: ScheduleSubjectDto | null}>} */
        let scheduleCells = [];

        // Recorremos horas y días, creando celdas planas con su contenido
        hours.forEach(hour => {
            days.forEach(day => {
                /** @type {{day: string, hour: string, subject: ScheduleSubjectDto | null}} */
                const cell = { day, hour, subject: null }; // Inicializamos cada celda con nulo

                subjects.forEach(subject => {
                    if (subject.classSchedule[day.toLowerCase()]) {
                        const [startTime, endTime] = subject.classSchedule[day.toLowerCase()];
                        const start = timeToMinutes(startTime);
                        const end = timeToMinutes(endTime);
                        const hourInMinutes = timeToMinutes(hour);

                        if (hourInMinutes >= start && hourInMinutes < end) {
                            cell.subject = subject; // Asignamos la materia a la celda correcta
                        }
                    }
                });

                scheduleCells.push(cell); // Agregamos la celda a la lista plana
            });
        });

        return scheduleCells;
    }

    // Procesar los datos antes de renderizar
    let scheduleCells = processSchedule(subjects);

    // Mantenemos un registro de las horas ocupadas
    let occupiedCells = new Set();

    /**
     * @param {string} day
     * @param {string} hour
     * @returns {boolean}
     */
    function isCellOccupied(day, hour) {
        return occupiedCells.has(`${day}-${hour}`);
    }

    /**
     * @param {string} day
     * @param {string} startTime
     * @param {string} endTime
     * @returns {string}
     */
    function markCellOccupied(day, startTime, endTime) {
        const startMinutes = timeToMinutes(startTime);
        const endMinutes = timeToMinutes(endTime);

        hours.forEach(hour => {
            const hourInMinutes = timeToMinutes(hour);
            if (hourInMinutes >= startMinutes && hourInMinutes < endMinutes) {
                occupiedCells.add(`${day}-${hour}`);
            }
        });

        return "";
    }
</script>
<Table.Root>
  <Table.Caption>Schedule</Table.Caption>
  <Table.Header>
    <Table.Row class="hover:bg-initial">
      <Table.Head class="w-1/12">Hour</Table.Head>
      {#each days as day}
        <Table.Head>{day}</Table.Head>
      {/each}
    </Table.Row>
  </Table.Header>

  <Table.Body>
    <!-- Iteramos sobre las celdas preprocesadas sin bucles anidados -->
    {#each hours as hour}
      <Table.Row class="!w-full hover:bg-initial">
        <!-- Mostrar la hora de la fila -->
        <Table.Cell class="font-medium">{hour}</Table.Cell>

        {#each days as day}
          {@const cell = scheduleCells.find(cell => cell.day === day && cell.hour === hour)}
          {#if cell && cell.subject && !isCellOccupied(day, hour)}
            <!-- Marcar las horas ocupadas por esta celda -->
            {markCellOccupied(day, cell.subject.classSchedule[day.toLowerCase()][0], cell.subject.classSchedule[day.toLowerCase()][1])}

            <!-- Usar Tailwind CSS para el color dinámico -->
            <Table.Cell rowspan={calculateRowSpan(cell.subject.classSchedule[day.toLowerCase()][0], cell.subject.classSchedule[day.toLowerCase()][1])}
                        class={`relative ${generateTailwindColorClass(cell.subject.id)}`}>
              <div class="p-1 flex flex-col items-center">
                <strong>{cell.subject.name}</strong>
                <span>Código: {cell.subject.code}</span>
                <span>Créditos: {cell.subject.credits}</span>
                <span>Semestre: {cell.subject.semester}</span>
                <span>Horas: {cell.subject.pHours}P {cell.subject.tHours}T {cell.subject.lHours}L</span>
              </div>
            </Table.Cell>
          {:else}
            <!-- Celda vacía si no hay materia o si ya fue ocupada -->
            <Table.Cell class="w-1/6"></Table.Cell>
          {/if}
        {/each}
      </Table.Row>
    {/each}
  </Table.Body>
</Table.Root>