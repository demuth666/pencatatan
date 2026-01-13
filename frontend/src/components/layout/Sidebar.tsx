import { Link, useLocation } from 'react-router-dom';
import { LayoutDashboard, ShoppingCart, PlusCircle, Menu, ChevronLeft, ChevronRight } from 'lucide-react';
import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Sheet, SheetContent, SheetTrigger } from '@/components/ui/sheet';
import { useState } from 'react';

const sidebarItems = [
  { icon: LayoutDashboard, label: 'Dashboard', href: '/' },
  { icon: ShoppingCart, label: 'Sales', href: '/sales' },
  { icon: PlusCircle, label: 'New Sale', href: '/sales/new' },
];

interface SidebarProps {
  isCollapsed?: boolean;
  toggleSidebar?: () => void;
  mobile?: boolean;
}

export function Sidebar({ isCollapsed = false, toggleSidebar, mobile = false }: SidebarProps) {
  const location = useLocation();

  return (
    <div className={cn(
      "flex flex-col h-full bg-black text-white transition-all duration-300",
      isCollapsed ? "w-20" : "w-64"
    )}>
      <div className="flex items-center justify-between p-4 h-16 border-b border-gray-800">
        {!isCollapsed && (
          <h2 className="text-lg font-bold tracking-tight whitespace-nowrap overflow-hidden">
            Pencatatan
          </h2>
        )}
        {/* Toggle Button for Desktop */}
        {!mobile && toggleSidebar && (
            <Button
              variant="ghost"
              size="icon"
              onClick={toggleSidebar}
              className="ml-auto hover:bg-gray-800 text-gray-400 hover:text-white"
            >
              {isCollapsed ? <ChevronRight className="h-4 w-4" /> : <ChevronLeft className="h-4 w-4" />}
            </Button>
        )}
      </div>

      <div className="flex-1 py-4 space-y-2 px-3">
        {sidebarItems.map((item) => (
          <Link key={item.href} to={item.href}>
            <Button
              variant={location.pathname === item.href ? 'secondary' : 'ghost'}
              className={cn(
                "w-full justify-start transition-all duration-200 mb-1",
                location.pathname === item.href
                  ? "bg-gray-800 text-white hover:bg-gray-700"
                  : "hover:bg-gray-900 hover:text-white text-gray-400",
                isCollapsed ? "justify-center px-2" : "px-4"
              )}
              title={isCollapsed ? item.label : undefined}
            >
              <item.icon className={cn("h-5 w-5", isCollapsed ? "mr-0" : "mr-3")} />
              {!isCollapsed && <span>{item.label}</span>}
            </Button>
          </Link>
        ))}
      </div>

      {/* Footer or extra content can go here */}
    </div>
  );
}

export function MobileSidebar() {
  const [open, setOpen] = useState(false);

  return (
    <Sheet open={open} onOpenChange={setOpen}>
      <SheetTrigger asChild>
        <Button variant="ghost" size="icon" className="md:hidden">
          <Menu className="h-6 w-6" />
        </Button>
      </SheetTrigger>
      <SheetContent side="left" className="p-0 w-72 bg-black border-r-gray-800 text-white">
         <Sidebar mobile={true} />
      </SheetContent>
    </Sheet>
  );
}
