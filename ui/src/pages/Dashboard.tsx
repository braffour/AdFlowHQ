import { BarChart3, Users, TrendingUp, Activity } from 'lucide-react'

export function Dashboard() {
  const stats = [
    { name: 'Active Workflows', value: '12', icon: Activity, change: '+2.5%', changeType: 'positive' },
    { name: 'Total Leads', value: '1,234', icon: Users, change: '+12.3%', changeType: 'positive' },
    { name: 'Conversion Rate', value: '8.2%', icon: TrendingUp, change: '+0.8%', changeType: 'positive' },
    { name: 'AI Analysis', value: '89%', icon: BarChart3, change: '+5.1%', changeType: 'positive' },
  ]

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
        <p className="mt-2 text-gray-600">Monitor your workflow performance and lead management</p>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4 mb-8">
        {stats.map((stat) => (
          <div key={stat.name} className="card">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <stat.icon className="h-8 w-8 text-primary-600" />
              </div>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-500">{stat.name}</p>
                <p className="text-2xl font-semibold text-gray-900">{stat.value}</p>
                <p className={`text-sm ${stat.changeType === 'positive' ? 'text-green-600' : 'text-red-600'}`}>
                  {stat.change}
                </p>
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Recent Activity */}
      <div className="card">
        <h2 className="text-lg font-semibold text-gray-900 mb-4">Recent Activity</h2>
        <div className="space-y-4">
          <div className="flex items-center space-x-3">
            <div className="w-2 h-2 bg-green-500 rounded-full"></div>
            <span className="text-sm text-gray-600">Workflow "Lead Processing" completed for tenant "acme"</span>
            <span className="text-xs text-gray-400">2 minutes ago</span>
          </div>
          <div className="flex items-center space-x-3">
            <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
            <span className="text-sm text-gray-600">AI analysis completed for lead "lead123"</span>
            <span className="text-xs text-gray-400">5 minutes ago</span>
          </div>
          <div className="flex items-center space-x-3">
            <div className="w-2 h-2 bg-yellow-500 rounded-full"></div>
            <span className="text-sm text-gray-600">Google Ads sync started for tenant "default"</span>
            <span className="text-xs text-gray-400">10 minutes ago</span>
          </div>
        </div>
      </div>
    </div>
  )
} 