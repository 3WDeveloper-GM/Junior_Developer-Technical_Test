import { BillRemoveView } from '@/layout/billRemove'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_bills/bills-remove')({
  component: () => <BillRemoveView /> })
