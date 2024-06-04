import { BillsReadView } from '@/layout/billRead'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_bills/bills-read')({
  component: () => <div>
<BillsReadView />
  </div>
})
