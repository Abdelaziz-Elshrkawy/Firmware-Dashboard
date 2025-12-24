import { Input } from "@/components/ui/input";
import { useAppContext } from "@/context/appContext";
import { Delete, RefreshCcw } from "lucide-react";
import { useState } from "react";
import AdminDialog from "./adminDialog";

export default function Products() {
  const [open, setOpen] = useState<number | null>(null);
  const {
    fetchProducts,
    state: { products },
  } = useAppContext();
  const [searchValue, setSearchValue] = useState("");

  return (
    <div className="w-full p-8">
      <h1 className="mb-4 text-2xl font-bold">Products</h1>
      <div className="flex items-center justify-between p-1 mb-4 border-2 rounded-md">
        <div className="flex items-center justify-between w-11/12 p-1 border-2 rounded-md">
          <Input
            placeholder="Search fro product"
            className="border-0 outline-0"
            value={searchValue}
            onChange={({ target: { value } }) => {
              setSearchValue(value);
            }}
          />
          <Delete
            className="text-red-400 cursor-pointer"
            onClick={() => {
              setSearchValue("");
            }}
          />
        </div>
        <RefreshCcw
          onClick={fetchProducts}
          className={`w-1/12 cursor-pointer ${
            products ? "" : "animate-spin text-green-400"
          }`}
        />
      </div>
      <ul className="p-2 mb-6 shadow-md">
        {products && products.length > 0
          ? (searchValue.length > 0
              ? products.filter((e) =>
                  e.name
                    .toLocaleLowerCase()
                    .includes(searchValue.toLocaleLowerCase())
                )
              : products
            )?.map((p) => (
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
