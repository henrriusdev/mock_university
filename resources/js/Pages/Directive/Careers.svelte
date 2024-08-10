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
  /** @type { Array<{id: number, name: string, leader?: string}>} */
  export let careers = [];

  let table = createTable(readable(careers), {
    page: addPagination(),
    sort: addSortBy(),
    filter: addTableFilter({
      /** @param {{filterValue: string, value: string}} param0 */
      fn: ({ filterValue, value }) =>
        value.toLowerCase().includes(filterValue.toLowerCase()),
    }),
  });

  const columns = table.createColumns([
    table.column({
      accessor: "id",
      header: "ID",
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
      accessor: "name",
      header: "Name",
      cell: ({ value }) => {
        return createRender(Avatar, {
          src: value.split(" ")[0],
          alt: "Avatar",
          name: value.split(" ").slice(1).join(" "),
        });
      },
    }),
    table.column({
      accessor: "leader",
      header: "Leader",
    }),
    table.column({
      accessor: ({ id }) => id,
      header: "Actions",
      cell: ({ value }) => {
        return createRender(DataTableActions, { id: value.toString() });
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
        Careers
      </h2>
      <Button
        variant="outline"
        class="flex justify-center gap-x-3 items-center"
      >
        Add career
        <ChevronRight class="w-6 h-6" />
      </Button>
    </div>
    <div class="w-full">
      {#if careers?.length === 0 || careers === null}
        <p class="text-center text-lg font-semibold text-muted-foreground p-10">
          No careers found. Add one.
        </p>
      {:else}
        <DataTable {table} {columns} />
      {/if}
    </div>
  </section>
</DirectiveLayout>
