import { FormFieldProps } from "../../type/formProps";

const FormField: React.FC<FormFieldProps> = ({
  type,
  placeholder,
  name,
  register,
  error,
  valueAsNumber,
  Setter,
}) => (
  <>
    <input
      className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
      type={type}
      placeholder={placeholder}
      {...register(name, { valueAsNumber })}
      onChange={(e) => Setter(e.currentTarget.value)}
    />
    {error && <span className="font-bold text-red-400">{error.message}</span>}
  </>
);
export default FormField;
