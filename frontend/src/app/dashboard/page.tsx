'use client';

import React, { useState, useEffect } from 'react';
import { Clock, TrendingUp, Users, FileText, AlertTriangle } from 'lucide-react';

// Define proper interfaces for our data
interface StatItem {
  id: string;
  name: string;
  value: string | number;
  icon: React.ReactNode;
  trend?: {
    value: string;
    isUp: boolean;
  };
}

interface Activity {
  id: string;
  type: 'login' | 'document' | 'user' | 'alert';
  content: string;
  timestamp: string;
  user?: {
    name: string;
    avatar: string;
  };
}

// Icon components
const LoginIcon = () => (
  <div className="p-2 bg-blue-100 rounded-full">
    <Clock className="h-5 w-5 text-blue-600" />
  </div>
);

const DocumentIcon = () => (
  <div className="p-2 bg-green-100 rounded-full">
    <FileText className="h-5 w-5 text-green-600" />
  </div>
);

const UserIcon = () => (
  <div className="p-2 bg-purple-100 rounded-full">
    <Users className="h-5 w-5 text-purple-600" />
  </div>
);

const AlertIcon = () => (
  <div className="p-2 bg-red-100 rounded-full">
    <AlertTriangle className="h-5 w-5 text-red-600" />
  </div>
);

// Stat card component
const StatCard = ({ stat }: { stat: StatItem }) => (
  <div className="bg-white overflow-hidden shadow rounded-lg">
    <div className="p-5">
      <div className="flex items-center">
        <div className="flex-shrink-0">{stat.icon}</div>
        <div className="ml-5 w-0 flex-1">
          <dl>
            <dt className="text-sm font-medium text-gray-500 truncate">{stat.name}</dt>
            <dd>
              <div className="text-lg font-medium text-gray-900">{stat.value}</div>
            </dd>
          </dl>
        </div>
      </div>
    </div>
    {stat.trend && (
      <div className="bg-gray-50 px-5 py-3">
        <div className="text-sm">
          <div className={`flex items-center ${stat.trend.isUp ? 'text-green-600' : 'text-red-600'}`}>
            {stat.trend.isUp ? (
              <TrendingUp className="h-4 w-4 mr-1" />
            ) : (
              <svg className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6" />
              </svg>
            )}
            <span className="sr-only">{stat.trend.isUp ? 'Increased' : 'Decreased'} by</span>
            {stat.trend.value}
          </div>
        </div>
      </div>
    )}
  </div>
);

// Activity component
const ActivityItem = ({ activity }: { activity: Activity }) => {
  const getIcon = () => {
    switch (activity.type) {
      case 'login':
        return <LoginIcon />;
      case 'document':
        return <DocumentIcon />;
      case 'user':
        return <UserIcon />;
      case 'alert':
        return <AlertIcon />;
      default:
        return null;
    }
  };

  return (
    <div className="relative pb-8">
      <div className="relative flex items-start space-x-3">
        {getIcon()}
        <div className="min-w-0 flex-1">
          <div>
            {activity.user && (
              <div className="text-sm font-medium text-gray-900">
                <span className="font-medium text-gray-900">
                  {activity.user.name}
                </span>
              </div>
            )}
            <p className="mt-0.5 text-sm text-gray-500">{activity.content}</p>
          </div>
          <div className="mt-2 text-sm text-gray-500">
            <p>{activity.timestamp}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

// Loading skeletons
const StatSkeleton = () => (
  <div className="bg-white overflow-hidden shadow rounded-lg animate-pulse">
    <div className="p-5">
      <div className="flex items-center">
        <div className="flex-shrink-0 bg-gray-200 h-10 w-10 rounded-full"></div>
        <div className="ml-5 w-0 flex-1">
          <div className="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
          <div className="h-6 bg-gray-200 rounded w-1/2"></div>
        </div>
      </div>
    </div>
    <div className="bg-gray-50 px-5 py-3">
      <div className="h-4 bg-gray-200 rounded w-1/4"></div>
    </div>
  </div>
);

const ActivitySkeleton = () => (
  <div className="relative pb-8 animate-pulse">
    <div className="relative flex items-start space-x-3">
      <div className="flex-shrink-0 bg-gray-200 h-10 w-10 rounded-full"></div>
      <div className="min-w-0 flex-1">
        <div className="h-4 bg-gray-200 rounded w-1/4 mb-2"></div>
        <div className="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
        <div className="h-4 bg-gray-200 rounded w-1/2 mt-2"></div>
      </div>
    </div>
  </div>
);

export default function Dashboard() {
  const [stats, setStats] = useState<StatItem[]>([]);
  const [recentActivity, setRecentActivity] = useState<Activity[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Simulate API call
    const fetchData = async () => {
      setIsLoading(true);
      // Simulate network delay
      await new Promise(resolve => setTimeout(resolve, 1500));
      
      // Mock data
      const statsData: StatItem[] = [
        {
          id: '1',
          name: 'Total Users',
          value: '24,331',
          icon: <div className="p-2 bg-blue-100 rounded-full"><Users className="h-5 w-5 text-blue-600" /></div>,
          trend: {
            value: '12% increase from last month',
            isUp: true
          }
        },
        {
          id: '2',
          name: 'Active Sessions',
          value: '1,324',
          icon: <div className="p-2 bg-green-100 rounded-full"><Clock className="h-5 w-5 text-green-600" /></div>,
          trend: {
            value: '7% increase from last week',
            isUp: true
          }
        },
        {
          id: '3',
          name: 'Documents Secured',
          value: '843,012',
          icon: <div className="p-2 bg-yellow-100 rounded-full"><FileText className="h-5 w-5 text-yellow-600" /></div>,
          trend: {
            value: '3% decrease from yesterday',
            isUp: false
          }
        },
        {
          id: '4',
          name: 'Security Alerts',
          value: '23',
          icon: <div className="p-2 bg-red-100 rounded-full"><AlertTriangle className="h-5 w-5 text-red-600" /></div>,
          trend: {
            value: '10% decrease from last week',
            isUp: true
          }
        }
      ];
      
      const activityData: Activity[] = [
        {
          id: '1',
          type: 'login',
          content: 'Successfully logged in from new device',
          timestamp: '5 minutes ago',
          user: {
            name: 'Alice Johnson',
            avatar: '/avatars/alice.jpg'
          }
        },
        {
          id: '2',
          type: 'document',
          content: 'Created document "2023 Security Compliance Report"',
          timestamp: '2 hours ago',
          user: {
            name: 'Bob Smith',
            avatar: '/avatars/bob.jpg'
          }
        },
        {
          id: '3',
          type: 'user',
          content: 'Added 15 new users to Engineering team',
          timestamp: 'Yesterday at 1:34 PM',
          user: {
            name: 'Carol Williams',
            avatar: '/avatars/carol.jpg'
          }
        },
        {
          id: '4',
          type: 'alert',
          content: 'Multiple failed login attempts detected',
          timestamp: '2 days ago'
        },
        {
          id: '5',
          type: 'document',
          content: 'Updated permissions for "Financial Records Q3"',
          timestamp: '3 days ago',
          user: {
            name: 'Dave Miller',
            avatar: '/avatars/dave.jpg'
          }
        }
      ];
      
      setStats(statsData);
      setRecentActivity(activityData);
      setIsLoading(false);
    };

    fetchData();
  }, []);

  return (
    <div className="py-6">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
        <h1 className="text-2xl font-semibold text-gray-900">Dashboard</h1>
      </div>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
        <div className="py-4">
          {/* Stats section */}
          <div className="mt-4">
            <h2 className="text-lg leading-6 font-medium text-gray-900 mb-4">Overview</h2>
            <div className="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
              {isLoading
                ? Array(4).fill(0).map((_, i) => <StatSkeleton key={i} />)
                : stats.map(stat => <StatCard key={stat.id} stat={stat} />)
              }
            </div>
          </div>
          
          {/* Recent activity section */}
          <div className="mt-10">
            <h2 className="text-lg leading-6 font-medium text-gray-900 mb-4">Recent Activity</h2>
            <div className="flow-root bg-white px-4 py-5 rounded-lg shadow">
              <ul className="-mb-8">
                {isLoading
                  ? Array(5).fill(0).map((_, i) => <li key={i}><ActivitySkeleton /></li>)
                  : recentActivity.map(activity => (
                      <li key={activity.id}>
                        <ActivityItem activity={activity} />
                      </li>
                    ))
                }
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
} 