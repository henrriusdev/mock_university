<script>
  import * as Table from "$lib/components/ui/table/index";
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
  import { Render, Subscribe } from "svelte-headless-table";

  /**
   * @typedef {import('svelte-headless-table').Table<any, any>} DataTable
   * @typedef {import('svelte-headless-table').Column<any, any>} DataTableColumn
   */

  /** @type {DataTableColumn[]} */
  export let columns;

  /** @type {DataTable} */
  export let table;

  const { headerRows, pageRows, tableAttrs, tableBodyAttrs } =
    table.createViewModel(columns);
</script>

<DirectiveLayout
  title="Students List | Directive | Mock University"
  description="Students List of the Mock University."
  keywords="students, list, directive"
>
  <div class="rounded-md border">
    <Table.Root {...$tableAttrs}>
      <Table.Header>
        {#each $headerRows as headerRow}
          <Subscribe rowAttrs={headerRow.attrs()}>
            <Table.Row>
              {#each headerRow.cells as cell (cell.id)}
                <Subscribe attrs={cell.attrs()} let:attrs props={cell.props()}>
                  <Table.Head {...attrs}>
                    <Render of={cell.render()} />
                  </Table.Head>
                </Subscribe>
              {/each}
            </Table.Row>
          </Subscribe>
        {/each}
      </Table.Header>
      <Table.Body {...$tableBodyAttrs}>
        {#each $pageRows as row (row.id)}
          <Table.Row {...row.attrs()}>
            {#each row.cells as cell (cell.id)}
              <Subscribe attrs={cell.attrs()} let:attrs>
                <Table.Cell {...attrs}>
                  <Render of={cell.render()} />
                </Table.Cell>
              </Subscribe>
            {/each}
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </div>
</DirectiveLayout>
