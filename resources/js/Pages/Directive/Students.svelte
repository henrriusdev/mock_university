<script>
  import DataTable from "$lib/datatable/data-table.svelte";
  import { createTable } from "svelte-headless-table";
  import { readable } from "svelte/store";

  /** @type { Array<{id: number, avatar: string, name: string, email: string, phone: string, career: string, totalAverage: number}>} */
  export let students = [];

  let table = createTable(readable(students));
  if (students?.length === 0 || students === null) {
    students = [
      {
        id: 1,
        avatar: "https://randomuser.me/api/portraits/men/1.jpg",
        name: "John Doe",
        email: "johndoe@gmail.com",
        phone: "1234567890",
        career: "Computer Science",
        totalAverage: 90,
      },
      {
        id: 2,
        avatar: "https://randomuser.me/api/portraits/women/1.jpg",
        name: "Jane Doe",
        email: "janedoe@gmail.com",
        phone: "1234567890",
        career: "Computer Science",
        totalAverage: 85,
      },
    ];
  }
  console.log(students?.length, students);

  table = createTable(readable(students));
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
