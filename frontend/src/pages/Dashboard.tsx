import { useQuery } from '@tanstack/react-query';
import { api } from '@/lib/api';
import { Sale, ApiResponse } from '@/types';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { DollarSign, ShoppingBag, CreditCard, TrendingUp } from 'lucide-react';
import { format, isToday } from 'date-fns';
import { Loader2 } from 'lucide-react';

export default function Dashboard() {
  const { data: sales, isLoading, error } = useQuery({
    queryKey: ['sales'],
    queryFn: async () => {
      const response = await api.get<ApiResponse<Sale[]>>('/sales');
      return response.data.data || [];
    },
  });

  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-64">
        <Loader2 className="h-8 w-8 animate-spin text-slate-500" />
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-center text-red-500 p-4">
        Error loading dashboard data.
      </div>
    );
  }

  const allSales = sales || [];

  const totalRevenue = allSales.reduce((acc, sale) => acc + sale.total, 0);
  const totalTransactions = allSales.length;
  const todaySales = allSales.filter(sale => isToday(new Date(sale.transaction_date || sale.created_at)));
  const todayRevenue = todaySales.reduce((acc, sale) => acc + sale.total, 0);

  // Actually "total" - "amount_received" might be the debt?
  // Let's assume is_debt means partially paid or unpaid.
  // Code Logic: If is_debt, ChangeAmount might be negative? Back end logic:
  // change_amount = req.AmountReceived - total
  // If AmountReceived < Total, ChangeAmount is negative (Debt amount).

  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>

      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Revenue</CardTitle>
            <DollarSign className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">Rp {totalRevenue.toLocaleString()}</div>
            <p className="text-xs text-muted-foreground">All time sales</p>
          </CardContent>
        </Card>

        <Card>
           <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Today's Revenue</CardTitle>
            <TrendingUp className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">Rp {todayRevenue.toLocaleString()}</div>
             <p className="text-xs text-muted-foreground">{todaySales.length} transactions today</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Transactions</CardTitle>
            <ShoppingBag className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{totalTransactions}</div>
             <p className="text-xs text-muted-foreground">Recorded sales</p>
          </CardContent>
        </Card>
      </div>

      <div className="grid gap-4 md:grid-cols-1">
        <Card className="col-span-1">
            <CardHeader>
                <CardTitle>Recent Sales</CardTitle>
            </CardHeader>
            <CardContent>
                 <div className="space-y-4">
                     {allSales.slice(0, 5).map(sale => (
                         <div key={sale.id} className="flex items-center justify-between border-b pb-2 last:border-0 last:pb-0">
                             <div>
                                 <p className="font-medium">
                                     {sale.name ? `${sale.name} - ` : ''}{sale.product}
                                 </p>
                                 <p className="text-sm text-muted-foreground">
                                     {format(new Date(sale.transaction_date || sale.created_at), 'MMM d, HH:mm')}
                                 </p>
                             </div>
                             <div className="font-bold">
                                 Rp {sale.total.toLocaleString()}
                             </div>
                         </div>
                     ))}
                     {allSales.length === 0 && <p className="text-center text-slate-500">No sales recorded yet.</p>}
                 </div>
            </CardContent>
        </Card>
      </div>
    </div>
  );
}
