<script>
    /** @typedef {{id: number, name: string, description: string, code: string, creditUnits: number, semester: number, practiceHours: number, theoryHours: number, labHours: number, classSchedule: {[key: string]: string[]}, professorId: number, professorName: string, careers: {id: number, name: string}[], prerequisites: {id: number, name: string}[]}} Subject */
    import { ChevronLeft } from "lucide-svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Card from "$lib/components/ui/card/index.js";
    import { Input, MaskInput } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
    import Textarea from "$lib/components/ui/textarea/textarea.svelte";
    import * as Select from "$lib/components/ui/select/index.js";
    import Combobox from "$lib/components/Combobox.svelte";
    import IMask from "imask";

    /** @type {Subject} */
    export let subject = {
        id: 0,
        name: '',
        description: '',
        code: '',
        creditUnits: 0,
        semester: 0,
        practiceHours: 0,
        theoryHours: 0,
        labHours: 0,
        classSchedule: {},
        professorId: 0,
        professorName: '',
        careers: [],
        prerequisites: []
    };

    /** @type {Array<{id: number, name: string}>} */
    export let careers = [];

    /** @type {Array<{id: number, name: string}>} */
    export let professors = [];

    /** @type {Array<{id: number, name: string, code: string, semester: number}>} */
    export let subjects = [];

    /** @type {string} */
    let career = subject?.careers?.map(c => c.id.toString()).join(',') ?? '';

    /** @type {string} */
    let professor = subject?.professorId?.toString() ?? '';

    /** @type {string} */
    let prereq = subject?.prerequisites?.map(p => p.id.toString()).join(',') ?? '';

    let practiceHours = subject?.practiceHours?.toString() ?? "0";
    let theoryHours = subject?.theoryHours?.toString() ?? "0";
    let labHours = subject?.labHours?.toString() ?? "0";
    let totalHours = Number(practiceHours) + Number(theoryHours) + Number(labHours);
    let semester = subject?.semester?.toString() === "0" ? "1" : subject?.semester?.toString();

    /** @type {Array<{day: string, from: string | undefined, to: string | undefined, daySelected: { value: string, label: string }}>} */
    let scheduleEntries = Object.keys(subject?.classSchedule || {}).map(day => {
        return {
            day,
            from: subject?.classSchedule[day][0],
            to: subject?.classSchedule[day][1],
            daySelected: { value: day, label: day.charAt(0).toUpperCase() + day.slice(1) }
        };
    });

    // Calcular horas totales dinámicamente
    $: totalHours = Number(practiceHours) + Number(theoryHours) + Number(labHours);
    $: semester = Number(semester) > 10 ? "10" : (Number(semester) === 0 && semester !== "" ? "1" : semester);
    $: if (totalHours || scheduleEntries){
        validateSchedule();
    }

    $: if (semester){
        subjects = subjects.filter(s => s.semester < Number(semester));
    }

    function validateSchedule() {
        let accumulatedMinutes = 0;

        if (scheduleEntries.length === 0) {
            addEmptyScheduleEntry();
        }

        // Recorremos todas las entradas del horario para sumar los minutos acumulados
        for (let entry of scheduleEntries) {
            if (!entry.from || !entry.to) continue;

            let [fromHour, fromMinute] = entry.from.split(":").map(Number);
            let [toHour, toMinute] = entry.to.split(":").map(Number);

            // Convertir a minutos para calcular la diferencia
            let fromTotalMinutes = fromHour * 60 + fromMinute;
            let toTotalMinutes = toHour * 60 + toMinute;
            let minutesInDay = toTotalMinutes - fromTotalMinutes;

            if (minutesInDay <= 0) {
                alert(`El horario de ${entry.day} es inválido.`);
                return;
            }

            accumulatedMinutes += minutesInDay;

            let totalAllowedMinutes = totalHours * 60;
            if (accumulatedMinutes > totalAllowedMinutes) {
                alert("La cantidad de horas programadas excede el total de horas permitidas.");
                return;
            }
        }

        let totalAllowedMinutes = totalHours * 60;
        let remainingMinutes = totalAllowedMinutes - accumulatedMinutes;

        if (remainingMinutes > 0 && scheduleEntries.every(entry => entry.from && entry.to)) {
            addEmptyScheduleEntry();
        }

        if (remainingMinutes === 0 && scheduleEntries.length > 1) {
            let lastEntry = scheduleEntries[scheduleEntries.length - 1];
            if (!lastEntry.from && !lastEntry.to && !lastEntry.day) {
                scheduleEntries.pop();
            }
        }
    }

    function addEmptyScheduleEntry() {
        scheduleEntries.push({ day: '', from: '', to: '', daySelected: { value: '', label: '' } });
    }

    /**
     * @param {Event} event
     */
    function prepareDataBeforeSubmit(event) {
        // Construir el objeto classSchedule a partir de scheduleEntries
        const classSchedule = {};
        scheduleEntries.forEach((entry) => {
            if (entry.daySelected.value && entry.from && entry.to) {
                // @ts-ignore
                classSchedule[entry.daySelected.value] = [entry.from, entry.to];
            }
        });

        // Crear un input oculto para enviar classSchedule como JSON
        console.log(classSchedule, scheduleEntries);
        const classScheduleInput = document.createElement('input');
        classScheduleInput.type = 'hidden';
        classScheduleInput.name = 'classSchedule';
        classScheduleInput.value = JSON.stringify(classSchedule);

        // Agregar el input al formulario
        // @ts-ignore
        event.target.appendChild(classScheduleInput);
    }
</script>


<DirectiveLayout
        title="{subject?.id ? `Update ${subject?.name}` : 'Add subject'}"
        description="Add or update a subject of the MockUniversity"
        keywords="add, update, inactive, mocku"
>
  <section class="bg-muted/40 flex w-full flex-col p-4">
    <main class="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
      <div class="mx-auto grid max-w-[59rem] flex-1 auto-rows-max gap-4">
        <div class="flex items-center gap-4 py-1">
          <Button
                  variant="ghost"
                  size="icon"
                  class="h-7 w-7"
                  href="/directive/subjects"
          >
            <ChevronLeft class="h-5 w-5"/>
            <span class="sr-only">Back</span>
          </Button>
          <h1
                  class="flex-1 shrink-0 whitespace-nowrap text-xl font-semibold tracking-tight sm:grow-0"
          >
            {#if subject?.id}
              Update {subject?.name}
            {:else}
              Add subject
            {/if}
          </h1>
        </div>
        <form method="post" action="/directive/subjects/view/submit" enctype="multipart/form-data" class="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8" on:submit={prepareDataBeforeSubmit}>
          <div
                  class="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8"
          >
            <Card.Root>
              <Card.Header>
                <Card.Title>General information</Card.Title>
                <Card.Description>
                  Fill the form with the subject's information.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-4">
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-2">
                      {#if subject?.id}
                          <input type="hidden" name="id" value="{subject?.id}"/>
                      {/if}
                    <Label for="name">Name</Label>
                      <Input type="text" id="name" bind:value={subject.name} maxlength="100" name="name" required minlength="5"
                      />
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-2">
                      <Label for="code">Code</Label>
                      <MaskInput id="code" bind:value={subject.code} imask="{{ mask: 'aaaa-000' }}" name="code" required minlength="8" maxlength="8"
                      />
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-4">
                      <Label for="description">Description</Label>
                      <Textarea id="description" bind:value={subject.description} name="description" required minlength="10" maxlength="500" rows="3"
                      />
                  </span>
                </div>
              </Card.Content>
            </Card.Root>
            <!-- Información académica -->
            <Card.Root>
              <Card.Header>
                <Card.Title>Academic information</Card.Title>
                <Card.Description>
                  Fill the form with the subject's academical information.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-5">
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-1">
                      <Label for="creditUnits">Credit Units</Label>
                      <MaskInput id="creditUnits" unmask="none" imask={{ mask: Number, min: 0, max: 10 }} bind:value={subject.creditUnits} required name="creditUnits"
                      />
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-1">
                      <Label for="practiceHours">Practice Hours</Label>
                      <MaskInput id="practiceHours" unmask="none" imask={{ mask: Number, min: 0, max: 5 }} bind:value={practiceHours} required name="practiceHours"
                      />
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-1">
                      <Label for="theoryHours">Theory Hours</Label>
                      <MaskInput id="theoryHours" unmask="none" imask={{ mask: Number, min: 0, max: 5 }} bind:value={theoryHours} required name="theoryHours"
                      />
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-1">
                      <Label for="labHours">Lab Hours</Label>
                      <MaskInput id="labHours" unmask="none" imask={{ mask: Number, min: 0, max: 5 }} bind:value={labHours} required name="labHours"
                      />
                    <input type="hidden" name="totalHours" value="{totalHours}"/>
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-1">
                      <Label for="semester">Semester</Label>
                      <Input id="semester" type="number" min="1" max="10" bind:value={semester} required name="semester"
                      />
                  </span>
                </div>
              </Card.Content>
            </Card.Root>

            <!-- Horario de clases -->
            <Card.Root>
              <Card.Header>
                <Card.Title>Class Schedule</Card.Title>
                <Card.Description>
                  Fill the form with the subject's class schedule, respecting the total number of hours.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-6">
                  {#each scheduleEntries as {daySelected, from, to}, index}
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-2">
                      <Label for="classScheduleDay{index}">Day</Label>
                      <Select.Root bind:selected={daySelected} on:change="{e => scheduleEntries[index].day = e.detail}">
                          <Select.Trigger>
                              <Select.Value placeholder="Select a day"/>
                          </Select.Trigger>
                          <Select.Content>
                              <Select.Item value="monday">Monday</Select.Item>
                              <Select.Item value="tuesday">Tuesday</Select.Item>
                              <Select.Item value="wednesday">Wednesday</Select.Item>
                              <Select.Item value="thursday">Thursday</Select.Item>
                              <Select.Item value="friday">Friday</Select.Item>
                              <Select.Item value="saturday">Saturday</Select.Item>
                          </Select.Content>
                      </Select.Root>
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-2">
                      <Label for="classScheduleFrom{index}">From</Label>
                      <MaskInput id="classScheduleFrom{index}" imask={{ mask: 'HH:MM', blocks: { HH: { mask: IMask.MaskedRange, from: 7, to: 18 }, MM: { mask: IMask.MaskedRange, from: 0, to: 59 } } }} bind:value={from}/>
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-2">
                      <Label for="classScheduleTo{index}">To</Label>
                      <MaskInput id="classScheduleTo{index}" imask={{ mask: 'HH:MM', blocks: { HH: { mask: IMask.MaskedRange, from: 7, to: 18 }, MM: { mask: IMask.MaskedRange, from: 0, to: 59 } } }} bind:value={to}/>
                  </span>
                  {/each}
                </div>
              </Card.Content>
            </Card.Root>
          </div>
          <div class="grid auto-rows-max items-start gap-4 lg:gap-8">
            <Card.Root>
              <Card.Header>
                <Card.Title>Professor, Careers & Prerequisites</Card.Title>
                <Card.Description>
                  Fill the form with the professor, prerequisites and careers (can be one or more) of the subject.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="flex flex-col gap-4">
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-2 flex flex-col justify-end gap-y-2">
                      <Label for="professor">Professor</Label>
                      <Combobox options={professors.map((p) => ({ value: p.id.toString(), label: p.name }))} bind:value={professor}/>
                    <input type="hidden" name="professorId" value="{professor}"/>
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-2 flex flex-col justify-end gap-y-2">
                      <Label for="career">Career</Label>
                      <Combobox multiple options={careers.map((c) => ({ value: c.id.toString(), label: c.name }))} bind:value={career}/>
                    <input type="hidden" name="careers" value="{career}"/>
                  </span>
                  <span class="text-sm font-semibold text-muted-foreground lg:col-span-2">
                      <Label for="prerequisites">Prerequisites</Label>
                      <Combobox multiple options={subjects.map((s) => ({ value: s.id.toString(), label: s.name }))} bind:value={prereq}/>
                    <input type="hidden" name="prerequisites" value="{prereq}"/>
                  </span>
                </div>
              </Card.Content>
            </Card.Root>
            <Card.Root>
              <Card.Header>
                <Card.Title>Actions</Card.Title>
                <Card.Description>
                  Save or cancel the changes made to the subject. Also inactive
                  the subject.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="flex flex-col gap-4">
                  <Button variant="ghost" type="reset">Cancel</Button>
                  <Button variant="default" type="submit">
                    {subject?.id ? "Update" : "Add"} subject
                  </Button>
                </div>
              </Card.Content>
            </Card.Root>
          </div>
        </form>
      </div>
    </main>
  </section>
</DirectiveLayout>
