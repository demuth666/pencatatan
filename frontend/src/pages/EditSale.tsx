import { useNavigate, useParams } from 'react-router-dom';
import { useMutation, useQueryClient, useQuery } from '@tanstack/react-query';
import { api } from '@/lib/api';
import { SaleForm } from '@/components/sales/SaleForm';
import { CreateSaleRequest, Sale, ApiResponse } from '@/types';
import { useToast } from '@/hooks/use-toast';
import { ArrowLeft, Loader2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Link } from 'react-router-dom';

export default function EditSale() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { toast } = useToast();
  const queryClient = useQueryClient();

  const { data: sale, isLoading } = useQuery({
      queryKey: ['sales', id],
      queryFn: async () => {
          const res = await api.get<ApiResponse<Sale>>(`/sales/${id}`);
          return res.data.data;
      },
      enabled: !!id
  });

  const mutation = useMutation({
    mutationFn: async (data: CreateSaleRequest) => { // Using Create request type as Update is similar subset
      await api.put(`/sales/${id}`, data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['sales'] });
      toast({
        title: "Success",
        description: "Sale updated successfully",
      });
      navigate('/sales');
    },
    onError: (error: any) => {
        const msg = error.response?.data?.error || "Failed to update sale";
        toast({
            title: "Error",
            description: msg,
            variant: "destructive",
        });
    },
  });

  if (isLoading) {
      return (
          <div className="flex justify-center items-center h-64">
              <Loader2 className="h-8 w-8 animate-spin text-slate-500" />
          </div>
      );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Link to="/sales">
           <Button variant="outline" size="icon">
             <ArrowLeft className="h-4 w-4"/>
           </Button>
        </Link>
        <h1 className="text-3xl font-bold tracking-tight">Edit Sale</h1>
      </div>

      <div className="rounded-md border bg-white p-6 shadow-sm max-w-2xl">
        {sale && (
            <SaleForm
                initialData={sale}
                onSubmit={async (values) => mutation.mutate(values)}
                isLoading={mutation.isPending}
                buttonText="Update Sale"
            />
        )}
      </div>
    </div>
  );
}
