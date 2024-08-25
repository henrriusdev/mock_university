<script>
  import { Button } from "$lib/components/ui/button/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import * as Table from "$lib/components/ui/table/index";
  import { ArrowUpDown } from "lucide-svelte";
  import { Render, Subscribe } from "svelte-headless-table";

  /** @typedef {import('svelte-headless-table').Table<any, any>} DataTable */
  /** @typedef {import('svelte-headless-table').Column<any, any>} DataTableColumn */

  /** @type {DataTableColumn[]} */
  export let columns;
  /** @type {DataTable} */
  export let table;

  const { headerRows, pageRows, tableAttrs, tableBodyAttrs, pluginStates } =
    table.createViewModel(columns);

  const hasNextPage = pluginStates.page?.hasNextPage;
  const hasPreviousPage = pluginStates.page?.hasPreviousPage;
  let pageIndex = pluginStates.page?.pageIndex;
  const filterValue = pluginStates.filter?.filterValue;
  const selectedDataIds = pluginStates.select?.selectedDataIds;
</script>

<div class="rounded-md border">
  <div class="flex items-center p-4">
    {#if filterValue !== undefined}
      <Input
        class="max-w-sm"
        placeholder="Filter names, email, phone, career and average..."
        type="text"
        bind:value="{$filterValue}"
      />
    {/if}
  </div>
  <Table.Root {...$tableAttrs}>
    <Table.Header>
      {#each $headerRows as headerRow}
        <Subscribe rowAttrs="{headerRow.attrs()}">
          <Table.Row>
            {#each headerRow.cells as cell (cell.id)}
              <Subscribe
                attrs="{cell.attrs()}"
                let:attrs
                props="{cell.props()}"
                let:props
              >
                <Table.Head {...attrs} class="[&:has([role=checkbox])]:pl-4">
                  {#if !props.sort.disabled}
                    <Button
                      variant="ghost"
                      on:click="{props.sort.toggle}"
                      class="flex items-center gap-x-2"
                    >
                      <ArrowUpDown class="w-4 h-4" />
                      <Render of="{cell.render()}" />
                    </Button>
                  {:else}
                    <Render of="{cell.render()}" />
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
        <Table.Row
          {...row.attrs()}
          data-state="{selectedDataIds?.[row.id] && 'selected'}"
        >
          {#each row.cells as cell (cell.id)}
            <Subscribe attrs="{cell.attrs()}" let:attrs>
              <Table.Cell {...attrs}>
                <Render of="{cell.render()}" />
              </Table.Cell>
            </Subscribe>
          {/each}
        </Table.Row>
      {/each}
    </Table.Body>
  </Table.Root>
</div>
<div class="flex items-center justify-end space-x-4 py-4">
  {#if hasPreviousPage !== undefined && hasNextPage !== undefined && pageIndex !== undefined}
    <Button
      variant="outline"
      size="sm"
      on:click="{() => ($pageIndex = $pageIndex - 1)}"
      disabled="{!hasPreviousPage}"
    >
      Previous
    </Button>
    <Button
      variant="outline"
      size="sm"
      disabled="{!hasNextPage}"
      on:click="{() => ($pageIndex = $pageIndex + 1)}"
    >
      Next
    </Button>
  {/if}
</div>
