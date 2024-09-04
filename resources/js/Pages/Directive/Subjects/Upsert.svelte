<script>
  /** @typedef {{id: number, name: string, description: string, code: string, creditUnits: number, semester: number, praticeHours: number, theoryHours: number, labHours: number, classSchedule: {[key: string]: string[]}, professorId: number, professorName: string, careers: {id: number, name: string}[]}} Subject */
  import { ChevronLeft } from "lucide-svelte";

  /** @typedef {Array<{id: number, name: string}>} Select */
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import { Input, MaskInput } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
  import Textarea from "$lib/components/ui/textarea/textarea.svelte";
  import * as Select from "$lib/components/ui/select/index.js";
  import Combobox from "$lib/components/Combobox.svelte";
  import IMask from "imask";

  /** @type {Subject | null} */
  export let subject = null;

    /** @type {Select} */
  export let careers = [];

    /** @type {Select} */
  export let professors = [];

  /** @type {string} */
  let career = '';

  /** @type {string} */
  let professor = '';

  console.log(subject);
</script>

<DirectiveLayout
  title="{subject?.id ? `Update ${subject?.name}` : 'Add subject'}"
  description="Add or update a subject of the MockUniversity"
  keywords="add, update, inactive, mocku"
>
  <section class="bg-muted/40 flex w-full flex-col p-4">
    <main class="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
      <div class="mx-auto grid max-w-[59rem] flex-1 auto-rows-max gap-4">
        <div class="flex items-center gap-4 py-1">
          <Button
                  variant="ghost"
                  size="icon"
                  class="h-7 w-7"
                  href="/directive/subjects"
          >
            <ChevronLeft class="h-5 w-5" />
            <span class="sr-only">Back</span>
          </Button>
          <h1
                  class="flex-1 shrink-0 whitespace-nowrap text-xl font-semibold tracking-tight sm:grow-0"
          >
            {#if subject?.id}
              Update {subject?.name}
            {:else}
              Add subject
            {/if}
          </h1>
        </div>
        <form method="post" action="/directive/subjects/view/submit" enctype="multipart/form-data"
              class="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8"
        >
          <div
                  class="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8"
          >
            <Card.Root>
              <Card.Header>
                <Card.Title>General information</Card.Title>
                <Card.Description>
                  Fill the form with the subject's information.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-4">
                  <span
                          class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                  >
                  {#if subject?.id}
                    <input type="hidden" name="id" value="{subject?.id}" />
                  {/if}
                    <Label for="name">Name</Label>
                    <Input
                            type="text"
                            id="name"
                            value="{subject?.name ?? ''}"
                            name="name"
                            maxlength="100"
                    />
                  </span>
                  <span
                          class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                  >
                    <Label for="code">Code</Label>
                    <MaskInput
                            id="code"
                            value="{subject?.code ?? ''}"
                            name="code"
                            imask="{{ mask: 'aaaa-000' }}"
                    />
                  </span>
                  <span
                          class="text-sm font-semibold text-muted-foreground lg:col-span-4"
                  >
                    <Label for="description">Description</Label>
                    <Textarea
                            id="description"
                            value="{subject?.description ?? ''}"
                            name="description"
                    />
                  </span>
                </div>
              </Card.Content>
            </Card.Root>
            <Card.Root>
              <Card.Header>
                <Card.Title>Academic information</Card.Title>
                <Card.Description>
                  Fill the form with the subject's academical information.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-5">
                  <span
                          class="text-sm font-semibold text-muted-foreground lg:col-span-1"
                  >
                    <Label for="creditUnits">Credit Units</Label>
                    <MaskInput
                            id="creditUnits"
                            unmask="none"
                            imask={{ mask: Number, min: 0, max: 22 }}
                            value="{subject?.creditUnits?.toString() ?? ''}"
                            name="creditUnits"
                    />
                  </span>
                  <span
                          class="text-sm font-semibold text-muted-foreground lg:col-span-1"
                  >
                    <Label for="practiceHours">Practice Hours</Label>
                    <MaskInput
                            id="practiceHours"
                            unmask="none"
                            imask={{ mask: Number, min: 0, max: 5 }}
                            value="{subject?.praticeHours?.toString() ?? ''}"
                            name="practiceHours"
                    />
                  </span>
                  <span
                          class="text-sm font-semibold text-muted-foreground lg:col-span-1"
                  >
                    <Label for="theoryHours">Theory Hours</Label>
                    <MaskInput
                            id="theoryHours"
                            unmask="none"
                            imask={{ mask: Number, min: 0, max: 5 }}
                            value="{subject?.theoryHours?.toString() ?? ''}"
                            name="theoryHours"
                    />
                  </span>
                  <span
                          class="text-sm font-semibold text-muted-foreground lg:col-span-1"
                  >
                    <Label for="labHours">Lab Hours</Label>
                    <MaskInput
                            id="labHours"
                            unmask="none"
                            imask={{ mask: Number, min: 0, max: 5 }}
                            value="{subject?.labHours?.toString() ?? ''}"
                            name="labHours"
                    />
                  </span>
                  <span
                          class="text-sm font-semibold text-muted-foreground lg:col-span-1"
                  >
                    <Label for="totalAverage">Semester</Label>
                    <Input
                            id="semester"
                            type="number"
                            min="1"
                            max="10"
                            value="{subject?.semester}"
                            name="semester"
                    />
                  </span>
                </div>
              </Card.Content>
            </Card.Root>
            <Card.Root>
              <Card.Header>
                <Card.Title>Class Schedule</Card.Title>
                <Card.Description>
                    Fill the form with the subject's class schedule, by the selected practice, theory and laboratory hours.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="grid gap-4 lg:grid-cols-6">
                  {#if Object.keys(subject?.classSchedule ?? {}).length === 0}
                    <span
                            class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                    >
                      <Label for="classSchedule">Class Schedule</Label>
                      <Select.Root>
                        <Select.Trigger>
                          <Select.Value placeholder="Select a day"/>
                        </Select.Trigger>
                        <Select.Content>
                          <Select.Item value="monday">Monday</Select.Item>
                          <Select.Item value="tuesday">Tuesday</Select.Item>
                          <Select.Item value="wednesday">Wednesday</Select.Item>
                          <Select.Item value="thursday">Thursday</Select.Item>
                          <Select.Item value="friday">Friday</Select.Item>
                          <Select.Item value="saturday">Saturday</Select.Item>
                          <Select.Item value="sunday">Sunday</Select.Item>
                        </Select.Content>
                      </Select.Root>
                    </span>
                    <span
                            class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                    >
                        <Label for="classSchedule">From</Label>
                        <MaskInput
                                id="classSchedule"
                                value=""
                                name="classSchedule"
                                imask={{ mask: 'HH:MM', blocks: { HH: { mask: IMask.MaskedRange, from: 7, to: 18 }, MM: { mask: IMask.MaskedRange, from: 0, to: 59 } } }}
                        />
                    </span>
                    <span
                            class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                    >
                        <Label for="classSchedule">To</Label>
                        <MaskInput
                                id="classSchedule"
                                value=""
                                name="classSchedule"
                                imask={{ mask: 'HH:MM', blocks: { HH: { mask: IMask.MaskedRange, from: 7, to: 18 }, MM: { mask: IMask.MaskedRange, from: 0, to: 59 } } }}
                        />
                    </span>
                  {:else}
                  {#each Object.keys(subject?.classSchedule ?? {}) as day}
                    <span
                            class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                    >
                      <Label for="classSchedule">Class Schedule</Label>
                      <Select.Root selected={{
                        value: day,
                        label: day.charAt(0).toUpperCase() + day.slice(1)
                      }}>
                        <Select.Trigger>
                          <Select.Value placeholder="Select a day"/>
                        </Select.Trigger>
                        <Select.Content>
                          <Select.Item value="monday">Monday</Select.Item>
                          <Select.Item value="tuesday">Tuesday</Select.Item>
                          <Select.Item value="wednesday">Wednesday</Select.Item>
                          <Select.Item value="thursday">Thursday</Select.Item>
                          <Select.Item value="friday">Friday</Select.Item>
                          <Select.Item value="saturday">Saturday</Select.Item>
                          <Select.Item value="sunday">Sunday</Select.Item>
                        </Select.Content>
                      </Select.Root>
                    </span>
                    <span
                            class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                    >
                        <Label for="classSchedule">From</Label>
                        <MaskInput
                                id="classSchedule"
                                value="{subject?.classSchedule[day][0] ?? ''}"
                                name="classSchedule"
                                imask={{ mask: 'HH:MM', blocks: { HH: { mask: IMask.MaskedRange, from: 0, to: 23 }, MM: { mask: IMask.MaskedRange, from: 0, to: 59 } } }}
                        />
                    </span>
                    <span
                            class="text-sm font-semibold text-muted-foreground lg:col-span-2"
                    >
                        <Label for="classSchedule">To</Label>
                        <MaskInput
                                id="classSchedule"
                                value="{subject?.classSchedule[day][1] ?? ''}"
                                name="classSchedule"
                                imask={{ mask: 'HH:MM', blocks: { HH: { mask: IMask.MaskedRange, from: 0, to: 23 }, MM: { mask: IMask.MaskedRange, from: 0, to: 59 } } }}
                        />
                    </span>
                  {/each}
                    {/if}

                </div>
              </Card.Content>
            </Card.Root>
          </div>
          <div class="grid auto-rows-max items-start gap-4 lg:gap-8">
            <Card.Root>
              <Card.Header>
                <Card.Title>Professor And Careers</Card.Title>
                <Card.Description>
                  Fill the form with the professor and careers (can be one or more) of the subject.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <!-- Make a image rounded of the avatar with a button for change it -->
                <div class="flex flex-col gap-4">
                    <span
                        class="text-sm font-semibold text-muted-foreground lg:col-span-2 flex flex-col justify-end gap-y-2"
                    >
                        <Label for="professor">Professor</Label>
                        <Combobox options={professors.map((p) => ({ value: p.id.toString(), label: p.name }))} bind:value={professor} />
                        <input type="hidden" id="professor" name="professor" bind:value={professor} />
                    </span>
                  <span
                      class="text-sm font-semibold text-muted-foreground lg:col-span-2 flex flex-col justify-end gap-y-2"
                  >
                    <Label for="career">Career</Label>
                    <Combobox multiple options={careers.map((c) => ({ value: c.id.toString(), label: c.name }))} bind:value={career} />
                    <input type="hidden" id="career" name="career" bind:value={career} />
                  </span>
                </div>
              </Card.Content>
            </Card.Root>
            <Card.Root>
              <Card.Header>
                <Card.Title>Actions</Card.Title>
                <Card.Description>
                  Save or cancel the changes made to the subject. Also inactive
                  the subject.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <div class="flex flex-col gap-4">
                  <Button variant="ghost" type="reset">Cancel</Button>
                  <Button variant="default" type="submit">
                    {subject?.id ? "Update" : "Add"} subject
                  </Button>
                  {#if subject?.id}
                    <Button variant="destructive" type="button">
                      Inactive subject
                    </Button>
                  {/if}
                </div>
              </Card.Content>
            </Card.Root>
          </div>
        </form>
      </div>
    </main>
  </section>
</DirectiveLayout>
