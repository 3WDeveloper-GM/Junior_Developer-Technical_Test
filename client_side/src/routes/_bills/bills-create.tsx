import { BillCreateView } from '@/layout/billCreate'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_bills/bills-create')({
  component: () => <BillCreateView />
})
