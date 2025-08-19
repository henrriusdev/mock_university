<script>
  import MainLayout from '../layouts/MainLayout.svelte';
  import { Chart, registerables } from 'chart.js';
  import { onMount } from 'svelte';
  
  export let user;
  
  let date = null;
  let searchText = '';
  let chartCanvas;
  let chartInstance;
  
  // Sample data for courses
  const courses = [
    { id: 1, code: 'CS101', name: 'Introduction to Computer Science', instructor: 'Dr. Smith', schedule: 'Mon/Wed 10:00-11:30' },
    { id: 2, code: 'MATH201', name: 'Calculus II', instructor: 'Dr. Johnson', schedule: 'Tue/Thu 13:00-14:30' },
    { id: 3, code: 'ENG105', name: 'Academic Writing', instructor: 'Prof. Williams', schedule: 'Fri 09:00-12:00' },
    { id: 4, code: 'PHYS150', name: 'Physics for Engineers', instructor: 'Dr. Brown', schedule: 'Mon/Wed 14:00-15:30' },
  ];

  // Chart data
  const chartData = {
    labels: ['Assignments', 'Quizzes', 'Midterm', 'Final', 'Participation'],
    datasets: [
      {
        label: 'Current Grade',
        data: [85, 78, 92, 0, 88],
        backgroundColor: [
          'rgba(54, 162, 235, 0.2)',
          'rgba(75, 192, 192, 0.2)',
          'rgba(255, 206, 86, 0.2)',
          'rgba(255, 99, 132, 0.2)',
          'rgba(153, 102, 255, 0.2)',
        ],
        borderColor: [
          'rgb(54, 162, 235)',
          'rgb(75, 192, 192)',
          'rgb(255, 206, 86)',
          'rgb(255, 99, 132)',
          'rgb(153, 102, 255)',
        ],
        borderWidth: 1,
      },
    ],
  };

  // Chart options
  const chartOptions = {
    scales: {
      y: {
        beginAtZero: true,
        max: 100,
      },
    },
    plugins: {
      legend: {
        position: 'bottom',
      },
    },
  };

  // Upcoming events
  const events = [
    { id: 1, title: 'Assignment Due: CS101', date: '2025-07-25' },
    { id: 2, title: 'Midterm Exam: MATH201', date: '2025-07-30' },
    { id: 3, title: 'Group Project Meeting', date: '2025-07-22' },
    { id: 4, title: 'Office Hours: Dr. Smith', date: '2025-07-21' },
  ];
  
  let activeTab = 0;
  let currentPage = 0;
  const rowsPerPage = 5;
  
  $: paginatedCourses = courses.slice(currentPage * rowsPerPage, (currentPage + 1) * rowsPerPage);
  $: totalPages = Math.ceil(courses.length / rowsPerPage);
  
  function nextPage() {
    if (currentPage < totalPages - 1) {
      currentPage++;
    }
  }
  
  function prevPage() {
    if (currentPage > 0) {
      currentPage--;
    }
  }
  
  function goToPage(page) {
    currentPage = page;
  }
  
  onMount(() => {
    Chart.register(...registerables);
    
    if (chartCanvas) {
      chartInstance = new Chart(chartCanvas, {
        type: 'bar',
        data: chartData,
        options: chartOptions
      });
    }
    
    return () => {
      if (chartInstance) {
        chartInstance.destroy();
      }
    };
  });
</script>

<Head title="Dashboard" />

<MainLayout {user}>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-6">Welcome, {user.name}</h1>
    
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
      <!-- Search and Calendar Card -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 md:col-span-1">
        <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-white">Quick Actions</h3>
        <div class="mb-4">
          <div class="relative">
            <i class="pi pi-search absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"></i>
            <input 
              type="text"
              bind:value={searchText}
              placeholder="Search courses..."
              class="w-full pl-10 pr-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>
        <div class="mb-4">
          <label class="block mb-2 text-gray-700 dark:text-gray-300">Select Date</label>
          <input 
            type="date"
            bind:value={date}
            class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
        <div class="flex justify-end">
          <button class="px-4 py-2 border border-blue-500 text-blue-500 rounded-md hover:bg-blue-50 dark:hover:bg-blue-900 flex items-center">
            <i class="pi pi-calendar mr-2"></i>
            Go to Calendar
          </button>
        </div>
      </div>
      
      <!-- Progress Card -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 md:col-span-2">
        <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-white">Academic Progress</h3>
        <canvas bind:this={chartCanvas} class="w-full h-64"></canvas>
      </div>
    </div>
    
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <!-- Courses Table -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 md:col-span-2">
        <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-white">My Courses</h3>
        <div class="overflow-x-auto">
          <table class="w-full">
            <thead>
              <tr class="border-b border-gray-200 dark:border-gray-700">
                <th class="text-left py-3 px-4 text-gray-700 dark:text-gray-300">Code</th>
                <th class="text-left py-3 px-4 text-gray-700 dark:text-gray-300">Course</th>
                <th class="text-left py-3 px-4 text-gray-700 dark:text-gray-300">Instructor</th>
                <th class="text-left py-3 px-4 text-gray-700 dark:text-gray-300">Schedule</th>
              </tr>
            </thead>
            <tbody>
              {#each paginatedCourses as course}
                <tr class="border-b border-gray-100 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700">
                  <td class="py-3 px-4 text-gray-900 dark:text-white font-medium">{course.code}</td>
                  <td class="py-3 px-4 text-gray-900 dark:text-white">{course.name}</td>
                  <td class="py-3 px-4 text-gray-700 dark:text-gray-300">{course.instructor}</td>
                  <td class="py-3 px-4 text-gray-700 dark:text-gray-300">{course.schedule}</td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
        
        <!-- Pagination -->
        <div class="flex items-center justify-between mt-4">
          <div class="text-sm text-gray-700 dark:text-gray-300">
            Showing {currentPage * rowsPerPage + 1} to {Math.min((currentPage + 1) * rowsPerPage, courses.length)} of {courses.length} entries
          </div>
          <div class="flex space-x-2">
            <button 
              on:click={prevPage}
              disabled={currentPage === 0}
              class="px-3 py-1 border border-gray-300 dark:border-gray-600 rounded text-gray-700 dark:text-gray-300 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50 dark:hover:bg-gray-700"
            >
              Previous
            </button>
            {#each Array(totalPages) as _, i}
              <button 
                on:click={() => goToPage(i)}
                class="px-3 py-1 border border-gray-300 dark:border-gray-600 rounded {currentPage === i ? 'bg-blue-500 text-white' : 'text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700'}"
              >
                {i + 1}
              </button>
            {/each}
            <button 
              on:click={nextPage}
              disabled={currentPage === totalPages - 1}
              class="px-3 py-1 border border-gray-300 dark:border-gray-600 rounded text-gray-700 dark:text-gray-300 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50 dark:hover:bg-gray-700"
            >
              Next
            </button>
          </div>
        </div>
      </div>
      
      <!-- Upcoming Events -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 md:col-span-1">
        <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-white">Upcoming Events</h3>
        <div class="border-b border-gray-200 dark:border-gray-700">
          <nav class="flex space-x-8">
            <button 
              on:click={() => activeTab = 0}
              class="py-2 px-1 border-b-2 font-medium text-sm {activeTab === 0 ? 'border-blue-500 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'}"
            >
              Events
            </button>
            <button 
              on:click={() => activeTab = 1}
              class="py-2 px-1 border-b-2 font-medium text-sm {activeTab === 1 ? 'border-blue-500 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'}"
            >
              Deadlines
            </button>
          </nav>
        </div>
        
        <div class="mt-4">
          {#if activeTab === 0}
            <ul class="space-y-3">
              {#each events as event}
                <li class="flex items-start py-2 border-b border-gray-100 dark:border-gray-700">
                  <span class="font-bold text-gray-900 dark:text-white mr-3 text-sm">{event.date}</span>
                  <span class="text-gray-700 dark:text-gray-300 text-sm">{event.title}</span>
                </li>
              {/each}
            </ul>
          {:else}
            <p class="text-gray-600 dark:text-gray-400 text-sm">View your upcoming assignment deadlines and exam dates.</p>
          {/if}
        </div>
      </div>
    </div>
  </div>
</MainLayout>
