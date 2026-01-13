export interface Sale {
  id: string;
  name: string;
  product: string;
  quantity: number;
  price: number;
  total: number;
  amount_received: number;
  change_amount: number;
  transaction_date: string;
  is_debt: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateSaleRequest {
  name?: string;
  product: string;
  quantity: number;
  price: number;
  amount_received: number;
  is_debt: boolean;
}

export interface UpdateSaleRequest {
  name?: string;
  product?: string;
  quantity?: number;
  price?: number;
  amount_received?: number;
  is_debt?: boolean;
}

export interface ApiResponse<T> {
  success: boolean;
  message?: string;
  data?: T;
  error?: string;
}
