import { useState } from 'react'
import { Play, Pause, Eye, Clock, CheckCircle, XCircle } from 'lucide-react'

interface Workflow {
  id: string
  name: string
  tenant: string
  leadId: string
  status: 'running' | 'completed' | 'failed' | 'pending'
  startedAt: string
  completedAt?: string
}

export function Workflows() {
  const [workflows] = useState<Workflow[]>([
    {
      id: 'acme_lead123_wf',
      name: 'Lead Processing',
      tenant: 'acme',
      leadId: 'lead123',
      status: 'completed',
      startedAt: '2024-01-15T10:30:00Z',
      completedAt: '2024-01-15T10:32:00Z',
    },
    {
      id: 'default_lead456_wf',
      name: 'Lead Processing',
      tenant: 'default',
      leadId: 'lead456',
      status: 'running',
      startedAt: '2024-01-15T10:35:00Z',
    },
    {
      id: 'acme_lead789_wf',
      name: 'Lead Processing',
      tenant: 'acme',
      leadId: 'lead789',
      status: 'failed',
      startedAt: '2024-01-15T10:25:00Z',
      completedAt: '2024-01-15T10:26:00Z',
    },
  ])

  const getStatusIcon = (status: Workflow['status']) => {
    switch (status) {
      case 'running':
        return <Clock className="h-5 w-5 text-blue-500" />
      case 'completed':
        return <CheckCircle className="h-5 w-5 text-green-500" />
      case 'failed':
        return <XCircle className="h-5 w-5 text-red-500" />
      case 'pending':
        return <Clock className="h-5 w-5 text-gray-500" />
    }
  }

  const getStatusColor = (status: Workflow['status']) => {
    switch (status) {
      case 'running':
        return 'bg-blue-100 text-blue-800'
      case 'completed':
        return 'bg-green-100 text-green-800'
      case 'failed':
        return 'bg-red-100 text-red-800'
      case 'pending':
        return 'bg-gray-100 text-gray-800'
    }
  }

  return (
    <div>
      <div className="mb-8 flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Workflows</h1>
          <p className="mt-2 text-gray-600">Monitor and manage your workflow executions</p>
        </div>
        <button className="btn-primary">
          <Play className="h-4 w-4 mr-2" />
          Start New Workflow
        </button>
      </div>

      <div className="card">
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Workflow
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Tenant
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Lead ID
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Status
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Started
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {workflows.map((workflow) => (
                <tr key={workflow.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="flex items-center">
                      {getStatusIcon(workflow.status)}
                      <div className="ml-3">
                        <div className="text-sm font-medium text-gray-900">{workflow.name}</div>
                        <div className="text-sm text-gray-500">{workflow.id}</div>
                      </div>
                    </div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {workflow.tenant}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {workflow.leadId}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${getStatusColor(workflow.status)}`}>
                      {workflow.status}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {new Date(workflow.startedAt).toLocaleString()}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button className="text-primary-600 hover:text-primary-900 mr-3">
                      <Eye className="h-4 w-4" />
                    </button>
                    {workflow.status === 'running' && (
                      <button className="text-red-600 hover:text-red-900">
                        <Pause className="h-4 w-4" />
                      </button>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
} 