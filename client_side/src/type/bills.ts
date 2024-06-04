export type Bill = {
  numeroRegistro?: number
  idFactura: string;
  fechaEmision: string;
  montoTotal: number;
  proveedor: {
    nombre: string;
    identificacion: string;
  };
  detalles: object;
  miscelaneo: object;
};

export const emptyBill: Bill = {
  idFactura: "",
  fechaEmision: "",
  proveedor: {
    nombre: "",
    identificacion: "",
  },
  montoTotal: 0,
  detalles: {},
  miscelaneo: {},
};
