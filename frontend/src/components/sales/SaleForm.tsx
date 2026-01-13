import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { Button } from '@/components/ui/button';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Checkbox } from '@/components/ui/checkbox';
import { Sale, CreateSaleRequest } from '@/types';
import { Loader2 } from 'lucide-react';
import { useEffect } from 'react';

const formSchema = z.object({
  name: z.string().optional(),
  product: z.string().min(1, 'Product name is required'),
  quantity: z.coerce.number().min(1, 'Quantity must be at least 1'),
  price: z.coerce.number().min(0, 'Price must be non-negative'),
  amount_received: z.coerce.number().min(0, 'Amount received must be non-negative'),
  is_debt: z.boolean().default(false),
});

type FormValues = z.infer<typeof formSchema>;

interface SaleFormProps {
  initialData?: Sale;
  onSubmit: (values: CreateSaleRequest) => Promise<void>;
  isLoading: boolean;
  buttonText: string;
}

export function SaleForm({ initialData, onSubmit, isLoading, buttonText }: SaleFormProps) {
  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema) as any, // Cast to any to avoid strict type mismatch with coercion
    defaultValues: {
      name: '',
      product: '',
      quantity: 1,
      price: 0,
      amount_received: 0,
      is_debt: false,
    },
  });

  useEffect(() => {
    if (initialData) {
      form.reset({
        name: initialData.name,
        product: initialData.product,
        quantity: initialData.quantity,
        price: initialData.price,
        amount_received: initialData.amount_received,
        is_debt: initialData.is_debt,
      });
    }
  }, [initialData, form]);

  const handleSubmit = async (values: FormValues) => {
    // Validate amount received covers total (if not debt)
    // Though backend validates, good to validate frontend too?
    // Let's rely on backend or simple check:
    const total = values.quantity * values.price;
    if (!values.is_debt && values.amount_received < total) {
       form.setError("amount_received", {
         type: "manual",
         message: `Insufficient amount. Total is ${total}`
       });
       return;
    }

    await onSubmit({
      name: values.name,
      product: values.product,
      quantity: values.quantity,
      price: values.price,
      amount_received: values.amount_received,
      is_debt: values.is_debt,
    });
  };

  const quantity = form.watch('quantity');
  const price = form.watch('price');
  const total = (quantity || 0) * (price || 0);

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(handleSubmit)} className="space-y-6 max-w-lg">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Customer Name (Optional)</FormLabel>
              <FormControl>
                <Input placeholder="e.g. Budi" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="product"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Product Name</FormLabel>
              <FormControl>
                <Input placeholder="e.g. Nasi Goreng" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <div className="grid grid-cols-2 gap-4">
          <FormField
            control={form.control}
            name="quantity"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Quantity</FormLabel>
                <FormControl>
                  <Input type="number" min={1} {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="price"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Price (per unit)</FormLabel>
                <FormControl>
                  <Input type="number" min={0} {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="p-4 bg-slate-50 rounded-md">
           <div className="text-sm text-slate-500">Total Price</div>
           <div className="text-2xl font-bold">Rp {total.toLocaleString()}</div>
        </div>

        <FormField
          control={form.control}
          name="amount_received"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Amount Received</FormLabel>
              <FormControl>
                <Input type="number" min={0} {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="is_debt"
          render={({ field }) => (
            <FormItem className="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
              <FormControl>
                <Checkbox
                  checked={field.value}
                  onCheckedChange={field.onChange}
                />
              </FormControl>
              <div className="space-y-1 leading-none">
                <FormLabel>
                  Mark as Debt (Belum Lunas)
                </FormLabel>
                <FormDescription>
                   Checking this allows payment to be less than total.
                </FormDescription>
              </div>
            </FormItem>
          )}
        />

        <Button type="submit" disabled={isLoading} className="w-full">
          {isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          {buttonText}
        </Button>
      </form>
    </Form>
  );
}
