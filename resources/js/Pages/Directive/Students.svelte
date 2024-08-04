<script>
  import DataTable from "$lib/datatable/data-table.svelte";
  import { createTable } from "svelte-headless-table";
  import { readable } from "svelte/store";

  /** @type { Array<{id: number, avatar: string, name: string, email: string, phone: string, career: string, totalAverage: number}>} */
  export let students = [];

  const table = createTable(readable(students));
  const columns = table.createColumns([
    table.column({
      accessor: "id",
      header: "ID",
    }),
    table.column({
      accessor: "avatar",
      header: " ",
    }),
    table.column({
      accessor: "name",
      header: "Name",
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
      accessor: ({ id }) => id,
      header: "Actions",
    }),
  ]);
</script>

<DataTable {table} {columns} />
