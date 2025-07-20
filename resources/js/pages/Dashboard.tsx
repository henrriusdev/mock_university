import { useState } from 'react';
import { Head } from '@inertiajs/react';
import MainLayout from '../layouts/MainLayout';

// PrimeReact components
import { Card } from 'primereact/card';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Button } from 'primereact/button';
import { Chart } from 'primereact/chart';
import { TabView, TabPanel } from 'primereact/tabview';
import { Calendar } from 'primereact/calendar';
import { InputText } from 'primereact/inputtext';

interface Course {
  id: number;
  code: string;
  name: string;
  instructor: string;
  schedule: string;
}

interface DashboardProps {
  user: {
    name: string;
    email: string;
  };
}

export default function Dashboard({ user }: DashboardProps) {
  const [date, setDate] = useState<Date | null>(null);
  const [searchText, setSearchText] = useState('');

  // Sample data for courses
  const courses: Course[] = [
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

  return (
    <MainLayout user={user}>
      <Head title="Dashboard" />
      
      <div className="p-4">
        <h1 className="text-2xl font-bold mb-6">Welcome, {user.name}</h1>
        
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
          {/* Search and Calendar Card */}
          <Card title="Quick Actions" className="md:col-span-1">
            <div className="mb-4">
              <span className="p-input-icon-left w-full">
                <i className="pi pi-search" />
                <InputText 
                  value={searchText} 
                  onChange={(e) => setSearchText(e.target.value)} 
                  placeholder="Search courses..." 
                  className="w-full"
                />
              </span>
            </div>
            <div className="mb-4">
              <label className="block mb-2">Select Date</label>
              <Calendar 
                value={date} 
                onChange={(e) => e.value !== undefined ? setDate(e.value) : setDate(null)} 
                showIcon 
                className="w-full"
              />
            </div>
            <div className="flex justify-end">
              <Button label="Go to Calendar" icon="pi pi-calendar" className="p-button-outlined" />
            </div>
          </Card>
          
          {/* Progress Card */}
          <Card title="Academic Progress" className="md:col-span-2">
            <Chart type="bar" data={chartData} options={chartOptions} />
          </Card>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          {/* Courses Table */}
          <Card title="My Courses" className="md:col-span-2">
            <DataTable 
              value={courses} 
              paginator 
              rows={5} 
              rowsPerPageOptions={[5, 10, 25]} 
              tableStyle={{ minWidth: '50rem' }}
            >
              <Column field="code" header="Code" sortable style={{ width: '15%' }} />
              <Column field="name" header="Course" sortable style={{ width: '35%' }} />
              <Column field="instructor" header="Instructor" sortable style={{ width: '25%' }} />
              <Column field="schedule" header="Schedule" style={{ width: '25%' }} />
            </DataTable>
          </Card>
          
          {/* Upcoming Events */}
          <Card title="Upcoming Events" className="md:col-span-1">
            <TabView>
              <TabPanel header="Events">
                <ul className="list-none p-0 m-0">
                  {events.map((event) => (
                    <li key={event.id} className="flex align-items-center py-2 border-bottom-1 border-gray-300">
                      <span className="mr-2 font-bold">{event.date}</span>
                      <span>{event.title}</span>
                    </li>
                  ))}
                </ul>
              </TabPanel>
              <TabPanel header="Deadlines">
                <p className="m-0">View your upcoming assignment deadlines and exam dates.</p>
              </TabPanel>
            </TabView>
          </Card>
        </div>
      </div>
    </MainLayout>
  );
}
