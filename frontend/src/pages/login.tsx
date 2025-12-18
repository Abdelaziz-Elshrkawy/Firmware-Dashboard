import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export default function Login() {
  return (
    <div className="flex items-center justify-center h-screen">
      <div className="p-6 bg-white rounded shadow w-96">
        <h1 className="mb-4 text-xl font-bold">Login</h1>
        <Input
          type="text"
          placeholder="Username"
          className="w-full mb-2 input"
        />
        <Input
          type="password"
          placeholder="Password"
          className="w-full mb-4 input"
        />
        <Button className="w-full cursor-pointer btn">Login</Button>
      </div>
    </div>
  );
}
