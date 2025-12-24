import AppContextProvider from "@/context/appContext";
import { RouterProvider } from "react-router";
import { router } from "./router";

function App() {
  return (
    <AppContextProvider>
      <RouterProvider router={router} />
    </AppContextProvider>
  );
}

export default App;
