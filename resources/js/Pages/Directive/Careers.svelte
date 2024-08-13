<script>
  import Avatar from "$lib/components/Avatar.svelte";
  import DataTableActions from "$lib/components/datatable/data-table-actions.svelte";
  import DataTable from "$lib/components/datatable/data-table.svelte";
  import Button from "$lib/components/ui/button/button.svelte";
  import { Checkbox } from "$lib/components/ui/checkbox";
  import { Input } from "$lib/components/ui/input";
  import { Label } from "$lib/components/ui/label";
  import { Textarea } from "$lib/components/ui/textarea";
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
  import { createRender, createTable } from "svelte-headless-table";
  import {
    addPagination,
    addSortBy,
    addTableFilter
  } from "svelte-headless-table/plugins";
  import { readable } from "svelte/store";
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

  const actions = [
    {
      label: "Edit",
      onClick: () => {
        console.log("Edit");
      },
    },
    {
      label: "Delete",
      onClick: () => {
        console.log("Delete");
      },
    },
  ];

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

  let checked = false;
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
    </div>
    <div class="flex justify-between items-start w-full gap-x-5">
      <div class="w-2/3">
        {#if careers?.length === 0 || careers === null}
          <p
            class="text-center text-lg font-semibold text-muted-foreground p-10"
          >
            No careers found. Add one.
          </p>
        {:else}
          <DataTable {table} {columns} />
        {/if}
      </div>
      <form
        method="post"
        action="/directive/careers/submit"
        class="w-1/3 p-10 flex flex-col gap-y-4 border-l-2 border-l-gray-300 dark:border-l-gray-800"
      >
        <h2 class="text-2xl font-bold text-center pb-3">Add Career</h2>
        <div class="flex flex-col gap-3">
          <Label for="name">Name</Label>
          <Input
            type="text"
            id="name"
            name="name"
            class="input"
            placeholder="Career name"
          />
        </div>
        <div class="flex flex-col gap-3">
          <Label for="description">Description</Label>
          <Textarea
            id="description"
            name="description"
            class="input"
            placeholder="Career description"
          />
        </div>
        <div class="flex flex-col gap-3">
          <div class="flex items-center space-x-2">
            <Checkbox id="leader" bind:checked aria-labelledby="leader-label" />
            <Label
              id="leader-label"
              for="leader"
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
            >
              Add a leader (Optional)
            </Label>
          </div>
          {#if checked}
            <Label for="leader">Leader</Label>
            <Input
              type="text"
              id="leader"
              name="leader"
              class="input"
              placeholder="Leader name"
            />
          {/if}
        </div>
        <Button
          type="submit"
          variant="default"
          class="w-full"
          on:click={() => (checked = !checked)}>
          Add Career
        </Button>
      </form>
    </div>
  </section>
</DirectiveLayout>
