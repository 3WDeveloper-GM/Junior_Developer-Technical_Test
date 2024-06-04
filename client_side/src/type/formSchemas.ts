import { z } from "zod";

export const BillsFetchSchema = z.object({
  startDate: z.coerce.string().date(),
  endDate: z.coerce.string().date(),
});

