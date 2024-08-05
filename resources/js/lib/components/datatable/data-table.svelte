<script>
  import * as Table from "$lib/components/ui/table/index";
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
  import { Render, Subscribe } from "svelte-headless-table";
  import { Button } from "$lib/components/ui/button/index.js";

  /**
   * @typedef {import('svelte-headless-table').Table<any, any>} DataTable
   * @typedef {import('svelte-headless-table').Column<any, any>} DataTableColumn
   */

  /** @type {DataTableColumn[]} */
  export let columns;

  /** @type {DataTable} */
  export let table;

  const { headerRows, pageRows, tableAttrs, tableBodyAttrs, pluginStates } =
    table.createViewModel(columns);

  const { hasNextPage, hasPreviousPage, pageIndex } = pluginStates.page;
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
                    {#if cell.id === "avatar"}
                      <div class="!w-8">
                        <Render of={cell.render()} />
                      </div>
                    {:else if cell.id === "totalAverage"}
                      <div class="text-right">
                        <Render of={cell.render()} />
                      </div>
                    {:else}
                      <Render of={cell.render()} />
                    {/if}
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
                  {#if cell.id === "avatar"}
                    <div class="!w-8">
                      <Render of={cell.render()} />
                    </div>
                  {:else if cell.id === "totalAverage"}
                    <div class="text-right">
                      <Render of={cell.render()} />
                    </div>
                  {:else}
                    <Render of={cell.render()} />
                  {/if}
                </Table.Cell>
              </Subscribe>
            {/each}
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </div>
  <div class="flex items-center justify-end space-x-4 py-4">
    <Button
      variant="outline"
      size="sm"
      on:click={() => ($pageIndex = $pageIndex - 1)}
      disabled={!$hasPreviousPage}>Previous</Button
    >
    <Button
      variant="outline"
      size="sm"
      disabled={!$hasNextPage}
      on:click={() => ($pageIndex = $pageIndex + 1)}>Next</Button
    >
  </div>
</DirectiveLayout>
