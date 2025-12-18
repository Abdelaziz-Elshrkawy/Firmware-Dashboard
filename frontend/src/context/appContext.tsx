import { serverEndpoint, serverUrl } from "@/helpers/helpers";
import { useState } from "react";

export default function AppContext() {
  const [products, setProducts] = useState<
    FetchedDataI<{ id: number; name: string }[]>
  >({
    view: "menu",
    data: undefined,
    ttl: undefined,
  });

  const fetchProducts = () => {
    fetch(serverUrl + serverEndpoint.product)
      .then(async (data) => {
        const res = await data.json();
        console.log(res.res);
        setProducts(res.res);
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const fetchDevices = () => {
    fetch(serverUrl + serverEndpoint.device)
      .then(async (data) => {
        const res = await data.json();
        console.log(res.res);
        setProducts(res.res);
      })
      .catch((e) => {
        console.log(e);
      });
  };
}

type ViewKey =
  | "menu"
  | "devicesList"
  | "deviceEdit"
  | "firmwaresList"
  | "firmwareEdit";

interface FetchedDataI<T = object> {
  view: ViewKey;
  data: T | undefined;
  ttl: Date | undefined;
}
