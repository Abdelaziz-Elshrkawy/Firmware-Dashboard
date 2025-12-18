import { serverEndpoint, serverUrl } from "@/helpers/helpers";
import { useEffect, useState } from "react";
import AdminDialog from "./adminDialog";

export default function Products() {
  const [open, setOpen] = useState<number | null>(null);
  const [products, setProducts] = useState<{ id: number; name: string }[]>([]);

  useEffect(() => {
    fetch(serverUrl + serverEndpoint.product)
      .then(async (data) => {
        const res = await data.json();
        console.log(res.res);
        setProducts(res.res);
      })
      .catch((e) => {
        console.log(e);
      });
  }, []);

  return (
    <div className="w-full p-8">
      <h1 className="mb-4 text-2xl font-bold">Products</h1>

      <ul className="p-2 mb-6 shadow-md">
        {products.length > 0
          ? products?.map((p) => (
              <li
                onClick={() => setOpen(p.id)}
                key={p.id}
                className="p-2 mb-2 border rounded shadow-md cursor-pointer"
              >
                {p.name}
              </li>
            ))
          : "loading..."}
      </ul>

      <AdminDialog open={open} setOpen={setOpen} />
    </div>
  );
}
