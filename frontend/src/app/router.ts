import NotFound from "@/components/general/notFound";
import Login from "@/pages/login";
import Products from "@/pages/products";
import { createBrowserRouter } from "react-router";

export const routes = Object.freeze({
  notFound: "*",
  login: "/",
  dashboard: "/dashboard",
});

export const router = createBrowserRouter([
  { path: routes.notFound, Component: NotFound },
  { path: routes.login, Component: Login },
  { path: routes.dashboard, Component: Products },
]);
