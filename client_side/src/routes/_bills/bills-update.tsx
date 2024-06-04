import { BillUpdateView } from '@/layout/billUpdate'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_bills/bills-update')({
  component: () => <BillUpdateView />
})
