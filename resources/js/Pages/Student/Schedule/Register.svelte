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

    // Generar un color único para cada materia
    /**
     * @param {number} id
     * @returns {string}
     */
    function generateColor(id) {
        const colors = ['#FF5733', '#33FF57', '#3357FF', '#FF33A1', '#FFA133', '#33FFA5', '#A533FF'];
        return colors[id % colors.length];
    }

    // Verificar si varias materias se deben mostrar en la hora y día actual
    /**
     * @param {ScheduleSubjectDto[]} subjects
     * @param {string} day
     * @param {string} currentHour
     * @returns {ScheduleSubjectDto[]}
     */
    function getSubjectsForTime(subjects, day, currentHour) {
        return subjects.filter(subject => {
            const schedule = subject.classSchedule[day.toLowerCase()];
            if (schedule && schedule.length === 2) {
                const [startTime, endTime] = schedule;

                // Verificar si el currentHour cae dentro del rango de startTime y endTime
                const currentTimeMinutes = timeToMinutes(currentHour);
                const startTimeMinutes = timeToMinutes(startTime);
                const endTimeMinutes = timeToMinutes(endTime);

                return currentTimeMinutes >= startTimeMinutes && currentTimeMinutes < endTimeMinutes;
            }
            return false;
        });
    }

    // Marcar las horas ocupadas por una materia en occupiedHours
    /**
     * @param {Set<string>} occupiedHours
     * @param {string} day
     * @param {string} startTime
     * @param {string} endTime
     * @returns {Set<string>}
     */
    function markOccupiedHours(occupiedHours, day, startTime, endTime) {
        const start = timeToMinutes(startTime);
        const end = timeToMinutes(endTime);

        hours.forEach(hour => {
            const time = timeToMinutes(hour);
            if (time >= start && time < end) {
                occupiedHours.add(day + hour);
            }
        });

        return occupiedHours;
    }
</script>

<StudentLayout title="Schedule Registration" description="Register for your subjects here"
               keywords="schedule, registration, student">
  <div class="flex flex-col space-y-4">
    <h1 class="text-2xl font-semibold">Schedule Registration</h1>
  </div>

  <Table.Root>
    <Table.Caption>Schedule</Table.Caption>
    <Table.Header>
      <Table.Row>
        <Table.Head class="w-1/12">Hour</Table.Head>
        {#each days as day}
          <Table.Head>{day}</Table.Head>
        {/each}
      </Table.Row>
    </Table.Header>

    <Table.Body>
      {#each hours as hour}
        <Table.Row>
          <!-- Mostrar la hora de la fila -->
          <Table.Cell class="font-medium">{hour}</Table.Cell>

          <!-- Lista para rastrear las horas ocupadas -->
          {@const occupiedHours = new Set()}

          {#each days as day}
            {#if !occupiedHours.has(day + hour)}
              {@const subjectInfos = getSubjectsForTime(subjects, day, hour)}

              {#if subjectInfos.length > 0}
                <!-- Marcar las horas cubiertas por cada materia como ocupadas -->
                {#each subjectInfos as subjectInfo}
                  {@const one = markOccupiedHours(occupiedHours, day, subjectInfo.classSchedule[day.toLowerCase()][0], subjectInfo.classSchedule[day.toLowerCase()][1])}

                  <!-- Mostrar la materia -->
                  <Table.Cell
                          rowspan={calculateRowSpan(subjectInfo.classSchedule[day.toLowerCase()][0], subjectInfo.classSchedule[day.toLowerCase()][1])}
                          class="relative" style="background-color: {generateColor(subjectInfo.id)};">
                    <div class="text-white font-bold p-2">
                      <strong>{subjectInfo.name}</strong> <br/>
                      <span>Código: {subjectInfo.code}</span> <br/>
                      <span>Créditos: {subjectInfo.credits}</span> <br/>
                      <span>Semestre: {subjectInfo.semester}</span>
                    </div>
                  </Table.Cell>
                {/each}
              {:else}
                <!-- Celda vacía si no hay materia -->
                <Table.Cell></Table.Cell>
              {/if}
            {:else}
              <!-- Celda vacía si está ocupada por un rowspan anterior -->
              <Table.Cell></Table.Cell>
            {/if}
          {/each}
        </Table.Row>
      {/each}
    </Table.Body>
  </Table.Root>
</StudentLayout>
