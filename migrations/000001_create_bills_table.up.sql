CREATE TABLE IF NOT EXISTS Bills (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  id_factura text NOT NULL,
  fecha_emision text NOT NULL,
  monto_total integer NOT NULL,
  nombre_proveedor text NOT NULL,
  id_proveedor text NOT NULL,
  version integer NOT NULL DEFAULT 1,
  detalles jsonb,
  miscelaneo json
);


