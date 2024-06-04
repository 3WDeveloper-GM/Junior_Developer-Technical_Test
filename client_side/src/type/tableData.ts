export type TableData = {

  billID: string;
  totalAmount: number;
  emissionDate: string;
  provider: {
    name: string;
    id: string;
  };
  details: object;
  miscelaneous: object;
};
