<script>
  import DataTable from "$lib/components/datatable/data-table.svelte";
  import { createTable, createRender } from "svelte-headless-table";
  import { addPagination } from "svelte-headless-table/plugins";
  import { readable } from "svelte/store";
  import DataTableActions from "$lib/components/datatable/data-table-actions.svelte";
  import Avatar from "$lib/components/Avatar.svelte";
  /** @type { Array<{id: number, avatar: string, name: string, email: string, phone: string, career: string, totalAverage: number}>} */
  export let students = [];

  let table = createTable(readable(students), {
    page: addPagination(),
  });
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

  table = createTable(readable(students), {
    page: addPagination(),
  });
  const columns = table.createColumns([
    table.column({
      accessor: "id",
      header: "ID",
    }),
    table.column({
      accessor: ({ avatar, name }) => ({ avatar, name }),
      header: "Student",
      cell: ({ value }) => {
        return createRender(Avatar, {
          src: value.avatar,
          alt: "Avatar",
          name: value.name,
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
      accessor: ({ id }) => id,
      header: "Actions",
      cell: ({ value }) => {
        return createRender(DataTableActions, { id: value.toString() });
      },
    }),
  ]);
</script>

<DataTable {table} {columns} />
