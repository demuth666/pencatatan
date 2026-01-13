import { Outlet } from 'react-router-dom';
import { Sidebar, MobileSidebar } from './Sidebar';
import { useState } from 'react';
import { cn } from '@/lib/utils';

export default function AppLayout() {
  const [isCollapsed, setIsCollapsed] = useState(false);

  return (
    <div className="flex h-screen overflow-hidden bg-slate-50">
      {/* Desktop Sidebar */}
      <aside className={cn(
        "hidden md:flex flex-col fixed inset-y-0 z-50 transition-all duration-300",
        isCollapsed ? "w-20" : "w-64"
      )}>
        <Sidebar
          isCollapsed={isCollapsed}
          toggleSidebar={() => setIsCollapsed(!isCollapsed)}
        />
      </aside>

      {/* Main Content */}
      <main className={cn(
        "flex-1 flex flex-col h-full overflow-hidden transition-all duration-300",
        isCollapsed ? "md:pl-20" : "md:pl-64"
      )}>
        {/* Header */}
        <header className="h-16 border-b bg-white flex items-center px-6 md:hidden">
          <MobileSidebar />
          <h1 className="ml-4 text-xl font-bold">Pencatatan</h1>
        </header>

        {/* Page Content */}
        <div className="flex-1 overflow-y-auto p-6">
          <Outlet />
        </div>
      </main>
    </div>
  );
}
