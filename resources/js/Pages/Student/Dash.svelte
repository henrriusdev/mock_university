<script>
    import DataTable from "$lib/components/datatable/data-table.svelte";
    import { createTable, createRender } from "svelte-headless-table";
    import {
        addPagination,
        addSortBy,
        addTableFilter,
        addSelectedRows,
    } from "svelte-headless-table/plugins";
    import { readable } from "svelte/store";
    import DataTableActions from "$lib/components/datatable/data-table-actions.svelte";
    import Avatar from "$lib/components/Avatar.svelte";
    import StudentLayout from "$lib/layouts/StudentLayout.svelte";
    import DataTableCheckbox from "$lib/components/datatable/data-table-checkbox.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import { ChevronRight } from "lucide-svelte";
    /** @typedef {{[key: string]: string[]}} Schedule */
    /** @type { Array<{id: number, name: string, creditUnits: number, semester: number, careers: Array<{id: number, name: string}>, code: string, practiceHours: number, theoryHours: number, labHours: number, totalHours: number, classSchedule: Schedule }>} */
    export let subjects = [];

    let table = createTable(readable(subjects), {
        page: addPagination(),
        sort: addSortBy(),
        filter: addTableFilter({
            /** @param {{filterValue: string, value: string}} param0 */
            fn: ({ filterValue, value }) =>
                value.toLowerCase().includes(filterValue.toLowerCase()),
        }),
        select: addSelectedRows(),
    });

    const actions = [
        {
            label: "Copy subject name",
            /** @param {{name: string}} param0 */
            onClick: ({ name }) => {
                window.navigator.clipboard.writeText(name);
            },
        },
        {
            label: "Copy subject code",
            /** @param {{code: string}} param0 */
            onClick: ({ code }) => {
                window.navigator.clipboard.writeText(code);
            },
        },
        {
            label: "Copy hours",
            /** @param {{totalHours: number, practiceHours: number, theoryHours: number, labHours: number}} param0 */
            onClick: ({ totalHours,  practiceHours, labHours, theoryHours}) => {
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
            onClick: ({ careers }) => {
                const careersNames = careers.map(({ name }) => name).join(", ");
                window.navigator.clipboard.writeText(careersNames);
            },
        },
        {
            label: "Copy class schedule",
            /** @param {{classSchedule: Schedule}} param0 */
            onClick: ({ classSchedule }) => {
                const schedule = Object.entries(classSchedule).map(
                    ([day, hours]) => `${day}: ${hours.join(", ")}`
                );
                window.navigator.clipboard.writeText(schedule.join("\n"));
            },
        },
        {
            label:"Edit",
            /** @param {{id: number}} param0 */
            onClick: ({ id }) => {
                window.location.href = `/directive/subjects/view?id=${id}`;
            }
        }
    ];

    table = createTable(readable(subjects), {
        page: addPagination(),
        sort: addSortBy(),
        filter: addTableFilter({
            /** @param {{filterValue: string, value: string}} param0 */
            fn: ({ filterValue, value }) =>
                value.toLowerCase().includes(filterValue.toLowerCase()),
        }),
        select: addSelectedRows(),
    });

    const columns = table.createColumns([
        table.column({
            accessor: "id",
            header: (_, { pluginStates }) => {
                const { allPageRowsSelected } = pluginStates.select;
                return createRender(DataTableCheckbox, {
                    checked: allPageRowsSelected,
                });
            },
            cell: ({ row }, { pluginStates }) => {
                const { getRowState } = pluginStates.select;
                const { isSelected } = getRowState(row);

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
            accessor: ({ name }) => name,
            header: "Subject",
        }),
        table.column({
            accessor: ({ code }) => code,
            header: "Code",
        }),
        table.column({
            accessor: ({ creditUnits }) => creditUnits,
            header: "C.U",
            plugins: {
                sort: {
                    disable: true,
                },
            }
        }),
        table.column({
            accessor: ({ semester }) => semester,
            header: "Semester",
        }),
        table.column({
            accessor: ({careers}) => {
                const careersNames = careers.map(({ name }) => name).join(", ");
                if (careersNames.length > 50) {
                    return careersNames.slice(0, 30) + "...";
                }
                return careersNames;
            },
            header: "Careers",
        }),
        table.column({
            accessor: ({ practiceHours }) => practiceHours,
            header: "P.H",
            plugins: {
                sort: {
                    disable: true,
                },
            }
        }),
        table.column({
            accessor: ({ theoryHours }) => theoryHours,
            header: "T.H",
            plugins: {
                sort: {
                    disable: true,
                },
            }
        }),
        table.column({
            accessor: ({ labHours }) => labHours,
            header: "L.H",
            plugins: {
                sort: {
                    disable: true,
                },
            }
        }),
        table.column({
            accessor: ({ totalHours }) => totalHours,
            header: "Total",
            plugins: {
                sort: {
                    disable: true,
                },
            }
        }),
        table.column({
            accessor: ({ classSchedule }) => {
                if (classSchedule === undefined) {
                    return "No schedule";
                }
                const schedule = Object.entries(classSchedule).map(
                    ([day, hours]) => `${day}: ${hours.join("-")}`
                );
                if (schedule.join(", ").length > 50) {
                    return schedule.join(", ").slice(0, 30) + "...";
                }

                return schedule.join(", ");
            },
            header: "Schedule",
            cell: ({ value }) => {
                return value.replace(/, /g, "\n");
            },
        }),
        table.column({
            accessor: (row) => row,
            header: "Actions",
            cell: ({ value }) => {
                return createRender(DataTableActions, { row: value, actions });
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
</script>


<StudentLayout title="Dashboard" description="Dashboard for Acme Inc." keywords="dashboard, acme, inc">

  <section
          class="flex flex-col !items-center justify-center max-w-full md:max-w-[90%] w-full sm:mx-auto p-3"
  >
    <div class="w-full flex justify-between items-center">
      <h2
              class="text-2xl md:text-3xl xl:text-5xl font-bold w-full text-left pb-1 md:pb-3"
      >
        Welcome
      </h2>
      <Button
              variant="outline"
              class="flex justify-center gap-x-3 items-center"
              href="/directive/subjects/view?id=add"
      >
        Add subject
        <ChevronRight class="w-6 h-6" />
      </Button>
    </div>
    <div class="w-full">
      {#if subjects?.length === 0 || subjects === null}
        <p class="text-center text-lg font-semibold text-muted-foreground p-10">
          No subjects found. Add one.
        </p>
      {:else}
        <DataTable {table} {columns} />
      {/if}
    </div>
  </section>
</StudentLayout>