import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import AppLayout from '@/components/layout/AppLayout';
import { Toaster } from '@/components/ui/toaster';

import SalesList from '@/pages/SalesList';

import CreateSale from '@/pages/CreateSale';
import EditSale from '@/pages/EditSale';

import Dashboard from '@/pages/Dashboard';

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route element={<AppLayout />}>
            <Route path="/" element={<Dashboard />} />
            <Route path="/sales" element={<SalesList />} />
            <Route path="/sales/new" element={<CreateSale />} />
            <Route path="/sales/:id/edit" element={<EditSale />} />
          </Route>
        </Routes>
        <Toaster />
      </BrowserRouter>
    </QueryClientProvider>
  );
}

export default App;
