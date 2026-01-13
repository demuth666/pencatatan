import { useNavigate } from 'react-router-dom';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { api } from '@/lib/api';
import { SaleForm } from '@/components/sales/SaleForm';
import { CreateSaleRequest } from '@/types';
import { useToast } from '@/hooks/use-toast';
import { ArrowLeft } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Link } from 'react-router-dom';

export default function CreateSale() {
  const navigate = useNavigate();
  const { toast } = useToast();
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: async (data: CreateSaleRequest) => {
      await api.post('/sales', data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['sales'] });
      toast({
        title: "Success",
        description: "Sale recorded successfully",
      });
      navigate('/sales');
    },
    onError: (error: any) => {
        const msg = error.response?.data?.error || "Failed to create sale";
        toast({
            title: "Error",
            description: msg,
            variant: "destructive",
        });
    },
  });

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Link to="/sales">
           <Button variant="outline" size="icon">
             <ArrowLeft className="h-4 w-4"/>
           </Button>
        </Link>
        <h1 className="text-3xl font-bold tracking-tight">New Sale</h1>
      </div>

      <div className="rounded-md border bg-white p-6 shadow-sm max-w-2xl">
        <SaleForm
            onSubmit={async (values) => mutation.mutate(values)}
            isLoading={mutation.isPending}
            buttonText="Record Sale"
        />
      </div>
    </div>
  );
}
