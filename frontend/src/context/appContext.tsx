import { serverEndpoint, serverUrl } from "@/helpers/helpers";
import {
  createContext,
  useContext,
  useEffect,
  useState,
  type ReactNode,
} from "react";

export interface AppContextI {
  state: {
    products: Product[] | undefined;
    devices: Device[] | undefined;
    firmwares: Firmware[] | undefined;
  };
  fetchProducts: () => void;
  fetchDevices: (product_id: number, id?: number) => void;
  fetchFirmwares: (product_id: number, id?: number) => void;
}

const AppContext = createContext<AppContextI | undefined>(undefined);

export default function AppContextProvider({
  children,
}: {
  children: ReactNode;
}) {
  const [products, setProducts] = useState<Product[] | undefined>(undefined);

  const [devices, setDevices] = useState<Device[] | undefined>(undefined);
  const [firmwares, setFirmwares] = useState<Firmware[] | undefined>(undefined);

  const fetchProducts = () => {
    setProducts(undefined);
    fetch(serverUrl + serverEndpoint.product)
      .then(async (data) => {
        const res = await data.json();
        console.log(res.res);
        setTimeout(() => setProducts(res.res), 4000);
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const fetchDevices = (product_id: number, id?: number) => {
    fetch(
      serverUrl + serverEndpoint.device + `?product_id=${product_id}` + id
        ? `id=${id}`
        : ""
    )
      .then(async (data) => {
        const res = await data.json();
        console.log(res.res);
        setDevices(res.res);
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const fetchFirmwares = (product_id: number, id?: number) => {
    fetch(
      serverUrl + serverEndpoint.firmware + `?product_id=${product_id}` + id
        ? `id=${id}`
        : ""
    )
      .then(async (data) => {
        const res = await data.json();
        console.log(res.res);
        setFirmwares(res.res);
      })
      .catch((e) => {
        console.log(e);
      });
  };

  useEffect(() => {
    queueMicrotask(() => {
      fetchProducts();
    });
  }, []);

  return (
    <AppContext
      value={{
        fetchProducts,
        fetchDevices,
        fetchFirmwares,
        state: {
          devices,
          firmwares,
          products,
        },
      }}
    >
      {children}
    </AppContext>
  );
}

export function useAppContext() {
  const ctx = useContext(AppContext);

  if (!ctx)
    throw new Error("App Context can be used only under AppContextProvider");
  return ctx;
}

interface FetchedDataI<T = object> {
  // view: ViewKey;
  data: T | undefined;
  ttl: number | undefined;
  lastUpdate: number;
}
