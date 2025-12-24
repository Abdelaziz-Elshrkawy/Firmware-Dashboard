declare interface Product {
  id: number;
  name: string;
}

declare interface Device {
  id: number;
  serial: string;
  product_id: number;
  firmware_id: number;
  api_key: string;
}

declare interface Firmware {
  id: number;
  version: string;
  product_id: number;
}
