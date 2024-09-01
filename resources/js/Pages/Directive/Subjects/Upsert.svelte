<script>
  /** @typedef {{id: number, name: string, description: string, code: string, creditUnits: number, semester: number, praticeHours: number, theoryHours: number, labHours: number, classSchedule: {[key: string]: string[]}, professorId: number, professorName: string, careers: {id: number, name: string}[]}} Subject */
  import { Camera, ChevronLeft } from "lucide-svelte";

  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import { Input, MaskInput } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  import DirectiveLayout from "$lib/layouts/DirectiveLayout.svelte";
  import DatePicker from "$lib/components/DatePicker.svelte";
  import Textarea from "$lib/components/ui/textarea/textarea.svelte";
  import { identityMask } from "$lib/utils";
  import * as Select from "$lib/components/ui/select/index.js";
  import Combobox from "$lib/components/Combobox.svelte";
  import { CalendarDate } from "@internationalized/date";

  /** @type {Subject | null} */
  export let subject = null;

  /** @type {Array<{id: number, name: string}>} */
  export let careers = [];

  // birth date is a string, so, parse it to a CalendarDate object, its format is 'YYYY-MM-DD'
</script>

<DirectiveLayout
  title="{subject?.id ? `Update ${subject?.name}` : 'Add Student'}"
  description="Add or update a student of the MockUniversity"
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
              Add Subject
            {/if}
          </h1>
        </div>
        <Card.Root>
          <Card.Header>
            <Card.Title>Subject Information</Card.Title>
          </Card.Header>
          <Card.Content>
            <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
              <div>
                <Label for="name">Name</Label>
                <Input
                  id="name"
                  name="name"
                  type="text"
                  bind:value="{subject?.name}"
                  required
                />
              </div>
              <div>
                <Label for="code">Code</Label>
                <Input
                  id="code"
                  name="code"
                  type="text"
                  bind:value="{subject?.code}"
                  required
                />
              </div>
              <div>
                <Label for="description">Description</Label>
                <Textarea
                  id="description"
                  name="description"
                  bind:value="{subject?.description}"
                  required
                />
              </div>
              <div>
                <Label for="creditUnits">Credit Units</Label>
                <Input
                  id="creditUnits"
                  name="creditUnits"
                  type="number"
                  bind:value="{subject?.creditUnits}"
                  required
                />
              </div>
              <div>
                <Label for="semester">Semester</Label>
                <Input
                  id="semester"
                  name="semester"
                  type="number"
                  bind:value="{subject?.semester}"
                  required
                />
              </div>
              <div>
                <Label for="praticeHours">Pratice Hours</Label>
                <Input
                  id="praticeHours"
                  name="praticeHours"
                  type="number"
                  bind:value="{subject?.praticeHours}"
                  required
                />
              </div>
              <div>
                <Label for="theoryHours">Theory Hours</Label>
                <Input
                  id="theoryHours"
                  name="theoryHours"
                  type="number"
                  bind:value="{subject?.theoryHours}"
                  required
                />
              </div>
              <div>
                <Label for="labHours">Lab Hours</Label>
                <Input
                  id="labHours"
                  name="labHours"
                  type="number"
                  bind:value="{subject?.labHours}"
                  required
                />
              </div>
              <div>
                <Label for="classSchedule">Class Schedule</Label>
                <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
                  {#each Object.entries(subject?.classSchedule || {}) as [day, hours]}
                    <div>
                      <Label for="day">{day}</Label>
                      <!-- <Combobox
                        id="day"
                        name="day"
                        bind:value="{hours}"
                        multiple
                        options="{[
                          '08:00',
                          '10:00',
                          '14:00',
                          '16:00',
                          '18:00',
                        ]}"
                      /> -->
                    </div>
                  {/each}
                </div>
              </div>
              <div>
                <Label for="professorId">Professor</Label>
                <!-- <Select
                  id="professorId"
                  name="professorId"
                  bind:value="{subject?.professorId}"
                  required
                >
                  <Select.Option value="" disabled
                    >Select a professor</Select.Option
                  >
                  <Select.Option value="1">Professor 1</Select.Option>
                  <Select.Option value="2">Professor 2</Select.Option>
                  <Select.Option value="3">Professor 3</Select.Option>
                </Select> -->
                <input
                  type="hidden"
                  name="professorId"
                  value="{subject?.professorId}"
                />
              </div>
            </div>
          </Card.Content>
        </Card.Root>
      </div>
    </main>
  </section>
</DirectiveLayout>
