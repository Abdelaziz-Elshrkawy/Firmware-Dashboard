import { routes } from "@/app/router";
import { Button } from "@/components/ui/button"; // Shadcn Button
import { ArrowLeftCircle } from "lucide-react";
import { useNavigate } from "react-router";

export default function NotFound() {
  const navigate = useNavigate();

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-background text-foreground">
      <div className="flex flex-col items-center justify-center space-y-4 text-center">
        <h1 className="text-5xl font-bold">404</h1>
        <p className="text-lg text-muted-foreground">
          Sorry, the page you are looking for does not exist.
        </p>
        <Button
          variant="outline"
          size="lg"
          className="flex items-center gap-2 mt-4 cursor-pointer group/back-btn"
          onClick={() =>
            navigate(routes.dashboard, {
              replace: true,
            })
          }
        >
          <ArrowLeftCircle className="w-5 h-5 group-hover/back-btn:-translate-x-1.25 transition-all " />
          Back to Dashboard
        </Button>
      </div>
    </div>
  );
}
