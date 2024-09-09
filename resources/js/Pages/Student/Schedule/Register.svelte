<script>
    import StudentLayout from "$lib/layouts/StudentLayout.svelte";
    import Schedule from "$lib/components/Schedule.svelte";

    const days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
    const hours = ["7:00", "7:30", "8:00", "8:30", "9:00", "9:30", "10:00", "10:30", "11:00", "11:30", "12:00", "12:30", "13:00", "13:30", "14:00", "14:30", "15:00", "15:30", "16:00", "16:30", "17:00", "17:30", "18:00", "18:30", "19:00"];

    /** @typedef {Object} ScheduleSubjectDto
     * @property {number} id
     * @property {string} name
     * @property {string} description
     * @property {string} code
     * @property {number} semester
     * @property {number} credits
     * @property {number} pHours
     * @property {number} tHours
     * @property {number} lHours
     * @property {Object<string, string[]>} classSchedule
     */

    /** @type {ScheduleSubjectDto[]} */
    export let subjects = [];
</script>

<StudentLayout title="Schedule Registration" description="Register for your subjects here"
               keywords="schedule, registration, student">
  <div class="flex flex-col space-y-4">
    <h1 class="text-2xl font-semibold">Schedule Registration</h1>
  </div>

  <Schedule subjects={subjects}/>

</StudentLayout>
