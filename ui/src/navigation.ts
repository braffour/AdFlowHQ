import { BarChart3, Workflow, Settings } from 'lucide-react'

export interface NavItem {
  name: string
  href: string
  icon: React.ComponentType<React.SVGProps<SVGSVGElement>>
}

export const navigation: NavItem[] = [
  { name: 'Dashboard', href: '/', icon: BarChart3 },
  { name: 'Workflows', href: '/workflows', icon: Workflow },
  { name: 'Settings', href: '/settings', icon: Settings },
]
