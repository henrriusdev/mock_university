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
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
  import DataTableCheckbox from "$lib/components/datatable/data-table-checkbox.svelte";
  import Button from "$lib/components/ui/button/button.svelte";
  import { ChevronRight } from "lucide-svelte";
  /** @type { Array<{id: number, avatar: string, name: string, email: string, phone: string, career: string, totalAverage: number}>} */
  export let students = [];

  let table = createTable(readable(students), {
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
      label: "Copy student name",
      /** @param {{name: string}} param0 */
      onClick: ({ name }) => {
        window.navigator.clipboard.writeText(name);
      },
    },
    {
      label: "Copy student email",
      /** @param {{email: string}} param0 */
      onClick: ({ email }) => {
        window.navigator.clipboard.writeText(email);
      },
    },
    {
      label: "Copy student phone",
      /** @param {{phone: string}} param0 */
      onClick: ({ phone }) => {
        window.navigator.clipboard.writeText(phone);
      },
    },
    {
      label: "Edit",
      /** @param {{id: number}} param0 */
      onClick: ({ id }) => {
        window.location.href = `/directive/students/view?id=${id}`;
      },
    },
  ];

  table = createTable(readable(students), {
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
      accessor: ({ avatar, name }) => {
        return `${avatar} ${name}`;
      },
      header: "Student",
      cell: ({ value }) => {
        return createRender(Avatar, {
          src: value.split(" ")[0],
          alt: "Avatar",
          name: value.split(" ").slice(1).join(" "),
        });
      },
    }),
    table.column({
      accessor: "email",
      header: "Email",
    }),
    table.column({
      accessor: "phone",
      header: "Phone",
    }),
    table.column({
      accessor: "career",
      header: "Career",
    }),
    table.column({
      accessor: "totalAverage",
      header: "Total Average",
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

<DirectiveLayout
  title="Students | Directive | Mock University"
  description="List of students registered in the system."
  keywords="Students, List, Table"
>
  <section
    class="flex flex-col !items-center justify-center max-w-full md:max-w-[90%] w-full sm:mx-auto p-3"
  >
    <div class="w-full flex justify-between items-center">
      <h2
        class="text-2xl md:text-3xl xl:text-5xl font-bold w-full text-left pb-1 md:pb-3"
      >
        Students
      </h2>
      <Button
        variant="outline"
        class="flex justify-center gap-x-3 items-center"
        href="/directive/students/view?id=add"
      >
        Add student
        <ChevronRight class="w-6 h-6" />
      </Button>
    </div>
    <div class="w-full">
      {#if students?.length === 0 || students === null}
        <p class="text-center text-lg font-semibold text-muted-foreground p-10">
          No students found. Add one.
        </p>
      {:else}
        <DataTable {table} {columns} />
      {/if}
    </div>
  </section>
</DirectiveLayout>
