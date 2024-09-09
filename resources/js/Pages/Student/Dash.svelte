<script>
    import DataTable from "$lib/components/datatable/data-table.svelte";
    import {createRender, createTable} from "svelte-headless-table";
    import {addPagination, addSelectedRows, addSortBy, addTableFilter,} from "svelte-headless-table/plugins";
    import {readable} from "svelte/store";
    import DataTableActions from "$lib/components/datatable/data-table-actions.svelte";
    import StudentLayout from "$lib/layouts/StudentLayout.svelte";
    import DataTableCheckbox from "$lib/components/datatable/data-table-checkbox.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import {ChevronRight} from "lucide-svelte";
    import {Badge} from "$lib/components/ui/badge/index.js";
    /** @typedef {{[key: string]: string[]}} Schedule */
    /** @type { Array<{id: number, subject: string, notes: number[], avg: number}> } */
    export let notes = [];
    export let scheduleRegistrationStart = "";
    export let scheduleRegistrationEnd = "";
    export let notesNumber = 0;
    export let userName = "";

    let table = createTable(readable(notes), {
        page: addPagination(),
        sort: addSortBy(),
        filter: addTableFilter({
            /** @param {{filterValue: string, value: string}} param0 */
            fn: ({filterValue, value}) =>
                value.toLowerCase().includes(filterValue.toLowerCase()),
        }),
        select: addSelectedRows(),
    });

    const actions = [
        {
            label: "Copy subject name",
            /** @param {{name: string}} param0 */
            onClick: ({name}) => {
                window.navigator.clipboard.writeText(name);
            },
        },
        {
            label: "Copy subject code",
            /** @param {{code: string}} param0 */
            onClick: ({code}) => {
                window.navigator.clipboard.writeText(code);
            },
        },
        {
            label: "Copy hours",
            /** @param {{totalHours: number, practiceHours: number, theoryHours: number, labHours: number}} param0 */
            onClick: ({totalHours, practiceHours, labHours, theoryHours}) => {
                let hours = `Practice: ${practiceHours}h
Theory: ${theoryHours}h
Lab: ${labHours}h
Total: ${totalHours}h`;
                window.navigator.clipboard.writeText(hours);
            },
        },
        {
            label: "Copy careers",
            /** @param {{careers: Array<{name: string}>}} param0 */
            onClick: ({careers}) => {
                const careersNames = careers.map(({name}) => name).join(", ");
                window.navigator.clipboard.writeText(careersNames);
            },
        },
        {
            label: "Copy class schedule",
            /** @param {{classSchedule: Schedule}} param0 */
            onClick: ({classSchedule}) => {
                const schedule = Object.entries(classSchedule).map(
                    ([day, hours]) => `${day}: ${hours.join(", ")}`
                );
                window.navigator.clipboard.writeText(schedule.join("\n"));
            },
        },
        {
            label: "Edit",
            /** @param {{id: number}} param0 */
            onClick: ({id}) => {
                window.location.href = `/directive/subjects/view?id=${id}`;
            }
        }
    ];

    table = createTable(readable(notes), {
        page: addPagination(),
        sort: addSortBy(),
        filter: addTableFilter({
            /** @param {{filterValue: string, value: string}} param0 */
            fn: ({filterValue, value}) =>
                value.toLowerCase().includes(filterValue.toLowerCase()),
        }),
        select: addSelectedRows(),
    });

    /** @type {import('svelte-headless-table').Column<{id: number, subject: string, notes: number[], avg: number}, {page: any, sort: any, filter: any, select: any }>[] } */
    const noteColumns = [];
    const cardinal = ["1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th", "10th"];
    for (let i = 0; i < notesNumber; i++) {
        noteColumns.push(
            table.column({
                accessor: ({notes}) => notes[i] ?? 0,
                header: `${cardinal[i]} note`,
            })
        );
    }

    const columns = table.createColumns([
        table.column({
            accessor: "id",
            header: (_, {pluginStates}) => {
                const {allPageRowsSelected} = pluginStates.select;
                return createRender(DataTableCheckbox, {
                    checked: allPageRowsSelected,
                });
            },
            cell: ({row}, {pluginStates}) => {
                const {getRowState} = pluginStates.select;
                const {isSelected} = getRowState(row);

                return createRender(DataTableCheckbox, {
                    checked: isSelected,
                });
            },
            plugins: {
                sort: {
                    disable: true,
                },
                filter: {
                    exclude: true,
                },
            },
        }),
        table.column({
            accessor: ({subject}) => subject,
            header: "Subject",
        }),
        ...noteColumns,
        table.column({
            accessor: (row) => row,
            header: "Actions",
            cell: ({value}) => {
                return createRender(DataTableActions, {row: value, actions});
            },
            plugins: {
                sort: {
                    disable: true,
                },
                filter: {
                    exclude: true,
                },
            },
        }),
    ]);

    const scheduleStart = new Date(scheduleRegistrationStart);
    const scheduleEnd = new Date(scheduleRegistrationEnd);
</script>

<StudentLayout title="Dashboard" description="Dashboard for Acme Inc." keywords="dashboard, acme, inc">
  <section
          class="flex flex-col !items-start justify-center max-w-full md:max-w-[90%] w-full sm:mx-auto p-3"
  >
    <div class="w-full flex justify-between items-start">
      <h2
              class="text-2xl md:text-3xl xl:text-4xl font-bold w-full text-left pb-1 md:pb-3 flex items-start flex-col gap-y-2"
      >
        <span class="text-capitalize">Welcome, {userName}!</span>
        {#if scheduleStart > new Date()}
          <Badge variant="ghost" class="ml-2">
            Schedule registration starts on {scheduleStart.toLocaleDateString()}
          </Badge>
        {:else if scheduleEnd < new Date()}
          <Badge variant="destructive" class="ml-2">
            Schedule registration ended on {scheduleEnd.toLocaleDateString()}
          </Badge>
        {:else}
          <Badge variant="default" class="ml-2">
            Schedule registration is open until {scheduleEnd.toLocaleDateString()}
          </Badge>
        {/if}

      </h2>
      <Button
              variant="outline"
              class="flex justify-center gap-x-3 items-center"
              href="/student/schedule"
      >
        Register my schedule
        <ChevronRight class="w-6 h-6"/>
      </Button>
    </div>
    <div class="w-full">
      {#if notes?.length === 0 || notes === null}
        <p class="text-center text-lg font-semibold text-muted-foreground p-10">
          No notes found, please register your schedule
        </p>
      {:else}
        <DataTable {table} {columns}/>
      {/if}
    </div>
  </section>
</StudentLayout>